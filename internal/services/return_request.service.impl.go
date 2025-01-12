package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type ReturnRequestServiceImpl struct {
	repo                *dataaccess.Queries
	blockchain          BlockchainService
	notificationService NotificationService
}

func NewReturnRequestServiceImpl(blockchain BlockchainService, notification NotificationService) ReturnRequestService {
	return &ReturnRequestServiceImpl{
		repo:                dataaccess.New(global.Db),
		blockchain:          blockchain,
		notificationService: notification,
	}
}

// GetByLandlordID implements ReturnRequestService.
func (r *ReturnRequestServiceImpl) GetByLandlordID(userID int) *responses.ResponseData {
	requests, err := r.repo.GetReturnRequestByLandlordID(context.Background(), int32(userID))

	if len(requests) == 0 {
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
			Data:       false,
		}
	}
	var detailedRequests []responses.GetReturnRequestByLandlordIDRes
	for _, request := range requests {
		sender, err := r.repo.GetUserByID(context.Background(), *request.CreatedUser)
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}
		var user responses.GetUserByIDRes
		common.MapStruct(sender, &user)

		detailedRequest := responses.GetReturnRequestByLandlordIDRes{
			ID:                 request.ID,
			ContractID:         request.ContractID,
			RoomID:             *request.RoomID,
			Reason:             request.Reason,
			ReturnDate:         request.ReturnDate,
			Status:             request.Status,
			DeductAmount:       request.DeductAmount,
			TotalReturnDeposit: request.TotalReturnDeposit,
			CreatedUser:        user,
			CreatedAt:          request.CreatedAt,
			UpdatedAt:          request.UpdatedAt,
		}

		detailedRequests = append(detailedRequests, detailedRequest)
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       detailedRequests,
	}
}

// Create implements ReturnRequestService.
func (r *ReturnRequestServiceImpl) Create(req requests.CreateReturnRequestParams, userID int) *responses.ResponseData {
	id := int32(userID)

	contract, err := r.blockchain.GetMContractByIDOnChain(int64(*req.ContractID))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	// if contract.PostRentalStatus != 0 {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    "Trạng thái hợp đồng không hợp lệ",
	// 		Data:       false,
	// 	}
	// }

	// user, _ := r.repo.GetUserByID(context.Background(), int32(userID))
	// _, err = r.blockchain.CreateReturnRequestOnChain(*user.PrivateKeyHex, int64(*req.ContractID))
	// if err != nil {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    err.Error(),
	// 		Data:       false,
	// 	}
	// }

	deductAmount := calculateDeductAmount(req.ReturnDate, int(contract.Deposit))

	params := dataaccess.CreateReturnRequestParams{
		ContractID:         req.ContractID,
		Reason:             req.Reason,
		ReturnDate:         req.ReturnDate,
		TotalReturnDeposit: common.IntToFloat64Ptr(int(contract.Deposit)),
		DeductAmount:       &deductAmount,
		CreatedUser:        &id,
	}

	returnRequestId, err := r.repo.CreateReturnRequest(context.Background(), params)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	reqId := int(returnRequestId)
	r.notificationService.SendNotification(int(contract.Landlord), "Bạn có một yêu cầu trả phòng mới", &reqId, "return_request")

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}

// GetByID implements ReturnRequestService.
func (r *ReturnRequestServiceImpl) GetByID(id int) *responses.ResponseData {
	returnRequest, err := r.repo.GetReturnRequestByID(context.Background(), int32(id))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	room, err := r.repo.GetRoomByID(context.Background(), *returnRequest.RoomID)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	sender, err := r.repo.GetUserByID(context.Background(), *returnRequest.CreatedUser)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       "nothing",
		}
	}

	contract, err := r.blockchain.GetMContractByIDOnChain(int64(*returnRequest.ContractID))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       "nothing",
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data: responses.GetReturnRequestByIDRes{
			ID:                 returnRequest.ID,
			Reason:             returnRequest.Reason,
			Room:               room,
			ContractID:         returnRequest.ContractID,
			ContractCode:       contract.Code,
			CreatedUser:        sender,
			ReturnDate:         returnRequest.ReturnDate,
			Status:             returnRequest.Status,
			DeductAmount:       returnRequest.DeductAmount,
			TotalReturnDeposit: returnRequest.TotalReturnDeposit,
			CreatedAt:          returnRequest.CreatedAt,
			UpdatedAt:          returnRequest.UpdatedAt,
		},
	}
}

// Aprrove implements ReturnRequestService.
func (r *ReturnRequestServiceImpl) Aprrove(id int, userID int) *responses.ResponseData {
	// Get return request
	returnRequest, err := r.repo.GetReturnRequestByID(context.Background(), int32(id))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusNotFound,
			Message:    "Yêu câu trả phòng không tồn tại",
			Data:       false,
		}
	}

	if *returnRequest.TotalReturnDeposit - *returnRequest.DeductAmount == 0 {
		user, _ := r.repo.GetUserByID(context.Background(), int32(userID))

		status := int32(2)
		params := dataaccess.ApproveReturnRequestParams{
			Status: &status,
			ID:     returnRequest.ID,
		}

		updateErr := r.repo.ApproveReturnRequest(context.Background(), params)
		if updateErr != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    updateErr.Error(),
				Data:       false,
			}
		}

		// get room by contract
		room, roomErr := r.repo.GetRoomByContractID(context.Background(), *returnRequest.ContractID)
		if roomErr != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    roomErr.Error(),
				Data:       false,
			}
		}

		if room.Owner != int32(userID) {
			return &responses.ResponseData{
				StatusCode: http.StatusForbidden,
				Message:    "Bạn không có quyền thực hiện thao tác này",
				Data:       false,
			}
		}

		// set room available
		updateRoomParam := dataaccess.UpdateRoomStatusParams{
			ID:     room.RoomID,
			IsRent: false,
		}
		_, updateRoomErr := r.repo.UpdateRoomStatus(context.Background(), updateRoomParam)
		if updateRoomErr != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    updateRoomErr.Error(),
				Data:       false,
			}
		}

		// inactive contract
		_, err = r.blockchain.DeclineMContractOnChain(*user.PrivateKeyHex, int64(*returnRequest.ContractID))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}
	} else {
		status := int32(1)
		params := dataaccess.ApproveReturnRequestParams{
			Status: &status,
			ID:     returnRequest.ID,
		}

		updateErr := r.repo.ApproveReturnRequest(context.Background(), params)
		if updateErr != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    updateErr.Error(),
				Data:       false,
			}
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}

func calculateDeductAmount(returnDate pgtype.Timestamp, deposit int) float64 {
	// Kiểm tra nếu ReturnDate có giá trị
	if !returnDate.Valid {
		return 0 // Nếu không hợp lệ, trả về 0
	}

	// Lấy giá trị kiểu time.Time từ pgtype.Timestamp
	returnTime := returnDate.Time
	currentDate := time.Now()

	// Tính số ngày giữa ReturnDate và ngày hiện tại
	daysDiff := int(returnTime.Sub(currentDate).Hours() / 24)

	// Nếu cách ngày hiện tại < 30, DeductAmount = deposit, ngược lại bằng 0
	if daysDiff < 30 {
		return float64(deposit)
	}
	return 0
}
