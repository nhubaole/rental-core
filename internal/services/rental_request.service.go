package services

import (
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type RentalRequestService interface {
	CreateRentalRequest(body *requests.CreateRentalRequest, myid int32) *responses.ResponseData
	DeleteRentalRequest(rentid int32, myid int32) *responses.ResponseData
	GetRentalRequestById(rentid int32, myid int32) *responses.ResponseData
	GetAllRentalRequest(phone string) *responses.ResponseData
}
