package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/responses"

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
			var rating responses.RatingInfo
			rating.Comment = *v.Comments
			rating.CreatedAt = pgtype.Timestamptz(v.CreatedAt)
			rating.RaterName = user.FullName
			rating.Rate = int(*v.OverallRating)
			userDetail.RatingInfo = append(userDetail.RatingInfo, rating)
			totalRating = totalRating + int(*v.OverallRating)
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
	} else if user.Role == 2 {
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
			var rating responses.RatingInfo
			rating.Comment = *v.Comments
			rating.CreatedAt = pgtype.Timestamptz(v.CreatedAt)
			rating.RaterName = user.FullName
			rating.Rate = int(*v.OverallRating)
			userDetail.RatingInfo = append(userDetail.RatingInfo, rating)
			totalRating = totalRating + int(*v.OverallRating)
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
