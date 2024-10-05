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
	rrc := wire.InitRentalRequestRouterHandler()
	pc := wire.InitProcessTrackingRouterHandler()
	ic := wire.InitIndexRouterHandler()
	bc := wire.InitBillingRouterHandler()
	returnRequestRoute := wire.InitReturnRequestRouterHandler()
	ratingRoute := wire.InitRatingRouterHandler()
	
	cc := wire.InitContractRouterHandler()
	r := routers.NewRouter(ac, uc, rc, rrc, pc, ic, bc, cc, returnRequestRoute, ratingRoute)

	r.Run()
}
