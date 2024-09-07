package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"

	"smart-rental/pkg/responses"

	"golang.org/x/crypto/bcrypt"
)

type AuthenServiceImpl struct {
	repo *dataaccess.Queries
}

func NewAuthenSerivceImpl() AuthenService {
	return &AuthenServiceImpl{
		repo: dataaccess.New(global.Db),
	}
}

func (as *AuthenServiceImpl) Register(user *dataaccess.CreateUserParams) *responses.ResponseData {
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(string(user.Password)), 10)
	if errHash != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Fail to hash password",
			Data:       false,
		}
	}

	user.Password = string(passwordHash)
	err := as.repo.CreateUser(context.Background(), *user)

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
		Data:       true,
	}

}
