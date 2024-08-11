package routers

import (
	"smart-rental/internal/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(uc *controllers.AuthenController) *gin.Engine {
	r := gin.Default()


	baseRouter := r.Group("/api/v1")
	authRouter := baseRouter.Group("/authen")
	authRouter.POST("", uc.Register)
	authRouter.POST("hello", uc.Hello)

	return r
}
