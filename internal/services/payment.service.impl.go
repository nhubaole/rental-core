package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type PaymentServiceImpl struct {
	repo *dataaccess.Queries
}



func NewPaymentServiceImpl() PaymentService {
	return &PaymentServiceImpl{
		repo: dataaccess.New(global.Db),
	}
}

// Create implements PaymentService.
func (p *PaymentServiceImpl) Create(req dataaccess.CreatePaymentParams) *responses.ResponseData {
	err := p.repo.CreatePayment(context.Background(), req)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       true,
	}

}

// GetByID implements PaymentService.
func (p *PaymentServiceImpl) GetByID(id int) *responses.ResponseData {
	payment, err := p.repo.GetPaymentByID(context.Background(), int32(id))
	if err != nil {
		if payment.ID == 0 {
			return &responses.ResponseData{
				StatusCode: http.StatusNoContent,
				Message:    responses.StatusNoData,
				Data:       nil,
			}
		}
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       payment,
	}
}
