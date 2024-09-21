package routers

import (
	"smart-rental/internal/controllers"
	"smart-rental/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	ac *controllers.AuthenController,
	uc *controllers.UserController,
	rc *controllers.RoomController,
	rrc *controllers.RentalRequestController,
	pc *controllers.ProcessTrackingController,
	id *controllers.IndexServiceController,
	bi *controllers.BillingServiceController,

) *gin.Engine {
	r := gin.Default()

	baseRouter := r.Group("/api/v1")
	authRouter := baseRouter.Group("/authen")
	authRouter.POST("register", ac.Register)
	authRouter.POST("login", ac.Login)
	authRouter.POST("verify-otp", ac.VerifyOTP)

	userRouter := baseRouter.Group("/users")
	userRouter.GET("", middlewares.AuthenMiddleware, uc.GetAll)
	userRouter.GET("/:id", middlewares.AuthenMiddleware, uc.GetUserByID)
	userRouter.PUT("/", middlewares.AuthenMiddleware, uc.Update)

	roomRouter := baseRouter.Group("/rooms")
	roomRouter.POST("", middlewares.AuthenMiddleware, rc.Create)
	roomRouter.GET("", middlewares.AuthenMiddleware, rc.GetAll)
	roomRouter.GET("/:id", middlewares.AuthenMiddleware, rc.GetByID)
	roomRouter.GET("/search-by-address", middlewares.AuthenMiddleware, rc.SearchByAddress)
	roomRouter.GET("/like/:id", middlewares.AuthenMiddleware, rc.Like)
	roomRouter.GET("/like", middlewares.AuthenMiddleware, rc.GetLikedRooms)
	roomRouter.GET("/status/:status", middlewares.AuthenMiddleware, rc.GetByStatus)

	rentalRequestRouter := baseRouter.Group("/requests")
	rentalRequestRouter.POST("", middlewares.AuthenMiddleware, rrc.Create)
	rentalRequestRouter.DELETE("/:id", middlewares.AuthenMiddleware, rrc.Delete)
	rentalRequestRouter.GET("", middlewares.AuthenMiddleware, rrc.GetAllRentalRequest)
	rentalRequestRouter.GET("/:id", middlewares.AuthenMiddleware, rrc.GetRentalRequestById)
	rentalRequestRouter.GET("/:id/review", middlewares.AuthenMiddleware, rrc.UpdateRentalRequestStatus)
	rentalRequestRouter.GET("/:id/tracking-process", middlewares.AuthenMiddleware, pc.GetProcessTrackingByRentalId)
	rentalRequestRouter.GET("/all/tracking-process", middlewares.AuthenMiddleware, pc.GetAllProcessTracking)

	billingRouter := baseRouter.Group("/billings")
	billingRouter.GET("/index/:year/:month", middlewares.AuthenMiddleware, id.GetIndexFromOwner)
	billingRouter.POST("/index", middlewares.AuthenMiddleware, id.CreateIndex)
	billingRouter.POST("/", middlewares.AuthenMiddleware, bi.CreateBill)
	billingRouter.GET("/:year/:month", middlewares.AuthenMiddleware, bi.GetBillByMonth)

	return r
}
