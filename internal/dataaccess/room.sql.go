// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
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

const createRoom = `-- name: CreateRoom :exec
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

func (q *Queries) CreateRoom(ctx context.Context, arg CreateRoomParams) error {
	_, err := q.db.Exec(ctx, createRoom,
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
	return err
}

const getRoomByID = `-- name: GetRoomByID :one
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
    AND id = $1
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
`

type GetRoomsRow struct {
	ID              int32              `json:"id"`
	Title           string             `json:"title"`
	Address         []string           `json:"address"`
	RoomNumber      int32              `json:"room_number"`
	RoomImages      []string           `json:"room_images"`
	Utilities       []string           `json:"utilities"`
	Description     string             `json:"description"`
	RoomType        *string            `json:"room_type"`
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

func (q *Queries) GetRooms(ctx context.Context) ([]GetRoomsRow, error) {
	rows, err := q.db.Query(ctx, getRooms)
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
    AND array_to_string(address, ', ') ILIKE '%' || $1 || '%'
`

type SearchRoomByAddressRow struct {
	ID              int32              `json:"id"`
	Title           string             `json:"title"`
	Address         []string           `json:"address"`
	RoomNumber      int32              `json:"room_number"`
	RoomImages      []string           `json:"room_images"`
	Utilities       []string           `json:"utilities"`
	Description     string             `json:"description"`
	RoomType        *string            `json:"room_type"`
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
