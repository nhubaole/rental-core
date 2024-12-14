-- name: CreateRoomRating :exec
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
($1, $2, $3, $4, $5, $6,$7, $8, $9, now());

-- name: CreateTenantRating :exec
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
($1, $2, $3, $4, $5, $6,$7, $8, $9, now());

-- name: CreateLandlordRating :exec
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
($1, $2, $3, $4, $5, $6,$7, $8, now());

-- name: GetRoomRatingByRoomID :one
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
            'rater_name', u.full_name,
            'created_at', rr.created_at,
            'rate', rr.overall_rating,
            'comment', rr.comments,
            'images', rr.images
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
WHERE 
    rr.room_id = $1
GROUP BY
    rr.room_id;




-- name: GetTenantRatingByID :many
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
WHERE tenant_id = $1;

-- name: GetLandlordRatingByID :many
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
WHERE landlord_id = $1;