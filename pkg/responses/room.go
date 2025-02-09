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
	AvailableFrom   pgtype.Timestamptz        `json:"available_from"`
	ListRoomNumbers map[int]interface{}       `json:"list_room_numbers"`
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
	Latitude        *float64                  `json:"latitude"`
	Longitude       *float64                  `json:"longitude"`
	Status          int32                     `json:"status"`
	IsRent          bool                      `json:"is_rent"`
	CreatedAt       pgtype.Timestamptz        `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz        `json:"updated_at"`
}

type MContractOnChainRes struct {
	ID                    int64  `json:"id"`
	Code                  string `json:"code"`
	Landlord              int64  `json:"landlord"`
	Tenant                int64  `json:"tenant"`
	RoomID                int64  `json:"room_id"`
	ActualPrice           int64  `json:"actual_price"`
	Deposit               int64  `json:"deposit"`
	BeginDate             int64  `json:"begin_date"`
	EndDate               int64  `json:"end_date"`
	PaymentMethod         string `json:"payment_method"`
	ElectricityMethod     string `json:"electricity_method"`
	ElectricityCost       int64  `json:"electricity_cost"`
	WaterMethod           string `json:"water_method"`
	WaterCost             int64  `json:"water_cost"`
	InternetCost          int64  `json:"internet_cost"`
	ParkingFee            int64  `json:"parking_fee"`
	ResponsibilityA       string `json:"responsibility_a"`
	ResponsibilityB       string `json:"responsibility_b"`
	GeneralResponsibility string `json:"general_responsibility"`
	SignatureA            string `json:"signature_a"`
	SignedTimeA           int64  `json:"signed_time_a"`
	SignatureB            string `json:"signature_b"`
	SignedTimeB           int64  `json:"signed_time_b"`
	ContractTemplateID    int64  `json:"contract_template_id"`
	PreRentalStatus       uint8  `json:"pre_rental_status"`
	RentalProcessStatus   uint8  `json:"rental_process_status"`
	PostRentalStatus      uint8  `json:"post_rental_status"`
	CreatedAt             int64  `json:"created_at"`
	UpdatedAt             int64  `json:"updated_at"`
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

type GetRoomsRes struct {
	ID          int32       `json:"id"`
	Title       string      `json:"title"`
	Address     []string    `json:"address"`
	RoomNumber  int32       `json:"room_number"`
	RoomImages  []string    `json:"room_images"`
	Utilities   []string    `json:"utilities"`
	Description string      `json:"description"`
	RoomType    *string     `json:"room_type"`
	Owner       int32       `json:"owner"`
	Area        float64     `json:"area"`
	TotalPrice  *float64    `json:"total_price"`
	Status      int32       `json:"status"`
	AvgRating   interface{} `json:"avg_rating"`
	TotalRating interface{} `json:"total_rating"`
	IsLiked     bool        `json:"is_liked"`
}
