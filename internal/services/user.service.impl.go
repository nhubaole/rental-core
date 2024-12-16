package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/responses"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
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

	updatedUserBank, err := u.repo.UpdateUserBank(context.Background(), *req)
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
	var userDetail responses.GetUserByIDRes

	common.MapStruct(user, &userDetail)

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "This user can't be found",
			Data:       nil,
		}
	}

	if user.Role == 1 {
		ratings, err := userRepo.repo.GetLandlordRatingByID(context.Background(), &user.ID)

		totalRating := 0

		if len(ratings) == 0 {
			userDetail.RatingInfo = nil
		}
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}
		for _, v := range ratings {
			user, err := userRepo.repo.GetUserByID(context.Background(), *v.RatedBy)
			if err != nil {
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
					Data:       nil,
				}
			}
			var happyTexts []string
			var unhappyTexts []string

			if v.FriendlinessRating != nil {
				if *v.FriendlinessRating == 5 {
					happyTexts = append(happyTexts, "thân thiện")
				} else if *v.FriendlinessRating == 1 {
					unhappyTexts = append(unhappyTexts, "thiếu thân thiện")
				}
			}
			if v.ProfessionalismRating != nil {
				if *v.ProfessionalismRating == 5 {
					happyTexts = append(happyTexts, "chuyên nghiệp")
				} else if *v.ProfessionalismRating == 1 {
					unhappyTexts = append(unhappyTexts, "thiếu chuyên nghiệp")
				}
			}
			if v.SupportRating != nil {
				if *v.SupportRating == 5 {
					happyTexts = append(happyTexts, "hỗ trợ tốt")
				} else if *v.SupportRating == 1 {
					unhappyTexts = append(unhappyTexts, "thiếu hỗ trợ")
				}
			}
			if v.TransparencyRating != nil {
				if *v.TransparencyRating == 5 {
					happyTexts = append(happyTexts, "minh bạch")
				} else if *v.TransparencyRating == 1 {
					unhappyTexts = append(unhappyTexts, "thiếu minh bạch")
				}
			}

			happyString := strings.Join(happyTexts, ", ")
			unhappyString := strings.Join(unhappyTexts, ", ")

			var rating responses.RatingInfo
			rating.Comment = *v.Comments
			rating.CreatedAt = pgtype.Timestamptz(v.CreatedAt)
			rating.RaterName = user.FullName
			rating.Rate = int(*v.OverallRating)
			rating.Happy = happyString
			rating.UnHappy = unhappyString

			userDetail.RatingInfo = append(userDetail.RatingInfo, rating)
			totalRating += int(*v.OverallRating)

		}
		rooms, err := userRepo.repo.GetRoomsByOwner(context.Background(), user.ID)

		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}

		userDetail.TotalRoom = len(rooms)
		userDetail.TotalRating = len(ratings)
		userDetail.AvgRating = float64(totalRating / len(ratings))
	} else if user.Role == 0 {
		totalRating := 0
		ratings, err := userRepo.repo.GetTenantRatingByID(context.Background(), &user.ID)
		if len(ratings) == 0 {
			userDetail.RatingInfo = nil
		}
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}

		for _, v := range ratings {
			user, err := userRepo.repo.GetUserByID(context.Background(), *v.RatedBy)
			if err != nil {
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    err.Error(),
					Data:       nil,
				}
			}

			var happyTexts []string
			var unhappyTexts []string

			if v.PaymentRating != nil {
				if *v.PaymentRating == 5 {
					happyTexts = append(happyTexts, "thanh toán đúng hạn")
				} else if *v.PaymentRating == 1 {
					unhappyTexts = append(unhappyTexts, "thanh toán trễ")
				}
			}
			if v.PropertyCareRating != nil {
				if *v.PropertyCareRating == 5 {
					happyTexts = append(happyTexts, "bảo quản tài sản tốt")
				} else if *v.PropertyCareRating == 1 {
					unhappyTexts = append(unhappyTexts, "không bảo quản tài sản")
				}
			}
			if v.NeighborhoodDisturbanceRating != nil {
				if *v.NeighborhoodDisturbanceRating == 5 {
					happyTexts = append(happyTexts, "hòa đồng với hàng xóm")
				} else if *v.NeighborhoodDisturbanceRating == 1 {
					unhappyTexts = append(unhappyTexts, "gây phiền toái hàng xóm")
				}
			}
			if v.ContractComplianceRating != nil {
				if *v.ContractComplianceRating == 5 {
					happyTexts = append(happyTexts, "tuân thủ hợp đồng")
				} else if *v.ContractComplianceRating == 1 {
					unhappyTexts = append(unhappyTexts, "vi phạm hợp đồng")
				}
			}

			happyString := strings.Join(happyTexts, ", ")
			unhappyString := strings.Join(unhappyTexts, ", ")

			var rating responses.RatingInfo
			rating.Comment = *v.Comments
			rating.CreatedAt = pgtype.Timestamptz(v.CreatedAt)
			rating.RaterName = user.FullName
			rating.Rate = int(*v.OverallRating)
			rating.Happy = happyString
			rating.UnHappy = unhappyString

			userDetail.RatingInfo = append(userDetail.RatingInfo, rating)
			totalRating += int(*v.OverallRating)
		}

		rooms, err := userRepo.repo.GetRoomByTenantID(context.Background(), user.ID)

		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}
		userDetail.TotalRoom = len(rooms)
		userDetail.TotalRating = len(ratings)
		userDetail.AvgRating = float64(totalRating / len(ratings))
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       userDetail,
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

// GetDeviceTokenByUserID implements UserService.
func (u *UserServiceImpl) GetDeviceTokenByUserID(id int) (string, error) {
	deviceToken, err := u.repo.GetDeviceTokenByUserID(context.Background(), int32(id))
	if err != nil {
		println(string(err.Error()))
		return err.Error(), err
	}
	return *deviceToken, err
}

// UpdateDeviceToken implements UserService.
func (u *UserServiceImpl) UpdateDeviceToken(userId int, deviceToken string) *responses.ResponseData {
	req := dataaccess.UpdateDeviceTokenParams{
		ID:          int32(userId),
		DeviceToken: &deviceToken,
	}
	err := u.repo.UpdateDeviceToken(context.Background(), req)
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
		Data:       nil,
	}
}