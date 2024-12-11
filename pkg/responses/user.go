package responses

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserResponse struct {
	ID          int32     `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	FullName    string    `json:"full_name"`
	Address     *string   `json:"address"`
	Role        int       `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
}

type RatingInfo struct {
	RaterName string             `json:"rater_name"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Rate      int                `json:"rate"`
	Comment   string             `json:"comment"`
	Happy     string             `json:"happy"`
	UnHappy   string             `json:"unhappy"`
}

type GetUserByIDRes struct {
	ID            int32              `json:"id"`
	PhoneNumber   string             `json:"phone_number"`
	AvatarUrl     *string            `json:"avatar_url"`
	Role          int32              `json:"role"`
	FullName      string             `json:"full_name"`
	TotalRoom     int                `json:"total_room"`
	TotalRating   int                `json:"total_rating"`
	AvgRating     float64                `json:"avg_rating"`
	Address       *string            `json:"address"`
	RatingInfo    []RatingInfo         `json:"rating_info"`
	WalletAddress *string            `json:"wallet_address"`
	PrivateKeyHex *string            `json:"private_key_hex"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
}
