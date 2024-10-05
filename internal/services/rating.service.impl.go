package services

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"smart-rental/global"
	"smart-rental/internal/constants"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"time"
)

type RatingServiceImpl struct {
	repo    *dataaccess.Queries
	storage StorageSerivce
}



func NewRatingServiceImpl(storage StorageSerivce) RatingService {
	return &RatingServiceImpl{
		repo:    dataaccess.New(global.Db),
		storage: storage,
	}
}

// CreateLandlordRating implements RatingService.
func (r *RatingServiceImpl) CreateLandlordRating(req requests.CreateLandlordRatingReq, userID int) *responses.ResponseData {
	var param dataaccess.CreateLandlordRatingParams
	common.MapStruct(req, &param)
	id := int32(userID)
	param.RatedBy = &id

	err := r.repo.CreateLandlordRating(context.Background(), param)
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

// CreateRoomRating implements RatingService.
func (r *RatingServiceImpl) CreateRoomRating(req requests.CreateRoomRatingReq, userID int) *responses.ResponseData {
	var urls []string
	for _, fileName := range req.Images {
		f, _ := fileName.Open()
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		fileExt := filepath.Ext(fileName.Filename)
		contentType := mime.TypeByExtension(fileExt)
		user := fmt.Sprintf("user_%d", userID)
		objKey := fmt.Sprintf("%s/%s/%d%s", constants.RATING_OBJ, user, timestamp, fileExt)

		url, err := r.storage.UploadFile(constants.BUCKET_NAME, objKey, f, contentType)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}
		urls = append(urls, url)

	}
	var param dataaccess.CreateRoomRatingParams
	common.MapStruct(req, &param)
	id := int32(userID)
	param.RatedBy = &id
	param.Images = urls

	err := r.repo.CreateRoomRating(context.Background(), param)
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

// CreateTenantRating implements RatingService.
func (r *RatingServiceImpl) CreateTenantRating(req requests.CreateTenantRatingReq, userID int) *responses.ResponseData {
	var urls []string
	for _, fileName := range req.Images {
		f, _ := fileName.Open()
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		fileExt := filepath.Ext(fileName.Filename)
		contentType := mime.TypeByExtension(fileExt)
		user := fmt.Sprintf("user_%d", userID)
		objKey := fmt.Sprintf("%s/%s/%d%s", constants.RATING_OBJ, user, timestamp, fileExt)

		url, err := r.storage.UploadFile(constants.BUCKET_NAME, objKey, f, contentType)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}
		urls = append(urls, url)

	}
	var param dataaccess.CreateTenantRatingParams
	common.MapStruct(req, &param)
	id := int32(userID)
	param.RatedBy = &id
	param.Images = urls

	err := r.repo.CreateTenantRating(context.Background(), param)
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

// GetRoomRatingByRoomID implements RatingService.
func (r *RatingServiceImpl) GetRoomRatingByRoomID(roomID int32) *responses.ResponseData {
	roomRatings, err := r.repo.GetRoomRatingByRoomID(context.Background(), &roomID)
	if err != nil {
		if len(roomRatings) == 0 {
			return &responses.ResponseData{
				StatusCode: http.StatusNoContent,
				Message:    responses.StatusNoData,
				Data:       nil,
			}
		}
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       roomRatings,
	}
}