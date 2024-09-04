package controllers

import (
	"fmt"
	"smart-rental/internal/dataaccess"
	"smart-rental/internal/services"
	"smart-rental/pkg/responses"

	"github.com/gin-gonic/gin"
)

type AuthenController struct {
	authenService services.AuthenService
}

func NewAuthController(service services.AuthenService) *AuthenController {
	return &AuthenController{
		authenService: service,
	}
}

func (uc AuthenController) Register(ctx *gin.Context) {
	newUser := dataaccess.CreateUserParams{}
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		fmt.Println(err)
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := uc.authenService.Register(&newUser)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)

}
