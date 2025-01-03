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
	UpdateDeviceToken(userId int, deviceToken string) *responses.ResponseData
	GetDeviceTokenByUserID(id int) (string, error)
	GetBankInfo(userID int32) *responses.ResponseData
}
