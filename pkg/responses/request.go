package responses

import "github.com/jackc/pgx/v5/pgtype"

type GetRequestByRoomIDRes struct {
	ID              int                `json:"id"`
	Code            string             `json:"code"`
	Sender          UserResponse       `json:"sender"`
	RoomID          int32              `json:"room_id"`
	SuggestedPrice  *float64           `json:"suggested_price"`
	NumOfPerson     *int32             `json:"num_of_person"`
	BeginDate       pgtype.Timestamptz `json:"begin_date"`
	EndDate         pgtype.Timestamptz `json:"end_date"`
	AdditionRequest *string            `json:"addition_request"`
	Status          int32              `json:"status"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
}

type GetReturnRequestByLandlordIDRes struct {
	ID                 int32            `json:"id"`
	ContractID         *int32           `json:"contract_id"`
	RoomID             int32            `json:"room_id"`
	Reason             *string          `json:"reason"`
	ReturnDate         pgtype.Timestamp `json:"return_date"`
	Status             *int32           `json:"status"`
	DeductAmount       *float64         `json:"deduct_amount"`
	TotalReturnDeposit *float64         `json:"total_return_deposit"`
	CreatedUser        UserResponse     `json:"created_user"`
	CreatedAt          pgtype.Timestamp `json:"created_at"`
	UpdatedAt          pgtype.Timestamp `json:"updated_at"`
}
