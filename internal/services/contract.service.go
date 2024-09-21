package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type ContractService interface {
	CreateTemplate(req dataaccess.CreateContractTemplateParams) *responses.ResponseData
	GetTemplateByAddress(address requests.GetTemplateByAddressRequest) *responses.ResponseData
	CreateContract(req requests.CreateContractRequest) *responses.ResponseData
	GetContractByID(id int) *responses.ResponseData
	ListContractByStatus(statusID int) *responses.ResponseData
}