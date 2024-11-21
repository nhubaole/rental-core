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

// CreateBankInfo implements UserService.
func (u *UserServiceImpl) CreateBankInfo(req *dataaccess.CreateUserBankParams) *responses.ResponseData {

	
	err := u.repo.CreateUserBank(context.Background(), *req)
	if err != nil {
		println(string(err.Error()))
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}

// UpdateBankInfo implements UserService.
func (u *UserServiceImpl) UpdateBankInfo(req *dataaccess.UpdateUserBankParams) *responses.ResponseData {
	
	updatedUserBank, err := u.repo.UpdateUserBank(context.Background(),*req)
	if err != nil {
		println(string(err.Error()))
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       updatedUserBank.UserID,
	}
}

// GetAll implements UserService.
func (u *UserServiceImpl) GetAll() *responses.ResponseData {
	users, err := u.repo.GetUsers(context.Background())
	var result []responses.UserResponse
	for _, v := range users {
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

func (userRepo *UserServiceImpl) Update(body *dataaccess.UpdateUserParams) *responses.ResponseData {
	user, err := userRepo.repo.UpdateUser(context.Background(), *body)
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
		Data:       user.ID,
	}
}
