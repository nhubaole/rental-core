package controllers

import (
	"smart-rental/internal/services"
	"smart-rental/pkg/common"
	"smart-rental/pkg/responses"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	services services.NotificationService
}

func NewNotificationController(services services.NotificationService) *NotificationController {
	return &NotificationController{
		services: services,
	}
}

func (nc NotificationController) GetAll(ctx *gin.Context) {
	user, err := common.GetCurrentUser(ctx)
	if err != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return

	}

	result := nc.services.GetByUserID(int(user.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}
