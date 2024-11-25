package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"

	"smart-rental/pkg/blockchain"
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

func (as *AuthenServiceImpl) Register(req *dataaccess.CreateUserParams) *responses.ResponseData {
	user, _ := as.repo.GetUserByPhone(context.Background(), req.PhoneNumber)
	if user.ID != 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusConflict,
			Message:    "user already exists",
			Data:       false,
		}
	}
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(string(req.Password)), 10)
	if errHash != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Fail to hash password",
			Data:       false,
		}
	}

	req.Password = string(passwordHash)
	opt := int32(common.GenerateDigitOTP())
	user.Otp = &opt
		// Generate Ethereum Wallet during registration
		walletAddress, _, errWallet := blockchain.CreateWallet(user.PhoneNumber)
		if errWallet != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    "Failed to create wallet",
				Data:       false,
			}
		}
	
		// Set wallet address in user params
		user.WalletAddress = &walletAddress
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
	user, err := as.repo.GetUserByPhone(context.Background(), req.PhoneNumber)

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
	common.MapStruct(user, &userResponse)
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
		Data: responses.LoginRes{
			AccessToken:  accessToken,
			RefreshToken: "not yet",
		},
	}
}

// VerifyOTP implements AuthenService.
func (as *AuthenServiceImpl) VerifyOTP(req *requests.VerifyOTPRequest) *responses.ResponseData {
	user, err := as.repo.GetUserByPhone(context.Background(), req.PhoneNumber)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	if user == (dataaccess.GetUserByPhoneRow{}) {
		return &responses.ResponseData{
			StatusCode: http.StatusNotFound,
			Message:    responses.StatusResourceNotFound,
			Data:       nil,
		}
	}

	if int(*user.Otp) != req.Otp {
		return &responses.ResponseData{
			StatusCode: http.StatusUnauthorized,
			Message:    responses.StatusAuthorizeFail,
			Data:       nil,
		}
	}

	var result responses.UserResponse
	common.MapStruct(user, &result)
	token, err := common.GenerateToken(result)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	var updateUser dataaccess.UpdateUserParams
	common.MapStruct(user, &updateUser)
	updateUser.Otp = nil

	_, updateErr := as.repo.UpdateUser(context.Background(), updateUser)
	if updateErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    updateErr.Error(),
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: 200,
		Message:    responses.StatusVerifySuccess,
		Data: responses.LoginRes{
			AccessToken:  token,
			RefreshToken: "not yet",
		},
	}
}
