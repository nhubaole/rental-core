package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type ConversationService interface {
	Create(req requests.CreateConversationReq, userID int) *responses.ResponseData
	GetByUserID(userID int) *responses.ResponseData
}