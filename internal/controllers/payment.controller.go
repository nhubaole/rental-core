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

func (controller PaymentController) GetAllBanks(ctx *gin.Context) {
	result := controller.services.GetAllBanks()
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (controller PaymentController) GetAll(ctx *gin.Context) {
	result := controller.services.GetAll()
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (c PaymentController) Confirm(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := c.services.Confirm(id)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (c PaymentController) GetPaymentInfo(ctx *gin.Context) {
	typeOfPayment := ctx.Query("type")

	if typeOfPayment != "contract" && typeOfPayment != "return" && typeOfPayment != "bill" {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
			return
	}
	id, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
			return
		}
	result := c.services.GetDetailInfo(typeOfPayment, id)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}