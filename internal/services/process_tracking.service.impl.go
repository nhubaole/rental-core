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

func (process *ProcessServiceImpl) CreateProcessTracking(body *dataaccess.CreateProcessTrackingParams) bool {
	_, er := process.repo.CreateProcessTracking(context.Background(), *body)
	if er != nil {
		fmt.Println(er.Error())
		return false
	}
	return true
}

func (process *ProcessServiceImpl) GetProcessTrackingByRentalId(userid int32, rentalId int32) *responses.ResponseData {

	var result []responses.GetProcessTracking
	processes, err := process.repo.GetProcessTrackingByRentalId(context.Background(), rentalId)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	for _, p := range processes {
		user, _ := process.repo.GetUserByID(context.Background(), p.Actor)

		result = append(result, responses.GetProcessTracking{
			ID: p.ID,
			Actor: user, 
			Action: p.Action,
			IssuedAt: p.IssuedAt,
			RequestID: p.RequestID,
		})
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
	}
}

func (process *ProcessServiceImpl) GetAllProcessTracking(userid int32) *responses.ResponseData {
	rs, er := process.repo.GetAllProcessTracking(context.Background(), userid)
	if er != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't find this Process",
			Data:       false,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Success",
		Data:       rs,
	}
}
