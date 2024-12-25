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
)

type RentalRequestServiceImpl struct {
	repo *dataaccess.Queries
}

func NewRentalRequestServiceImpl() RentalRequestService {
	return &RentalRequestServiceImpl{repo: dataaccess.New(global.Db)}
}

// GetRentalRequestByRoomID implements RentalRequestService.
func (rentalService *RentalRequestServiceImpl) GetRentalRequestByRoomID(roomID int) *responses.ResponseData {
	requests, err := rentalService.repo.GetRequestByRoomID(context.Background(), int32(roomID))

	if len(requests) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusNoContent,
			Message:    responses.StatusResourceNotFound,
			Data:       nil,
		}
	}
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	var detailedRequests []responses.GetRequestByRoomIDRes
	for _, request := range requests {
		// Lấy thông tin chi tiết của người gửi dựa trên sender_id
		sender, err := rentalService.repo.GetUserByID(context.Background(), request.SenderID)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}
		var user responses.UserResponse
		common.MapStruct(sender, &user)

		// Xây dựng đối tượng chi tiết
		detailedRequest := responses.GetRequestByRoomIDRes{
			ID:              int(request.ID),
			Code:            request.Code,
			Sender:          user,
			RoomID:          request.RoomID,
			SuggestedPrice:  request.SuggestedPrice,
			NumOfPerson:     request.NumOfPerson,
			BeginDate:       request.BeginDate,
			EndDate:         request.EndDate,
			AdditionRequest: request.AdditionRequest,
			Status:          request.Status,
			CreatedAt:       request.CreatedAt,
			UpdatedAt:       request.UpdatedAt,
		}

		// Thêm đối tượng vào danh sách kết quả
		detailedRequests = append(detailedRequests, detailedRequest)
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       detailedRequests,
	}
}
func (rentalService *RentalRequestServiceImpl) CreateRentalRequest(body *requests.CreateRentalRequest, userid int32) *responses.ResponseData {
	// check if the room reqId existed
	rs, checkEr := rentalService.repo.GetRoomByID(context.Background(), body.RoomID)
	if checkEr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Phòng không tồn tại",
			Data:       false,
		}
	}
	if rs.Owner == userid {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Bạn không thể thực hiện thao tác này",
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
				Message:    "Bạn đã gửi yêu cầu thuê cho phòng này",
				Data:       false,
			}
		}
		fmt.Println(rentStatus.DeletedAt)
		if !rentStatus.DeletedAt.Valid {
			return &responses.ResponseData{
				StatusCode: http.StatusNotAcceptable,
				Message:    "Bạn đã gửi yêu cầu thuê cho phòng này",
				Data:       false,
			}
		}
	}
	// parse to the new body
	parseBody := dataaccess.CreateRentalRequestParams{
		SenderID:        userid,
		RoomID:          body.RoomID,
		SuggestedPrice:  body.SuggestedPrice,
		NumOfPerson:     body.NumOfPerson,
		BeginDate:       body.BeginDate,
		EndDate:         body.EndDate,
		AdditionRequest: body.AdditionRequest,
		Status:          1,
	}
	parseBody.Code = common.GenerateCode("YC")

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
	request, er := rentalService.repo.GetRequestByID(context.Background(), rentid)
	if er != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "You do not have this rental request",
			Data:       "nothing",
		}
	}
	// check room owner and user
	room, err := rentalService.repo.GetRoomByID(context.Background(), request.RoomID)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	if request.SenderID != userid && room.Owner != userid {
		return &responses.ResponseData{
			StatusCode: http.StatusOK,
			Message:    "Nothing found",
			Data:       nil,
		}
	}

	sender, err := rentalService.repo.GetUserByID(context.Background(), request.SenderID)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       "nothing",
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data: responses.GetRentalRequestByIDRes{
			ID:              request.ID,
			Code:            request.Code,
			Sender:          sender,
			Room:            room,
			SuggestedPrice:  request.SuggestedPrice,
			NumOfPerson:     request.NumOfPerson,
			BeginDate:       request.BeginDate,
			EndDate:         request.EndDate,
			AdditionRequest: request.AdditionRequest,
			Status:          request.Status,
			CreatedAt:       request.CreatedAt,
			UpdatedAt:       request.UpdatedAt,
		},
	}
}

