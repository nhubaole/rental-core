package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type ReturnRequestServiceImpl struct {
	repo *dataaccess.Queries
}



func NewReturnRequestServiceImpl() ReturnRequestService {
	return &ReturnRequestServiceImpl{
		repo: dataaccess.New(global.Db),
	}
}

// Create implements ReturnRequestService.
func (r *ReturnRequestServiceImpl) Create(req dataaccess.CreateReturnRequestParams, userID int) *responses.ResponseData {
	id := int32(userID)
	req.CreatedUser = &id

	err := r.repo.CreateReturnRequest(context.Background(), req)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
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

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       returnRequest,
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
		ID: room.RoomID,
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

	// update status = 4 - expire in contract
	updateContractErr := r.repo.SetExpiredContract(context.Background(), *returnRequest.ContractID)
	if updateContractErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    updateContractErr.Error(),
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