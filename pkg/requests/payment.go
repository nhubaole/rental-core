package requests

import (
	"mime/multipart"

)

type CreatePaymentReq struct {
	Code            string                `form:"code"`
	BillID          *int32                `form:"bill_id"`
	ContractID      *int32                `form:"contract_id"`
	Amount          float64               `form:"amount"`
	Status          int32                 `form:"status"`
	ReturnRequestID *int32                `form:"return_request_id"`
	TransferContent *string               `form:"transfer_content"`
	EvidenceImage   *multipart.FileHeader `form:"evidence_image"`
}
