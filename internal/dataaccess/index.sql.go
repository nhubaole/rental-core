// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: index.sql

package dataaccess

import (
	"context"
)

const getAllIndex = `-- name: GetAllIndex :many
SELECT 
    ro.id AS room_id, 
    t.id, 
    t.prev_month, 
    t.curr_month, 
    t.prev_water, 
    t.curr_water, 
    t.prev_electricity, 
    t.curr_electricity, 
    t.year
FROM (
    SELECT 
        id, 
        LAG(i.month) OVER (PARTITION BY i.room_id ORDER BY i.year, i.month) AS prev_month, 
        i.month AS curr_month,
        LAG(i.water_index) OVER (PARTITION BY i.room_id ORDER BY i.year, i.month) AS prev_water, 
        i.water_index AS curr_water, 
        LAG(i.electricity_index) OVER (PARTITION BY i.room_id ORDER BY i.year, i.month) AS prev_electricity, 
        i.electricity_index AS curr_electricity, 
        i.room_id, 
        i.year
    FROM public.index AS i
) AS t
LEFT JOIN public.rooms AS ro ON t.room_id = ro.id
LEFT JOIN public.index AS idx ON t.id = idx.id
WHERE ro.id = $1
  AND idx.month = $2
  AND idx.year = $3
ORDER BY t.year, t.curr_month
`

type GetAllIndexParams struct {
	ID    int32 `json:"id"`
	Month int32 `json:"month"`
	Year  int32 `json:"year"`
}

type GetAllIndexRow struct {
	RoomID          *int32      `json:"room_id"`
	ID              int32       `json:"id"`
	PrevMonth       interface{} `json:"prev_month"`
	CurrMonth       int32       `json:"curr_month"`
	PrevWater       interface{} `json:"prev_water"`
	CurrWater       *float64    `json:"curr_water"`
	PrevElectricity interface{} `json:"prev_electricity"`
	CurrElectricity *float64    `json:"curr_electricity"`
	Year            int32       `json:"year"`
}

func (q *Queries) GetAllIndex(ctx context.Context, arg GetAllIndexParams) ([]GetAllIndexRow, error) {
	rows, err := q.db.Query(ctx, getAllIndex, arg.ID, arg.Month, arg.Year)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllIndexRow
	for rows.Next() {
		var i GetAllIndexRow
		if err := rows.Scan(
			&i.RoomID,
			&i.ID,
			&i.PrevMonth,
			&i.CurrMonth,
			&i.PrevWater,
			&i.CurrWater,
			&i.PrevElectricity,
			&i.CurrElectricity,
			&i.Year,
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

const getIndexById = `-- name: GetIndexById :one
SELECT id, water_index, electricity_index, room_id, month, year 
FROM PUBLIC.INDEX
WHERE id = $1
`

func (q *Queries) GetIndexById(ctx context.Context, id int32) (Index, error) {
	row := q.db.QueryRow(ctx, getIndexById, id)
	var i Index
	err := row.Scan(
		&i.ID,
		&i.WaterIndex,
		&i.ElectricityIndex,
		&i.RoomID,
		&i.Month,
		&i.Year,
	)
	return i, err
}

const getIndexByOwnerIdShort = `-- name: GetIndexByOwnerIdShort :many
SELECT idx.id, idx.room_id, idx.water_index, idx.electricity_index, idx.month, idx.year
FROM  PUBLIC.INDEX AS idx 
LEFT JOIN public.ROOMS ro ON ro.id = idx.room_id
WHERE  ro.owner = $1
`

type GetIndexByOwnerIdShortRow struct {
	ID               int32    `json:"id"`
	RoomID           int32    `json:"room_id"`
	WaterIndex       *float64 `json:"water_index"`
	ElectricityIndex *float64 `json:"electricity_index"`
	Month            int32    `json:"month"`
	Year             int32    `json:"year"`
}

func (q *Queries) GetIndexByOwnerIdShort(ctx context.Context, owner int32) ([]GetIndexByOwnerIdShortRow, error) {
	rows, err := q.db.Query(ctx, getIndexByOwnerIdShort, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetIndexByOwnerIdShortRow
	for rows.Next() {
		var i GetIndexByOwnerIdShortRow
		if err := rows.Scan(
			&i.ID,
			&i.RoomID,
			&i.WaterIndex,
			&i.ElectricityIndex,
			&i.Month,
			&i.Year,
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

const getIndexByRoomId = `-- name: GetIndexByRoomId :many
SELECT id, water_index, electricity_index, room_id, month, year
FROM PUBLIC.INDEX
WHERE room_id = $1
`

func (q *Queries) GetIndexByRoomId(ctx context.Context, roomID int32) ([]Index, error) {
	rows, err := q.db.Query(ctx, getIndexByRoomId, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Index
	for rows.Next() {
		var i Index
		if err := rows.Scan(
			&i.ID,
			&i.WaterIndex,
			&i.ElectricityIndex,
			&i.RoomID,
			&i.Month,
			&i.Year,
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

const upsertIndex = `-- name: UpsertIndex :one
INSERT INTO PUBLIC.INDEX(
  water_index,
  electricity_index,
  room_id,
  month,
  year
)
VALUES(
    $1, $2, $3, $4, $5
  )
ON CONFLICT (room_id, month, year)
DO UPDATE
SET
  water_index = COALESCE(EXCLUDED.water_index, PUBLIC.INDEX.water_index),
  electricity_index = COALESCE(EXCLUDED.electricity_index, PUBLIC.INDEX.electricity_index)
RETURNING id, water_index, electricity_index, room_id, month, year
`

type UpsertIndexParams struct {
	WaterIndex       *float64 `json:"water_index"`
	ElectricityIndex *float64 `json:"electricity_index"`
	RoomID           int32    `json:"room_id"`
	Month            int32    `json:"month"`
	Year             int32    `json:"year"`
}

func (q *Queries) UpsertIndex(ctx context.Context, arg UpsertIndexParams) (Index, error) {
	row := q.db.QueryRow(ctx, upsertIndex,
		arg.WaterIndex,
		arg.ElectricityIndex,
		arg.RoomID,
		arg.Month,
		arg.Year,
	)
	var i Index
	err := row.Scan(
		&i.ID,
		&i.WaterIndex,
		&i.ElectricityIndex,
		&i.RoomID,
		&i.Month,
		&i.Year,
	)
	return i, err
}
