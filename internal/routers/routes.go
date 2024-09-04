package routers

import (
	"smart-rental/internal/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(ac *controllers.AuthenController, uc *controllers.UserController) *gin.Engine {
	r := gin.Default()

	baseRouter := r.Group("/api/v1")
	authRouter := baseRouter.Group("/authen")
	authRouter.POST("", ac.Register)

	userRouter := baseRouter.Group("/users")
	userRouter.GET("", uc.GetAll)

	return r
}
