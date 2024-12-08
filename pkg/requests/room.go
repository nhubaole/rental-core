package requests

import "mime/multipart"

type CreateRoomForm struct {
	Title           string                  `form:"title"`
	Address         []string                `form:"address"`
	RoomNumber      int32                   `form:"room_number"`
	RoomImages      []*multipart.FileHeader `form:"room_images"`
	Utilities       []string                `form:"utilities"`
	Description     string                  `form:"description"`
	RoomType        *string                 `form:"room_type"`
	Owner           int32                   `form:"owner"`
	Capacity        int32                   `form:"capacity"`
	Gender          *int32                  `form:"gender"`
	Area            float64                 `form:"area"`
	TotalPrice      *float64                `form:"total_price"`
	Deposit         float64                 `form:"deposit"`
	ElectricityCost float64                 `form:"electricity_cost"`
	WaterCost       float64                 `form:"water_cost"`
	InternetCost    float64                 `form:"internet_cost"`
	IsParking       bool                    `form:"is_parking"`
	ParkingFee      *float64                `form:"parking_fee"`
	Status          int32                   `form:"status"`
	IsRent          bool                    `form:"is_rent"`
}

type UpdateRoomRequest struct {
	ID              int32                   `form:"id"`
	Title           string                  `form:"title"`
	Address         []string                `form:"address"`
	RoomNumber      int32                   `form:"room_number"`
	RoomImages      []*multipart.FileHeader `form:"room_images"`
	Utilities       []string                `form:"utilities"`
	Description     string                  `form:"description"`
	RoomType        *string                 `form:"room_type"`
	Capacity        int32                   `form:"capacity"`
	Gender          *int32                  `form:"gender"`
	Area            float64                 `form:"area"`
	TotalPrice      *float64                `form:"total_price"`
	Deposit         float64                 `form:"deposit"`
	ElectricityCost float64                 `form:"electricity_cost"`
	WaterCost       float64                 `form:"water_cost"`
	InternetCost    float64                 `form:"internet_cost"`
	IsParking       bool                    `form:"is_parking"`
	ParkingFee      *float64                `form:"parking_fee"`
	Status          int32                   `form:"status"`
	IsRent          bool                    `form:"is_rent"`
	DeleteFiles     []string                `form:"delete_files"`
}

type CreateRoomOnChainReq struct {
	RoomID     int64
	TotalPrice int
	Deposit    int64
	Status     int64
	IsRent     bool
}

type CreateMContractOnChainReq struct {
	ContractId            int64
	ContractCode          string
	LandlordId            int64
	TenantId              int64
	RoomId                int64
	ActualPrice           int64
	Deposit               int64
	BeginDate             int64
	EndDate               int64
	PaymentMethod         string
	ElectricityMethod     string
	ElectricityCost       int64
	WaterMethod           string
	WaterCost             int64
	InternetCost          int64
	ParkingFee            int64
	ResponsibilityA       string
	ResponsibilityB       string
	GeneralResponsibility string
	SignatureA            string
	SignedTimeA           int64
	SignatureB            string
	SignedTimeB           int64
	ContractTemplateId    int64
}

type SignMContractOnChainReq struct {
	ContractId int64
	SignatureB string
}
