package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type RoomServiceImpl struct {
	repo *dataaccess.Queries
	storageService StorageSerivce
}


func NewRoomServiceImpl(storage StorageSerivce) RoomService {
	return &RoomServiceImpl{
		repo: dataaccess.New(global.Db),
		storageService: storage,
	}
}


// CreateRoom implements RoomService.
func (r *RoomServiceImpl) CreateRoom(req dataaccess.CreateRoomParams) *responses.ResponseData {
	err := r.repo.CreateRoom(context.Background(),req)
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