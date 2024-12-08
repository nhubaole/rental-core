package services

import (
	"context"
	"fmt"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type IndexServiceImpl struct {
	query *dataaccess.Queries
	blockchain BlockchainService
}

func NewIndexServiceImpl(blockchain BlockchainService) IndexService {
	return &IndexServiceImpl{
		query: dataaccess.New(global.Db),
		blockchain: blockchain,
	}
}

func (is *IndexServiceImpl) CreateIndex(userid int32, body *dataaccess.CreateIndexParams) *responses.ResponseData {
	// Find room if match the owner
	rs, er := is.query.GetRoomByID(context.Background(), body.RoomID)

	//TODO: get contract id??
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

	result, error := is.query.CreateIndex(context.Background(), *body)
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

	if !(matchedContract.PreRentalStatus == 2 && matchedContract.RentalProcessStatus != 0) {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Trạng thái hợp đồng không hợp lệ",
			Data:       false,
		}
	}

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

func (is *IndexServiceImpl) GetAllIndex(userid int32, month int32, year int32) *responses.ResponseData {
	// Find rooms whose the owner's id is userid
	param := dataaccess.GetIndexByOwnerIdParams{
		Owner: userid,
		Year:  month,
		Month: year,
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
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Ok",
		Data:       rs,
	}
}
