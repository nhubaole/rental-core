package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)


type BillingService interface{
	CreateBill(userID int32, body dataaccess.CreateBillParams) *responses.ResponseData
	GetBillByMonth(userID int32,month int32, year int32) *responses.ResponseData
	GetBillByID(id int32) *responses.ResponseData
	GetBillMetrics(req dataaccess.GetAllMetric4BillByRoomIDParams) *responses.ResponseData
	GetBillByStatus(statusID int32) *responses.ResponseData
	GetBillOfRentedRoomByOwnerID(currentUser int) *responses.ResponseData
}