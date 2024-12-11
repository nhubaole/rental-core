// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: billing.sql

package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBill = `-- name: CreateBill :exec
INSERT INTO PUBLIC.BILLING
(
    code,
    contract_id, --1
    old_water_index, --2
    old_electricity_index, --3
    new_water_index, --4
    new_electricity_index, --5
    total_water_cost, --6
    total_electricity_cost, --7
    addition_fee,  --8
    addition_note, --9
    total_amount, --10
    month, --11
    year, --12
    created_at,  --13
    updated_at --15
) VALUES
(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,  now(), now()
)
`

type CreateBillParams struct {
	Code                 string   `json:"code"`
	ContractID           int32    `json:"contract_id"`
	OldWaterIndex        *int32   `json:"old_water_index"`
	OldElectricityIndex  *int32   `json:"old_electricity_index"`
	NewWaterIndex        *int32   `json:"new_water_index"`
	NewElectricityIndex  *int32   `json:"new_electricity_index"`
	TotalWaterCost       *float64 `json:"total_water_cost"`
	TotalElectricityCost *float64 `json:"total_electricity_cost"`
	AdditionFee          *int32   `json:"addition_fee"`
	AdditionNote         *string  `json:"addition_note"`
	TotalAmount          float64  `json:"total_amount"`
	Month                int32    `json:"month"`
	Year                 int32    `json:"year"`
}

func (q *Queries) CreateBill(ctx context.Context, arg CreateBillParams) error {
	_, err := q.db.Exec(ctx, createBill,
		arg.Code,
		arg.ContractID,
		arg.OldWaterIndex,
		arg.OldElectricityIndex,
		arg.NewWaterIndex,
		arg.NewElectricityIndex,
		arg.TotalWaterCost,
		arg.TotalElectricityCost,
		arg.AdditionFee,
		arg.AdditionNote,
		arg.TotalAmount,
		arg.Month,
		arg.Year,
	)
	return err
}

const getAllMetric4BillByRoomID = `-- name: GetAllMetric4BillByRoomID :one
SELECT t.room_id,
       t.prev_month,
       t.curr_month,
       t.prev_water,
       t.curr_water,
       t.prev_electricity,
       t.curr_electricity, 
       t.year
FROM (
	SELECT id , LAG(i.month) OVER(ORDER BY month, year) AS prev_month , month as curr_month,
	LAG(i.water_index) OVER(ORDER BY month, year) as prev_water , water_index as curr_water, 
	LAG(i.electricity_index) OVER(ORDER BY month,year) as prev_electricity, electricity_index as curr_electricity, 
	room_id, year
	FROM public.index as i
) AS t
LEFT JOIN public.INDEX idx ON t.id = idx.id
WHERE idx.room_id = $1
AND idx.month = $2
AND idx.year = $3
`

type GetAllMetric4BillByRoomIDParams struct {
	RoomID int32 `json:"room_id"`
	Month  int32 `json:"month"`
	Year   int32 `json:"year"`
}

type GetAllMetric4BillByRoomIDRow struct {
	RoomID          int32       `json:"room_id"`
	PrevMonth       interface{} `json:"prev_month"`
	CurrMonth       int32       `json:"curr_month"`
	PrevWater       interface{} `json:"prev_water"`
	CurrWater       float64     `json:"curr_water"`
	PrevElectricity interface{} `json:"prev_electricity"`
	CurrElectricity float64     `json:"curr_electricity"`
	Year            int32       `json:"year"`
}

func (q *Queries) GetAllMetric4BillByRoomID(ctx context.Context, arg GetAllMetric4BillByRoomIDParams) (GetAllMetric4BillByRoomIDRow, error) {
	row := q.db.QueryRow(ctx, getAllMetric4BillByRoomID, arg.RoomID, arg.Month, arg.Year)
	var i GetAllMetric4BillByRoomIDRow
	err := row.Scan(
		&i.RoomID,
		&i.PrevMonth,
		&i.CurrMonth,
		&i.PrevWater,
		&i.CurrWater,
		&i.PrevElectricity,
		&i.CurrElectricity,
		&i.Year,
	)
	return i, err
}

