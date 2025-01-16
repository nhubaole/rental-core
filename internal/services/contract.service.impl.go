package services

import (
	"context"
	"fmt"
	"strings"

	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/blockchain/gen"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type ContractServiceImpl struct {
	repo                *dataaccess.Queries
	blockchain          BlockchainService
	notificationService NotificationService
}

func NewContractServiceImpl(blockchain BlockchainService, notification NotificationService) ContractService {
	return &ContractServiceImpl{
		repo:                dataaccess.New(global.Db),
		blockchain:          blockchain,
		notificationService: notification,
	}
}

// GetContractByUser implements ContractService.
func (c *ContractServiceImpl) GetContractByUser(userID int) *responses.ResponseData {
	// user, err := c.repo.GetUserByID(context.Background(), int32(userID))
	// contracts, err := c.blockchain.GetMContractByIDOnChain(1)
	// if err != nil {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    err.Error(),
	// 		Data:       false,
	// 	}
	// }

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       nil,
	}
}

// CreateTemplate implements ContractService.
func (c *ContractServiceImpl) CreateTemplate(req dataaccess.CreateContractTemplateParams) *responses.ResponseData {
	template, _ := c.repo.GetContractTemplateByAddress(context.Background(), req.Address)
	if template.ID != 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Mẫu hợp đồng cho địa chỉ này đã tồn tại",
			Data:       false,
		}
	}
	createErr := c.repo.CreateContractTemplate(context.Background(), req)
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

// GetTemplateByAddress implements ContractService.
func (c *ContractServiceImpl) GetTemplateByAddress(address requests.GetTemplateByAddressRequest) *responses.ResponseData {
	template, err := c.repo.GetContractTemplateByAddress(context.Background(), address.Address)
	if err != nil {
		if template.ID == 0 {
			return &responses.ResponseData{
				StatusCode: http.StatusOK,
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
		Data:       template,
	}
}
func (c *ContractServiceImpl) CreateContract(req requests.CreateContractRequest, userID int) *responses.ResponseData {
	template, _ := c.repo.GetContractTemplateByAddress(context.Background(), req.Address)
	// if err != nil {
	// 	return &responses.ResponseData{
	// 		StatusCode: http.StatusInternalServerError,
	// 		Message:    "Không tìm thấy mẫu hợp đồng cho địa chỉ này",
	// 		Data:       false,
	// 	}
	// }
	signOfA, encryptedErrA := common.EncryptBase64AES(req.SignatureA, global.Config.JWT.AESKey)
	if encryptedErrA != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Lỗi ký hợp đồng",
			Data:       false,
		}
	}
	parkingFee := common.IfNullInt64(req.ParkingFee, common.Float64PtrToInt64Ptr(&template.ParkingFee))
	generalResponsibility := common.IfNullStr(req.GeneralResponsibility, &template.GeneralResponsibility)

	var params = &dataaccess.CreateContractParams{
		RoomID:     &req.RoomID,
		SignatureA: &signOfA,
	}

	contractId, err := c.repo.CreateContract(context.Background(), *params)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	contract := &requests.CreateMContractOnChainReq{
		ContractId:            int64(contractId),                                                                                  // ID duy nhất của hợp đồng
		ContractCode:          common.GenerateCode("HD"),                                                                          // Mã hợp đồng
		LandlordId:            int64(req.PartyA),                                                                                  // ID của chủ nhà
		TenantId:              int64(req.PartyB),                                                                                  // ID của người thuê
		RoomId:                int64(req.RoomID),                                                                                  // ID của phòng
		ActualPrice:           int64(req.ActualPrice),                                                                             // Giá thực tế của hợp đồng
		Deposit:               int64(req.Deposit),                                                                                 // Tiền đặt cọc
		BeginDate:             int64(req.BeginDate.Time.Unix()),                                                                   // Thời gian bắt đầu hợp đồng (Unix timestamp)
		EndDate:               int64(req.EndDate.Time.Unix()),                                                                     // Thời gian kết thúc hợp đồng (Unix timestamp)
		PaymentMethod:         *req.PaymentMethod,                                                                                 // Phương thức thanh toán
		ElectricityMethod:     common.IfNullStr(&req.ElectricityMethod, &template.ElectricityMethod),                              // Phương thức tính điện
		ElectricityCost:       common.IfNullInt64((&req.ElectricityCost), common.Float64PtrToInt64Ptr(&template.ElectricityCost)), // Giá điện
		WaterMethod:           common.IfNullStr(&req.WaterMethod, &template.WaterMethod),                                          // Phương thức tính nước
		WaterCost:             common.IfNullInt64(&req.WaterCost, common.Float64PtrToInt64Ptr(&template.WaterCost)),               // Giá nước
		InternetCost:          common.IfNullInt64(&req.InternetCost, common.Float64PtrToInt64Ptr(&template.InternetCost)),         // Giá internet
		ParkingFee:            parkingFee,                                                                                         // Phí gửi xe
		ResponsibilityA:       common.IfNullStr(&req.ResponsibilityA, &template.ResponsibilityA),                                  // Trách nhiệm bên A
		ResponsibilityB:       common.IfNullStr(&req.ResponsibilityB, &template.ResponsibilityB),                                  // Trách nhiệm bên B
		GeneralResponsibility: generalResponsibility,                                                                              // Trách nhiệm chung
		SignatureA:            "",                                                                                                 // Chữ ký của bên A
		SignedTimeA:           req.SignedTimeA.Time.Unix(),                                                                        // Thời gian ký của bên A
		SignatureB:            "",                                                                                                 // Chữ ký của bên B
		SignedTimeB:           int64(0),                                                                                           // Thời gian ký của bên B
		ContractTemplateId:    int64(template.ID),                                                                                 // ID mẫu hợp đồng
	}

	user, _ := c.repo.GetUserByID(context.Background(), int32(userID))
	if _, err := c.blockchain.CreateMContractOnChain(*user.PrivateKeyHex, *contract); err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	id := int(contractId)

	c.notificationService.SendNotification(int(req.PartyB), "Hợp đồng thuê trọ của bạn đã được chủ nhà tạo", &id, "contract")

	trackingParam := dataaccess.CreateProcessTrackingParams{
		Actor:     int32(userID),
		Action:    "Chủ nhà đã tạo hợp đồng thuê trọ",
		RequestID: req.RequestID,
	}
	_, er := c.repo.CreateProcessTracking(context.Background(), trackingParam)
	if er != nil {
	}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       contractId,
	}
}

