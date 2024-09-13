package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"

	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
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
// Login implements AuthenService.
func (as *AuthenServiceImpl) Login(req *requests.LoginRequest) *responses.ResponseData {
	user, err := as.repo.GetUserByPhone(context.Background(),req.PhoneNumber)

	if err != nil {	
		return &responses.ResponseData{
			StatusCode: 401,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	// Compare hashed password
	success := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if success != nil {
		return &responses.ResponseData{
			StatusCode: 401,
			Message:    responses.StatusAuthorizeFail,
			Data:       nil,
		}
	}

	var userResponse responses.UserResponse
	accessToken, errToken := common.GenerateToken(userResponse)

	if errToken != nil {
		return &responses.ResponseData{
			StatusCode: 401,
			Message:    "Failed to create token",
			Data:       nil,
		}
	}
	return &responses.ResponseData{
		StatusCode: 200,
		Message:    responses.StatusSuccess,
		Data:       responses.LoginRes{
			AccessToken: accessToken,
			RefreshToken: "not yet",
		},
	}
}
