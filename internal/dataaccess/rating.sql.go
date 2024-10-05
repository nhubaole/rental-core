// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: rating.sql

package dataaccess

import (
	"context"
)

const createLandlordRating = `-- name: CreateLandlordRating :exec
INSERT INTO public.landlord_ratings(
    landlord_id, --1
    rated_by, --2
    friendliness_rating, --3
    professionalism_rating, --4
    support_rating, --5
    transparency_rating, --6
    overall_rating, --7
    comments, --8
    created_at)
VALUES
($1, $2, $3, $4, $5, $6,$7, $8, now())
`

type CreateLandlordRatingParams struct {
	LandlordID            *int32  `json:"landlord_id"`
	RatedBy               *int32  `json:"rated_by"`
	FriendlinessRating    *int32  `json:"friendliness_rating"`
	ProfessionalismRating *int32  `json:"professionalism_rating"`
	SupportRating         *int32  `json:"support_rating"`
	TransparencyRating    *int32  `json:"transparency_rating"`
	OverallRating         *int32  `json:"overall_rating"`
	Comments              *string `json:"comments"`
}

func (q *Queries) CreateLandlordRating(ctx context.Context, arg CreateLandlordRatingParams) error {
	_, err := q.db.Exec(ctx, createLandlordRating,
		arg.LandlordID,
		arg.RatedBy,
		arg.FriendlinessRating,
		arg.ProfessionalismRating,
		arg.SupportRating,
		arg.TransparencyRating,
		arg.OverallRating,
		arg.Comments,
	)
	return err
}

const createRoomRating = `-- name: CreateRoomRating :exec
INSERT INTO public.room_ratings(
    room_id, -- 1
    rated_by, --2
    amenities_rating, --3
    location_rating, --4
    cleanliness_rating, --5
    price_rating, --6
    overall_rating, --7
    comments, --8
    images, --9
    created_at
)
VALUES
($1, $2, $3, $4, $5, $6,$7, $8, $9, now())
`

type CreateRoomRatingParams struct {
	RoomID            *int32   `json:"room_id"`
	RatedBy           *int32   `json:"rated_by"`
	AmenitiesRating   *int32   `json:"amenities_rating"`
	LocationRating    *int32   `json:"location_rating"`
	CleanlinessRating *int32   `json:"cleanliness_rating"`
	PriceRating       *int32   `json:"price_rating"`
	OverallRating     *int32   `json:"overall_rating"`
	Comments          *string  `json:"comments"`
	Images            []string `json:"images"`
}

func (q *Queries) CreateRoomRating(ctx context.Context, arg CreateRoomRatingParams) error {
	_, err := q.db.Exec(ctx, createRoomRating,
		arg.RoomID,
		arg.RatedBy,
		arg.AmenitiesRating,
		arg.LocationRating,
		arg.CleanlinessRating,
		arg.PriceRating,
		arg.OverallRating,
		arg.Comments,
		arg.Images,
	)
	return err
}

const createTenantRating = `-- name: CreateTenantRating :exec
INSERT INTO public.tenant_ratings(
    tenant_id, --1
    rated_by,  --2
    payment_rating, --3 
    property_care_rating, --4 
    neighborhood_disturbance_rating, --5
    contract_compliance_rating, --6
    overall_rating,  --7
    comments, --8
    images, --9
    created_at)
VALUES
($1, $2, $3, $4, $5, $6,$7, $8, $9, now())
`

type CreateTenantRatingParams struct {
	TenantID                      *int32   `json:"tenant_id"`
	RatedBy                       *int32   `json:"rated_by"`
	PaymentRating                 *int32   `json:"payment_rating"`
	PropertyCareRating            *int32   `json:"property_care_rating"`
	NeighborhoodDisturbanceRating *int32   `json:"neighborhood_disturbance_rating"`
	ContractComplianceRating      *int32   `json:"contract_compliance_rating"`
	OverallRating                 *int32   `json:"overall_rating"`
	Comments                      *string  `json:"comments"`
	Images                        []string `json:"images"`
}

func (q *Queries) CreateTenantRating(ctx context.Context, arg CreateTenantRatingParams) error {
	_, err := q.db.Exec(ctx, createTenantRating,
		arg.TenantID,
		arg.RatedBy,
		arg.PaymentRating,
		arg.PropertyCareRating,
		arg.NeighborhoodDisturbanceRating,
		arg.ContractComplianceRating,
		arg.OverallRating,
		arg.Comments,
		arg.Images,
	)
	return err
}

const getRoomRatingByRoomID = `-- name: GetRoomRatingByRoomID :many
SELECT id,
 room_id, 
 rated_by, 
 amenities_rating, 
 location_rating, 
 cleanliness_rating, 
 price_rating, 
 overall_rating, 
 comments, 
 images, 
 created_at
FROM public.room_ratings
WHERE room_id = $1
`

func (q *Queries) GetRoomRatingByRoomID(ctx context.Context, roomID *int32) ([]RoomRating, error) {
	rows, err := q.db.Query(ctx, getRoomRatingByRoomID, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RoomRating
	for rows.Next() {
		var i RoomRating
		if err := rows.Scan(
			&i.ID,
			&i.RoomID,
			&i.RatedBy,
			&i.AmenitiesRating,
			&i.LocationRating,
			&i.CleanlinessRating,
			&i.PriceRating,
			&i.OverallRating,
			&i.Comments,
			&i.Images,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}