// GetContractByID implements ContractService.
func (c *ContractServiceImpl) GetContractByID(id int) *responses.ResponseData {
	contract, err := c.blockchain.GetMContractByIDOnChain(int64(id))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	landlord, err := c.repo.GetUserByID(context.Background(), int32(contract.Landlord))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	tenant, err := c.repo.GetUserByID(context.Background(), int32(contract.Tenant))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	roomDetails, err := c.repo.GetRoomByID(context.Background(), int32(contract.RoomID))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	contractDB, err := c.repo.GetContractById(context.Background(), int32(id))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	signatureA := ""
	if contractDB.SignatureA != nil {
		signatureA, _ = common.DecryptBase64AES(*contractDB.SignatureA, global.Config.JWT.AESKey)
	}

	signatureB := ""
	if contractDB.SignatureB != nil {
		signatureB, _ = common.DecryptBase64AES(*contractDB.SignatureB, global.Config.JWT.AESKey)
	}

	var status int
	switch contract.PreRentalStatus {
	case 0, 1:
		status = 0
	case 2:
		status = 1
	case 3:
		status = 2
	default:
		status = -1
	}

	responseData := map[string]interface{}{
		"status":          status,
		"date_created":    contract.CreatedAt,
		"address_created": "Thành phố Hồ Chí Minh",
		"party_a": map[string]string{
			"name":             landlord.FullName,
			"dob":              landlord.Dob.Time.String(),
			"registered_place": "97 đường số 11, phường Trường Thọ, quận Thủ Đức, TP HCM",
			"cccd":             "079304032010",
			"issue_date":       "06/09/2022",
			"issue_by":         "Cục cảnh sát quản lý hành chính và trật tự xã hội",
			"phone":            landlord.PhoneNumber,
		},
		"party_b": map[string]string{
			"name":             tenant.FullName,
			"dob":              tenant.Dob.Time.String(),
			"registered_place": "26 ấp Thị, xã Hương Mỹ, huyện Mỏ Cày Nam, tỉnh Bến Tre",
			"cccd":             "083405012453",
			"issue_date":       "12/10/2021",
			"issue_by":         "Cục cảnh sát quản lý hành chính và trật tự xã hội",
			"phone":            tenant.PhoneNumber,
		},
		"room_address":     roomDetails.Address,
		"room_price":       contract.Deposit,
		"method":           contract.PaymentMethod,
		"electric_cost":    contract.ElectricityCost,
		"water_cost":       contract.WaterCost,
		"deposit":          contract.Deposit,
		"start_date":       contract.BeginDate,
		"end_date":         contract.EndDate,
		"responsibility_a": contract.ResponsibilityA,
		"responsibility_b": contract.ResponsibilityB,
		"general":          contract.GeneralResponsibility,
		"signature_a":      signatureA,
		"signature_b":      signatureB,
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       responseData,
	}
}

