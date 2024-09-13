package services

import (
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type AuthenService interface {
	Register(user *dataaccess.CreateUserParams) *responses.ResponseData
	Login(req *requests.LoginRequest) *responses.ResponseData
}
