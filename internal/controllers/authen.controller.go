package controllers

import (
	"net/http"
	"smart-rental/internal/models"
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
	newUser := models.User{}
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := uc.authenService.Register(&newUser)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)

}
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
func (uc AuthenController) Hello(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}