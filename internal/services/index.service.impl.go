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
}

func NewIndexServiceImpl() IndexService {
	return &IndexServiceImpl{
		query: dataaccess.New(global.Db),
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
