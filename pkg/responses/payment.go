package responses

type GetPaymentInfoRes struct {
	BankName       string  `json:"bank_name"`
	AccountName    string  `json:"account_name"`
	AccountNumber  string  `json:"account_number"`
	Amount         float64 `json:"amount"`
	TranferContent string  `json:"tranfer_content"`
	QrUrl          string  `json:"qr_url"`
}
