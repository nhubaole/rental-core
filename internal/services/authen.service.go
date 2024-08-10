package services

import (
	"smart-rental/internal/models"
	"smart-rental/pkg/responses"
)

type AuthenService interface {
	Register(user *models.User) *responses.ResponseData
}