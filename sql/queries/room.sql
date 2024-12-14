-- name: CreateRoom :one
INSERT INTO "rooms" 
(
  "title", 
  "address", 
  "room_number", 
  "room_images", 
  "utilities", 
  "description", 
  "room_type", 
  "owner", 
  "capacity", 
  "gender", 
  "area", 
  "total_price", 
  "deposit", 
  "electricity_cost", 
  "water_cost", 
  "internet_cost", 
  "is_parking", 
  "parking_fee", 
  "status", 
  "is_rent", 
  "created_at", 
  "updated_at"
) 
VALUES 
(
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, now(), now()
)
RETURNING id;

-- name: GetRooms :many
SELECT 
    r.id, 
    r.title, 
    r.address, 
    r.room_number, 
    r.room_images, 
    r.utilities, 
    r.description, 
    r.room_type, 
    r.owner, 
    r.area, 
    r.total_price, 
    r.status, 
    COALESCE(AVG(rt.overall_rating), 0) AS avg_rating,  -- Tính trung bình rating
    COALESCE(COUNT(rt.id), 0) AS total_rating,  -- Đếm tổng số lượng rating
     EXISTS (
        SELECT 1 
        FROM PUBLIC."like" l
        WHERE l.room_id = r.id AND l.user_id = $1 AND l.deleted_at IS NULL
    ) AS is_liked 
FROM 
    public.rooms r
LEFT JOIN 
    public.room_ratings rt ON r.id = rt.room_id
WHERE 
    r.deleted_at IS NULL
GROUP BY 
    r.id, 
    r.title, 
    r.address, 
    r.room_number, 
    r.room_images, 
    r.utilities, 
    r.description, 
    r.room_type, 
    r.owner, 
    r.area, 
    r.total_price, 
    r.status;


-- name: GetRoomByID :one
SELECT 
    r.id, 
    r.title, 
    r.address, 
    r.room_number, 
    r.room_images, 
    r.utilities, 
    r.description, 
    r.room_type, 
    r.available_from,
    (SELECT jsonb_object_agg(room_number, id)::text
    FROM PUBLIC.rooms 
    WHERE deleted_at IS NULL AND address = r.address) AS list_room_numbers,
    r.owner, 
    r.capacity, 
    r.gender, 
    r.area, 
    r.total_price, 
    r.deposit, 
    r.electricity_cost, 
    r.water_cost, 
    r.internet_cost, 
    r.is_parking, 
    r.parking_fee, 
    r.status, 
    r.is_rent, 
    r.created_at, 
    r.updated_at
FROM 
    PUBLIC.rooms r
WHERE 
    deleted_at IS NULL 
    AND r.id = $1;

-- name: SearchRoomByAddress :many
SELECT 
    r.id, 
    r.title, 
    r.address, 
    r.room_number, 
    r.room_images, 
    r.utilities, 
    r.description, 
    r.room_type, 
    r.available_from,
    r.owner, 
    r.area, 
    r.total_price, 
    r.status, 
    COALESCE(AVG(rt.overall_rating), 0) AS avg_rating,  -- Tính trung bình rating
    COALESCE(COUNT(rt.id), 0) AS total_rating  -- Đếm tổng số lượng rating
FROM 
    public.rooms r
LEFT JOIN 
    public.room_ratings rt ON r.id = rt.room_id
WHERE 
    r.deleted_at IS NULL
    AND array_to_string(r.address, ', ') ILIKE '%' || $1 || '%'
GROUP BY 
    r.id, 
    r.title, 
    r.address, 
    r.room_number, 
    r.room_images, 
    r.utilities, 
    r.description, 
    r.room_type, 
    r.owner, 
    r.area, 
    r.total_price, 
    r.status;;  


-- name: LikeRoom :exec
INSERT INTO PUBLIC."like"
(
    room_id,
    user_id,
    created_at,
    updated_at
) VALUES
(
    $1, $2, NOW(), NOW()
);

-- name: UnlikeRoom :exec
DELETE FROM PUBLIC."like"
WHERE room_id = $1 AND user_id = $2;

-- name: CheckUserLikedRoom :one
SELECT 1
FROM PUBLIC."like"
WHERE room_id = $1 AND user_id = $2 AND deleted_at IS NULL;

-- name: GetLikedRooms :many
SELECT r.*
FROM PUBLIC."like" l
JOIN PUBLIC.rooms r ON l.room_id = r.id
WHERE l.user_id = $1 AND l.deleted_at IS NULL;

-- name: GetRoomsByStatus :many
SELECT *
FROM PUBLIC.rooms
WHERE status = $1;

-- name: UpdateRoom :one
UPDATE 
    PUBLIC.rooms
SET 
    title = COALESCE($2, title),
    address = COALESCE($3, address),
    room_number = COALESCE($4, room_number),
    room_images = COALESCE($5, room_images),
    utilities = COALESCE($6, utilities),
    description = COALESCE($7, description),
    room_type = COALESCE($8, room_type),
    capacity = COALESCE($9, capacity),
    gender = COALESCE($10, gender),
    area = COALESCE($11, area),
    total_price = COALESCE($12, total_price),
    deposit = COALESCE($13, deposit),
    electricity_cost = COALESCE($14, electricity_cost),
    water_cost = COALESCE($15, water_cost),
    internet_cost = COALESCE($16, internet_cost),
    is_parking = COALESCE($17, is_parking),
    parking_fee = COALESCE($18, parking_fee),
    status = COALESCE($19, status),
    is_rent = COALESCE($20, is_rent),
    updated_at = NOW() 
WHERE 
    deleted_at IS NULL
    AND id = $1
    RETURNING id;

-- name: GetRoomsByOwner :many
SELECT *
FROM PUBLIC.rooms
where owner = $1;

-- name: GetRoomByContractID :one
SELECT r.id AS room_id,
		r.title,
		r.address,
		r.room_number,
        r.owner
FROM PUBLIC.rooms r
LEFT JOIN public.contracts c ON r.id = c.room_id
WHERE c.id = $1;

