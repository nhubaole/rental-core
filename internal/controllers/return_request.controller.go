package controllers

import (
	"net/http"
	"smart-rental/internal/dataaccess"
	"smart-rental/internal/services"
	"smart-rental/pkg/common"
	"smart-rental/pkg/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReturnRequestController struct {
	service services.ReturnRequestService
}

func NewReturnRequestController(service services.ReturnRequestService) *ReturnRequestController {
	return &ReturnRequestController{
		service: service,
	}
}

func(rrc *ReturnRequestController) Create(ctx *gin.Context) {
	var body dataaccess.CreateReturnRequestParams
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	currentUser, parseUserErr := common.GetCurrentUser(ctx)
	if parseUserErr != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	result := rrc.service.Create(body, int(currentUser.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (rrc *ReturnRequestController) GetReturnRequestByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	
	result := rrc.service.GetByID(id)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (rrc *ReturnRequestController) GetReturnRequestByLandlordID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	
	result := rrc.service.GetByLandlordID(id)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (rrc *ReturnRequestController) ApproveReturnRequest(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	currentUser, parseUserErr := common.GetCurrentUser(ctx)
	if parseUserErr != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	result := rrc.service.Aprrove(id, int(currentUser.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}