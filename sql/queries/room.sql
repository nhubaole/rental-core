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
    id, 
    title, 
    address, 
    room_number, 
    room_images, 
    utilities, 
    description, 
    room_type, 
    owner, 
    capacity, 
    gender, 
    area, 
    total_price, 
    deposit, 
    electricity_cost, 
    water_cost, 
    internet_cost, 
    is_parking, 
    parking_fee, 
    status, 
    is_rent, 
    created_at, 
    updated_at
FROM 
    PUBLIC.rooms
WHERE 
    deleted_at IS NULL;

-- name: GetRoomByID :one
SELECT 
    id, 
    title, 
    address, 
    room_number, 
    room_images, 
    utilities, 
    description, 
    room_type, 
    owner, 
    capacity, 
    gender, 
    area, 
    total_price, 
    deposit, 
    electricity_cost, 
    water_cost, 
    internet_cost, 
    is_parking, 
    parking_fee, 
    status, 
    is_rent, 
    created_at, 
    updated_at
FROM 
    PUBLIC.rooms
WHERE 
    deleted_at IS NULL 
    AND id = $1;

-- name: SearchRoomByAddress :many
SELECT 
    id, 
    title, 
    address, 
    room_number, 
    room_images, 
    utilities, 
    description, 
    room_type, 
    owner, 
    capacity, 
    gender, 
    area, 
    total_price, 
    deposit, 
    electricity_cost, 
    water_cost, 
    internet_cost, 
    is_parking, 
    parking_fee, 
    status, 
    is_rent, 
    created_at, 
    updated_at
FROM 
    PUBLIC.rooms
WHERE 
    deleted_at IS NULL
    AND array_to_string(address, ', ') ILIKE '%' || $1 || '%';  


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

