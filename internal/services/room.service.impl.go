package services

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/common/constants"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"time"
)

type RoomServiceImpl struct {
	repo           *dataaccess.Queries
	storageService StorageSerivce
}

func NewRoomServiceImpl(storage StorageSerivce) RoomService {
	return &RoomServiceImpl{
		repo:           dataaccess.New(global.Db),
		storageService: storage,
	}
}

// CreateRoom implements RoomService.
func (r *RoomServiceImpl) CreateRoom(req requests.CreateRoomForm) *responses.ResponseData {
	if exist, _ := r.storageService.IsBucketExists(constants.BUCKET_NAME); !exist {
		err := r.storageService.CreateBucket(constants.BUCKET_NAME)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	}

	var urls []string
	for _, fileName := range req.RoomImages {
		f, _ := fileName.Open()
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		fileExt := filepath.Ext(fileName.Filename)
		contentType := mime.TypeByExtension(fileExt)
		objKey := fmt.Sprintf("room-images/%s/%d%s", "test", timestamp, fileExt)

		url, err := r.storageService.UploadFile(constants.BUCKET_NAME, objKey, f, contentType)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}
		urls = append(urls, url)

	}

	params := dataaccess.CreateRoomParams{
		// Title:           req.Title,
		// Address:         req.Address,
		// RoomNumber:      req.RoomNumber,
		// RoomImages:      urls,
		// Utilities:       req.Utilities,
		// Description:     req.Description,
		// RoomType:        req.RoomType,
		// Owner:           req.Owner,
		// Capacity:        req.Capacity,
		// Gender:          req.Gender,
		// Area:            req.Area,
		// TotalPrice:      req.TotalPrice,
		// Deposit:         req.Deposit,
		// ElectricityCost: req.ElectricityCost,
		// WaterCost:       req.WaterCost,
		// InternetCost:    req.InternetCost,
		// IsParking:       req.IsParking,
		// ParkingFee:      req.ParkingFee,
		// Status:          req.Status,
		// IsRent:          req.IsRent,
	}

	common.MapStruct(req, &params)
	params.RoomImages = urls
	err := r.repo.CreateRoom(context.Background(), params)
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

// GetRooms implements RoomService.
func (r *RoomServiceImpl) GetRooms() *responses.ResponseData {
	rooms, err := r.repo.GetRooms(context.Background())

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       rooms,
	}
}

// GetRoomByID implements RoomService.
func (r *RoomServiceImpl) GetRoomByID(id int) *responses.ResponseData {
	room, err := r.repo.GetRoomByID(context.Background(), int32(id))

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       room,
	}
}