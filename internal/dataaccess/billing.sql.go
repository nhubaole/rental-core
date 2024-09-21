// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: billing.sql

package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBill = `-- name: CreateBill :one
INSERT INTO PUBLIC.BILLING
(
    code,
    contract_id,
    index_id,
    addition_fee,
    addition_note,
    total_amount,
    month,
    year,
    paid_time,
    created_at,
    updated_at
) VALUES
(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, now(), now()
)
RETURNING id, code, contract_id, index_id, addition_fee, addition_note, total_amount, month, year, paid_time, created_at
`

type CreateBillParams struct {
	Code         string             `json:"code"`
	ContractID   int32              `json:"contract_id"`
	IndexID      int32              `json:"index_id"`
	AdditionFee  *int32             `json:"addition_fee"`
	AdditionNote *string            `json:"addition_note"`
	TotalAmount  float64            `json:"total_amount"`
	Month        int32              `json:"month"`
	Year         int32              `json:"year"`
	PaidTime     pgtype.Timestamptz `json:"paid_time"`
}

type CreateBillRow struct {
	ID           int32              `json:"id"`
	Code         string             `json:"code"`
	ContractID   int32              `json:"contract_id"`
	IndexID      int32              `json:"index_id"`
	AdditionFee  *int32             `json:"addition_fee"`
	AdditionNote *string            `json:"addition_note"`
	TotalAmount  float64            `json:"total_amount"`
	Month        int32              `json:"month"`
	Year         int32              `json:"year"`
	PaidTime     pgtype.Timestamptz `json:"paid_time"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) CreateBill(ctx context.Context, arg CreateBillParams) (CreateBillRow, error) {
	row := q.db.QueryRow(ctx, createBill,
		arg.Code,
		arg.ContractID,
		arg.IndexID,
		arg.AdditionFee,
		arg.AdditionNote,
		arg.TotalAmount,
		arg.Month,
		arg.Year,
		arg.PaidTime,
	)
	var i CreateBillRow
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.ContractID,
		&i.IndexID,
		&i.AdditionFee,
		&i.AdditionNote,
		&i.TotalAmount,
		&i.Month,
		&i.Year,
		&i.PaidTime,
		&i.CreatedAt,
	)
	return i, err
}

const getBillByMonth = `-- name: GetBillByMonth :many
SELECT bi.id, bi.code, contract_id, index_id, addition_fee, addition_note, total_amount, month, year, paid_time, bi.created_at, bi.updated_at, bi.deleted_at, ct.id, ct.code, party_a, party_b, request_id, room_id, actual_price, payment_method, electricity_method, electricity_cost, water_method, water_cost, internet_cost, parking_fee, deposit, begin_date, end_date, responsibility_a, responsibility_b, general_responsibility, signature_a, signed_time_a, signature_b, signed_time_b, ct.created_at, ct.updated_at, ct.deleted_at
FROM PUBLIC.BILLING as bi
left join public.contracts as ct on ct.id = bi.contract_id
WHERE year = $1 and month=$2 and (party_a = $3 or party_b = $3)
`

type GetBillByMonthParams struct {
	Year   int32 `json:"year"`
	Month  int32 `json:"month"`
	PartyA int32 `json:"party_a"`
}

type GetBillByMonthRow struct {
	ID                    int32              `json:"id"`
	Code                  string             `json:"code"`
	ContractID            int32              `json:"contract_id"`
	IndexID               int32              `json:"index_id"`
	AdditionFee           *int32             `json:"addition_fee"`
	AdditionNote          *string            `json:"addition_note"`
	TotalAmount           float64            `json:"total_amount"`
	Month                 int32              `json:"month"`
	Year                  int32              `json:"year"`
	PaidTime              pgtype.Timestamptz `json:"paid_time"`
	CreatedAt             pgtype.Timestamptz `json:"created_at"`
	UpdatedAt             pgtype.Timestamptz `json:"updated_at"`
	DeletedAt             pgtype.Timestamptz `json:"deleted_at"`
	ID_2                  *int32             `json:"id_2"`
	Code_2                *string            `json:"code_2"`
	PartyA                *int32             `json:"party_a"`
	PartyB                *int32             `json:"party_b"`
	RequestID             *int32             `json:"request_id"`
	RoomID                *int32             `json:"room_id"`
	ActualPrice           *float64           `json:"actual_price"`
	PaymentMethod         *string            `json:"payment_method"`
	ElectricityMethod     *string            `json:"electricity_method"`
	ElectricityCost       *float64           `json:"electricity_cost"`
	WaterMethod           *string            `json:"water_method"`
	WaterCost             *float64           `json:"water_cost"`
	InternetCost          *float64           `json:"internet_cost"`
	ParkingFee            *float64           `json:"parking_fee"`
	Deposit               *float64           `json:"deposit"`
	BeginDate             pgtype.Date        `json:"begin_date"`
	EndDate               pgtype.Date        `json:"end_date"`
	ResponsibilityA       *string            `json:"responsibility_a"`
	ResponsibilityB       *string            `json:"responsibility_b"`
	GeneralResponsibility *string            `json:"general_responsibility"`
	SignatureA            *string            `json:"signature_a"`
	SignedTimeA           pgtype.Timestamptz `json:"signed_time_a"`
	SignatureB            *string            `json:"signature_b"`
	SignedTimeB           pgtype.Timestamptz `json:"signed_time_b"`
	CreatedAt_2           pgtype.Timestamptz `json:"created_at_2"`
	UpdatedAt_2           pgtype.Timestamptz `json:"updated_at_2"`
	DeletedAt_2           pgtype.Timestamp   `json:"deleted_at_2"`
}

func (q *Queries) GetBillByMonth(ctx context.Context, arg GetBillByMonthParams) ([]GetBillByMonthRow, error) {
	rows, err := q.db.Query(ctx, getBillByMonth, arg.Year, arg.Month, arg.PartyA)
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
			&i.IndexID,
			&i.AdditionFee,
			&i.AdditionNote,
			&i.TotalAmount,
			&i.Month,
			&i.Year,
			&i.PaidTime,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.DeletedAt,
			&i.ID_2,
			&i.Code_2,
			&i.PartyA,
			&i.PartyB,
			&i.RequestID,
			&i.RoomID,
			&i.ActualPrice,
			&i.PaymentMethod,
			&i.ElectricityMethod,
			&i.ElectricityCost,
			&i.WaterMethod,
			&i.WaterCost,
			&i.InternetCost,
			&i.ParkingFee,
			&i.Deposit,
			&i.BeginDate,
			&i.EndDate,
			&i.ResponsibilityA,
			&i.ResponsibilityB,
			&i.GeneralResponsibility,
			&i.SignatureA,
			&i.SignedTimeA,
			&i.SignatureB,
			&i.SignedTimeB,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.DeletedAt_2,
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
