package requests

type CreateBankInfoReq struct {
	BankID        int     `json:"bank_id"`
	AccountNumber string  `json:"account_number"`
	AccountName   string  `json:"account_name"`
	CardNumber    *string `json:"card_number"`
	Currency      *string `json:"currency"`
}

type UpdateBankInfoReq struct {
	BankID        int     `json:"bank_id"`
	AccountNumber string  `json:"account_number"`
	AccountName   string  `json:"account_name"`
	CardNumber    *string `json:"card_number"`
	Currency      *string `json:"currency"`
}

type UpdateDeviceTokenReq struct {
	DeviceToken string `json:"device_token"`
}