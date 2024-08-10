package repo

import "smart-rental/internal/models"

type AuthenRepo interface {
	CreateUser(user *models.User) (*int, error)
}