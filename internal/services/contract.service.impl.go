package services

import (
	"context"

	"net/http"
	"smart-rental/global"
	"smart-rental/internal/dataaccess"
	"smart-rental/pkg/common"
	"smart-rental/pkg/requests"
	"smart-rental/pkg/responses"
)

type ContractServiceImpl struct {
	repo *dataaccess.Queries
}


func NewContractServiceImpl() ContractService {
	return &ContractServiceImpl{
		repo: dataaccess.New(global.Db),
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
	parkingFee := common.IfNullFloat64(req.ParkingFee, &template.ParkingFee)
	generalResponsibility := common.IfNullStr(req.GeneralResponsibility, &template.GeneralResponsibility)
	contract := &dataaccess.CreateContractParams{
		Code:                  req.Code,
		PartyA:                req.PartyA,
		PartyB:                req.PartyB,
		RequestID:             req.RequestID,
		RoomID:                req.RoomID,
		ActualPrice:           req.ActualPrice,
		PaymentMethod:         req.PaymentMethod,
		ElectricityMethod:     common.IfNullStr(&req.ElectricityMethod, &template.ElectricityMethod),
		ElectricityCost:       common.IfNullFloat64(&req.ElectricityCost, &template.ElectricityCost),
		WaterMethod:           common.IfNullStr(&req.WaterMethod, &template.WaterMethod),
		WaterCost:             common.IfNullFloat64(&req.WaterCost, &template.WaterCost),
		InternetCost:          common.IfNullFloat64(&req.InternetCost, &template.InternetCost),
		ParkingFee:            &parkingFee,
		Deposit:               req.Deposit,
		BeginDate:             req.BeginDate,
		EndDate:               req.EndDate,
		ResponsibilityA:       common.IfNullStr(&req.ResponsibilityA, &template.ResponsibilityA),
		ResponsibilityB:       common.IfNullStr(&req.ResponsibilityB, &template.ResponsibilityB),
		GeneralResponsibility: &generalResponsibility,
		SignatureA:            req.SignatureA,
		SignedTimeA:           req.SignedTimeA,
		SignatureB:            req.SignatureB,
		SignedTimeB:           req.SignedTimeB,
		ContractTemplateID:    &template.ID,
	}

	if err := c.repo.CreateContract(context.Background(), *contract); err != nil {
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
	if (len(contracts) == 0) {
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