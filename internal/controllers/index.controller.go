package controllers

import (
	"net/http"
	"smart-rental/internal/dataaccess"
	"smart-rental/internal/services"
	"smart-rental/pkg/common"
	"smart-rental/pkg/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IndexServiceController struct {
	indexService services.IndexService
}

func NewIndexServiceController(service services.IndexService) *IndexServiceController {
	return &IndexServiceController{
		indexService: service,
	}
}

func (controller IndexServiceController) GetIndexFromOwner(ctx *gin.Context) {
	myuser, _ := common.GetCurrentUser(ctx)
	year, err := strconv.Atoi(ctx.Param("year"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	month, err := strconv.Atoi(ctx.Param("month"))
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	result := controller.indexService.GetAllIndex(myuser.ID, int32(month), int32(year))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (controller IndexServiceController) CreateIndex(ctx *gin.Context) {
	myuser, _ := common.GetCurrentUser(ctx)
	var body *dataaccess.CreateIndexParams
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		responses.APIResponse(ctx, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	result := controller.indexService.CreateIndex(myuser.ID, body)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}
