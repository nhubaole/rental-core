package services

import (
	"net/http"
	"smart-rental/internal/models"
	"smart-rental/internal/repo"
	"smart-rental/pkg/responses"

	"golang.org/x/crypto/bcrypt"
)

type AuthenServiceImpl struct {
	repo repo.AuthenRepo
}

func NewAuthenSerivceImpl(repo repo.AuthenRepo) AuthenService{
	return &AuthenServiceImpl{
		repo: repo,
	}
}

func(as *AuthenServiceImpl)Register(user *models.User) *responses.ResponseData{
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errHash != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Fail to hash password",
			Data:       false,
		}
	}

	user.Password = string(passwordHash)

	result, err := as.repo.CreateUser(user)

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       result,
	}

}