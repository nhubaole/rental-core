package services

import (
	"context"
	"fmt"
	"mime"
	"net/http"
	"net/url"
	"path/filepath"
	"smart-rental/global"
	"smart-rental/internal/constants"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"strconv"
	"time"
)

type PaymentServiceImpl struct {
	repo           *dataaccess.Queries
	storageService StorageSerivce
	blockchain     BlockchainService
}

func NewPaymentServiceImpl(storage StorageSerivce, blockchain BlockchainService) PaymentService {
	return &PaymentServiceImpl{
		repo:           dataaccess.New(global.Db),
		storageService: storage,
		blockchain:     blockchain,
	}
}

// GetDetailInfo implements PaymentService.
func (p *PaymentServiceImpl) GetDetailInfo(typeOfPayment string, id int32) *responses.ResponseData {
	var amount = 0.0
	contractId := int32(0)

	if typeOfPayment == "contract" {
		contractId = id
	} else if typeOfPayment == "bill" {
		bill, err := p.repo.GetBillByID(context.Background(), id)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    responses.StatusInternalError,
				Data:       nil,
			}
		}
		amount = bill.TotalAmount
		contractId = bill.ContractID
	} else if typeOfPayment == "return" {
		returnRequest, err := p.repo.GetReturnRequestByID(context.Background(), id)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    responses.StatusInternalError,
				Data:       nil,
			}
		}
		amount = *returnRequest.TotalReturnDeposit
		contractId = *returnRequest.ContractID
	}

	contract, err := p.blockchain.GetMContractByIDOnChain(int64(contractId))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	if typeOfPayment == "contract" {
		amount = float64(contract.Deposit)
	}

	fmt.Printf("Contract ID: %d\n", contractId)

	userId := int32(1)

	userBankInfo, err := p.repo.GetBankInfoByUserID(context.Background(), userId)

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    responses.StatusInternalError,
			Data:       nil,
		}
	}

	bankInfo, err := p.repo.GetBankByID(context.Background(), userBankInfo.BankID)

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    responses.StatusInternalError,
			Data:       nil,
		}
	}


	var response responses.GetPaymentInfoRes
	prefix := fmt.Sprintf("SR%d", id)
	response.TranferContent = common.GenerateCode(prefix)
	response.BankName = bankInfo.BankName
	response.AccountName = userBankInfo.AccountName
	response.AccountNumber = userBankInfo.AccountNumber
	encodedName := url.QueryEscape(userBankInfo.AccountName)
	response.QrUrl = "https://img.vietqr.io/image/" + bankInfo.BankCode + "-" + userBankInfo.AccountNumber + "-compact2.png?amount=" + strconv.Itoa(int(amount)) + "&addInfo=" + response.TranferContent + "&accountName=" + encodedName
	response.Amount = amount

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       response,
	}
}

// Confirm implements PaymentService.
func (p *PaymentServiceImpl) Confirm(id int, userID int) *responses.ResponseData {
	payment, err := p.repo.GetPaymentByID(context.Background(), int32(id))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	user, _ := p.repo.GetUserByID(context.Background(), int32(userID))
	if payment.ContractID != nil {
		contract, err := p.blockchain.GetMContractByIDOnChain(int64(*payment.ContractID))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
		if contract.PreRentalStatus != 1 {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    "Trạng thái hợp đồng không hợp lệ",
				Data:       false,
			}
		}
		_, err = p.blockchain.PayDepositOnChain(*user.PrivateKeyHex, int64(*payment.ContractID))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	} else if payment.BillID != nil {
		bill, err := p.repo.GetBillByID(context.Background(), *payment.BillID)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    responses.StatusInternalError,
				Data:       nil,
			}
		}
		contract, err := p.blockchain.GetMContractByIDOnChain(int64(bill.ContractID))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
		if contract.RentalProcessStatus != 2 {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    "Trạng thái hợp đồng không hợp lệ",
				Data:       false,
			}
		}
		_, err = p.blockchain.PayBillOnChain(*user.PrivateKeyHex, int64(bill.ContractID))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	} else if payment.ReturnRequestID != nil {
		returnRequest, err := p.repo.GetReturnRequestByID(context.Background(), *payment.ReturnRequestID)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    responses.StatusInternalError,
				Data:       nil,
			}
		}
		contract, err := p.blockchain.GetMContractByIDOnChain(int64(*returnRequest.ContractID))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
		if contract.PostRentalStatus != 1 {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    "Trạng thái hợp đồng không hợp lệ",
				Data:       false,
			}
		}
		_, err = p.blockchain.ApproveReturnRequestOnChain(*user.PrivateKeyHex, int64(*returnRequest.ContractID))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	}

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

	paymentId, createErr := p.repo.CreatePayment(context.Background(), params)

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
		Data:       paymentId,
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
