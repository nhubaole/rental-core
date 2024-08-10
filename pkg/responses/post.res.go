package responses

import "time"

type PostResponse struct {
	Id           uint         `json:"id"`
	TextContent  string       `json:"textContent"`
	ImageContent []string       `json:"imageContent"`
	Author       UserResponse `json:"author"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
}
