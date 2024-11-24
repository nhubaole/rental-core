package services

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"path/filepath"
	"smart-rental/global"
	"smart-rental/internal/constants"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"time"
)

type PaymentServiceImpl struct {
	repo *dataaccess.Queries
	storageService StorageSerivce
}



func NewPaymentServiceImpl(storage StorageSerivce) PaymentService {
	return &PaymentServiceImpl{
		repo: dataaccess.New(global.Db),
		storageService: storage,
	}
}

// Create implements PaymentService.
func (p *PaymentServiceImpl) Create(req requests.CreatePaymentReq, userID int32) *responses.ResponseData {
	f, _ := req.EvidenceImage.Open() 
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	fileExt := filepath.Ext(req.EvidenceImage.Filename)
	contentType := mime.TypeByExtension(fileExt)
	objKey := fmt.Sprintf("%s/%s/%d%s",constants.PAYMENT_OBJ, "payment", timestamp, fileExt)

	url, err := p.storageService.UploadFile(constants.BUCKET_NAME, objKey, f, contentType)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}


	var params dataaccess.CreatePaymentParams
	common.MapStruct(req, &params)
	params.EvidenceImage = &url
	params.SenderID = userID

	createErr := p.repo.CreatePayment(context.Background(), params)

	if createErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    createErr.Error(),
			Data:       false,
		}
	}
	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       true,
	}

}

// GetByID implements PaymentService.
func (p *PaymentServiceImpl) GetByID(id int) *responses.ResponseData {
	payment, err := p.repo.GetPaymentByID(context.Background(), int32(id))
	if err != nil {
		if payment.ID == 0 {
			return &responses.ResponseData{
				StatusCode: http.StatusNoContent,
				Message:    responses.StatusNoData,
				Data:       nil,
			}
		}
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       payment,
	}
}


// GetAllBanks implements PaymentService.
func (p *PaymentServiceImpl) GetAllBanks() *responses.ResponseData {
	banks, err := p.repo.GetAllBanks(context.Background())
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
		Data:       banks,
	}
}
