package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type ContractService interface {
	CreateTemplate(req dataaccess.CreateContractTemplateParams) *responses.ResponseData
	GetTemplateByAddress(address requests.GetTemplateByAddressRequest) *responses.ResponseData
	CreateContract(req requests.CreateContractRequest,userID int) *responses.ResponseData
	GetContractByID(id int) *responses.ResponseData
	ListContractByStatus(statusID int, userId int, isLandlord bool) *responses.ResponseData
	SignContract(req requests.SignContractParams, userID int) *responses.ResponseData
	DeclineContract(id int, userID int) *responses.ResponseData
	GetContractByUser(userID int) *responses.ResponseData
	GetTemplateByOwner(userID int32) *responses.ResponseData
}