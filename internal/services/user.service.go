package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type UserService interface {
	GetAll() *responses.ResponseData
	GetUserByID(ID int) *responses.ResponseData 
	Update(user *dataaccess.UpdateUserParams) *responses.ResponseData
	CreateBankInfo(req *dataaccess.CreateUserBankParams) *responses.ResponseData
	UpdateBankInfo(req *dataaccess.UpdateUserBankParams) *responses.ResponseData
}
