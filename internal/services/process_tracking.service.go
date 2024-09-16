package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type ProcessService interface {
	CreateProcessTracking(body *dataaccess.CreateProcessTrackingParams) bool
	GetProcessTrackingByRentalId(userid int32, rentalId int32) *responses.ResponseData
	GetAllProcessTracking(userid int32) *responses.ResponseData
}
