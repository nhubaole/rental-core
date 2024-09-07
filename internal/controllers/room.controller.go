package controllers

import (
	"fmt"
	"smart-rental/internal/dataaccess"
	"smart-rental/internal/services"
	"smart-rental/pkg/responses"

	"github.com/gin-gonic/gin"
)

type RoomController struct {
	roomService services.RoomService
}

func NewRoomController(service services.RoomService) *RoomController {
	return &RoomController{
		roomService: service,
	}
}

func (rc RoomController) Create(ctx *gin.Context) {
	newRoom := dataaccess.CreateRoomParams{}
	if err := ctx.ShouldBindJSON(&newRoom); err != nil {
		fmt.Println(err)
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := rc.roomService.CreateRoom(newRoom)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)

}
