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

	fmt.Println(userid)
	rs, er := process.repo.GetRequestByUserID(context.Background(), userid)
	if er != nil {
		fmt.Println(er.Error() + "11111")
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't find this process",
			Data:       false,
		}
	}
	for _, r := range rs {
		if r.ID == rentalId {
			result, error := process.repo.GetProcessTrackingByRentalId(context.Background(), rentalId)
			if error != nil {
				fmt.Println(error.Error())
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
	fmt.Println("ERROR GetProcessTrackingByRentalId Service Impl ??")
	return &responses.ResponseData{
		StatusCode: http.StatusBadRequest,
		Message:    "Can't find this process",
		Data:       false,
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