const getBillByID = `-- name: GetBillByID :one
SELECT  b.id,
        b.code,
        b.contract_id,
        b.addition_fee,
        b.addition_note,
        b.total_amount,
        b.month,
        b.year,
        b.old_water_index, 
        b.old_electricity_index, 
        b.new_water_index, 
        b.new_electricity_index, 
        b.total_water_cost, 
        b.total_electricity_cost,
        b.status,
        b.created_at,
        b.updated_at
FROM PUBLIC.BILLING b
WHERE deleted_at IS NULL 
      AND id = $1
`

type GetBillByIDRow struct {
	ID                   int32              `json:"id"`
	Code                 string             `json:"code"`
	ContractID           int32              `json:"contract_id"`
	AdditionFee          *int32             `json:"addition_fee"`
	AdditionNote         *string            `json:"addition_note"`
	TotalAmount          float64            `json:"total_amount"`
	Month                int32              `json:"month"`
	Year                 int32              `json:"year"`
	OldWaterIndex        *int32             `json:"old_water_index"`
	OldElectricityIndex  *int32             `json:"old_electricity_index"`
	NewWaterIndex        *int32             `json:"new_water_index"`
	NewElectricityIndex  *int32             `json:"new_electricity_index"`
	TotalWaterCost       *float64           `json:"total_water_cost"`
	TotalElectricityCost *float64           `json:"total_electricity_cost"`
	Status               *int32             `json:"status"`
	CreatedAt            pgtype.Timestamptz `json:"created_at"`
	UpdatedAt            pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) GetBillByID(ctx context.Context, id int32) (GetBillByIDRow, error) {
	row := q.db.QueryRow(ctx, getBillByID, id)
	var i GetBillByIDRow
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.ContractID,
		&i.AdditionFee,
		&i.AdditionNote,
		&i.TotalAmount,
		&i.Month,
		&i.Year,
		&i.OldWaterIndex,
		&i.OldElectricityIndex,
		&i.NewWaterIndex,
		&i.NewElectricityIndex,
		&i.TotalWaterCost,
		&i.TotalElectricityCost,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getBillByMonth = `-- name: GetBillByMonth :many
SELECT  b.id,
        b.code,
        b.contract_id,
        b.addition_fee,
        b.addition_note,
        b.total_amount,
        b.month,
        b.year,
        b.old_water_index, 
        b.old_electricity_index, 
        b.new_water_index, 
        b.new_electricity_index, 
        b.total_water_cost, 
        b.total_electricity_cost,
        b.status,
        b.created_at,
        b.updated_at
FROM PUBLIC.BILLING as b
WHERE b.year = $1 
    AND b.month=$2
`

type GetBillByMonthParams struct {
	Year  int32 `json:"year"`
	Month int32 `json:"month"`
}

type GetBillByMonthRow struct {
	ID                   int32              `json:"id"`
	Code                 string             `json:"code"`
	ContractID           int32              `json:"contract_id"`
	AdditionFee          *int32             `json:"addition_fee"`
	AdditionNote         *string            `json:"addition_note"`
	TotalAmount          float64            `json:"total_amount"`
	Month                int32              `json:"month"`
	Year                 int32              `json:"year"`
	OldWaterIndex        *int32             `json:"old_water_index"`
	OldElectricityIndex  *int32             `json:"old_electricity_index"`
	NewWaterIndex        *int32             `json:"new_water_index"`
	NewElectricityIndex  *int32             `json:"new_electricity_index"`
	TotalWaterCost       *float64           `json:"total_water_cost"`
	TotalElectricityCost *float64           `json:"total_electricity_cost"`
	Status               *int32             `json:"status"`
	CreatedAt            pgtype.Timestamptz `json:"created_at"`
	UpdatedAt            pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) GetBillByMonth(ctx context.Context, arg GetBillByMonthParams) ([]GetBillByMonthRow, error) {
	rows, err := q.db.Query(ctx, getBillByMonth, arg.Year, arg.Month)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBillByMonthRow
	for rows.Next() {
		var i GetBillByMonthRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.ContractID,
			&i.AdditionFee,
			&i.AdditionNote,
			&i.TotalAmount,
			&i.Month,
			&i.Year,
			&i.OldWaterIndex,
			&i.OldElectricityIndex,
			&i.NewWaterIndex,
			&i.NewElectricityIndex,
			&i.TotalWaterCost,
			&i.TotalElectricityCost,
			&i.Status,
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

const getBillByStatus = `-- name: GetBillByStatus :many
SELECT  b.id,
        b.code,
        b.contract_id,
        b.addition_fee,
        b.addition_note,
        b.total_amount,
        b.month,
        b.year,
        b.old_water_index, 
        b.old_electricity_index, 
        b.new_water_index, 
        b.new_electricity_index, 
        b.total_water_cost, 
        b.total_electricity_cost,
        b.status,
        b.created_at,
        b.updated_at
FROM PUBLIC.BILLING b
WHERE deleted_at IS NULL 
      AND status = $1
`

type GetBillByStatusRow struct {
	ID                   int32              `json:"id"`
	Code                 string             `json:"code"`
	ContractID           int32              `json:"contract_id"`
	AdditionFee          *int32             `json:"addition_fee"`
	AdditionNote         *string            `json:"addition_note"`
	TotalAmount          float64            `json:"total_amount"`
	Month                int32              `json:"month"`
	Year                 int32              `json:"year"`
	OldWaterIndex        *int32             `json:"old_water_index"`
	OldElectricityIndex  *int32             `json:"old_electricity_index"`
	NewWaterIndex        *int32             `json:"new_water_index"`
	NewElectricityIndex  *int32             `json:"new_electricity_index"`
	TotalWaterCost       *float64           `json:"total_water_cost"`
	TotalElectricityCost *float64           `json:"total_electricity_cost"`
	Status               *int32             `json:"status"`
	CreatedAt            pgtype.Timestamptz `json:"created_at"`
	UpdatedAt            pgtype.Timestamptz `json:"updated_at"`
}

func (q *Queries) GetBillByStatus(ctx context.Context, status *int32) ([]GetBillByStatusRow, error) {
	rows, err := q.db.Query(ctx, getBillByStatus, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBillByStatusRow
	for rows.Next() {
		var i GetBillByStatusRow
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.ContractID,
			&i.AdditionFee,
			&i.AdditionNote,
			&i.TotalAmount,
			&i.Month,
			&i.Year,
			&i.OldWaterIndex,
			&i.OldElectricityIndex,
			&i.NewWaterIndex,
			&i.NewElectricityIndex,
			&i.TotalWaterCost,
			&i.TotalElectricityCost,
			&i.Status,
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

const getBillOfRentedRoomByOwnerID = `-- name: GetBillOfRentedRoomByOwnerID :many

SELECT r.id AS room_id,
		r.room_number,
		b.id,
		b.status,
		b.total_amount,
		u.full_name
FROM rooms r 
LEFT JOIN contracts c ON r.id = c.room_id 
LEFT JOIN billing b ON c.id = b.contract_id 
RIGHT JOIN users u ON c.party_b = u.id 
WHERE r."owner" = $1
	AND r.is_rent = true
`

type GetBillOfRentedRoomByOwnerIDRow struct {
	RoomID      int32    `json:"room_id"`
	RoomNumber  int32    `json:"room_number"`
	ID          *int32   `json:"id"`
	Status      *int32   `json:"status"`
	TotalAmount *float64 `json:"total_amount"`
	FullName    string   `json:"full_name"`
}

func (q *Queries) GetBillOfRentedRoomByOwnerID(ctx context.Context, owner int32) ([]GetBillOfRentedRoomByOwnerIDRow, error) {
	rows, err := q.db.Query(ctx, getBillOfRentedRoomByOwnerID, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBillOfRentedRoomByOwnerIDRow
	for rows.Next() {
		var i GetBillOfRentedRoomByOwnerIDRow
		if err := rows.Scan(
			&i.RoomID,
			&i.RoomNumber,
			&i.ID,
			&i.Status,
			&i.TotalAmount,
			&i.FullName,
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
