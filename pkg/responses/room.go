package responses

import (
	"smart-rental/internal/dataaccess"

	"github.com/jackc/pgx/v5/pgtype"
)

type RoomOnChainRes struct {
	ID         int64
	TotalPrice int
	Deposit    int64
	Status     int64
	IsRent     bool
	CreatedAt  int64
	UpdatedAt  int64
}

type GetRoomByIDRes struct {
	ID              int32                     `json:"id"`
	Title           string                    `json:"title"`
	Address         []string                  `json:"address"`
	RoomNumber      int32                     `json:"room_number"`
	RoomImages      []string                  `json:"room_images"`
	Utilities       []string                  `json:"utilities"`
	Description     string                    `json:"description"`
	RoomType        *string                   `json:"room_type"`
	Owner           dataaccess.GetUserByIDRow `json:"owner"`
	Capacity        int32                     `json:"capacity"`
	Gender          *int32                    `json:"gender"`
	Area            float64                   `json:"area"`
	TotalPrice      *float64                  `json:"total_price"`
	Deposit         float64                   `json:"deposit"`
	ElectricityCost float64                   `json:"electricity_cost"`
	WaterCost       float64                   `json:"water_cost"`
	InternetCost    float64                   `json:"internet_cost"`
	IsParking       bool                      `json:"is_parking"`
	ParkingFee      *float64                  `json:"parking_fee"`
	Status          int32                     `json:"status"`
	IsRent          bool                      `json:"is_rent"`
	CreatedAt       pgtype.Timestamptz        `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz        `json:"updated_at"`
}

type MContractOnChainRes struct {
	ID                    int64  // ID của hợp đồng
	Code                  string // Mã hợp đồng
	Landlord              int64  // ID của chủ nhà
	Tenant                int64  // ID của người thuê
	RoomID                int64  // ID của phòng
	ActualPrice           int64  // Giá thực tế của hợp đồng
	Deposit               int64  // Tiền đặt cọc
	BeginDate             int64  // Thời gian bắt đầu hợp đồng (timestamp)
	EndDate               int64  // Thời gian kết thúc hợp đồng (timestamp)
	PaymentMethod         string // Phương thức thanh toán
	ElectricityMethod     string // Phương thức tính điện
	ElectricityCost       int64  // Giá điện
	WaterMethod           string // Phương thức tính nước
	WaterCost             int64  // Giá nước
	InternetCost          int64  // Giá internet
	ParkingFee            int64  // Phí gửi xe
	ResponsibilityA       string // Trách nhiệm bên A
	ResponsibilityB       string // Trách nhiệm bên B
	GeneralResponsibility string // Trách nhiệm chung
	SignatureA            string // Chữ ký của bên A
	SignedTimeA           int64  // Thời gian ký của bên A (timestamp)
	SignatureB            string // Chữ ký của bên B
	SignedTimeB           int64  // Thời gian ký của bên B (timestamp)
	ContractTemplateID    int64  // ID mẫu hợp đồng
	PreRentalStatus       uint8  // Trạng thái trước khi thuê
	RentalProcessStatus   uint8  // Trạng thái trong quá trình thuê
	PostRentalStatus      uint8  // Trạng thái sau khi thuê
	CreatedAt             int64  // Thời gian tạo hợp đồng (timestamp)
	UpdatedAt             int64  // Thời gian cập nhật hợp đồng (timestamp)
}

type GetAllMetric4BillByRoomID struct {
	RoomID          int32       `json:"room_id"`
	PrevMonth       interface{} `json:"prev_month"`
	CurrMonth       int32       `json:"curr_month"`
	PrevWater       interface{} `json:"prev_water"`
	CurrWater       float64     `json:"curr_water"`
	PrevElectricity interface{} `json:"prev_electricity"`
	CurrElectricity float64     `json:"curr_electricity"`
	Year            int32       `json:"year"`
	ContractID      int32       `json:"contract_id"`      // Contract ID
	ActualPrice     int64       `json:"actual_price"`     // Actual price of the contract
	WaterCost       int64       `json:"water_cost"`       // Water cost
	ElectricityCost int64       `json:"electricity_cost"` // Electricity cost
	InternetCost    int64       `json:"internet_cost"`    // Internet cost
	ParkingFee      int64       `json:"parking_fee"`      // Parking fee
}
