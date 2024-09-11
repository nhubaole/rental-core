package controllers

import (
	"smart-rental/internal/services"
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