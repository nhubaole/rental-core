package controllers

import (
	"net/http"
	"smart-rental/internal/services"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ConversationController struct {
	services services.ConversationService
}

func NewConversationController(services services.ConversationService) *ConversationController {
	return &ConversationController{
		services: services,
	}
}

func(cc ConversationController) CreateConversation(ctx *gin.Context) {
	req := requests.CreateConversationReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	currentUser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return
	}
	result := cc.services.Create(req, int(currentUser.ID))
	responses.APIResponse(ctx, result.StatusCode,result.Message, result.Data)
}

func(cc ConversationController) GetConversationByCurrentUser(ctx *gin.Context) {
	currentUser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return
	}

	result := cc.services.GetByUserID(int(currentUser.ID))
	responses.APIResponse(ctx, result.StatusCode,result.Message, result.Data)
}

func(cc ConversationController) GetConversationByUserID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	result := cc.services.GetByUserID(int(userID))
	responses.APIResponse(ctx, result.StatusCode,result.Message, result.Data)
}