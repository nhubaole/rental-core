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
	rrc := wire.InitRentalRequestRouterHandler()
	pc := wire.InitProcessTrackingRouterHandler()
	ic := wire.InitIndexRouterHandler()
	bc := wire.InitBillingRouterHandler()
	returnRequestRoute := wire.InitReturnRequestRouterHandler()
	ratingRoute := wire.InitRatingRouterHandler()
	ms := wire.InitMessageRouterHandler()
	conversationRoute := wire.InitConversationRouterHandler()
	payment := wire.InitPaymentRouterHandler()
	notification := wire.InitNotificationRouterHandler()
	
	cc := wire.InitContractRouterHandler()
	r := routers.NewRouter(ac, uc, rc, rrc, pc, ic, bc, cc, returnRequestRoute, ratingRoute, ms, conversationRoute, payment, notification)

	r.Run()
}
