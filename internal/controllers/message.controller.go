package controllers

import (
	"smart-rental/internal/services"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	service services.SocketIOService
}

func NewMessageController(service services.SocketIOService) *MessageController{
	return &MessageController{
		service: service,
	}
}

func(ms *MessageController)SendMessage(ctx *gin.Context) {
	var body requests.MessageReq
	if err := ctx.ShouldBindJSON(&body); err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := ms.service.SendMessage(body)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}