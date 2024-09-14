package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type RentalRequestService interface {
	CreateRentalRequest(body *requests.CreateRentalRequest, myid int32) *responses.ResponseData
}
