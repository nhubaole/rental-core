//go:build wireinject

package wire

import (
	//"smart-rental/initialize"
	"smart-rental/internal/controllers"
	//"smart-rental/internal/dataaccess"
	"smart-rental/internal/services"

	"github.com/google/wire"
)

func InitAuthenRouterHandler() *controllers.AuthenController {
	wire.Build(
		services.NewAuthenSerivceImpl,
		controllers.NewAuthController,
	)

	return &controllers.AuthenController{}
}

func InitUserRouterHandler() *controllers.UserController {
	wire.Build(
		services.NewUserServiceImpl,
		controllers.NewUserController,
	)

	return &controllers.UserController{}
}

func InitRoomRouterHandler() *controllers.RoomController {
	wire.Build(
		services.NewRoomServiceImpl,
		services.NewStorageServiceImpl,
		controllers.NewRoomController,
	)
	return &controllers.RoomController{}
}
func InitRentalRequestRouterHandler() *controllers.RentalRequestController {
	wire.Build(
		services.NewRentalRequestServiceImpl,
		controllers.NewRentalRequestController,
	)
	return &controllers.RentalRequestController{}
}

func InitContractRouterHandler() *controllers.ContractController {
	wire.Build(
		services.NewContractServiceImpl,
		controllers.NewContractController,
	)
	return &controllers.ContractController{}
}

func InitProcessTrackingRouterHandler() *controllers.ProcessTrackingController {
	wire.Build(
		services.NewProcessServiceImpl,
		controllers.NewProcessTrackingController,
	)
	return &controllers.ProcessTrackingController{}
}

func InitIndexRouterHandler() *controllers.IndexController {
	wire.Build(
		services.NewIndexServiceImpl,
		controllers.NewIndexController,
	)
	return &controllers.IndexController{}
}

func InitBillingRouterHandler() *controllers.BillingController {
	wire.Build(
		services.NewBillingServiceImpl,
		controllers.NewBillingController,
	)
	return &controllers.BillingController{}
}

func InitReturnRequestRouterHandler() *controllers.ReturnRequestController {
	wire.Build(
		services.NewReturnRequestServiceImpl,
		controllers.NewReturnRequestController,
	)
	return &controllers.ReturnRequestController{}
}

func InitRatingRouterHandler() *controllers.RatingController {
	wire.Build(
		services.NewRatingServiceImpl,
		services.NewStorageServiceImpl,
		controllers.NewRatingController,
	)
	return &controllers.RatingController{}
}

func InitMessageRouterHandler() *controllers.MessageController {
	wire.Build(
		services.NewSocketIOServiceImpl,
		controllers.NewMessageController,
	)
	return &controllers.MessageController{}
}