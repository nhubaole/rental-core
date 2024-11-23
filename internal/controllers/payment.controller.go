package controllers

import (
	"smart-rental/internal/services"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	services services.PaymentService
}

func NewPaymentController(services services.PaymentService) *PaymentController {
	return &PaymentController{
		services: services,
	}
}

func (c PaymentController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := c.services.GetByID(id)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (controller PaymentController) Create(ctx *gin.Context) {
	user, err := common.GetCurrentUser(ctx)
	if err != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return

	}
	var formData requests.CreatePaymentReq
	if err := ctx.ShouldBind(&formData); err != nil {
		responses.APIResponse(ctx, 400, err.Error(), nil)
		return
	}
	result := controller.services.Create(formData, user.ID)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}
