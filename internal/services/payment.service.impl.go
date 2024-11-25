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
	repo           *dataaccess.Queries
	repo           *dataaccess.Queries
	storageService StorageSerivce
}



func NewPaymentServiceImpl(storage StorageSerivce) PaymentService {
	return &PaymentServiceImpl{
		repo:           dataaccess.New(global.Db),
		repo:           dataaccess.New(global.Db),
		storageService: storage,
	}
}

// GetDetailInfo implements PaymentService.
func (p *PaymentServiceImpl) GetDetailInfo(typeOfPayment string, id int) *responses.ResponseData {
	var response responses.GetPaymentInfoRes

	// switch typeOfPayment {
	// case "contract":
	// 	_, depositAmount, err := getContractAmount(client, contractAddress, callerAddress)
	// 	if err != nil {
	// 		return &responses.ResponseData{
	// 			StatusCode: http.StatusInternalServerError,
	// 			Message:    responses.StatusInternalError,
	// 			Data:       nil,
	// 		}
	// 	}
	// 	response.Amount = depositAmount.String() // Sử dụng depositAmount
	// case "bill":
	// 	billingId := big.NewInt(int64(id)) // ID hóa đơn
	// 	totalAmount, err := getBillingAmount(client, contractAddress, billingId)
	// 	if err != nil {
	// 		return &responses.ResponseData{
	// 			StatusCode: http.StatusInternalServerError,
	// 			Message:    responses.StatusInternalError,
	// 			Data:       nil,
	// 		}
	// 	}
	// 	response.Amount = totalAmount.String() // Sử dụng totalAmount
	// case "return":
	// 	_, totalReturnDeposit, err := getReturnRequestAmount(client, contractAddress, callerAddress)
	// 	if err != nil {
	// 		return &responses.ResponseData{
	// 			StatusCode: http.StatusInternalServerError,
	// 			Message:    responses.StatusInternalError,
	// 			Data:       nil,
	// 		}
	// 	}
	// 	response.Amount = totalReturnDeposit.String() // Sử dụng totalReturnDeposit
	// default:
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusBadRequest,
	// 		Message:    "Invalid type of payment",
	// 		Data:       nil,
	// 	}
	// }
	
	prefix := fmt.Sprintf("SR%d", id)
	response.TranferContent = common.GenerateCode(prefix)
	response.BankName = "Techcombank"
	response.AccountName = "Le Bao Nhu"
	response.AccountNumber = "092321234123"
	response.QrUrl = "httpasdasdakjdnakasdnasdk"

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       response,
	}
}

// Confirm implements PaymentService.
func (p *PaymentServiceImpl) Confirm(id int) *responses.ResponseData {
	paymentIdUPdated, err := p.repo.ConfirmPayment(context.Background(), int32(id))
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
		Data:       paymentIdUPdated,
	}
}

// GetAll implements PaymentService.
func (p *PaymentServiceImpl) GetAll() *responses.ResponseData {
	payments, err := p.repo.GetAllPayments(context.Background())
	if len(payments) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusNoContent,
			Message:    responses.StatusNoData,
			Data:       nil,
		}
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
		Data:       payments,
	}
}

// Create implements PaymentService.
func (p *PaymentServiceImpl) Create(req requests.CreatePaymentReq, userID int32) *responses.ResponseData {
	f, _ := req.EvidenceImage.Open()
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	fileExt := filepath.Ext(req.EvidenceImage.Filename)
	contentType := mime.TypeByExtension(fileExt)
	objKey := fmt.Sprintf("%s/%s/%d%s", constants.PAYMENT_OBJ, "payment", timestamp, fileExt)

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
	params.Status = 0

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
