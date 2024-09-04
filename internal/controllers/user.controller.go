package controllers

import (
	"smart-rental/internal/services"
	"smart-rental/pkg/responses"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us services.UserService
}

func NewUserController(us services.UserService) *UserController {
	return &UserController{
		us: us,
	}
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	result := uc.us.GetAll()
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}
