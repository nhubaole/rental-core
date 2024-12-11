package responses

import "github.com/jackc/pgx/v5/pgtype"

type GetBillByMonthRes struct {
	ID                   int                `json:"id"`
	Code                 string             `json:"code"`
	ContractID           int32              `json:"contract_id"`
	RoomID               int32              `json:"room_id"`
	LandlordID           int64              `json:"landlord_id"`
	TenantID             int64              `json:"tenant_id"`
	AdditionFee          *int32             `json:"addition_fee"`
	AdditionNote         *string            `json:"addition_note"`
	TotalAmount          float64            `json:"total_amount"`
	Month                int32              `json:"month"`
	Year                 int32              `json:"year"`
	OldWaterIndex        *int32             `json:"old_water_index"`
	OldElectricityIndex  *int32             `json:"old_electricity_index"`
	NewWaterIndex        *int32             `json:"new_water_index"`
	NewElectricityIndex  *int32             `json:"new_electricity_index"`
	TotalWaterCost       *float64           `json:"total_water_cost"`
	TotalElectricityCost *float64           `json:"total_electricity_cost"`
	InternetCost         float64            `json:"internet_cost"`
	ParkingFee           float64            `json:"parking_fee"`
	Status               *int32             `json:"status"`
	CreatedAt            pgtype.Timestamptz `json:"created_at"`
	UpdatedAt            pgtype.Timestamptz `json:"updated_at"`
}
