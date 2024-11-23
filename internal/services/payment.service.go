package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type PaymentService interface {
	GetByID(id int) *responses.ResponseData
	Create(req dataaccess.CreatePaymentParams)*responses.ResponseData
}