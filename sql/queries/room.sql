-- name: CreateRoom :exec
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
);

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