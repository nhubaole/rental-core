package main

import (
	"smart-rental/initialize"
	"smart-rental/internal/routers"
	"smart-rental/internal/wire"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	initialize.DatabaseConnection()
	// db := dataaccess.New(initialize.DB)
	// as := services.NewAuthenSerivceImpl(db)
	// ac := controllers.NewAuthController(as)

	ac := wire.InitAuthenRouterHandler()
	uc := wire.InitUserRouterHandler()
	r := routers.NewRouter(ac, uc)
	r.Run()
}
