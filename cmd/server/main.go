package main

import (
	"fmt"

	"smart-rental/global"
	"smart-rental/internal/initialize"
	"smart-rental/internal/routers"
	"smart-rental/internal/wire"

	_ "github.com/joho/godotenv/autoload"
	"github.com/redis/go-redis/v9"
)

func main() {
	initialize.Run()
	client := redis.NewClient(&redis.Options{
        Addr:	  global.Config.DB.DBHost,
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
	fmt.Println(client, "=====")
    fmt.Println("DB ConnString:", global.Db)
	ac := wire.InitAuthenRouterHandler()
	uc := wire.InitUserRouterHandler()
	rc := wire.InitRoomRouterHandler()
	r := routers.NewRouter(ac, uc, rc)
	fmt.Println("============",global.Config.Server.Port)

	r.Run()
}
