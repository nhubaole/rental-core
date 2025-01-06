package services

import (
	"context"
	"fmt"
	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
	"strings"
)

type BillingServiceImpl struct {
	query               *dataaccess.Queries
	blockchain          BlockchainService
	notificationService NotificationService
}

func NewBillingServiceImpl(blockchain BlockchainService, notification NotificationService) BillingService {
	return &BillingServiceImpl{
		query:               dataaccess.New(global.Db),
		blockchain:          blockchain,
		notificationService: notification,
	}
}

func (service *BillingServiceImpl) CreateBill(userID int32, req requests.CreateBill) *responses.ResponseData {
	param := dataaccess.GetAllMetric4BillByRoomIDParams{
		RoomID: req.RoomID,
		Month:  req.Month,
		Year:   req.Year,
	}
	metric, err := service.query.GetAllMetric4BillByRoomID(context.Background(), param)
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

	waterUsage := *metric.CurrWater - metric.PrevWater.(float64)
	electricityUsage := *metric.CurrElectricity - metric.PrevElectricity.(float64)

	totalWater := waterUsage * float64(matchedContract.WaterCost)
	totalElectric := electricityUsage * float64(matchedContract.ElectricityCost)

	fmt.Println(float64(matchedContract.ActualPrice))
	fmt.Println(float64(matchedContract.Deposit))
	fmt.Println(float64(matchedContract.InternetCost))
	fmt.Println(float64(matchedContract.ParkingFee))

	totalAmount := float64(matchedContract.Deposit) +
		totalWater +
		totalElectric +
		float64(matchedContract.InternetCost) +
		float64(matchedContract.ParkingFee)

	body := dataaccess.CreateBillParams{
		Code:                 common.GenerateCode("HD"),
		ContractID:           int32(matchedContract.ID),
		OldWaterIndex:        common.Float64ToInt32Ptr(metric.PrevWater.(float64)),
		OldElectricityIndex:  common.Float64ToInt32Ptr(metric.PrevElectricity.(float64)),
		NewWaterIndex:        common.Float64ToInt32Ptr(*metric.CurrWater),
		NewElectricityIndex:  common.Float64ToInt32Ptr(*metric.CurrElectricity),
		TotalWaterCost:       &totalWater,
		TotalElectricityCost: &totalElectric,
		AdditionFee:          req.AdditionFee,
		AdditionNote:         req.AdditionNote,
		TotalAmount:          totalAmount,
		Month:                req.Month,
		Year:                 req.Year,
	}

	// if !(matchedContract.RentalProcessStatus == 0 || matchedContract.RentalProcessStatus == 1) {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    "Trạng thái hợp đồng không hợp lệ",
	// 		Data:       false,
	// 	}
	// }

	billId, err := service.query.CreateBill(context.Background(), body)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	// user, _ := service.query.GetUserByID(context.Background(), int32(userID))
	// _, err = service.blockchain.CreateBillOnChain(*user.PrivateKeyHex, int64(body.ContractID))
	// if err != nil {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    err.Error(),
	// 		Data:       false,
	// 	}
	// }

	id := int(billId)
	service.notificationService.SendNotification(int(matchedContract.Tenant), "Bạn có hoá đơn thu tiền mới. Vui lòng kiểm tra và thanh toán đúng hạn", &id, "bill")

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

	result := make([]map[string]interface{}, 0)

	addressMap := make(map[string][]map[string]interface{})
	for _, bill := range bills {
		if bill.ContractID == nil {
			continue
		}

		contract, err := service.blockchain.GetMContractByIDOnChain(int64(*bill.ContractID))
		if err != nil {
			continue
		}

		if contract.PreRentalStatus != 2 {
			continue
		}

		tenant, err := service.query.GetUserByID(context.Background(), int32(contract.Tenant))
		if err != nil {
			continue
		}

		tenantName := tenant.FullName
		tenantAvt := ""
		if tenant.AvatarUrl != nil {
			tenantAvt = *tenant.AvatarUrl
		}

		billData := map[string]interface{}{
			"id":           bill.BillID,
			"avatar":       tenantAvt,
			"status":       bill.BillStatus,
			"room_number":  bill.RoomNumber,
			"room_id":      bill.RoomID,
			"tenant_name":  tenantName,
			"payment_id":   bill.PaymentID,
			"total_amount": bill.TotalAmount,
		}

		address := strings.Join(bill.Address, ", ")
		addressMap[address] = append(addressMap[address], billData)
	}

	for address, listBill := range addressMap {
		result = append(result, map[string]interface{}{
			"address":   address,
			"list_bill": listBill,
		})
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

	responseData := map[string]interface{}{
		"id":                    bill.ID,
		"code":                  bill.Code,
		"created_at":            bill.CreatedAt,
		"paid_at":               nil,
		"payment_id":            bill.PaymentID,
		"status":                bill.Status,
		"room_price":            contract.Deposit,
		"old_water_index":       bill.OldWaterIndex,
		"old_electricity_index": bill.OldElectricityIndex,
		"new_water_index":       bill.NewWaterIndex,
		"new_electricity_index": bill.NewElectricityIndex,
		"water_cost":            contract.WaterCost,
		"electricity_cost":      contract.ElectricityCost,
		"internet_cost":         contract.InternetCost,
		"parking_fee":           contract.ParkingFee,
		"addition_fee":          bill.AdditionFee,
		"addition_note":         bill.AdditionNote,
		"total_amount":          bill.TotalAmount,
		"info": map[string]interface{}{
			"tenant_name":  tenantInfo.FullName,
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

	waterUsage := *metric.CurrWater - metric.PrevWater.(float64)
	electricityUsage := *metric.CurrElectricity - metric.PrevElectricity.(float64)
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
		"actual_price":          matchedContract.ActualPrice,
		"old_water_index":       metric.PrevWater,
		"old_electricity_index": metric.PrevElectricity,
		"new_water_index":       metric.CurrWater,
		"new_electricity_index": metric.CurrElectricity,
		"water_cost":            matchedContract.WaterCost,
		"electricity_cost":      matchedContract.ElectricityCost,
		"internet_cost":         matchedContract.InternetCost,
		"parking_fee":           matchedContract.ParkingFee,
		"total_amount":          totalAmount,
		"info": map[string]interface{}{
			"tenant_name":  tenant.FullName,
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
