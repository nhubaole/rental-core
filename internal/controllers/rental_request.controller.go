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

type RentalRequestController struct {
	rentalService services.RentalRequestService
}

func NewRentalRequestController(service services.RentalRequestService) *RentalRequestController {
	return &RentalRequestController{
		rentalService: service,
	}
}

func (controller RentalRequestController) Create(ctx *gin.Context) {

	var body *requests.CreateRentalRequest
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	myuser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	result := controller.rentalService.CreateRentalRequest(body, myuser.ID)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (controller RentalRequestController) Delete(ctx *gin.Context) {
	rentid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	myuser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	result := controller.rentalService.DeleteRentalRequest(int32(rentid), myuser.ID)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}
func (controller RentalRequestController) GetRentalRequestById(ctx *gin.Context) {
	rentid, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	myuser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	result := controller.rentalService.GetRentalRequestById(int32(rentid), myuser.ID)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (controller RentalRequestController) GetAllRentalRequest(ctx *gin.Context) {
	myuser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	result := controller.rentalService.GetAllRentalRequest(myuser.PhoneNumber)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (controller RentalRequestController) UpdateRentalRequestStatus(ctx *gin.Context) {
	myuser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	action := ctx.Query("action")

	if action == "approve" || action == "decline" {
		rentid, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
			return
		}
		result := controller.rentalService.ReviewRentalRequest(action, int32(rentid), myuser.ID)
		responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
		return
	}
	responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request parameter", nil)
}

func (controller RentalRequestController) GetRentalRequestByRoomId(ctx *gin.Context) {
	roomID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	result := controller.rentalService.GetRentalRequestByRoomID(roomID)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}
