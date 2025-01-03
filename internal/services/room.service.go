package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type RoomService interface {
	CreateRoom(req requests.CreateRoomForm, userID int) *responses.ResponseData
	GetRooms(userID int) *responses.ResponseData
	GetRoomByID(id int) *responses.ResponseData
	SearchRoomByAddress(address string) *responses.ResponseData
	LikeRoom(roomID int, userID int) *responses.ResponseData
	GetLikedRooms(userID int) *responses.ResponseData
	GetRoomByStatus(status int) *responses.ResponseData
	GetRoomByOwner(userID int) *responses.ResponseData
	UpdateRoom(req requests.UpdateRoomRequest) *responses.ResponseData
	CheckUserLikedRoom(roomId int, userId int) *responses.ResponseData
}
