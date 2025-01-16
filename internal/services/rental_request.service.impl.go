package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type RentalRequestServiceImpl struct {
	repo                *dataaccess.Queries
	notificationService NotificationService
}

func NewRentalRequestServiceImpl(notification NotificationService) RentalRequestService {
	return &RentalRequestServiceImpl{
		repo:                dataaccess.New(global.Db),
		notificationService: notification,
	}
}

// GetRentalRequestByRoomID implements RentalRequestService.
func (rentalService *RentalRequestServiceImpl) GetRentalRequestByRoomID(roomID int) *responses.ResponseData {
	requests, err := rentalService.repo.GetRequestByRoomID(context.Background(), int32(roomID))

	if len(requests) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusOK,
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
		if rentStatus.Status != 3 && rentStatus.Status != 4 && !rentStatus.DeletedAt.Valid {
			return &responses.ResponseData{
				StatusCode: http.StatusNotAcceptable,
				Message:    "Bạn đã gửi yêu cầu thuê cho phòng này",
				Data:       false,
			}
		}
		fmt.Println(rentStatus.DeletedAt)
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
		// rentalService.repo.DeleteRequest(context.Background(), res.ID)

		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to finish your request",
			Data:       nil,
		}
	}

	id := int(res.ID)
	rentalService.notificationService.SendNotification(int(rs.Owner), "Bạn có một yêu cầu thuê phòng mới", &id, "rental_request")

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
			StatusCode: http.StatusOK,
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
		StatusCode: http.StatusOK,
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

	var result []map[string]interface{}

	if user.Role == 1 {
		// Role = 1: Owner
		requests, err := rentalService.repo.GetRequestByOwnerID(context.Background(), int32(userID))
		if err != nil {
			fmt.Println(err.Error())
			return &responses.ResponseData{
				StatusCode: http.StatusOK,
				Message:    "Nothing found",
				Data:       "nothing",
			}
		}

		for _, req := range requests {
			// Decode request_info JSON
			var requestInfo []map[string]interface{}
			if err := json.Unmarshal([]byte(req.RequestInfo), &requestInfo); err != nil {
				fmt.Println("Error decoding request_info:", err)
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    "Failed to decode request info",
					Data:       nil,
				}
			}

			room, err := rentalService.repo.GetRoomByID(context.Background(), *req.RoomID)
			if err != nil {
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
					Data:       nil,
				}
			}

			// Build response for each room
			result = append(result, map[string]interface{}{
				"room": map[string]interface{}{
					"id":      room.ID,
					"title":   room.Title,
					"price":   room.TotalPrice,
					"address": room.Address,
				},
				"request_count": req.RequestCount,
				"request_info":  requestInfo,
			})
		}
	} else if user.Role == 0 {
		// Role = 0: Sender
		requests, err := rentalService.repo.GetRequestBySenderID(context.Background(), int32(userID))
		if err != nil {
			fmt.Println(err.Error())
			return &responses.ResponseData{
				StatusCode: http.StatusOK,
				Message:    "Nothing found",
				Data:       "nothing",
			}
		}

		for _, req := range requests {
			// Decode request_info JSON
			var requestInfo []map[string]interface{}
			if err := json.Unmarshal([]byte(req.RequestInfo), &requestInfo); err != nil {
				fmt.Println("Error decoding request_info:", err)
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    "Failed to decode request info",
					Data:       nil,
				}
			}

			room, err := rentalService.repo.GetRoomByID(context.Background(), *req.RoomID)
			if err != nil {
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
					Data:       nil,
				}
			}

			// Build response for each room
			result = append(result, map[string]interface{}{
				"room": map[string]interface{}{
					"id":      room.ID,
					"title":   room.Title,
					"price":   room.TotalPrice,
					"address": room.Address,
				},
				"request_count": req.RequestCount,
				"request_info":  requestInfo,
			})
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data: result,
	}
}

func (rentalService *RentalRequestServiceImpl) ReviewRentalRequest(result string, reqId int32, userid int32) *responses.ResponseData {
	request, err := rentalService.repo.GetRequestByID(context.Background(), reqId)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	var status int
	var trackingParam dataaccess.CreateProcessTrackingParams
	id := int(reqId)

	if result == "approve" {
		status = 2
		trackingParam = dataaccess.CreateProcessTrackingParams{
			Actor:     userid,
			Action:    "Chủ nhà đã chấp nhận yêu cầu thuê phòng",
			RequestID: reqId,
		}

		rentalService.notificationService.SendNotification(int(request.SenderID), "Chủ nhà đã chấp nhận yêu cầu thuê phòng của bạn", &id, "rental_request")
	} else if result == "decline" {
		actor, err := rentalService.repo.GetUserByID(context.Background(), userid)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}

		if actor.Role == 1 {
			status = 3
			trackingParam = dataaccess.CreateProcessTrackingParams{
				Actor:     userid,
				Action:    "Chủ nhà đã từ chối yêu cầu thuê phòng",
				RequestID: reqId,
			}

			rentalService.notificationService.SendNotification(int(request.SenderID), "Chủ nhà đã từ chối yêu cầu thuê phòng của bạn", &id, "rental_request")
		} else {
			status = 4
			trackingParam = dataaccess.CreateProcessTrackingParams{
				Actor:     userid,
				Action:    "Người thuê đã huỷ yêu cầu thuê phòng",
				RequestID: reqId,
			}
			roomDetail, err := rentalService.repo.GetRoomByID(context.Background(), request.RoomID)
			if err != nil {
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
					Data:       nil,
				}
			}

			rentalService.notificationService.SendNotification(int(roomDetail.Owner), "Người thuê đã huỷ yêu cầu thuê phòng", &id, "rental_request")
		}
	}

	param := dataaccess.UpdateRequestStatusByIdParams{
		Status: int32(status),
		ID:     reqId,
	}

	err = rentalService.repo.UpdateRequestStatusById(context.Background(), param)

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	_, er := rentalService.repo.CreateProcessTracking(context.Background(), trackingParam)
	if er != nil {
		// fmt.Println("ERROR Create process tracking failed!!")
		// rentalService.repo.DeleteRequest(context.Background(), reqId)

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
