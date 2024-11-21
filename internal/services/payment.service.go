package services

import "smart-rental/pkg/responses"

type PaymentService interface {
	GetByID(id int) *responses.ResponseData
	// CreatePayment()
}