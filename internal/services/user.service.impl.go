package services

import (
	"context"
	"net/http"
	"smart-rental/initialize"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type UserServiceImpl struct {
	repo *dataaccess.Queries
}

func NewUserServiceImpl() UserService {
	return &UserServiceImpl{
		repo: dataaccess.New(initialize.DB),
	}
}

// GetAll implements UserService.
func (u *UserServiceImpl) GetAll() *responses.ResponseData {
	users, err := u.repo.GetUsers(context.Background())

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       users,
	}
}
