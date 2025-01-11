// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: room.sql

package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkUserLikedRoom = `-- name: CheckUserLikedRoom :one
SELECT 1
FROM PUBLIC."like"
WHERE room_id = $1 AND user_id = $2 AND deleted_at IS NULL
`

type CheckUserLikedRoomParams struct {
	RoomID int32 `json:"room_id"`
	UserID int32 `json:"user_id"`
}

func (q *Queries) CheckUserLikedRoom(ctx context.Context, arg CheckUserLikedRoomParams) (int32, error) {
	row := q.db.QueryRow(ctx, checkUserLikedRoom, arg.RoomID, arg.UserID)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const createRoom = `-- name: CreateRoom :one
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
RETURNING id
`

type CreateRoomParams struct {
	Title           string   `json:"title"`
	Address         []string `json:"address"`
	RoomNumber      int32    `json:"room_number"`
	RoomImages      []string `json:"room_images"`
	Utilities       []string `json:"utilities"`
	Description     string   `json:"description"`
	RoomType        *string  `json:"room_type"`
	Owner           int32    `json:"owner"`
	Capacity        int32    `json:"capacity"`
	Gender          *int32   `json:"gender"`
	Area            float64  `json:"area"`
	TotalPrice      *float64 `json:"total_price"`
	Deposit         float64  `json:"deposit"`
	ElectricityCost float64  `json:"electricity_cost"`
	WaterCost       float64  `json:"water_cost"`
	InternetCost    float64  `json:"internet_cost"`
	IsParking       bool     `json:"is_parking"`
	ParkingFee      *float64 `json:"parking_fee"`
	Status          int32    `json:"status"`
	IsRent          bool     `json:"is_rent"`
}

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) (int32, error) {
	row := q.db.QueryRow(ctx, createRoom,
		arg.Title,
		arg.Address,
		arg.RoomNumber,
		arg.RoomImages,
		arg.Utilities,
		arg.Description,
		arg.RoomType,
		arg.Owner,
		arg.Capacity,
		arg.Gender,
		arg.Area,
		arg.TotalPrice,
		arg.Deposit,
		arg.ElectricityCost,
		arg.WaterCost,
		arg.InternetCost,
		arg.IsParking,
		arg.ParkingFee,
		arg.Status,
		arg.IsRent,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getLikedRooms = `-- name: GetLikedRooms :many
SELECT r.id, r.title, r.address, r.room_number, r.room_images, r.utilities, r.description, r.room_type, r.owner, r.capacity, r.gender, r.area, r.total_price, r.deposit, r.electricity_cost, r.water_cost, r.internet_cost, r.is_parking, r.parking_fee, r.status, r.is_rent, r.created_at, r.updated_at, r.deleted_at, r.available_from
FROM PUBLIC."like" l
JOIN PUBLIC.rooms r ON l.room_id = r.id
WHERE l.user_id = $1 AND l.deleted_at IS NULL
`

func (q *Queries) GetLikedRooms(ctx context.Context, userID int32) ([]Room, error) {
	rows, err := q.db.Query(ctx, getLikedRooms, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Room
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Address,
			&i.RoomNumber,
			&i.RoomImages,
			&i.Utilities,
			&i.Description,
			&i.RoomType,
			&i.Owner,
			&i.Capacity,
			&i.Gender,
			&i.Area,
			&i.TotalPrice,
			&i.Deposit,
			&i.ElectricityCost,
			&i.WaterCost,
			&i.InternetCost,
			&i.IsParking,
			&i.ParkingFee,
			&i.Status,
			&i.IsRent,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.AvailableFrom,
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

const getRoomByContractID = `-- name: GetRoomByContractID :one
SELECT r.id AS room_id,
		r.title,
		r.address,
		r.room_number,
        r.owner
FROM PUBLIC.rooms r
LEFT JOIN public.contracts c ON r.id = c.room_id
WHERE c.id = $1
`

type GetRoomByContractIDRow struct {
	RoomID     int32    `json:"room_id"`
	Title      string   `json:"title"`
	Address    []string `json:"address"`
	RoomNumber int32    `json:"room_number"`
	Owner      int32    `json:"owner"`
}

func (q *Queries) GetRoomByContractID(ctx context.Context, id int32) (GetRoomByContractIDRow, error) {
	row := q.db.QueryRow(ctx, getRoomByContractID, id)
	var i GetRoomByContractIDRow
	err := row.Scan(
		&i.RoomID,
		&i.Title,
		&i.Address,
		&i.RoomNumber,
		&i.Owner,
	)
	return i, err
}

const getRoomByID = `-- name: GetRoomByID :one
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
    AND r.id = $1
`

type GetRoomByIDRow struct {
	ID              int32              `json:"id"`
	Title           string             `json:"title"`
	Address         []string           `json:"address"`
	RoomNumber      int32              `json:"room_number"`
	RoomImages      []string           `json:"room_images"`
	Utilities       []string           `json:"utilities"`
	Description     string             `json:"description"`
	RoomType        *string            `json:"room_type"`
	AvailableFrom   pgtype.Timestamptz `json:"available_from"`
	ListRoomNumbers string             `json:"list_room_numbers"`
	Owner           int32              `json:"owner"`
	Capacity        int32              `json:"capacity"`
	Gender          *int32             `json:"gender"`
	Area            float64            `json:"area"`
	TotalPrice      *float64           `json:"total_price"`
	Deposit         float64            `json:"deposit"`
	ElectricityCost float64            `json:"electricity_cost"`
	WaterCost       float64            `json:"water_cost"`
	InternetCost    float64            `json:"internet_cost"`
	IsParking       bool               `json:"is_parking"`
	ParkingFee      *float64           `json:"parking_fee"`
	Status          int32              `json:"status"`
	IsRent          bool               `json:"is_rent"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) GetRoomByID(ctx context.Context, id int32) (GetRoomByIDRow, error) {
	row := q.db.QueryRow(ctx, getRoomByID, id)
	var i GetRoomByIDRow
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Address,
		&i.RoomNumber,
		&i.RoomImages,
		&i.Utilities,
		&i.Description,
		&i.RoomType,
		&i.AvailableFrom,
		&i.ListRoomNumbers,
		&i.Owner,
		&i.Capacity,
		&i.Gender,
		&i.Area,
		&i.TotalPrice,
		&i.Deposit,
		&i.ElectricityCost,
		&i.WaterCost,
		&i.InternetCost,
		&i.IsParking,
		&i.ParkingFee,
		&i.Status,
		&i.IsRent,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getRooms = `-- name: GetRooms :many
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
    r.status
ORDER BY 
    r.created_at DESC
`

type GetRoomsRow struct {
	ID          int32       `json:"id"`
	Title       string      `json:"title"`
	Address     []string    `json:"address"`
	RoomNumber  int32       `json:"room_number"`
	RoomImages  []string    `json:"room_images"`
	Utilities   []string    `json:"utilities"`
	Description string      `json:"description"`
	RoomType    *string     `json:"room_type"`
	Owner       int32       `json:"owner"`
	Area        float64     `json:"area"`
	TotalPrice  *float64    `json:"total_price"`
	Status      int32       `json:"status"`
	AvgRating   interface{} `json:"avg_rating"`
	TotalRating interface{} `json:"total_rating"`
	IsLiked     bool        `json:"is_liked"`
}

func (q *Queries) GetRooms(ctx context.Context, userID int32) ([]GetRoomsRow, error) {
	rows, err := q.db.Query(ctx, getRooms, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRoomsRow
	for rows.Next() {
		var i GetRoomsRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Address,
			&i.RoomNumber,
			&i.RoomImages,
			&i.Utilities,
			&i.Description,
			&i.RoomType,
			&i.Owner,
			&i.Area,
			&i.TotalPrice,
			&i.Status,
			&i.AvgRating,
			&i.TotalRating,
			&i.IsLiked,
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

const getRoomsByOwner = `-- name: GetRoomsByOwner :many
SELECT id, title, address, room_number, room_images, utilities, description, room_type, owner, capacity, gender, area, total_price, deposit, electricity_cost, water_cost, internet_cost, is_parking, parking_fee, status, is_rent, created_at, updated_at, deleted_at, available_from
FROM PUBLIC.rooms
where owner = $1
`

func (q *Queries) GetRoomsByOwner(ctx context.Context, owner int32) ([]Room, error) {
	rows, err := q.db.Query(ctx, getRoomsByOwner, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Room
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Address,
			&i.RoomNumber,
			&i.RoomImages,
			&i.Utilities,
			&i.Description,
			&i.RoomType,
			&i.Owner,
			&i.Capacity,
			&i.Gender,
			&i.Area,
			&i.TotalPrice,
			&i.Deposit,
			&i.ElectricityCost,
			&i.WaterCost,
			&i.InternetCost,
			&i.IsParking,
			&i.ParkingFee,
			&i.Status,
			&i.IsRent,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.AvailableFrom,
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

const getRoomsByStatus = `-- name: GetRoomsByStatus :many
SELECT id, title, address, room_number, room_images, utilities, description, room_type, owner, capacity, gender, area, total_price, deposit, electricity_cost, water_cost, internet_cost, is_parking, parking_fee, status, is_rent, created_at, updated_at, deleted_at, available_from
FROM PUBLIC.rooms
WHERE status = $1
`

func (q *Queries) GetRoomsByStatus(ctx context.Context, status int32) ([]Room, error) {
	rows, err := q.db.Query(ctx, getRoomsByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Room
	for rows.Next() {
		var i Room
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Address,
			&i.RoomNumber,
			&i.RoomImages,
			&i.Utilities,
			&i.Description,
			&i.RoomType,
			&i.Owner,
			&i.Capacity,
			&i.Gender,
			&i.Area,
			&i.TotalPrice,
			&i.Deposit,
			&i.ElectricityCost,
			&i.WaterCost,
			&i.InternetCost,
			&i.IsParking,
			&i.ParkingFee,
			&i.Status,
			&i.IsRent,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.AvailableFrom,
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

const likeRoom = `-- name: LikeRoom :exec
INSERT INTO PUBLIC."like"
(
    room_id,
    user_id,
    created_at,
    updated_at
) VALUES
(
    $1, $2, NOW(), NOW()
)
`

type LikeRoomParams struct {
	RoomID int32 `json:"room_id"`
	UserID int32 `json:"user_id"`
}

func (q *Queries) LikeRoom(ctx context.Context, arg LikeRoomParams) error {
	_, err := q.db.Exec(ctx, likeRoom, arg.RoomID, arg.UserID)
	return err
}

const searchRoomByAddress = `-- name: SearchRoomByAddress :many
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
    r.gender,
    r.capacity,
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
    r.available_from,
    r.gender,
    r.capacity,
    r.owner, 
    r.area, 
    r.total_price, 
    r.status
ORDER BY 
    r.created_at DESC
`

type SearchRoomByAddressRow struct {
	ID            int32              `json:"id"`
	Title         string             `json:"title"`
	Address       []string           `json:"address"`
	RoomNumber    int32              `json:"room_number"`
	RoomImages    []string           `json:"room_images"`
	Utilities     []string           `json:"utilities"`
	Description   string             `json:"description"`
	RoomType      *string            `json:"room_type"`
	AvailableFrom pgtype.Timestamptz `json:"available_from"`
	Gender        *int32             `json:"gender"`
	Capacity      int32              `json:"capacity"`
	Owner         int32              `json:"owner"`
	Area          float64            `json:"area"`
	TotalPrice    *float64           `json:"total_price"`
	Status        int32              `json:"status"`
	AvgRating     interface{}        `json:"avg_rating"`
	TotalRating   interface{}        `json:"total_rating"`
}

func (q *Queries) SearchRoomByAddress(ctx context.Context, dollar_1 *string) ([]SearchRoomByAddressRow, error) {
	rows, err := q.db.Query(ctx, searchRoomByAddress, dollar_1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SearchRoomByAddressRow
	for rows.Next() {
		var i SearchRoomByAddressRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Address,
			&i.RoomNumber,
			&i.RoomImages,
			&i.Utilities,
			&i.Description,
			&i.RoomType,
			&i.AvailableFrom,
			&i.Gender,
			&i.Capacity,
			&i.Owner,
			&i.Area,
			&i.TotalPrice,
			&i.Status,
			&i.AvgRating,
			&i.TotalRating,
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

const unlikeRoom = `-- name: UnlikeRoom :exec
DELETE FROM PUBLIC."like"
WHERE room_id = $1 AND user_id = $2
`

type UnlikeRoomParams struct {
	RoomID int32 `json:"room_id"`
	UserID int32 `json:"user_id"`
}

func (q *Queries) UnlikeRoom(ctx context.Context, arg UnlikeRoomParams) error {
	_, err := q.db.Exec(ctx, unlikeRoom, arg.RoomID, arg.UserID)
	return err
}

const updateRoom = `-- name: UpdateRoom :one
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
    RETURNING id
`

type UpdateRoomParams struct {
	ID              int32    `json:"id"`
	Title           string   `json:"title"`
	Address         []string `json:"address"`
	RoomNumber      int32    `json:"room_number"`
	RoomImages      []string `json:"room_images"`
	Utilities       []string `json:"utilities"`
	Description     string   `json:"description"`
	RoomType        *string  `json:"room_type"`
	Capacity        int32    `json:"capacity"`
	Gender          *int32   `json:"gender"`
	Area            float64  `json:"area"`
	TotalPrice      *float64 `json:"total_price"`
	Deposit         float64  `json:"deposit"`
	ElectricityCost float64  `json:"electricity_cost"`
	WaterCost       float64  `json:"water_cost"`
	InternetCost    float64  `json:"internet_cost"`
	IsParking       bool     `json:"is_parking"`
	ParkingFee      *float64 `json:"parking_fee"`
	Status          int32    `json:"status"`
	IsRent          bool     `json:"is_rent"`
}

func (q *Queries) UpdateRoom(ctx context.Context, arg UpdateRoomParams) (int32, error) {
	row := q.db.QueryRow(ctx, updateRoom,
		arg.ID,
		arg.Title,
		arg.Address,
		arg.RoomNumber,
		arg.RoomImages,
		arg.Utilities,
		arg.Description,
		arg.RoomType,
		arg.Capacity,
		arg.Gender,
		arg.Area,
		arg.TotalPrice,
		arg.Deposit,
		arg.ElectricityCost,
		arg.WaterCost,
		arg.InternetCost,
		arg.IsParking,
		arg.ParkingFee,
		arg.Status,
		arg.IsRent,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}
