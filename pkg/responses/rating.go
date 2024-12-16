package responses

type GetRoomRatingByRoomIDRes struct {
	TotalRating int64                    `json:"total_rating"`
	DetailCount interface{}              `json:"detail_count"`
	AvgRating   interface{}              `json:"avg_rating"`
	RatingInfo  []map[string]interface{} `json:"rating_info"`
}
