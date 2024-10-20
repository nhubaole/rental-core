package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type SocketIOService interface {
	SendMessage(req requests.MessageReq) *responses.ResponseData
	GetMessages() *responses.ResponseData
	GetMessageByID(id int) *responses.ResponseData
	GetMessageByConversationID(conversationID int)*responses.ResponseData
}