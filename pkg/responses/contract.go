package responses

type ContractOnChainRes struct {
    ContractAddress string `json:"contract_address"`
    Landlord        string `json:"landlord"`
    Tenant          string `json:"tenant"`
    RoomID          int64  `json:"room_id"`
    ActualPrice     int64  `json:"actual_price"`
    DepositAmount   int64  `json:"deposit_amount"`
    BeginDate       int64  `json:"begin_date"`
    EndDate         int64  `json:"end_date"`
    ContractCode    string `json:"contract_code"`
}
