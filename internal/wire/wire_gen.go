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

func InitContractRouterHandler() *controllers.ContractController {
	contractService := services.NewContractServiceImpl()
	contractController := controllers.NewContractController(contractService)
	return contractController
}
