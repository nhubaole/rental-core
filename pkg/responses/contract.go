package responses

// type ContractOnChainRes struct {
//     ContractAddress string `json:"contract_address"`
//     Landlord        string `json:"landlord"`
//     Tenant          string `json:"tenant"`
//     RoomID          int64  `json:"room_id"`
//     ActualPrice     int64  `json:"actual_price"`
//     DepositAmount   int64  `json:"deposit_amount"`
//     BeginDate       int64  `json:"begin_date"`
//     EndDate         int64  `json:"end_date"`
//     ContractCode    string `json:"contract_code"`
// }

type ContractOnChainRes struct {
	ID                    int32
	Code                  string
	PartyA                int32
	PartyB                int32
	RequestID             int32
	RoomID                int32
	ActualPrice           float64
	PaymentMethod         string
	ElectricityMethod     string
	ElectricityCost       float64
	WaterMethod           string
	WaterCost             float64
	InternetCost          float64
	ParkingFee            float64
	Deposit               float64
	BeginDate             int64   // UNIX timestamp
	EndDate               int64   // UNIX timestamp
	ResponsibilityA       string
	ResponsibilityB       string
	GeneralResponsibility string
	SignatureA            string
	SignedTimeA           int64   // UNIX timestamp
	SignatureB            string
	SignedTimeB           int64   // UNIX timestamp
	CreatedAt             int64   // UNIX timestamp
	UpdatedAt             int64   // UNIX timestamp
	ContractTemplateID    int32
}