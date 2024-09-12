package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type UserServiceImpl struct {
	repo *dataaccess.Queries
}

func NewUserServiceImpl() UserService {
	return &UserServiceImpl{
		repo: dataaccess.New(global.Db),
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

func (userRepo *UserServiceImpl) GetUserByID(id int) *responses.ResponseData {
	user, err := userRepo.repo.GetUserByID(context.Background(), int32(id))

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "This user can't be found",
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       user,
	}
}

func (userRepo *UserServiceImpl) Update(body *dataaccess.UpdateUserParam) *responses.ResponseData {
	err := userRepo.repo.UpdateUser(context.Background(), body)
	if err != nil {
		println(string(err.Error()))
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to finish your request",
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusNoContent,
		Message:    responses.StatusSuccess,
	}
}
