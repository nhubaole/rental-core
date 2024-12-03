package services

import (
	"context"
	"time"

	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"

	"github.com/jackc/pgx/v5/pgtype"
)

type ContractServiceImpl struct {
	repo       *dataaccess.Queries
	blockchain BlockchainService
}

func NewContractServiceImpl(blockchain BlockchainService) ContractService {
	return &ContractServiceImpl{
		repo:       dataaccess.New(global.Db),
		blockchain: blockchain,
	}
}

// GetContractByUser implements ContractService.
func (c *ContractServiceImpl) GetContractByUser(userID int) *responses.ResponseData {
	// user, err := c.repo.GetUserByID(context.Background(), int32(userID))
	contracts, err := c.blockchain.GetMContractByIDOnChain(3)
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
		Data:       contracts,
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
		Data:       template,
	}
}
func (c *ContractServiceImpl) CreateContract(req requests.CreateContractRequest, userID int) *responses.ResponseData {
	template, err := c.repo.GetContractTemplateByAddress(context.Background(), req.Address)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Không tìm thấy mẫu hợp đồng cho địa chỉ này",
			Data:       false,
		}
	}
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

	contract := &requests.CreateMContractOnChainReq{
		ContractId:            int64(3),                                                                                           // ID duy nhất của hợp đồng
		ContractCode:          req.Code,                                                                                           // Mã hợp đồng
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
		SignatureA:            signOfA,                                                                                            // Chữ ký của bên A
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

	return &responses.ResponseData{
		StatusCode: http.StatusCreated,
		Message:    responses.StatusSuccess,
		Data:       true,
	}
}

// GetContractByID implements ContractService.
func (c *ContractServiceImpl) GetContractByID(id int) *responses.ResponseData {
	contract, err := c.repo.GetContractByID(context.Background(), int32(id))
	contract.SignatureA, _ = common.DecryptBase64AES(contract.SignatureA, global.Config.JWT.AESKey)
	signB, _ := common.DecryptBase64AES(*contract.SignatureB, global.Config.JWT.AESKey)
	contract.SignatureB = &signB
	if err != nil {
		if (contract == dataaccess.GetContractByIDRow{}) {
			return &responses.ResponseData{
				StatusCode: http.StatusNoContent,
				Message:    responses.StatusNoData,
				Data:       nil,
			}
		}
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    responses.StatusInternalError,
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       contract,
	}
}

// ListContractByStatus implements ContractService.
func (c *ContractServiceImpl) ListContractByStatus(statusID int) *responses.ResponseData {
	status := int32(statusID)
	contracts, err := c.repo.ListContractByStatus(context.Background(), &status)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       nil,
		}
	}
	if len(contracts) == 0 {
		return &responses.ResponseData{
			StatusCode: http.StatusNoContent,
			Message:    responses.StatusNoData,
			Data:       nil,
		}
	}

	return &responses.ResponseData{
		StatusCode: http.StatusOK,
		Message:    responses.StatusSuccess,
		Data:       contracts,
	}

}

// SignContract implements ContractService.
func (c *ContractServiceImpl) SignContract(req dataaccess.SignContractParams, userID int) *responses.ResponseData {
	contract, err := c.blockchain.GetMContractByIDOnChain(3)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	if userID != int(contract.Tenant) {
		return &responses.ResponseData{
			StatusCode: http.StatusForbidden,
			Message:    "Bạn không có quyền thực hiện thao tác này",
			Data:       false,
		}
	}
	signB, encryptErr := common.EncryptBase64AES(*req.SignatureB, global.Config.JWT.AESKey)
	req.SignatureB = &signB
	if encryptErr != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Lỗi ký hợp đồng",
			Data:       false,
		}
	}
	// err := c.repo.SignContract(context.Background(), req)
	params := &requests.SignMContractOnChainReq{
		ContractId: int64(req.ID),
		SignatureB: *req.SignatureB,
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

	createTenantParam := dataaccess.CreateTenantParams{
		RoomID:   int32(contract.RoomID),
		TenantID: int32(contract.Tenant),
		BeginDate: pgtype.Timestamptz{
			Time:  time.Unix(contract.BeginDate, 0),
			Valid: true,
		},
	}

	errUpdateTenant := c.repo.CreateTenant(context.Background(), createTenantParam)
	if errUpdateTenant != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    errUpdateTenant.Error(),
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
	contract, _ := c.repo.GetContractByID(context.Background(), int32(id))
	if userID != int(contract.PartyB) {
		return &responses.ResponseData{
			StatusCode: http.StatusForbidden,
			Message:    "Bạn không có quyền thực hiện thao tác này",
			Data:       false,
		}
	}

	err := c.repo.DeclineContract(context.Background(), contract.ID)
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
