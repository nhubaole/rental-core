package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type BillingServiceImpl struct {
	query *dataaccess.Queries
}



func NewBillingServiceImpl() BillingService {
	return &BillingServiceImpl{
		query: dataaccess.New(global.Db),
	}
}

func (service *BillingServiceImpl) CreateBill(userID int32, body dataaccess.CreateBillParams) *responses.ResponseData {
	err := service.query.CreateBill(context.Background(), body)
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

func (service *BillingServiceImpl) GetBillByMonth(userID int32, month int32, year int32) *responses.ResponseData {
	param := dataaccess.GetBillByMonthParams{
		Year:   year,
		Month:  month,
		PartyA: userID,
	}
	bill, err := service.query.GetBillByMonth(context.Background(), param)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't find your bill",
			Data:       false,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Ok",
		Data:       bill,
	}
}

// GetBillByID implements BillingService.
func (service *BillingServiceImpl) GetBillByID(id int32) *responses.ResponseData {
	bill, err := service.query.GetBillByID(context.Background(), id)
	if err != nil {
		if (bill == dataaccess.GetBillByIDRow{}) {
			return &responses.ResponseData{
				StatusCode: http.StatusNoContent,
				Message:    "No bill found",
				Data:       false,
			}
		}
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       bill,
	}
}

// GetBillMetric implements BillingService.
func (service *BillingServiceImpl) GetBillMetrics(req dataaccess.GetAllMetric4BillByRoomIDParams) *responses.ResponseData {
	metric, err := service.query.GetAllMetric4BillByRoomID(context.Background(), req)
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
		Data:       metric,
	}
}

// GetBillByStatus implements BillingService.
func (service *BillingServiceImpl) GetBillByStatus(statusID int32) *responses.ResponseData {
	bills, err := service.query.GetBillByStatus(context.Background(), &statusID)

	if err != nil {
		if len(bills) == 0 {
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
		Data:       bills,
	}
}

// GetBillOfRentedRoomByOwnerID implements BillingService.
func (service *BillingServiceImpl) GetBillOfRentedRoomByOwnerID(currentUser int) *responses.ResponseData {
	bills, err := service.query.GetBillOfRentedRoomByOwnerID(context.Background(), int32(currentUser))
	if err != nil {
		if len(bills) == 0 {
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
		Data:       bills,
	}
}