package requests

import "mime/multipart"

type CreateRoomRatingReq struct {
	RoomID            *int32                  `form:"room_id"`
	AmenitiesRating   *int32                  `form:"amenities_rating"`
	LocationRating    *int32                  `form:"location_rating"`
	CleanlinessRating *int32                  `form:"cleanliness_rating"`
	PriceRating       *int32                  `form:"price_rating"`
	OverallRating     *int32                  `form:"overall_rating"`
	Comments          *string                 `form:"comments"`
	Images            []*multipart.FileHeader `form:"images"`
}

type CreateTenantRatingReq struct {
	TenantID                      *int32                  `form:"tenant_id"`
	PaymentRating                 *int32                  `form:"payment_rating"`
	PropertyCareRating            *int32                  `form:"property_care_rating"`
	NeighborhoodDisturbanceRating *int32                  `form:"neighborhood_disturbance_rating"`
	ContractComplianceRating      *int32                  `form:"contract_compliance_rating"`
	OverallRating                 *int32                  `form:"overall_rating"`
	Comments                      *string                 `form:"comments"`
	Images                        []*multipart.FileHeader `form:"images"`
}

type CreateLandlordRatingReq struct {
	LandlordID            *int32  `json:"landlord_id"`
	FriendlinessRating    *int32  `json:"friendliness_rating"`
	ProfessionalismRating *int32  `json:"professionalism_rating"`
	SupportRating         *int32  `json:"support_rating"`
	TransparencyRating    *int32  `json:"transparency_rating"`
	OverallRating         *int32  `json:"overall_rating"`
	Comments              *string `json:"comments"`
}