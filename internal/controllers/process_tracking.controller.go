package controllers

import (
	"net/http"
	"smart-rental/internal/services"
	"smart-rental/pkg/common"
	"smart-rental/pkg/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProcessTrackingController struct {
	processService services.ProcessService
}

func NewProcessTrackingController(service services.ProcessService) *ProcessTrackingController {
	return &ProcessTrackingController{
		processService: service,
	}
}

func (controller ProcessTrackingController) GetProcessTrackingByRentalId(ctx *gin.Context) {
	myuser, _ := common.GetCurrentUser(ctx)

	rentid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	result := controller.processService.GetProcessTrackingByRentalId(myuser.ID, int32(rentid))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (controller ProcessTrackingController) GetAllProcessTracking(ctx *gin.Context) {
	myuser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	result := controller.processService.GetAllProcessTracking(myuser.ID)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}
