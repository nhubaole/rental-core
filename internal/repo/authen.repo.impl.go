package repo

import (
	"gorm.io/gorm"
	"smart-rental/internal/models"
)

type AuthenRepoImpl struct {
	db *gorm.DB
}

func NewAuthenRepoImpl(Db *gorm.DB) AuthenRepo {
	return &AuthenRepoImpl{db: Db}
}

// Register implements AuthenRepo.
func (ar *AuthenRepoImpl) CreateUser(user *models.User) (*int, error) {

	result := ar.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user.ID, nil
}
