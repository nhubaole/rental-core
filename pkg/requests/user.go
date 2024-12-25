package requests

import (
	"mime/multipart"

)

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

type UpdateUserReq struct {
	ID          int32                 `form:"id"`
	PhoneNumber string                `form:"phone_number"`
	FullName    string                `form:"full_name"`
	Address     *string               `form:"address"`
	Role        int32                 `form:"role"`
	Otp         *int32                `form:"otp"`
	Gender      *int32                `form:"gender"`
	Dob         string           `form:"dob"`
	AvatarUrl   *multipart.FileHeader `form:"avatar_url"`
	DeleteFile  *string               `form:"delete_file"`
}
