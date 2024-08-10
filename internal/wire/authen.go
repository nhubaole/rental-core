//go:build wireinject

package wire

import (
	"smart-rental/internal/controllers"
	"smart-rental/internal/database"
	"smart-rental/internal/repo"
	"smart-rental/internal/services"

	"github.com/google/wire"
)

func InitAuthenRouterHandler() *controllers.AuthenController{
	wire.Build(database.DatabaseConnection,repo.NewAuthenRepoImpl, services.NewAuthenSerivceImpl, controllers.NewAuthController)

	return &controllers.AuthenController{}
}