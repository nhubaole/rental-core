package requests

import "github.com/jackc/pgx/v5/pgtype"

type CreateContractRequest struct {
	Address               []string           `json:"address"`
	Code                  string             `json:"code"`
	PartyA                int32              `json:"party_a"`
	PartyB                int32              `json:"party_b"`
	RequestID             int32              `json:"request_id"`
	RoomID                int32              `json:"room_id"`
	ActualPrice           float64            `json:"actual_price"`
	PaymentMethod         *string            `json:"payment_method"`
	ElectricityMethod     string             `json:"electricity_method"`
	ElectricityCost       float64            `json:"electricity_cost"`
	WaterMethod           string             `json:"water_method"`
	WaterCost             float64            `json:"water_cost"`
	InternetCost          float64            `json:"internet_cost"`
	ParkingFee            *float64           `json:"parking_fee"`
	Deposit               float64            `json:"deposit"`
	BeginDate             pgtype.Date        `json:"begin_date"`
	EndDate               pgtype.Date        `json:"end_date"`
	ResponsibilityA       string             `json:"responsibility_a"`
	ResponsibilityB       string             `json:"responsibility_b"`
	GeneralResponsibility *string            `json:"general_responsibility"`
	SignatureA            string             `json:"signature_a"`
	SignedTimeA           pgtype.Timestamptz `json:"signed_time_a"`
	SignatureB            string             `json:"signature_b"`
	SignedTimeB           pgtype.Timestamptz `json:"signed_time_b"`
}

type GetTemplateByAddressRequest struct {
	Address []string `json:"address"`
}