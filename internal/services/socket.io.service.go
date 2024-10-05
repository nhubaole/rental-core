package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type SocketIOService interface {
	SendMessage(req requests.MessageReq) *responses.ResponseData
}