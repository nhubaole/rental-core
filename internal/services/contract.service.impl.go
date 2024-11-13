package services

import (
	"context"

	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"

	"github.com/jackc/pgx/v5/pgtype"
)

type ContractServiceImpl struct {
	repo *dataaccess.Queries
	blockchain BlockchainService
}

func NewContractServiceImpl(blockchain BlockchainService) ContractService {
	return &ContractServiceImpl{
		repo: dataaccess.New(global.Db),
		blockchain: blockchain,
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
func (c *ContractServiceImpl) CreateContract(req requests.CreateContractRequest) *responses.ResponseData {
	template, err := c.repo.GetContractTemplateByAddress(context.Background(), req.Address)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Không tìm thấy mẫu hợp đồng cho địa chỉ này",
			Data:       false,
		}
	}
	signOfA, encryptedErrA := common.EncryptBase64AES(req.SignatureA, global.Config.JWT.AESKey)
	//signOfB, encryptedErrB := common.EncryptBase64AES(req.SignatureB, global.Config.JWT.AESKey)
	if encryptedErrA != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    "Lỗi ký hợp đồng",
			Data:       false,
		}
	}
	parkingFee := common.IfNullInt64(req.ParkingFee,common.Float64PtrToInt64Ptr(&template.ParkingFee))
	generalResponsibility := common.IfNullStr(req.GeneralResponsibility, &template.GeneralResponsibility)
	contract := &requests.CreateLeaseAgreementOnChainReq{
		ContractCode:      req.Code,
		TenantAddress:     "",
		RoomID:            int64(req.RoomID),    // Converting int32 to int64
		ActualPrice:       int(req.ActualPrice), // Converting float64 to int
		PaymentMethod:     *req.PaymentMethod,
		ElectricityMethod: common.IfNullStr(&req.ElectricityMethod, &template.ElectricityMethod),
		ElectricityCost:   common.IfNullInt64((&req.ElectricityCost),common.Float64PtrToInt64Ptr(&template.ElectricityCost)),
		WaterMethod:       common.IfNullStr(&req.WaterMethod, &template.WaterMethod),
		WaterCost:         common.IfNullInt64(&req.WaterCost,common.Float64PtrToInt64Ptr(&template.WaterCost)),
		InternetCost:      common.IfNullInt64(&req.InternetCost,common.Float64PtrToInt64Ptr(&template.InternetCost)),
		ParkingFee:        parkingFee,
		DepositAmount:         int(req.Deposit),          // Assuming Deposit is a float64 in CreateLeaseAgreementOnChainReq as well
		BeginDate:             req.BeginDate.Time.Unix(), // Assuming pgtype.Date type compatibility in CreateLeaseAgreementOnChainReq
		EndDate:               req.EndDate.Time.Unix(),   // Assuming pgtype.Date type compatibility in CreateLeaseAgreementOnChainReq
		ResponsibilityA:       common.IfNullStr(&req.ResponsibilityA, &template.ResponsibilityA),
		ResponsibilityB:       common.IfNullStr(&req.ResponsibilityB, &template.ResponsibilityB),
		GeneralResponsibility: generalResponsibility,
		SignatureA:            signOfA,                   // Assuming you will handle converting the signature to a [6]byte type
		SignedTimeA:           req.SignedTimeA.Time.Unix(), // Assuming pgtype.Timestamptz type compatibility in CreateLeaseAgreementOnChainReq
		ContractTemplateID:    int64(template.ID),          // Converting template.ID to int64
	}

	if _,err := c.blockchain.CreateLeaseAgreementProducerContract("", *contract); err != nil {
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
	contract, _ := c.repo.GetContractByID(context.Background(), int32(req.ID))
	if userID != int(contract.PartyB) {
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
	err := c.repo.SignContract(context.Background(), req)
	if err != nil {
		return &responses.ResponseData{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
			Data:       false,
		}
	}
	createTenantParam := dataaccess.CreateTenantParams{
		RoomID:    contract.RoomID,
		TenantID:  contract.PartyB,
		BeginDate: pgtype.Timestamptz(contract.BeginDate),
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
