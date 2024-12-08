package requests

import "github.com/jackc/pgx/v5/pgtype"

type CreateContractRequest struct {
	Address               []string           `json:"address"`
	Code                  string             `json:"code"`
	PartyA                int32              `json:"party_a"`
	PartyB                int32              `json:"party_b"`
	RequestID             int32              `json:"request_id"`
	RoomID                int32              `json:"room_id"`
	ActualPrice           int64              `json:"actual_price"`
	PaymentMethod         *string            `json:"payment_method"`
	ElectricityMethod     string             `json:"electricity_method"`
	ElectricityCost       int64              `json:"electricity_cost"`
	WaterMethod           string             `json:"water_method"`
	WaterCost             int64              `json:"water_cost"`
	InternetCost          int64              `json:"internet_cost"`
	ParkingFee            *int64             `json:"parking_fee"`
	Deposit               int64              `json:"deposit"`
	BeginDate             pgtype.Date        `json:"begin_date"`
	EndDate               pgtype.Date        `json:"end_date"`
	ResponsibilityA       string             `json:"responsibility_a"`
	ResponsibilityB       string             `json:"responsibility_b"`
	GeneralResponsibility *string            `json:"general_responsibility"`
	SignatureA            string             `json:"signature_a"`
	SignedTimeA           pgtype.Timestamptz `json:"signed_time_a"`
	// SignatureB            string             `json:"signature_b"`
	// SignedTimeB           pgtype.Timestamptz `json:"signed_time_b"`
}

type GetTemplateByAddressRequest struct {
	Address []string `json:"address"`
}

type CreateLeaseAgreementOnChainReq struct {
	TenantAddress         string // Address of the tenant
	RoomID                int64  // ID of the room
	ActualPrice           int    // Actual price of the lease
	DepositAmount         int    // Deposit amount for the lease
	BeginDate             int64  // Start date of the lease (timestamp)
	EndDate               int64  // End date of the lease (timestamp)
	ContractCode          string // Unique code for the contract
	SignatureA            string // Signature of the landlord
	SignedTimeA           int64  // Timestamp of landlord's signature
	PaymentMethod         string // Method of payment (e.g., "monthly", "yearly")
	ElectricityMethod     string // Method of electricity payment
	ElectricityCost       int64  // Cost of electricity
	WaterMethod           string // Method of water payment
	WaterCost             int64  // Cost of water
	InternetCost          int64  // Cost of internet
	ParkingFee            int64  // Parking fee
	ResponsibilityA       string // Responsibilities of the landlord
	ResponsibilityB       string // Responsibilities of the tenant
	GeneralResponsibility string // General responsibilities for both parties
	ContractTemplateID    int64  // ID of the contract template used
}

type CreateContractOnChainReq struct {
	Code                  string `json:"code"`    // Mã hợp đồng
	RoomID                int64  `json:"room_id"` // ID phòng
	Tenant                string `json:"tenant"`
	TotalPrice            int64  `json:"total_price"`            // Giá thuê thực tế
	Deposit               int64  `json:"deposit"`                // Số tiền đặt cọc
	BeginDate             int64  `json:"begin_date"`             // Ngày bắt đầu (timestamp)
	EndDate               int64  `json:"end_date"`               // Ngày kết thúc (timestamp)
	PaymentMethod         string `json:"payment_method"`         // Phương thức thanh toán
	ElectricityMethod     string `json:"electricity_method"`     // Phương thức thanh toán điện
	ElectricityCost       int64  `json:"electricity_cost"`       // Chi phí điện
	WaterMethod           string `json:"water_method"`           // Phương thức thanh toán nước
	WaterCost             int64  `json:"water_cost"`             // Chi phí nước
	InternetCost          int64  `json:"internet_cost"`          // Chi phí Internet
	ParkingFee            int64  `json:"parking_fee"`            // Phí đỗ xe
	ResponsibilityA       string `json:"responsibility_a"`       // Trách nhiệm của bên A (chủ nhà)
	ResponsibilityB       string `json:"responsibility_b"`       // Trách nhiệm của bên B (người thuê)
	GeneralResponsibility string `json:"general_responsibility"` // Trách nhiệm chung
	SignatureA            string `json:"signature_a"`            // Chữ ký của bên A
	SignedTimeA           int64  `json:"signed_time_a"`          // Thời gian ký của bên A (timestamp)
	ContractTemplateId    int64  `json:"contract_template_id"`   // ID mẫu hợp đồng
}

type SignContractParams struct {
	Id         int32  `json:"id"`
	SignatureB string `json:"signature_b"`
}