// ListContractByStatus implements ContractService.
func (c *ContractServiceImpl) ListContractByStatus(statusID int, userId int, isLandlord bool) *responses.ResponseData {
	contractIds, _ := c.repo.ListContractIds(context.Background())
	var contracts []gen.ContractManagementMContract
	var err error

	switch statusID {
	case 0:
		contracts, err = c.blockchain.GetListMContractByStatus(contractIds, 0, int64(userId), isLandlord)
		contracts1, _ := c.blockchain.GetListMContractByStatus(contractIds, 1, int64(userId), isLandlord)

		contracts = append(contracts, contracts1...)
	case 1:
		contracts, err = c.blockchain.GetListMContractByStatus(contractIds, 2, int64(userId), isLandlord)
	case 2:
		contracts, err = c.blockchain.GetListMContractByStatus(contractIds, 3, int64(userId), isLandlord)
	default:
		return &responses.ResponseData{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid statusID",
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

	if len(contracts) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusOK,
			Message:    responses.StatusNoData,
			Data:       nil,
		}
	}

	var result []map[string]interface{}

	for _, contract := range contracts {
		roomDetails, err := c.repo.GetRoomByID(context.Background(), int32(contract.RoomId.Int64()))
		// if err != nil {
		// 	return &responses.ResponseData{
		// 		StatusCode: http.StatusInternalServerError,
		// 		Message:    err.Error(),
		// 		Data:       nil,
		// 	}
		// }

		landlord, err := c.repo.GetUserByID(context.Background(), int32(contract.Landlord.Int64()))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}

		tenant, err := c.repo.GetUserByID(context.Background(), int32(contract.Tenant.Int64()))
		if err != nil {
			return &responses.ResponseData{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
				Data:       nil,
			}
		}

		result = append(result, map[string]interface{}{
			"id":               contract.Id,
			"code":             contract.Code,
			"room_address":     strings.Join(roomDetails.Address, ", "),
			"room_number":      roomDetails.RoomNumber,
			"landlord_name":    landlord.FullName,
			"tenant_name":      tenant.FullName,
			"signature_time_a": contract.SignedTimeA,
			"signature_time_b": contract.SignedTimeB,
			"created_at":       contract.CreatedAt,
			"expired_at":       contract.CreatedAt.Int64() + 7*24*60*60, //1 tuan sau khi tao
		})
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       result,
	}
}

// SignContract implements ContractService.
func (c *ContractServiceImpl) SignContract(req requests.SignContractParams, userID int) *responses.ResponseData {
	contract, err := c.blockchain.GetMContractByIDOnChain(int64(req.Id))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	signB, encryptErr := common.EncryptBase64AES(req.SignatureB, global.Config.JWT.AESKey)
	req.SignatureB = signB
	if encryptErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Lỗi ký hợp đồng",
			Data:       false,
		}
	}

	err = c.repo.UpdateSignatureB(context.Background(), dataaccess.UpdateSignatureBParams{ID: req.Id, SignatureB: &req.SignatureB})

	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	if contract.PreRentalStatus != 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Trạng thái hợp đồng không hợp lệ",
			Data:       false,
		}
	}

	params := &requests.SignMContractOnChainReq{
		ContractId: int64(req.Id),
		SignatureB: "",
	}
	user, _ := c.repo.GetUserByID(context.Background(), int32(userID))
	_, err = c.blockchain.SignMContractOnChain(*user.PrivateKeyHex, *params)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}

// DeclineContract implements ContractService.
func (c *ContractServiceImpl) DeclineContract(id int, userID int) *responses.ResponseData {
	user, _ := c.repo.GetUserByID(context.Background(), int32(userID))

	contract, _ := c.blockchain.GetMContractByIDOnChain(int64(id))

	fmt.Print(contract)
	if userID != int(contract.Tenant) {
		return &responses.ResponseData{
			StatusCode: http.StatusForbidden,
			Message:    "Bạn không có quyền thực hiện thao tác này",
			Data:       false,
		}
	}

	_, err := c.blockchain.DeclineMContractOnChain(*user.PrivateKeyHex, int64(id))
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}

	c.notificationService.SendNotification(int(contract.Landlord), "Người thuê đã từ chối ký hợp đồng thuê trọ", &id, "contract")

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}
