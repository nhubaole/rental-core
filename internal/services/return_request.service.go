package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type ReturnRequestService interface {
	Create(req requests.CreateReturnRequestParams, userID int) *responses.ResponseData
	GetByID(id int) *responses.ResponseData
	Aprrove(id int, userID int) *responses.ResponseData
	GetByLandlordID(userID int) *responses.ResponseData
	GetByTenantID(userID int) *responses.ResponseData
}
