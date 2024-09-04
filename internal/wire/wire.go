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
		//initialize.DatabaseConnection,
		//	initialize.NewQueries,
		//dataaccess.New,
		services.NewUserServiceImpl,
		controllers.NewUserController,
	)

	return &controllers.UserController{}
}
