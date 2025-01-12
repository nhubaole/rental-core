package responses

import (
	"smart-rental/internal/dataaccess"

	"github.com/jackc/pgx/v5/pgtype"
)

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
	CreatedUser        GetUserByIDRes     `json:"created_user"`
	CreatedAt          pgtype.Timestamp `json:"created_at"`
	UpdatedAt          pgtype.Timestamp `json:"updated_at"`
}
type GetRentalRequestByIDRes struct {
	ID              int32                     `json:"id"`
	Code            string                    `json:"code"`
	Sender          dataaccess.GetUserByIDRow `json:"sender"`
	Room            dataaccess.GetRoomByIDRow `json:"room"`
	SuggestedPrice  *float64                  `json:"suggested_price"`
	NumOfPerson     *int32                    `json:"num_of_person"`
	BeginDate       pgtype.Timestamptz        `json:"begin_date"`
	EndDate         pgtype.Timestamptz        `json:"end_date"`
	AdditionRequest *string                   `json:"addition_request"`
	Status          int32                     `json:"status"`
	CreatedAt       pgtype.Timestamptz        `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz        `json:"updated_at"`
	DeletedAt       pgtype.Timestamptz        `json:"deleted_at"`
}

type GetReturnRequestByIDRes struct {
	ID                 int32                     `json:"id"`
	Reason             *string                   `json:"reason"`
	ContractID         *int32                    `json:"contract_id"`
	ContractCode       string                    `json:"contract_code"`
	Room               dataaccess.GetRoomByIDRow `json:"room"`
	ReturnDate         pgtype.Timestamp          `json:"return_date"`
	Status             *int32                    `json:"status"`
	DeductAmount       *float64                  `json:"deduct_amount"`
	TotalReturnDeposit *float64                  `json:"total_return_deposit"`
	CreatedUser        dataaccess.GetUserByIDRow `json:"created_user"`
	CreatedAt          pgtype.Timestamp          `json:"created_at"`
	UpdatedAt          pgtype.Timestamp          `json:"updated_at"`
}

type GetProcessTracking struct {
	ID        int32                     `json:"id"`
	Actor     dataaccess.GetUserByIDRow `json:"actor"`
	Action    string                    `json:"action"`
	IssuedAt  pgtype.Timestamptz        `json:"issued_at"`
	RequestID int32                     `json:"request_id"`
}
