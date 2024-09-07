package routers

import (
	"smart-rental/internal/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter(ac *controllers.AuthenController, uc *controllers.UserController, rc *controllers.RoomController) *gin.Engine {
	r := gin.Default()

	baseRouter := r.Group("/api/v1")
	authRouter := baseRouter.Group("/authen")
	authRouter.POST("register", ac.Register)

	userRouter := baseRouter.Group("/users")
	userRouter.GET("", uc.GetAll)

	roomRouter := baseRouter.Group("/rooms")
	roomRouter.POST("", rc.Create)

	return r
}
