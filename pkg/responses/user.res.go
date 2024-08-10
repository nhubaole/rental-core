package responses

import "time"

type UserResponse struct {
	Id        uint      `json:"id"`
	UserName  string    `json:"userName"`
	FullName  string    `json:"fullName"`
	Email     string    `json:"email"`
	Title     string    `json:"title"`
	Gender    string    `json:"gender"`
	AvatarUrl string    `json:"avatarUrl"`
	CreatedAt time.Time `json:"createdAt"`
}
