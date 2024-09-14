package controllers

import (
	"fmt"
	"net/http"
	"smart-rental/internal/services"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"

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
	fmt.Print("28")
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body1", nil)
		return
	}
	fmt.Print("33")
	myuser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body2", nil)
		return
	}
	fmt.Print("39")
	result := controller.rentalService.CreateRentalRequest(body, myuser.ID)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)

}
