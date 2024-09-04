package responses

import "time"

type UserResponse struct {
	Id        uint      `json:"id"`
	PhoneNumber  string    `json:"phoneNumber"`
	FullName  string    `json:"fullName"`
	Address     string    `json:"address"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"createdAt"`
}
