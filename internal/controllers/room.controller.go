package controllers

import (
	"smart-rental/internal/services"
	"strconv"

	"smart-rental/pkg/common"
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

func (rc RoomController) GetAll(ctx *gin.Context) {
	result := rc.roomService.GetRooms()
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (rc RoomController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := rc.roomService.GetRoomByID(id)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (rc RoomController) SearchByAddress(ctx *gin.Context) {
	searchParam := ctx.Query("search")

	result := rc.roomService.SearchRoomByAddress(searchParam)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (rc RoomController) Like(ctx *gin.Context) {
	roomId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}
	user, err := common.GetCurrentUser(ctx)
	if err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
		
	}

	result := rc.roomService.LikeRoom(int(roomId), int(user.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (rc RoomController) GetLikedRooms(ctx *gin.Context) {
	userId := 1 //TODO: get current user id
	result := rc.roomService.GetLikedRooms(userId)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}
