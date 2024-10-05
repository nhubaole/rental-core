package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type ReturnRequestService interface {
	Create(req dataaccess.CreateReturnRequestParams, userID int) *responses.ResponseData
	GetByID(id int) *responses.ResponseData
	Aprrove(id int, userID int) *responses.ResponseData
}