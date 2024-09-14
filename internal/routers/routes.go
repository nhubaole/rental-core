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
	authRouter.POST("login", ac.Login)
	authRouter.POST("verify-otp", ac.VerifyOTP)

	userRouter := baseRouter.Group("/users")
	userRouter.GET("", uc.GetAll)
	userRouter.GET("/:id", uc.GetUserByID)
	userRouter.PUT("/", uc.Update)

	roomRouter := baseRouter.Group("/rooms")
	roomRouter.POST("", rc.Create)
	roomRouter.GET("", rc.GetAll)
	roomRouter.GET("/:id", rc.GetByID)
	roomRouter.GET("/search-by-address", rc.SearchByAddress)
	roomRouter.GET("/like/:id", rc.Like)

	return r
}
