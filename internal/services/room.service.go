package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type RoomService interface {
	CreateRoom(req requests.CreateRoomForm) *responses.ResponseData
	GetRooms() *responses.ResponseData
	GetRoomByID(id int) *responses.ResponseData
}
