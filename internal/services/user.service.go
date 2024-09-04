package services

import "smart-rental/pkg/responses"

type UserService interface {
	GetAll() *responses.ResponseData
}
