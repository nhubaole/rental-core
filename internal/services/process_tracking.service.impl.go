package services

import (
	"context"
	"fmt"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type ProcessServiceImpl struct {
	repo *dataaccess.Queries
}

func NewProcessServiceImpl() ProcessService {
	return &ProcessServiceImpl{
		repo: dataaccess.New(global.Db),
	}
}

func (process *ProcessServiceImpl) CreateProcessTracking(body *dataaccess.CreateProgressTrackingParams) bool {
	_, er := process.repo.CreateProgressTracking(context.Background(), *body)
	if er != nil {
		fmt.Println(er.Error())
		return false
	}
	return true
}

func (process *ProcessServiceImpl) GetProcessTrackingByRentalId(userid int32, rentalId int32) *responses.ResponseData {
	rentalService := new(RentalRequestServiceImpl)
	rs, er := rentalService.repo.GetRequestByUserID(context.Background(), userid)
	if er != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't find this process",
			Data:       false,
		}
	}
	for _, r := range rs {
		if r.ID == rentalId {
			result, error := process.repo.GetProgressTrackingByRentalId(context.Background(), rentalId)
			if error != nil {
				return &responses.ResponseData{
					StatusCode: http.StatusBadRequest,
					Message:    "Can't find this process",
					Data:       false,
				}
			}
			return &responses.ResponseData{
				StatusCode: http.StatusOK,
				Message:    "Success",
				Data:       result,
			}
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusBadRequest,
		Message:    "Can't find this process",
		Data:       false,
	}
}

func (process *ProcessServiceImpl) GetAllProcessTracking(userid int32) *responses.ResponseData {
	rs, er := process.repo.GetAllProgressTracking(context.Background(), userid)
	if er != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't find this progress",
			Data:       false,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       rs,
	}
}
