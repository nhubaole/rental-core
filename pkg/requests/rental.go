package requests

import "github.com/jackc/pgx/v5/pgtype"

type CreateRentalRequest struct {
	RoomID          int32              `json:"room_id"`
	SuggestedPrice  *float64           `json:"suggested_price"`
	NumOfPerson     *int32             `json:"num_of_person"`
	BeginDate       pgtype.Timestamptz `json:"begin_date"`
	EndDate         pgtype.Timestamptz `json:"end_date"`
	AdditionRequest *string            `json:"addition_request"`
}

