// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"smart-rental/internal/controllers"
	"smart-rental/internal/services"
)

// Injectors from wire.go:

func InitPaymentRouterHandler() *controllers.PaymentController {
	storageSerivce := services.NewStorageServiceImpl()
	blockchainService := services.NewBlockchainServiceImpl()
	notificationService := services.NewNotificationServiceImpl()
	paymentService := services.NewPaymentServiceImpl(storageSerivce, blockchainService, notificationService)
	paymentController := controllers.NewPaymentController(paymentService)
	return paymentController
}

func InitAuthenRouterHandler() *controllers.AuthenController {
	authenService := services.NewAuthenSerivceImpl()
	authenController := controllers.NewAuthController(authenService)
	return authenController
}

func InitUserRouterHandler() *controllers.UserController {
	userService := services.NewUserServiceImpl()
	userController := controllers.NewUserController(userService)
	return userController
}

func InitRoomRouterHandler() *controllers.RoomController {
	storageSerivce := services.NewStorageServiceImpl()
	blockchainService := services.NewBlockchainServiceImpl()
	roomService := services.NewRoomServiceImpl(storageSerivce, blockchainService)
	roomController := controllers.NewRoomController(roomService)
	return roomController
}

func InitRentalRequestRouterHandler() *controllers.RentalRequestController {
	notificationService := services.NewNotificationServiceImpl()
	rentalRequestService := services.NewRentalRequestServiceImpl(notificationService)
	rentalRequestController := controllers.NewRentalRequestController(rentalRequestService)
	return rentalRequestController
}

func InitContractRouterHandler() *controllers.ContractController {
	blockchainService := services.NewBlockchainServiceImpl()
	notificationService := services.NewNotificationServiceImpl()
	contractService := services.NewContractServiceImpl(blockchainService, notificationService)
	contractController := controllers.NewContractController(contractService)
	return contractController
}

func InitProcessTrackingRouterHandler() *controllers.ProcessTrackingController {
	processService := services.NewProcessServiceImpl()
	processTrackingController := controllers.NewProcessTrackingController(processService)
	return processTrackingController
}

func InitIndexRouterHandler() *controllers.IndexController {
	blockchainService := services.NewBlockchainServiceImpl()
	indexService := services.NewIndexServiceImpl(blockchainService)
	indexController := controllers.NewIndexController(indexService)
	return indexController
}

func InitBillingRouterHandler() *controllers.BillingController {
	blockchainService := services.NewBlockchainServiceImpl()
	notificationService := services.NewNotificationServiceImpl()
	billingService := services.NewBillingServiceImpl(blockchainService, notificationService)
	billingController := controllers.NewBillingController(billingService)
	return billingController
}

func InitReturnRequestRouterHandler() *controllers.ReturnRequestController {
	blockchainService := services.NewBlockchainServiceImpl()
	notificationService := services.NewNotificationServiceImpl()
	returnRequestService := services.NewReturnRequestServiceImpl(blockchainService, notificationService)
	returnRequestController := controllers.NewReturnRequestController(returnRequestService)
	return returnRequestController
}

func InitRatingRouterHandler() *controllers.RatingController {
	storageSerivce := services.NewStorageServiceImpl()
	ratingService := services.NewRatingServiceImpl(storageSerivce)
	ratingController := controllers.NewRatingController(ratingService)
	return ratingController
}

func InitMessageRouterHandler() *controllers.MessageController {
	socketIOService := services.NewSocketIOServiceImpl()
	messageController := controllers.NewMessageController(socketIOService)
	return messageController
}

func InitConversationRouterHandler() *controllers.ConversationController {
	conversationService := services.NewConversationServiceImpl()
	conversationController := controllers.NewConversationController(conversationService)
	return conversationController
}

func InitNotificationRouterHandler() *controllers.NotificationController {
	notificationService := services.NewNotificationServiceImpl()
	notificationController := controllers.NewNotificationController(notificationService)
	return notificationController
}
