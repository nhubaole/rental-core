package services

import (
	"smart-rental/pkg/responses"
)

type NotificationService interface {
	GetByUserID(userID int) *responses.ResponseData
}