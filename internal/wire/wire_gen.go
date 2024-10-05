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
	roomService := services.NewRoomServiceImpl(storageSerivce)
	roomController := controllers.NewRoomController(roomService)
	return roomController
}

func InitRentalRequestRouterHandler() *controllers.RentalRequestController {
	rentalRequestService := services.NewRentalRequestServiceImpl()
	rentalRequestController := controllers.NewRentalRequestController(rentalRequestService)
	return rentalRequestController
}

func InitContractRouterHandler() *controllers.ContractController {
	contractService := services.NewContractServiceImpl()
	contractController := controllers.NewContractController(contractService)
	return contractController
}

func InitProcessTrackingRouterHandler() *controllers.ProcessTrackingController {
	processService := services.NewProcessServiceImpl()
	processTrackingController := controllers.NewProcessTrackingController(processService)
	return processTrackingController
}

func InitIndexRouterHandler() *controllers.IndexController {
	indexService := services.NewIndexServiceImpl()
	indexController := controllers.NewIndexController(indexService)
	return indexController
}

func InitBillingRouterHandler() *controllers.BillingController {
	billingService := services.NewBillingServiceImpl()
	billingController := controllers.NewBillingController(billingService)
	return billingController
}

func InitReturnRequestRouterHandler() *controllers.ReturnRequestController {
	returnRequestService := services.NewReturnRequestServiceImpl()
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
