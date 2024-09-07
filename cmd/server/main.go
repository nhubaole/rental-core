package main

import (

	"smart-rental/internal/initialize"
	"smart-rental/internal/routers"
	"smart-rental/internal/wire"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	initialize.Run()

	ac := wire.InitAuthenRouterHandler()
	uc := wire.InitUserRouterHandler()
	rc := wire.InitRoomRouterHandler()
	r := routers.NewRouter(ac, uc, rc)

	r.Run()
}