func (rentalService *RentalRequestServiceImpl) GetAllRentalRequest(userID int) *responses.ResponseData {
	user, err := rentalService.repo.GetUserByID(context.Background(), int32(userID))
	if err != nil {
		fmt.Println(err.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusOK,
			Message:    "Nothing found",
			Data:       "nothing",
		}
	}
	var result []responses.GetRentalRequestByIDRes
	if user.Role == 1 {

		requests, err := rentalService.repo.GetRequestByOwnerID(context.Background(), int32(userID))
		if err != nil {
			fmt.Println(err.Error())
			return &responses.ResponseData{
				StatusCode: http.StatusOK,
				Message:    "Nothing found",
				Data:       "nothing",
			}
		}

		for _, v := range requests {
			room, err := rentalService.repo.GetRoomByID(context.Background(), v.RoomID)
			if err != nil {
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
					Data:       nil,
				}
			}
			sender, err := rentalService.repo.GetUserByID(context.Background(), v.SenderID)
			if err != nil {
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
					Data:       "nothing",
				}
			}
			req := responses.GetRentalRequestByIDRes{
				ID:              v.ID,
				Code:            v.Code,
				Sender:          sender,
				Room:            room,
				SuggestedPrice:  v.SuggestedPrice,
				NumOfPerson:     v.NumOfPerson,
				BeginDate:       v.BeginDate,
				EndDate:         v.EndDate,
				AdditionRequest: v.AdditionRequest,
				Status:          v.Status,
				CreatedAt:       v.CreatedAt,
				UpdatedAt:       v.UpdatedAt,
			}

			result = append(result, req)

		}
	} else if user.Role == 0 {
		requests, err := rentalService.repo.GetRequestBySenderID(context.Background(), int32(userID))
		if err != nil {
			fmt.Println(err.Error())
			return &responses.ResponseData{
				StatusCode: http.StatusOK,
				Message:    "Nothing found",
				Data:       "nothing",
			}
		}

		for _, v := range requests {
			room, err := rentalService.repo.GetRoomByID(context.Background(), v.RoomID)
			if err != nil {
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
					Data:       nil,
				}
			}
			sender, err := rentalService.repo.GetUserByID(context.Background(), v.SenderID)
			if err != nil {
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
					Data:       "nothing",
				}
			}
			req := responses.GetRentalRequestByIDRes{
				ID:              v.ID,
				Code:            v.Code,
				Sender:          sender,
				Room:            room,
				SuggestedPrice:  v.SuggestedPrice,
				NumOfPerson:     v.NumOfPerson,
				BeginDate:       v.BeginDate,
				EndDate:         v.EndDate,
				AdditionRequest: v.AdditionRequest,
				Status:          v.Status,
				CreatedAt:       v.CreatedAt,
				UpdatedAt:       v.UpdatedAt,
			}

			result = append(result, req)

		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data: map[string]interface{}{
			"count":    len(result),
			"requests": result,
		},
	}

}

func (rentalService *RentalRequestServiceImpl) ReviewRentalRequest(result string, reqId int32, userid int32) *responses.ResponseData {
	var status int
	var trackingParam dataaccess.CreateProcessTrackingParams
	if result == "approve" {
		status = 2
		trackingParam = dataaccess.CreateProcessTrackingParams{
			Actor:     userid,
			Action:    "Chủ nhà đã chấp nhận yêu cầu thuê phòng",
			RequestID: reqId,
		}
	} else if result == "decline" {
		status = 3
		trackingParam = dataaccess.CreateProcessTrackingParams{
			Actor:     userid,
			Action:    "Chủ nhà đã từ chối yêu cầu thuê phòng",
			RequestID: reqId,
		}
	}

	param := dataaccess.UpdateRequestStatusByIdParams{
		Status: int32(status),
		ID:     reqId,
	}

	err := rentalService.repo.UpdateRequestStatusById(context.Background(), param)

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}


	_, er := rentalService.repo.CreateProcessTracking(context.Background(), trackingParam)
	if er != nil {
		fmt.Println("ERROR Create process tracking failed!!")

		rentalService.repo.DeleteRequest(context.Background(), reqId)

		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to finish your request",
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       reqId,
	}
}
