package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type NotificationService interface {
	GetByUserID(userID int) *responses.ResponseData
	Create(req dataaccess.CreateNotificationParams) (string, error)
	SendNotification(userId int, message string, referenceId *int, referenceType string) error
}