// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"smart-rental/internal/controllers"
	"smart-rental/internal/database"
	"smart-rental/internal/repo"
	"smart-rental/internal/services"
)

// Injectors from authen.go:

func InitAuthenRouterHandler() *controllers.AuthenController {
	db := database.DatabaseConnection()
	authenRepo := repo.NewAuthenRepoImpl(db)
	authenService := services.NewAuthenSerivceImpl(authenRepo)
	authenController := controllers.NewAuthController(authenService)
	return authenController
}
