package main

import (
	"smart-rental/internal/routers"
	"smart-rental/internal/wire"
)

func main() {

	ac := wire.InitAuthenRouterHandler()
	r := routers.NewRouter(ac)
	r.Run()
}
