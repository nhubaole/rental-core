package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type UserService interface {
	GetAll() *responses.ResponseData
	GetUserByID(ID int) *responses.ResponseData
	Update(user *requests.UpdateUserReq) *responses.ResponseData
	CreateBankInfo(req *dataaccess.CreateUserBankParams) *responses.ResponseData
	UpdateBankInfo(req *dataaccess.UpdateUserBankParams) *responses.ResponseData
	GetBankInfo(userID int32) *responses.ResponseData
}
