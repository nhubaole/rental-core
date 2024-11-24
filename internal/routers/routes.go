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
	ptc *controllers.ProcessTrackingController,
	ic *controllers.IndexController,
	bc *controllers.BillingController,
	cc *controllers.ContractController,
	returnRequestController *controllers.ReturnRequestController,
	ratingController *controllers.RatingController,
	ms *controllers.MessageController,
	conversation *controllers.ConversationController,
	payment *controllers.PaymentController,

) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	baseRouter := r.Group("/api/v1")
	authRouter := baseRouter.Group("/authen")
	authRouter.POST("register", ac.Register)
	authRouter.POST("login", ac.Login)
	authRouter.POST("verify-otp", ac.VerifyOTP)

	userRouter := baseRouter.Group("/users")
	userRouter.GET("", middlewares.AuthenMiddleware, uc.GetAll)
	userRouter.GET("/:id", middlewares.AuthenMiddleware, uc.GetUserByID)
	userRouter.GET("/current-user", middlewares.AuthenMiddleware, uc.GetCurrentUser)
	userRouter.PUT("/", middlewares.AuthenMiddleware, uc.Update)
	userRouter.POST("/bank-info", middlewares.AuthenMiddleware, uc.CreateUserBank)
	userRouter.PUT("/bank-info", middlewares.AuthenMiddleware, uc.UpdateUserBank)

	roomRouter := baseRouter.Group("/rooms")
	roomRouter.POST("", middlewares.AuthenMiddleware, rc.Create)
	roomRouter.GET("", middlewares.AuthenMiddleware, rc.GetAll)
	roomRouter.GET("/:id", middlewares.AuthenMiddleware, rc.GetByID)
	roomRouter.GET("/search-by-address", middlewares.AuthenMiddleware, rc.SearchByAddress)
	roomRouter.GET("/like/:id", middlewares.AuthenMiddleware, rc.Like)
	roomRouter.GET("/like", middlewares.AuthenMiddleware, rc.GetLikedRooms)
	roomRouter.GET("/status/:status", middlewares.AuthenMiddleware, rc.GetByStatus)
	roomRouter.PUT("", middlewares.AuthenMiddleware, rc.UpdateRoom)

	rentalRequestRouter := baseRouter.Group("/requests")
	rentalRequestRouter.POST("", middlewares.AuthenMiddleware, rrc.Create)
	rentalRequestRouter.DELETE("/:id", middlewares.AuthenMiddleware, rrc.Delete)
	rentalRequestRouter.GET("", middlewares.AuthenMiddleware, rrc.GetAllRentalRequest)
	rentalRequestRouter.GET("/:id", middlewares.AuthenMiddleware, rrc.GetRentalRequestById)
	rentalRequestRouter.GET("/:id/review", middlewares.AuthenMiddleware, rrc.UpdateRentalRequestStatus)
	rentalRequestRouter.GET("/:id/tracking-process", middlewares.AuthenMiddleware, ptc.GetProcessTrackingByRentalId)
	rentalRequestRouter.GET("/all/tracking-process", middlewares.AuthenMiddleware, ptc.GetAllProcessTracking)

	billingRouter := baseRouter.Group("/billings")
	billingRouter.GET("/index/:year/:month", middlewares.AuthenMiddleware, ic.GetIndexFromOwner)
	billingRouter.POST("/index", middlewares.AuthenMiddleware, ic.CreateIndex)
	billingRouter.POST("/", middlewares.AuthenMiddleware, bc.CreateBill)
	billingRouter.GET("/", middlewares.AuthenMiddleware, bc.GetBillByMonth)
	billingRouter.GET("/:id", middlewares.AuthenMiddleware, bc.GetBillByID)
	billingRouter.POST("/get-metrics", bc.GetBillMetric)
	billingRouter.GET("/status/:statusID", bc.GetBillByStatusID)
	billingRouter.GET("/get-bill-of-rented-rooms", middlewares.AuthenMiddleware, bc.GetBillOfRentedRoomByOwnerID)

	contractRouter := baseRouter.Group("/contracts")
	contractRouter.POST("/template", cc.CreateTemplate)
	contractRouter.POST("/template/get-by-address", cc.GetTemplateByAddress)
	contractRouter.POST("", cc.Create)
	contractRouter.GET("/:id", cc.GetByID)
	contractRouter.GET("/status/:statusID", cc.GetByStatus)
	contractRouter.PUT("/sign", middlewares.AuthenMiddleware, cc.SignContract)
	contractRouter.PUT("/decline/:id", middlewares.AuthenMiddleware, cc.DeclineContract)

	returnRequestRouter := baseRouter.Group("/return-requests")
	returnRequestRouter.POST("", middlewares.AuthenMiddleware, returnRequestController.Create)
	returnRequestRouter.GET("/:id", middlewares.AuthenMiddleware, returnRequestController.GetReturnRequestByID)
	returnRequestRouter.GET("/confirm/:id", middlewares.AuthenMiddleware, returnRequestController.ApproveReturnRequest)

	ratingRouter := baseRouter.Group("/ratings")
	ratingRouter.POST("create-room-rating", middlewares.AuthenMiddleware, ratingController.CreateRoomRating)
	ratingRouter.POST("create-tenant-rating", middlewares.AuthenMiddleware, ratingController.CreateTenantRating)
	ratingRouter.POST("create-landlord-rating", middlewares.AuthenMiddleware, ratingController.CreateLandlordRating)
	ratingRouter.GET("/:roomID", middlewares.AuthenMiddleware, ratingController.GetRoomRatingByRoomID)

	messageRouter := baseRouter.Group("/messages")
	messageRouter.POST("", ms.SendMessage)
	messageRouter.GET("/conversation/:conversationID",middlewares.AuthenMiddleware, ms.GetMessagesByConversationID)

	conversationRouter := baseRouter.Group("/conversations")
	conversationRouter.POST("",middlewares.AuthenMiddleware, conversation.CreateConversation)
	conversationRouter.GET("/get-by-current-user", middlewares.AuthenMiddleware, conversation.GetConversationByUserID)

	paymentRouter := baseRouter.Group("/payments")
	paymentRouter.GET("/:id", middlewares.AuthenMiddleware, payment.GetByID)
	paymentRouter.POST("", middlewares.AuthenMiddleware, payment.Create)
	paymentRouter.GET("/banks", middlewares.AuthenMiddleware, payment.GetAllBanks)
	return r
}
