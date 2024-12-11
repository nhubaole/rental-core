package services

import (
	"context"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/responses"
)

type BillingServiceImpl struct {
	query      *dataaccess.Queries
	blockchain BlockchainService
}

func NewBillingServiceImpl(blockchain BlockchainService) BillingService {
	return &BillingServiceImpl{
		query:      dataaccess.New(global.Db),
		blockchain: blockchain,
	}
}

func (service *BillingServiceImpl) CreateBill(userID int32, body dataaccess.CreateBillParams) *responses.ResponseData {
	contract, err := service.blockchain.GetMContractByIDOnChain(int64(body.ContractID))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	if !(contract.RentalProcessStatus == 0 || contract.RentalProcessStatus == 1) {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Trạng thái hợp đồng không hợp lệ",
			Data:       false,
		}
	}

	err = service.query.CreateBill(context.Background(), body)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	user, _ := service.query.GetUserByID(context.Background(), int32(userID))
	_, err = service.blockchain.CreateBillOnChain(*user.PrivateKeyHex, int64(body.ContractID))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}

func (service *BillingServiceImpl) GetBillByMonth(userID int32, month int32, year int32) *responses.ResponseData {
	param := dataaccess.GetBillByMonthParams{
		Year:  year,
		Month: month,
	}

	bills, err := service.query.GetBillByMonth(context.Background(), param)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Can't find your bill",
			Data:       false,
		}
	}

	var filteredBills []responses.GetBillByMonthRes

	for _, bill := range bills {
		contract, err := service.blockchain.GetMContractByIDOnChain(int64(bill.ContractID))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error retrieving contract for bill",
				Data:       false,
			}
		}
		var detailBill responses.GetBillByMonthRes

		common.MapStruct(bill, &detailBill)
		detailBill.TenantID = contract.Tenant
		detailBill.LandlordID = contract.Landlord
		detailBill.InternetCost = float64(contract.InternetCost)
		detailBill.ParkingFee = float64(contract.ParkingFee)
		detailBill.RoomID = int32(contract.RoomID)
		detailBill.CreatedAt = bill.CreatedAt
		detailBill.UpdatedAt = bill.UpdatedAt

		if int32(contract.Landlord) == userID || int32(contract.Tenant) == userID {
			filteredBills = append(filteredBills, detailBill)
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       filteredBills,
	}
}

// GetBillByID implements BillingService.
func (service *BillingServiceImpl) GetBillByID(id int32) *responses.ResponseData {
	bill, err := service.query.GetBillByID(context.Background(), id)
	if err != nil {
		if (bill == dataaccess.GetBillByIDRow{}) {
			return &responses.ResponseData{
				StatusCode: http.StatusNoContent,
				Message:    "No bill found",
				Data:       false,
			}
		}
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       bill,
	}
}

// GetBillMetric implements BillingService.
func (service *BillingServiceImpl) GetBillMetrics(req dataaccess.GetAllMetric4BillByRoomIDParams) *responses.ResponseData {
	metric, err := service.query.GetAllMetric4BillByRoomID(context.Background(), req)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	contracts, err := service.query.ListContractByRoomId(context.Background(), &req.RoomID)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	var matchedContract responses.MContractOnChainRes
	for _, contract := range contracts {
		onChainContract, err := service.blockchain.GetMContractByIDOnChain(int64(contract))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       false,
			}
		}

		if onChainContract.PreRentalStatus == 2 {
			matchedContract = *onChainContract
			break
		}
	}

	combinedResponse := responses.GetAllMetric4BillByRoomID{
		RoomID:          metric.RoomID,
		PrevMonth:       metric.PrevMonth,
		CurrMonth:       metric.CurrMonth,
		PrevWater:       metric.PrevWater,
		CurrWater:       metric.CurrWater,
		PrevElectricity: metric.PrevElectricity,
		CurrElectricity: metric.CurrElectricity,
		Year:            metric.Year,
		ContractID:      int32(matchedContract.ID),
		ActualPrice:     matchedContract.ActualPrice,
		WaterCost:       matchedContract.WaterCost,
		ElectricityCost: matchedContract.ElectricityCost,
		InternetCost:    matchedContract.InternetCost,
		ParkingFee:      matchedContract.ParkingFee,
	}

	// Return the combined response
	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       combinedResponse,
	}
}

// GetBillByStatus implements BillingService.
func (service *BillingServiceImpl) GetBillByStatus(statusID int32) *responses.ResponseData {
	bills, err := service.query.GetBillByStatus(context.Background(), &statusID)

	if err != nil {
		if len(bills) == 0 {
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
		Data:       bills,
	}
}

// GetBillOfRentedRoomByOwnerID implements BillingService.
func (service *BillingServiceImpl) GetBillOfRentedRoomByOwnerID(currentUser int) *responses.ResponseData {
	bills, err := service.query.GetBillOfRentedRoomByOwnerID(context.Background(), int32(currentUser))
	if err != nil {
		if len(bills) == 0 {
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
		Data:       bills,
	}
}
