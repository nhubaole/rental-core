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

type BillingServiceController struct {
	service services.BillingService
}

func NewBillingServiceController(service services.BillingService) *BillingServiceController {
	return &BillingServiceController{
		service: service,
	}
}

func (controller BillingServiceController) CreateBill(ctx *gin.Context) {
	myuser, _ := common.GetCurrentUser(ctx)
	var body *dataaccess.CreateBillParams
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	result := controller.service.CreateBill(myuser.ID, body)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}


func (controller BillingServiceController) GetBillByMonth(ctx *gin.Context) {
	myuser, _ := common.GetCurrentUser(ctx)
	year, err := strconv.Atoi(ctx.Param("year"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	month, err := strconv.Atoi(ctx.Param("month"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	result := controller.service.GetBillByMonth(myuser.ID, int32(month), int32(year))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}