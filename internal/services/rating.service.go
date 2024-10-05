package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type RatingService interface {
	CreateRoomRating(req requests.CreateRoomRatingReq, userID int) *responses.ResponseData
	CreateTenantRating(req requests.CreateTenantRatingReq, userID int) *responses.ResponseData
	CreateLandlordRating(req requests.CreateLandlordRatingReq, userID int) *responses.ResponseData
	GetRoomRatingByRoomID(roomID int32) *responses.ResponseData
}
