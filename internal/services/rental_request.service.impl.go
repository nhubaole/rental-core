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

func (rentalService *RentalRequestServiceImpl) CreateRentalRequest(body *requests.CreateRentalRequest, userid int32) *responses.ResponseData {
	// check if the room reqId existed
	rs, checkEr := rentalService.repo.GetRoomByID(context.Background(), body.RoomID)
	if checkEr != nil {
		fmt.Println(checkEr.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "We can't find this room",
			Data:       "nothing here",
		}
	}
	if rs.Owner == userid {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "You can't rent your room!! You have already owned it",
			Data:       false,
		}
	}
	request := dataaccess.CheckRequestExistedParams{
		SenderID: userid,
		RoomID:   body.RoomID,
	}
	// check if there is already a rental request
	rentStatus, checkError2 := rentalService.repo.CheckRequestExisted(context.Background(), request)
	// create retal request
	if checkError2 == nil {
		if rentStatus.Status != 3 {
			return &responses.ResponseData{
				StatusCode: http.StatusNotAcceptable,
				Message:    "You have to wait until the owner responds",
				Data:       "nothing here",
			}
		}
		fmt.Println(rentStatus.DeletedAt)
		if !rentStatus.DeletedAt.Valid {
			return &responses.ResponseData{
				StatusCode: http.StatusNotAcceptable,
				Message:    "You have to wait until the owner responds",
				Data:       "nothing here",
			}
		}
	}
	// parse to the new body
	parseBody := dataaccess.CreateRentalRequestParams{
		SenderID:        userid,
		Code:            "",
		RoomID:          body.RoomID,
		SuggestedPrice:  body.SuggestedPrice,
		NumOfPerson:     body.NumOfPerson,
		BeginDate:       body.BeginDate,
		EndDate:         body.EndDate,
		AdditionRequest: body.AdditionRequest,
		Status:          1,
	}
	// add things
	parseBody.Status = 1
	parseBody.SenderID = userid
	mytime := int(time.Now().UnixNano() / int64(time.Microsecond))
	parseBody.Code = common.GenerateCode("RR", int(userid), int(body.RoomID), mytime)

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

	trackingParam := dataaccess.CreateProcessTrackingParams{
		Actor:     userid,
		Action:    "Người thuê tạo yêu cầu thuê phòng thành công",
		RequestID: res.ID,
	}
	_, er := rentalService.repo.CreateProcessTracking(context.Background(), trackingParam)
	if er != nil {
		fmt.Println("ERROR Create process tracking failed!!")

		// ignore error because it's already an error
		rentalService.repo.DeleteRequest(context.Background(), res.ID)

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

func (rentalService *RentalRequestServiceImpl) DeleteRentalRequest(rentid int32, userid int32) *responses.ResponseData {
	// Find rental request
	result, er := rentalService.repo.GetRequestByID(context.Background(), rentid)
	if er != nil || result.SenderID != userid {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "You do not have this rental request",
			Data:       "nothing",
		}
	}
	if result.Status != 1 {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "You can not delete this rental request",
			Data:       "nothing",
		}
	}
	if result.DeletedAt.Valid {
		return &responses.ResponseData{
			StatusCode: http.StatusNoContent,
			Message:    "Already Deleted",
			Data:       "nothing",
		}
	}

	err := rentalService.repo.DeleteRequest(context.Background(), rentid)
	if err != nil {
		fmt.Println(err.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Something wrong happened",
			Data:       "nothing",
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusNoContent,
		Message:    "Deleted",
		Data:       "nothing",
	}
}

func (rentalService *RentalRequestServiceImpl) GetRentalRequestById(rentid int32, userid int32) *responses.ResponseData {
	// Find rental request
	result, er := rentalService.repo.GetRequestByID(context.Background(), rentid)
	if er != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "You do not have this rental request",
			Data:       "nothing",
		}
	}
	// check room owner and user
	result2, err := rentalService.repo.GetRoomByID(context.Background(), result.RoomID)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusOK,
			Message:    "Nothing found",
			Data:       "nothing",
		}
	}
	if result.SenderID != userid && result2.Owner != userid {
		return &responses.ResponseData{
			StatusCode: http.StatusOK,
			Message:    "Nothing found",
			Data:       "nothing",
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       result,
	}
}

func (rentalService *RentalRequestServiceImpl) GetAllRentalRequest(phoneNumber string) *responses.ResponseData {
	// Find rental request and check if the owner or renter is inside it

	result, er := rentalService.repo.GetUserByPhone(context.Background(), phoneNumber)
	if er != nil {
		fmt.Println(er.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Who are you buddy?",
			Data:       "nothing",
		}
	}

	finalResult, err := rentalService.repo.GetRequestByUserID(context.Background(), result.ID)
	if err != nil {
		fmt.Println(err.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusOK,
			Message:    "Nothing found",
			Data:       "nothing",
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       finalResult,
	}

}

func (rentalService *RentalRequestServiceImpl) ReviewRentalRequest(result string, reqId int32, userid int32) *responses.ResponseData {
	// check if this user is the owner
	checkRoom, er := rentalService.repo.GetRequestByUserID(context.Background(), userid)
	if er != nil {
		fmt.Println(er.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "No rental request found",
			Data:       false,
		}
	}
	for _, room := range checkRoom {
		if room.ID == reqId {
			if room.SenderID != userid {

				temp := dataaccess.UpdateRequestStatusByIdParams{
					ID: reqId,
				}
				if result == "approve" {
					temp.Status = 2
				} else if result == "decline" {
					temp.Status = 3
				}
				err := rentalService.repo.UpdateRequestStatusById(context.Background(), temp)
				str := ""
				if temp.Status == 2 {
					str = "Người cho thuê đã đồng ý yêu cầu thuê"
				} else {
					str = "Người cho thuê không đồng ý yêu cầu thuê"
				}

				// log to process tracking
				trackingParam := dataaccess.CreateProcessTrackingParams{
					Actor:     userid,
					Action:    str,
					RequestID: reqId,
				}
				_, er := rentalService.repo.CreateProcessTracking(context.Background(), trackingParam)

				if er != nil {
					fmt.Println("ERROR Create process tracking failed!!")
					return &responses.ResponseData{
						StatusCode: http.StatusInternalServerError,
						Message:    "Bad",
						Data:       "An error occured",
					}
				}

				if err != nil {
					fmt.Println(err.Error())
					return &responses.ResponseData{
						StatusCode: http.StatusInternalServerError,
						Message:    "Bad",
						Data:       "An error occured",
					}
				}
				return &responses.ResponseData{
					StatusCode: http.StatusAccepted,
					Message:    "Success",
					Data:       true,
				}
			}
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusBadRequest,
		Message:    "We can't find this room",
		Data:       false,
	}

}
