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

type UserController struct {
	userService services.UserService
}

func NewUserController(us services.UserService) *UserController {
	return &UserController{
		userService: us,
	}
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	result := uc.userService.GetAll()
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (uc *UserController) GetUserByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := uc.userService.GetUserByID(id)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (uc *UserController) Update(ctx *gin.Context) {
	// Parse request body
	var updateUserParam *dataaccess.UpdateUserParams
	err := ctx.ShouldBindJSON(&updateUserParam)
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	// Call service layer Update function
	result := uc.userService.Update(updateUserParam)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (uc *UserController) GetCurrentUser(ctx *gin.Context) {
	user, err := common.GetCurrentUser(ctx)
	if err != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return

	}

	result := uc.userService.GetUserByID(int(user.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (uc *UserController) UpdateUserBank(ctx *gin.Context) {
	user, err := common.GetCurrentUser(ctx)
	if err != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return

	}
	var updateUserBankBody dataaccess.UpdateUserBankParams
	updateUserBankBody.UserID = user.ID
	errParse := ctx.ShouldBindJSON(&updateUserBankBody)
	if errParse != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	result := uc.userService.UpdateBankInfo(&updateUserBankBody)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (uc *UserController) CreateUserBank(ctx *gin.Context) {
	user, err := common.GetCurrentUser(ctx)
	if err != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return

	}
	var createUserBankBody dataaccess.CreateUserBankParams
	createUserBankBody.UserID = user.ID
	errParse := ctx.ShouldBindJSON(&createUserBankBody)
	if errParse != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}

	result := uc.userService.CreateBankInfo(&createUserBankBody)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (uc *UserController) UpdateDeviceToken(ctx *gin.Context) {
	user, err := common.GetCurrentUser(ctx)
	if err != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return

	}
	var updateUserDeviceToken requests.UpdateDeviceTokenReq
	errParse := ctx.ShouldBindJSON(&updateUserDeviceToken)
	if errParse != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request body", nil)
		return
	}
	
	result := uc.userService.UpdateDeviceToken(int(user.ID), updateUserDeviceToken.DeviceToken)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}