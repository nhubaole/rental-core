package services

import (
	"context"
	"fmt"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"time"
)

type RentalRequestServiceImpl struct {
	repo *dataaccess.Queries
}

func NewRentalRequestServiceImpl() RentalRequestService {
	return &RentalRequestServiceImpl{repo: dataaccess.New(global.Db)}
}

func (rentalService *RentalRequestServiceImpl) CreateRentalRequest(body *requests.CreateRentalRequest, myid int32) *responses.ResponseData {
	// check if there is already a rental request and if the room is created
	request := dataaccess.CheckRequestExistedParams{
		SenderID: myid,
		RoomID:   body.RoomID,
	}
	rentStatus, checkError2 := rentalService.repo.CheckRequestExisted(context.Background(), request)
	// create retal request
	if checkError2 == nil {
		if rentStatus != 3 {
			return &responses.ResponseData{
				StatusCode: http.StatusNotAcceptable,
				Message:    "You have to wait until the owner responds",
				Data:       "nothing here",
			}
		}
	} else {
		fmt.Println(checkError2.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "We can't find this room",
			Data:       "nothing here",
		}
	}
	// parse to the new body
	parseBody := dataaccess.CreateRentalRequestParams{
		SenderID:        myid,
		Code:            "",
		RoomID:          body.RoomID,
		SuggestedPrice:  body.SuggestedPrice,
		NumOfPerson:     body.NumOfPerson,
		BeginDate:       body.BeginDate,
		EndDate:         body.EndDate,
		AdditionRequest: body.AdditionRequest,
		Status:          1,
	}
	parseBody.Status = 1
	parseBody.SenderID = myid
	mytime := int(time.Now().UnixNano() / int64(time.Microsecond))
	parseBody.Code = common.GenerateCode("RR", int(myid), int(body.RoomID), mytime)

	// push to database
	res, err := rentalService.repo.CreateRentalRequest(context.Background(), parseBody)
	if err != nil {
		println(string(err.Error()))
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to finish your request",
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       res,
	}
}
