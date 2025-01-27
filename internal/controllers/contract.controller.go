package controllers

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/internal/services"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContractController struct {
	services services.ContractService
}

func NewContractController(services services.ContractService) *ContractController {
	return &ContractController{
		services: services,
	}
}

func(cc ContractController) CreateTemplate(ctx *gin.Context) {
	req := dataaccess.CreateContractTemplateParams{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := cc.services.CreateTemplate(req)
	responses.APIResponse(ctx, result.StatusCode,result.Message, result.Data)
}

func(cc ContractController) GetTemplateByAddress(ctx *gin.Context) {
	var req	requests.GetTemplateByAddressRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := cc.services.GetTemplateByAddress(req)
	responses.APIResponse(ctx, result.StatusCode,result.Message, result.Data)
}

func(cc ContractController) GetTemplateByOwner(ctx *gin.Context) {
	user, err := common.GetCurrentUser(ctx)
	if err != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return

	}

	result := cc.services.GetTemplateByOwner(user.ID)
	responses.APIResponse(ctx, result.StatusCode,result.Message, result.Data)
}

func(cc ContractController) Create(ctx *gin.Context) {
	var req	requests.CreateContractRequest
	user, err := common.GetCurrentUser(ctx)
	if err != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return

	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}
	
	result := cc.services.CreateContract(req, int(user.ID))
	responses.APIResponse(ctx, result.StatusCode,result.Message, result.Data)
}

func (cc ContractController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}

	result := cc.services.GetContractByID(id)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (cc ContractController) GetByStatus(ctx *gin.Context) {
	statusID, err := strconv.Atoi(ctx.Param("statusID"))
	if err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}
	currentUser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return
	}
	isLandlord := currentUser.Role == 1

	result := cc.services.ListContractByStatus(statusID, int(currentUser.ID), isLandlord)
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (cc ContractController) SignContract(ctx *gin.Context) {
	var req	requests.SignContractParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}
	currentUser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return
	}

	result := cc.services.SignContract(req, int(currentUser.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (cc ContractController) DeclineContract(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		responses.APIResponse(ctx, 400, "Bad request", nil)
		return
	}
	currentUser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return
	}

	result := cc.services.DeclineContract(id, int(currentUser.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}

func (cc ContractController) GetByUser(ctx *gin.Context) {
	
	currentUser, errr := common.GetCurrentUser(ctx)
	if errr != nil {
		responses.APIResponse(ctx, 401, "Unauthorized", nil)
		return
	}

	result := cc.services.GetContractByUser(int(currentUser.ID))
	responses.APIResponse(ctx, result.StatusCode, result.Message, result.Data)
}