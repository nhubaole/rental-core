package controllers

import (
	"net/http"
	"smart-rental/internal/services"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"strconv"

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
		responses.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	result := ms.service.SendMessage(body)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func(ms *MessageController) GetMessagesByConversationID(ctx *gin.Context) {
	conversationID, err := strconv.Atoi(ctx.Param("conversationID"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Bad request", nil)
		return
	}

	result := ms.service.GetMessageByConversationID(int(conversationID))
	responses.APIResponse(ctx, result.StatusCode,result.Message, result.Data)
}