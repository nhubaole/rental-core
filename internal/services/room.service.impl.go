package services

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"smart-rental/global"
	"smart-rental/internal/constants"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"strings"
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
	var params dataaccess.CreateRoomParams

	common.MapStruct(req, &params)
	id, err := r.repo.CreateRoom(context.Background(), params)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	var urls []string
	for _, fileName := range req.RoomImages {
		f, _ := fileName.Open() 
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		fileExt := filepath.Ext(fileName.Filename)
		contentType := mime.TypeByExtension(fileExt)
		roomID := fmt.Sprintf("room_%d", id)
		objKey := fmt.Sprintf("%s/%s/%d%s",constants.ROOM_OBJ, roomID, timestamp, fileExt)

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

	var updateRoom dataaccess.UpdateRoomParams
	updateRoom.ID = id
	updateRoom.RoomImages = urls
	_, updateErr := r.repo.UpdateRoom(context.Background(), updateRoom)
	if updateErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    updateErr.Error(),
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

// SearchRoomByAddress implements RoomService.
func (r *RoomServiceImpl) SearchRoomByAddress(address string) *responses.ResponseData {
	rooms, err := r.repo.SearchRoomByAddress(context.Background(), &address)

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

// LikeRoom implements RoomService.
func (r *RoomServiceImpl) LikeRoom(roomID int, userID int) *responses.ResponseData {
	param := dataaccess.CheckUserLikedRoomParams{
		RoomID: int32(roomID),
		UserID: int32(userID),
	}
	_, err := r.repo.CheckUserLikedRoom(context.Background(), param)
	var isLiked bool

	if err != nil {
		isLiked = true
		param := dataaccess.LikeRoomParams{
			RoomID: int32(roomID),
			UserID: int32(userID),
		}
		err := r.repo.LikeRoom(context.Background(), param)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	} else {
		isLiked = false
		param := dataaccess.UnlikeRoomParams{
			RoomID: int32(roomID),
			UserID: int32(userID),
		}
		err := r.repo.UnlikeRoom(context.Background(), param)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	}
	res := map[string]bool{"isLiked": isLiked}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       res,
	}
}

// GetLikedRooms implements RoomService.
func (r *RoomServiceImpl) GetLikedRooms(userID int) *responses.ResponseData {
	rooms, err := r.repo.GetLikedRooms(context.Background(), int32(userID))

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

// GetRoomByStatus implements RoomService.
func (r *RoomServiceImpl) GetRoomByStatus(status int) *responses.ResponseData {
	rooms, err := r.repo.GetRoomsByStatus(context.Background(), int32(status))

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

// UpdateRoom implements RoomService.
func (r *RoomServiceImpl) UpdateRoom(req requests.UpdateRoomRequest) *responses.ResponseData {
	// delete file from s3
	for _, urlString :=range req.DeleteFiles{
		parsedURL, err := url.Parse(urlString)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}
		objKey := parsedURL.Path

		objKey = strings.TrimPrefix(objKey, "/")
		deleteErr := r.storageService.DeleteObject(constants.BUCKET_NAME, objKey)

		if deleteErr != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    "failed to delete object from s3",
				Data:       nil,
			}
		}
	}

	// update file to s3
	var urls []string
	for _, fileName := range req.RoomImages {
		f, _ := fileName.Open() 
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		fileExt := filepath.Ext(fileName.Filename)
		contentType := mime.TypeByExtension(fileExt)
		roomID := fmt.Sprintf("room_%d", req.ID)
		objKey := fmt.Sprintf("%s/%s/%d%s",constants.ROOM_OBJ, roomID, timestamp, fileExt)

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

	// update data in db
	var param dataaccess.UpdateRoomParams
	common.MapStruct(req,&param)
	param.RoomImages = urls

	_, updateErr := r.repo.UpdateRoom(context.Background(), param)
	if updateErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    updateErr.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}
