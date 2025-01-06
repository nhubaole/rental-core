package responses

import (
	"smart-rental/internal/dataaccess"

	"github.com/jackc/pgx/v5/pgtype"
)

type ConversationRes struct {
	ID          int32              `json:"id"`
	UserA       int32              `json:"user_a"`
	UserB       int32              `json:"user_b"`
	LastMessage *dataaccess.GetMessageByIDRow `json:"last_message"`
	CreatedAt   pgtype.Timestamp   `json:"created_at"`
}
