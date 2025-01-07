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

	// commented
	// contracts, err := is.query.ListContractByRoomId(context.Background(), &body.RoomID)
	// if err != nil {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    err.Error(),
	// 		Data:       false,
	// 	}
	// }

	// var matchedContract responses.MContractOnChainRes
	// for _, contract := range contracts {
	// 	onChainContract, err := is.blockchain.GetMContractByIDOnChain(int64(contract))
	// 	if err != nil {
	// 		return &responses.ResponseData{
	// 			StatusCode: http.StatusInternalServerError,
	// 			Message:    err.Error(),
	// 			Data:       false,
	// 		}
	// 	}

	// 	if onChainContract.PreRentalStatus == 2 {
	// 		matchedContract = *onChainContract
	// 		break
	// 	}
	// }

	// if !(matchedContract.PreRentalStatus == 2 && matchedContract.RentalProcessStatus != 0) {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    "Trạng thái hợp đồng không hợp lệ",
	// 		Data:       false,
	// 	}
	// }

	//commented
	// user, _ := is.query.GetUserByID(context.Background(), userid)
	// _, err = is.blockchain.InputMeterReadingOnChain(*user.PrivateKeyHex, matchedContract.ID)
	// if err != nil {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    "blockchain error",
	// 		Data:       false,
	// 	}
	// }

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    "Created",
		Data:       result,
	}
}

func (is *IndexServiceImpl) GetAllIndex(userid int32, month int32, year int32, mType string) *responses.ResponseData {
	roomDetails, err := is.query.GetRoomsByOwner(context.Background(), userid)
	if err != nil {
		fmt.Printf("Error fetching rooms by owner: %v\n", err)
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	groupedResults := make(map[string][]map[string]interface{})

	for _, room := range roomDetails {
		// Create local variables for month and year for each room
		roomMonth := month
		roomYear := year
		address := strings.Join(room.Address, ", ")

		indexParam := dataaccess.GetAllIndexParams{
			ID:    room.ID,
			Year:  roomYear,
			Month: roomMonth,
		}
		rs, er := is.query.GetAllIndex(context.Background(), indexParam)

		fmt.Printf("RoomID: %d, Month: %d, Year: %d, Records Found: %d\n", room.ID, roomMonth, roomYear, len(rs))

		if er != nil || len(rs) == 0 {
			// Try the previous month if no record is found
			if roomMonth == 1 {
				roomMonth = 12
				roomYear--
			} else {
				roomMonth--
			}
			indexParam.Year = roomYear
			indexParam.Month = roomMonth
			rs, er = is.query.GetAllIndex(context.Background(), indexParam)

			fmt.Printf("RoomID: %d, Month: %d, Year: %d, Records Found: %d (Previous Month)\n", room.ID, roomMonth, roomYear, len(rs))

			if er != nil || len(rs) == 0 {
				// No records found even for the previous month
				groupedResults[address] = append(groupedResults[address], map[string]interface{}{
					"room_id":     room.ID,
					"room_number": room.RoomNumber,
					"old_index":   nil,
					"new_index":   nil,
					"used":        nil,
				})
				continue
			} else {
				// Previous month's record found
				record := rs[len(rs)-1]
				if mType == "water" {
					groupedResults[address] = append(groupedResults[address], map[string]interface{}{
						"room_id":     room.ID,
						"room_number": room.RoomNumber,
						"old_index":   record.CurrWater,
						"new_index":   nil,
						"used":        nil,
					})
				} else {
					groupedResults[address] = append(groupedResults[address], map[string]interface{}{
						"room_id":     room.ID,
						"room_number": room.RoomNumber,
						"old_index":   record.CurrElectricity,
						"new_index":   nil,
						"used":        nil,
					})
				}
				continue
			}
		}

		// Process the found records
		for _, record := range rs {
			indexInfo := map[string]interface{}{
				"room_id":     room.ID,
				"room_number": room.RoomNumber,
			}
			if mType == "water" {
				indexInfo["old_index"] = record.PrevWater
				indexInfo["new_index"] = record.CurrWater
				if record.CurrWater != nil && record.PrevWater != nil {
					indexInfo["used"] = *record.CurrWater - record.PrevWater.(float64)
				} else {
					indexInfo["used"] = nil
				}
			} else {
				indexInfo["old_index"] = record.PrevElectricity
				indexInfo["new_index"] = record.CurrElectricity
				if record.CurrElectricity != nil && record.PrevElectricity != nil {
					indexInfo["used"] = *record.CurrElectricity - record.PrevElectricity.(float64)
				} else {
					indexInfo["used"] = nil
				}
			}
			groupedResults[address] = append(groupedResults[address], indexInfo)
		}
	}

	// Convert grouped results to the desired format
	var result []map[string]interface{}
	for address, indexInfos := range groupedResults {
		result = append(result, map[string]interface{}{
			"address":    address,
			"index_info": indexInfos,
		})
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Ok",
		Data:       result,
	}
}
