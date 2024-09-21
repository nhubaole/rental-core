package services

import (
	"context"
	"fmt"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
)

type BillingServiceImpl struct {
	query *dataaccess.Queries
}

func NewBillingServiceImpl() BillingService {
	return &BillingServiceImpl{
		query: dataaccess.New(global.Db),
	}
}

func (service *BillingServiceImpl) CreateBill(userid int32, body *dataaccess.CreateBillParams) *responses.ResponseData {
	// Check user is the owner
	rs, er := service.query.GetIndexByOwnerIdShort(context.Background(), userid)
	if er != nil {
		fmt.Println(er.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Some error occured",
			Data:       false,
		}
	}
	for _, r := range rs {
		if r.ID == body.IndexID {
			result, error := service.query.CreateBill(context.Background(), *body)
			if error != nil {
				fmt.Println(error.Error())
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    "Some error occured",
					Data:       false,
				}
			}
			// r.RoomID -> rentalrequest.id
			rss, err := service.query.GetRentalRequestSuccessByRoomId(context.Background(), r.RoomID)
			if err != nil {
				fmt.Println(err.Error())
				return &responses.ResponseData{
					StatusCode: http.StatusInternalServerError,
					Message:    "Some error occured",
					Data:       false,
				}
			}

			track := dataaccess.CreateProcessTrackingParams{
				Actor:     userid,
				Action:    "Tạo hóa đơn",
				RequestID: rss.ID,
			}
			_, errr := service.query.CreateProcessTracking(context.Background(), track)
			if errr != nil {
				fmt.Println(errr.Error())
				fmt.Println("ERROR Create Process Tracking while creating bill!!")
				// return &responses.ResponseData{
				// 	StatusCode: http.StatusInternalServerError,
				// 	Message:    "Some error occured",
				// 	Data:       false,
				// }
			}

			return &responses.ResponseData{
				StatusCode: http.StatusOK,
				Message:    "Ok",
				Data:       result,
			}

		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusBadRequest,
		Message:    "Can't find your index",
		Data:       false,
	}
}

func (service *BillingServiceImpl) GetBillByMonth(userid int32, month int32, year int32) *responses.ResponseData {
	//
	rat := dataaccess.GetBillByMonthParams{
		Year:   year,
		Month:  month,
		PartyA: userid,
	}
	rs, er := service.query.GetBillByMonth(context.Background(), rat)
	if er != nil {
		fmt.Println(er.Error())
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't find your bill",
			Data:       false,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    "Ok",
		Data:       rs,
	}

}
