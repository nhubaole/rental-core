package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/responses"
)

type ReturnRequestServiceImpl struct {
	repo                *dataaccess.Queries
	blockchain          BlockchainService
	notificationService NotificationService
}

func NewReturnRequestServiceImpl(blockchain BlockchainService, notification NotificationService) ReturnRequestService {
	return &ReturnRequestServiceImpl{
		repo:                dataaccess.New(global.Db),
		blockchain:          blockchain,
		notificationService: notification,
	}
}

// GetByLandlordID implements ReturnRequestService.
func (r *ReturnRequestServiceImpl) GetByLandlordID(userID int) *responses.ResponseData {
	requests, err := r.repo.GetReturnRequestByLandlordID(context.Background(), int32(userID))

	if len(requests) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusNoContent,
			Message:    responses.StatusNoData,
			Data:       nil,
		}
	}
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	var detailedRequests []responses.GetReturnRequestByLandlordIDRes
	for _, request := range requests {
		// Lấy thông tin chi tiết của người gửi dựa trên sender_id
		sender, err := r.repo.GetUserByID(context.Background(), *request.CreatedUser)
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
		detailedRequest := responses.GetReturnRequestByLandlordIDRes{
			ID:                 request.ID,
			ContractID:         request.ContractID,
			RoomID:             *request.RoomID,
			Reason:             request.Reason,
			ReturnDate:         request.ReturnDate,
			Status:             request.Status,
			DeductAmount:       request.DeductAmount,
			TotalReturnDeposit: request.TotalReturnDeposit,
			CreatedUser:        user,
			CreatedAt:          request.CreatedAt,
			UpdatedAt:          request.UpdatedAt,
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

// Create implements ReturnRequestService.
func (r *ReturnRequestServiceImpl) Create(req dataaccess.CreateReturnRequestParams, userID int) *responses.ResponseData {
	id := int32(userID)
	req.CreatedUser = &id

	contract, err := r.blockchain.GetMContractByIDOnChain(int64(*req.ContractID))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	if contract.PostRentalStatus != 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Trạng thái hợp đồng không hợp lệ",
			Data:       false,
		}
	}

	user, _ := r.repo.GetUserByID(context.Background(), int32(userID))
	_, err = r.blockchain.CreateReturnRequestOnChain(*user.PrivateKeyHex, int64(*req.ContractID))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	returnRequestId, err := r.repo.CreateReturnRequest(context.Background(), req)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	reqId := int(returnRequestId)
	r.notificationService.SendNotification(int(contract.Landlord), "Bạn có một yêu cầu trả phòng mới", &reqId, "return_request")

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}

// GetByID implements ReturnRequestService.
func (r *ReturnRequestServiceImpl) GetByID(id int) *responses.ResponseData {
	returnRequest, err := r.repo.GetReturnRequestByID(context.Background(), int32(id))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	room, err := r.repo.GetRoomByID(context.Background(), *returnRequest.RoomID)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	sender, err := r.repo.GetUserByID(context.Background(), *returnRequest.CreatedUser)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       "nothing",
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data: responses.GetReturnRequestByIDRes{
			ID:                 returnRequest.ID,
			Reason:             returnRequest.Reason,
			Room:               room,
			ContractID:         returnRequest.ContractID,
			CreatedUser:        sender,
			ReturnDate:         returnRequest.ReturnDate,
			Status:             returnRequest.Status,
			DeductAmount:       returnRequest.DeductAmount,
			TotalReturnDeposit: returnRequest.TotalReturnDeposit,
			CreatedAt:          returnRequest.CreatedAt,
			UpdatedAt:          returnRequest.UpdatedAt,
		},
	}
}

// Aprrove implements ReturnRequestService.
func (r *ReturnRequestServiceImpl) Aprrove(id int, userID int) *responses.ResponseData {
	// Get return request
	returnRequest, err := r.repo.GetReturnRequestByID(context.Background(), int32(id))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusNotFound,
			Message:    "Yêu câu trả phòng không tồn tại",
			Data:       false,
		}
	}

	if returnRequest.TotalReturnDeposit != nil && *returnRequest.TotalReturnDeposit == float64(0) {
		contract, err := r.blockchain.GetMContractByIDOnChain(int64(*returnRequest.ContractID))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
		if contract.PostRentalStatus != 1 {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    "Trạng thái hợp đồng không hợp lệ",
				Data:       false,
			}
		}
		user, _ := r.repo.GetUserByID(context.Background(), int32(userID))
		_, err = r.blockchain.ApproveReturnRequestOnChain(*user.PrivateKeyHex, int64(*returnRequest.ContractID))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	}

	// update status = 2 in return request table
	updateErr := r.repo.ApproveReturnRequest(context.Background(), returnRequest.ID)
	if updateErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    updateErr.Error(),
			Data:       false,
		}
	}

	// get room by contract
	room, roomErr := r.repo.GetRoomByContractID(context.Background(), *returnRequest.ContractID)
	if roomErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    roomErr.Error(),
			Data:       false,
		}
	}

	if room.Owner != int32(userID) {
		return &responses.ResponseData{
			StatusCode: http.StatusForbidden,
			Message:    "Bạn không có quyền thực hiện thao tác này",
			Data:       false,
		}
	}

	// set room available
	updateRoomParam := dataaccess.UpdateRoomParams{
		ID:     room.RoomID,
		IsRent: false,
	}
	_, updateRoomErr := r.repo.UpdateRoom(context.Background(), updateRoomParam)
	if updateRoomErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    updateRoomErr.Error(),
			Data:       false,
		}
	}

	user, _ := r.repo.GetUserByID(context.Background(), int32(userID))
	_, err = r.blockchain.DeclineMContractOnChain(*user.PrivateKeyHex, int64(id))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	// update tenants table
	updateTenantErr := r.repo.DeleteTenantByRoomID(context.Background(), room.RoomID)
	if updateTenantErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    updateTenantErr.Error(),
			Data:       false,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}
