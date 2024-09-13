package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
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
	var result []responses.UserResponse
	for _,v := range users {
		var item responses.UserResponse
		common.MapStruct(v, &item)
		result = append(result, item)
	}
	

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
		Data:       result,
	}
}
