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
    fmt.Println("DB ConnString:", global.Db)
	ac := wire.InitAuthenRouterHandler()
	uc := wire.InitUserRouterHandler()
	r := routers.NewRouter(ac, uc)
	fmt.Println("============",global.Config.Server.Port)

	r.Run()
}
