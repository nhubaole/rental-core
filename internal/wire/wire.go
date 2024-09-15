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
		//initialize.DatabaseConnection,
		//	initialize.NewQueries,
		//dataaccess.New,
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

func InitContractRouterHandler() *controllers.ContractController {
	wire.Build(
		services.NewContractServiceImpl,
		controllers.NewContractController,
	)
	return &controllers.ContractController{}
}
