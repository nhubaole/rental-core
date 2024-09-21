package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)


type BillingService interface{
	CreateBill(userid int32, body *dataaccess.CreateBillParams) *responses.ResponseData
	GetBillByMonth(userid int32,month int32, year int32) *responses.ResponseData
}