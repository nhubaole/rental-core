package responses

type GetBillByMonthRes struct {
	Address  string                   `json:"address"`
	ListBill []map[string]interface{} `json:"list_bill"`
}
