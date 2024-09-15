package routers

import (
	"smart-rental/internal/controllers"
	"smart-rental/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter(ac *controllers.AuthenController,
	uc *controllers.UserController,
	rc *controllers.RoomController,
	cc *controllers.ContractController ) *gin.Engine {
	r := gin.Default()

	baseRouter := r.Group("/api/v1")
	authRouter := baseRouter.Group("/authen")
	authRouter.POST("register", ac.Register)
	authRouter.POST("login", ac.Login)
	authRouter.POST("verify-otp", ac.VerifyOTP)

	userRouter := baseRouter.Group("/users")
	userRouter.GET("",middlewares.AuthenMiddleware, uc.GetAll)
	userRouter.GET("/:id",middlewares.AuthenMiddleware, uc.GetUserByID)
	userRouter.PUT("/", middlewares.AuthenMiddleware,uc.Update)

	roomRouter := baseRouter.Group("/rooms")
	roomRouter.POST("",middlewares.AuthenMiddleware, rc.Create)
	roomRouter.GET("",middlewares.AuthenMiddleware, rc.GetAll)
	roomRouter.GET("/:id",middlewares.AuthenMiddleware, rc.GetByID)
	roomRouter.GET("/search-by-address",middlewares.AuthenMiddleware, rc.SearchByAddress)
	roomRouter.GET("/like/:id",middlewares.AuthenMiddleware, rc.Like)
	roomRouter.GET("/like",middlewares.AuthenMiddleware, rc.GetLikedRooms)
	roomRouter.GET("/status/:status",middlewares.AuthenMiddleware, rc.GetByStatus)

	contractRouter := baseRouter.Group("/contracts")
	contractRouter.POST("/template", cc.CreateTemplate)
	contractRouter.POST("/template/get-by-address", cc.GetTemplateByAddress)

	return r
}
