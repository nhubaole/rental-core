package responses

type GetPaymentInfoRes struct {
	BankName       string  `json:"bank_name"`
	ShortName      string  `json:"short_name"`
	Logo           string  `json:"logo"`
	AccountName    string  `json:"account_name"`
	AccountNumber  string  `json:"account_number"`
	Amount         float64 `json:"amount"`
	TranferContent string  `json:"tranfer_content"`
	QrUrl          string  `json:"qr_url"`
}
