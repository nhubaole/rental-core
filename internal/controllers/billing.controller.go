package controllers

import (
	"net/http"
	"smart-rental/internal/dataaccess"
	"smart-rental/internal/services"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BillingController struct {
	service services.BillingService
}

func NewBillingController(service services.BillingService) *BillingController {
	return &BillingController{
		service: service,
	}
}

func (controller BillingController) CreateBill(ctx *gin.Context) {
	currentUser, _ := common.GetCurrentUser(ctx)
	var body requests.CreateBill
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	result := controller.service.CreateBill(currentUser.ID, body)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (controller BillingController) GetBillByMonth(ctx *gin.Context) {
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

func (controller BillingController) GetBillByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	result := controller.service.GetBillByID(int32(id))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (controller BillingController) GetBillMetric(ctx *gin.Context) {
	var body dataaccess.GetAllMetric4BillByRoomIDParams
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	result := controller.service.GetBillMetrics(body)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}
func (controller BillingController) GetBillByStatusID(ctx *gin.Context) {
	user, err := common.GetCurrentUser(ctx)
	if err != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return

	}

	id, err := strconv.Atoi(ctx.Param("statusID"))
	if (id < 0 && id > 2) {
		responses.APIResponse(ctx, http.StatusBadRequest, "Status must be between 0 and 2", nil)
		return
	}
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	result := controller.service.GetBillByStatus(user.ID, int32(id))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (controller BillingController) GetBillOfRentedRoomByOwnerID(ctx *gin.Context) {
	currentUser, _ := common.GetCurrentUser(ctx)
	
	result := controller.service.GetBillOfRentedRoomByOwnerID(int(currentUser.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}