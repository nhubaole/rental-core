package services

import (
	"context"
	"encoding/json"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/responses"
	"strings"
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
		Owner: userID,
	}

	bills, err := service.query.GetBillByMonth(context.Background(), param)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       false,
		}
	}

	var result []responses.GetBillByMonthRes

	for _, v := range bills {
		var bill responses.GetBillByMonthRes
		var listBillJson []map[string]interface{}

		json.Unmarshal([]byte(v.ListBill), &listBillJson)
		bill.Address = strings.Join(v.Address, ", ")
		bill.ListBill = listBillJson

		result = append(result, bill)

	}


	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
	}
}

// GetBillByID implements BillingService.
func (service *BillingServiceImpl) GetBillByID(id int32) *responses.ResponseData {
	bill, err := service.query.GetBillByID(context.Background(), id)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	contract, err := service.blockchain.GetMContractByIDOnChain(int64(bill.ContractID))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error retrieving contract for bill",
			Data:       false,
		}
	}

	roomInfo, err := service.query.GetRoomByID(context.Background(), int32(contract.RoomID))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to fetch room details",
			Data:       false,
		}
	}

	tenantInfo, err := service.query.GetUserByID(context.Background(), int32(contract.Tenant))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Unable to fetch tenant info",
			Data:       false,
		}
	}

	totalAmount := contract.ActualPrice +
		(int64(*bill.NewWaterIndex)-int64(*bill.OldWaterIndex))*contract.WaterCost +
		(int64(*bill.NewElectricityIndex)-int64(*bill.OldElectricityIndex))*contract.ElectricityCost +
		contract.InternetCost +
		contract.ParkingFee +
		int64(*bill.AdditionFee)

	responseData := map[string]interface{}{
		"id":                  bill.ID,
		"code":                bill.Code,
		"created_at":          bill.CreatedAt,
		"paid_at":             nil,
		"status":              bill.Status,
		"room_price":          contract.ActualPrice,
		"old_water_index":     bill.OldWaterIndex,
		"old_electricity_index": bill.OldElectricityIndex,
		"new_water_index":     bill.NewWaterIndex,
		"new_electricity_index": bill.NewElectricityIndex,
		"water_cost":          contract.WaterCost,
		"electricity_cost":    contract.ElectricityCost,
		"internet_cost":       contract.InternetCost,
		"parking_fee":         contract.ParkingFee,
		"addition_fee":        bill.AdditionFee,
		"addition_note":       bill.AdditionNote,
		"total_amount":        totalAmount,
		"info": map[string]interface{}{
			"tenant_name": tenantInfo.FullName,
			"phone_number": tenantInfo.PhoneNumber,
			"room_number":  roomInfo.RoomNumber,
			"address":      strings.Join(roomInfo.Address, ", "),
			"month":        bill.Month,
			"year":         bill.Year,
		},
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       responseData,
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

	roomInfo, err := service.query.GetRoomByID(context.Background(), req.RoomID)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	waterUsage := metric.CurrWater - metric.PrevWater.(float64)
	electricityUsage := metric.CurrElectricity - metric.PrevElectricity.(float64)
	totalAmount := float64(matchedContract.ActualPrice) +
		waterUsage*float64(matchedContract.WaterCost) +
		electricityUsage*float64(matchedContract.ElectricityCost) +
		float64(matchedContract.InternetCost) +
		float64(matchedContract.ParkingFee)
	
	tenant, err := service.query.GetUserByID(context.Background(), int32(matchedContract.Tenant))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}

	responseData := map[string]interface{}{
		"actual_price":         matchedContract.ActualPrice,
		"old_water_index":      metric.PrevWater,
		"old_electricity_index": metric.PrevElectricity,
		"new_water_index":      metric.CurrWater,
		"new_electricity_index": metric.CurrElectricity,
		"water_cost":           matchedContract.WaterCost,
		"electricity_cost":     matchedContract.ElectricityCost,
		"internet_cost":        matchedContract.InternetCost,
		"parking_fee":          matchedContract.ParkingFee,
		"total_amount":         totalAmount,
		"info": map[string]interface{}{
			"tenant_name": tenant.FullName,
			"phone_number": tenant.PhoneNumber,
			"room_number":  roomInfo.RoomNumber,
			"address":      strings.Join(roomInfo.Address, ", "),
			"month":        metric.CurrMonth,
			"year":         metric.Year,
		},
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       responseData,
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
