package controllers

import (

	"smart-rental/internal/services"

	"smart-rental/pkg/requests"
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
	//newRoom := dataaccess.CreateRoomParams{}
	var formData requests.CreateRoomForm
	if err := ctx.ShouldBind(&formData); err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	// var params dataaccess.CreateRoomParams
	// common.MapStruct(formData, &params)

	result := rc.roomService.CreateRoom(formData)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)

}

