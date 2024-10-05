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

type RatingController struct {
	service services.RatingService
}

func NewRatingController(service services.RatingService) *RatingController {
	return &RatingController{
		service: service,
	}
}

func(ratingController *RatingController) CreateRoomRating(ctx *gin.Context) {
	var body requests.CreateRoomRatingReq
	err := ctx.ShouldBind(&body)
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	currentUser, parseUserErr := common.GetCurrentUser(ctx)
	if parseUserErr != nil {
		responses.APIResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	result := ratingController.service.CreateRoomRating(body, int(currentUser.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func(ratingController *RatingController) CreateTenantRating(ctx *gin.Context) {
	var body requests.CreateTenantRatingReq
	err := ctx.ShouldBind(&body)
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	currentUser, parseUserErr := common.GetCurrentUser(ctx)
	if parseUserErr != nil {
		responses.APIResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	result := ratingController.service.CreateTenantRating(body, int(currentUser.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func(ratingController *RatingController) CreateLandlordRating(ctx *gin.Context) {
	var body requests.CreateLandlordRatingReq
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	currentUser, parseUserErr := common.GetCurrentUser(ctx)
	if parseUserErr != nil {
		responses.APIResponse(ctx, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	result := ratingController.service.CreateLandlordRating(body, int(currentUser.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func(ratingController *RatingController) GetRoomRatingByRoomID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("roomID"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	result := ratingController.service.GetRoomRatingByRoomID(int32(id))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}