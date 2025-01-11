// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
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

const getLandlordRatingByID = `-- name: GetLandlordRatingByID :many
SELECT id,
 landlord_id, 
 rated_by, 
 friendliness_rating, 
 professionalism_rating, 
 support_rating, 
 transparency_rating, 
 overall_rating, 
 comments, 
 created_at
FROM public.landlord_ratings
WHERE landlord_id = $1
`

func (q *Queries) GetLandlordRatingByID(ctx context.Context, landlordID *int32) ([]LandlordRating, error) {
	rows, err := q.db.Query(ctx, getLandlordRatingByID, landlordID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LandlordRating
	for rows.Next() {
		var i LandlordRating
		if err := rows.Scan(
			&i.ID,
			&i.LandlordID,
			&i.RatedBy,
			&i.FriendlinessRating,
			&i.ProfessionalismRating,
			&i.SupportRating,
			&i.TransparencyRating,
			&i.OverallRating,
			&i.Comments,
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

const getRoomRatingByRoomID = `-- name: GetRoomRatingByRoomID :one
WITH happy_unhappy AS (
    SELECT
        rr.room_id,
        rr.rated_by,
        string_agg(
            CASE
                WHEN rr.amenities_rating = 5 THEN 'tiện nghi'
                WHEN rr.location_rating = 5 THEN 'vị trí tốt'
                WHEN rr.cleanliness_rating = 5 THEN 'sạch sẽ'
                WHEN rr.price_rating = 5 THEN 'giá cả hợp lý'
                ELSE NULL
            END, ','
        ) AS happy,
        string_agg(
            CASE
                WHEN rr.amenities_rating = 1 THEN 'thiếu tiện nghi'
                WHEN rr.location_rating = 1 THEN 'vị trí không tốt'
                WHEN rr.cleanliness_rating = 1 THEN 'không sạch sẽ'
                WHEN rr.price_rating = 1 THEN 'giá cả không hợp lý'
                ELSE NULL
            END, ','
        ) AS unhappy
    FROM
        public.room_ratings rr
    WHERE
        rr.room_id = $1
    GROUP BY
        rr.room_id, rr.rated_by
)
SELECT 
    COUNT(*) AS total_rating,  -- Tổng số lượng rating
    COALESCE(
        jsonb_object_agg(
            subquery.overall_rating, 
            subquery.rating_count
        ), 
        '{}'::jsonb
    ) AS detail_count, -- Đếm số lượng rating theo từng mức
    COALESCE(AVG(rr.overall_rating), 0) AS avg_rating, -- Trung bình rating
    jsonb_agg(
        jsonb_build_object(
            'rater_id', u.id,
            'rater_name', u.full_name,
            'rater_avatar', u.avatar_url,
            'created_at', rr.created_at,
            'rate', rr.overall_rating,
            'comment', rr.comments,
            'images', rr.images,
            'happy', hu.happy,
            'unhappy', hu.unhappy
        )
    )::text AS rating_info 
FROM 
    public.room_ratings rr
LEFT JOIN 
    public.users u ON rr.rated_by = u.id
LEFT JOIN (
    SELECT 
        r.room_id, 
        r.overall_rating, 
        COUNT(*) AS rating_count
    FROM 
        public.room_ratings r
    WHERE 
        r.room_id = $1
    GROUP BY 
        r.room_id, r.overall_rating
) AS subquery ON subquery.room_id = rr.room_id AND subquery.overall_rating = rr.overall_rating
LEFT JOIN 
    happy_unhappy hu ON hu.room_id = rr.room_id AND hu.rated_by = rr.rated_by
WHERE 
    rr.room_id = $1
GROUP BY
    rr.room_id
`

type GetRoomRatingByRoomIDRow struct {
	TotalRating int64       `json:"total_rating"`
	DetailCount interface{} `json:"detail_count"`
	AvgRating   interface{} `json:"avg_rating"`
	RatingInfo  string      `json:"rating_info"`
}

func (q *Queries) GetRoomRatingByRoomID(ctx context.Context, roomID *int32) (GetRoomRatingByRoomIDRow, error) {
	row := q.db.QueryRow(ctx, getRoomRatingByRoomID, roomID)
	var i GetRoomRatingByRoomIDRow
	err := row.Scan(
		&i.TotalRating,
		&i.DetailCount,
		&i.AvgRating,
		&i.RatingInfo,
	)
	return i, err
}

const getTenantRatingByID = `-- name: GetTenantRatingByID :many
SELECT id,
 tenant_id, 
 rated_by, 
 payment_rating, 
 property_care_rating, 
 neighborhood_disturbance_rating, 
 contract_compliance_rating, 
 overall_rating, 
 comments, 
 images, 
 created_at
FROM public.tenant_ratings
WHERE tenant_id = $1
`

func (q *Queries) GetTenantRatingByID(ctx context.Context, tenantID *int32) ([]TenantRating, error) {
	rows, err := q.db.Query(ctx, getTenantRatingByID, tenantID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []TenantRating
	for rows.Next() {
		var i TenantRating
		if err := rows.Scan(
			&i.ID,
			&i.TenantID,
			&i.RatedBy,
			&i.PaymentRating,
			&i.PropertyCareRating,
			&i.NeighborhoodDisturbanceRating,
			&i.ContractComplianceRating,
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
