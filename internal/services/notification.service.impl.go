package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type NotificationServiceImpl struct {
	repo *dataaccess.Queries
}

func NewNotificationServiceImpl() NotificationService {
	return &NotificationServiceImpl{
		repo: dataaccess.New(global.Db),
	}
}

func (n *NotificationServiceImpl) GetByUserID(userID int) *responses.ResponseData {
	notifications, err := n.repo.GetNotificationsByUserID(context.Background(), int32(userID))

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	if len(notifications) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusNoContent,
			Message:    responses.StatusNoData,
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       notifications,
	}
}
