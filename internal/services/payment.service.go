package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type PaymentService interface {
	GetByID(id int) *responses.ResponseData
	Create(req requests.CreatePaymentReq)*responses.ResponseData
}