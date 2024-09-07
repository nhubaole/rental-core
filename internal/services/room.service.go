package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type RoomService interface {
	CreateRoom(req dataaccess.CreateRoomParams) *responses.ResponseData
}