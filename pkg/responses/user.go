package responses

import "time"

type UserResponse struct {
	ID          int32     `json:"id"`
	PhoneNumber string    `json:"phone_number"`
	FullName    string    `json:"full_name"`
	Address     *string   `json:"address"`
	Role        int       `json:"role"`
	CreatedAt   time.Time `json:"created_at"`
}
