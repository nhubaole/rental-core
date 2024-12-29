package services

import (
	"context"
	"fmt"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
	"strings"
)

type IndexServiceImpl struct {
	query      *dataaccess.Queries
	blockchain BlockchainService
}

func NewIndexServiceImpl(blockchain BlockchainService) IndexService {
	return &IndexServiceImpl{
		query:      dataaccess.New(global.Db),
		blockchain: blockchain,
	}
}

func (is *IndexServiceImpl) CreateIndex(userid int32, body *dataaccess.UpsertIndexParams) *responses.ResponseData {
	rs, er := is.query.GetRoomByID(context.Background(), body.RoomID)

	if er != nil {
		fmt.Println(er.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to find the room!",
			Data:       false,
		}
	}
	if rs.Owner != userid {
		fmt.Println("ERROR user is not the owner")
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to find the room!",
			Data:       false,
		}
	}

	result, error := is.query.UpsertIndex(context.Background(), *body)
	if error != nil {
		fmt.Println(error.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Some error occured",
			Data:       false,
		}
	}

	contracts, err := is.query.ListContractByRoomId(context.Background(), &body.RoomID)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	var matchedContract responses.MContractOnChainRes
	for _, contract := range contracts {
		onChainContract, err := is.blockchain.GetMContractByIDOnChain(int64(contract))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}

		if onChainContract.PreRentalStatus == 2 {
			matchedContract = *onChainContract
			break
		}
	}

	// if !(matchedContract.PreRentalStatus == 2 && matchedContract.RentalProcessStatus != 0) {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    "Trạng thái hợp đồng không hợp lệ",
	// 		Data:       false,
	// 	}
	// }

	user, _ := is.query.GetUserByID(context.Background(), userid)
	_, err = is.blockchain.InputMeterReadingOnChain(*user.PrivateKeyHex, matchedContract.ID)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "blockchain error",
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    "Created",
		Data:       result,
	}
}

func (is *IndexServiceImpl) GetAllIndex(userid int32, month int32, year int32, mType string) *responses.ResponseData {
	param := dataaccess.GetIndexByOwnerIdParams{
		Owner: userid,
		Year:  year,
		Month: month,
	}

	rs, er := is.query.GetIndexByOwnerId(context.Background(), param)
	if er != nil {
		fmt.Println(er.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Unable to find indexes!",
			Data:       false,
		}
	}

	roomData := make(map[int32]map[string]interface{})
	for _, record := range rs {
		roomDetails, err := is.query.GetRoomByID(context.Background(), *record.RoomID)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}

		if _, exists := roomData[*record.RoomID]; !exists {
			roomData[*record.RoomID] = map[string]interface{}{
				"room_id":    record.RoomID,
				"address":    strings.Join(roomDetails.Address, ", "),
				"index_info": []map[string]interface{}{},
			}
		}

		if mType == "water" {
			indexInfo := map[string]interface{}{
				"room_number": roomDetails.RoomNumber,
				"old_index":   record.PrevWater.(float64),
				"new_index":   record.CurrWater,
				"used":        *record.CurrWater - record.PrevWater.(float64),
			}
			roomData[*record.RoomID]["index_info"] = append(roomData[*record.RoomID]["index_info"].([]map[string]interface{}), indexInfo)
		} else {
			indexInfo := map[string]interface{}{
				"room_number": roomDetails.RoomNumber,
				"old_index":   record.PrevElectricity.(float64),
				"new_index":   record.CurrElectricity,
				"used":        *record.CurrElectricity - record.PrevElectricity.(float64),
			}
			roomData[*record.RoomID]["index_info"] = append(roomData[*record.RoomID]["index_info"].([]map[string]interface{}), indexInfo)
		}
	}

	var result []map[string]interface{}
	for _, data := range roomData {
		result = append(result, data)
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Ok",
		Data:       result,
	}
}
