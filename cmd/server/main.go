package main

import (
	"fmt"
	"smart-rental/global"
	"smart-rental/internal/initialize"
	"smart-rental/internal/routers"
	"smart-rental/internal/wire"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	initialize.Run()
	fmt.Println("====", global.S3)
	ac := wire.InitAuthenRouterHandler()
	uc := wire.InitUserRouterHandler()
	rc := wire.InitRoomRouterHandler()
	cc := wire.InitContractRouterHandler()
	r := routers.NewRouter(ac, uc, rc, cc)

	r.Run()
}
