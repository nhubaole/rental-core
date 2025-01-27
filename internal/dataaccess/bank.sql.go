// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: bank.sql

package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBank = `-- name: CreateBank :exec
INSERT INTO PUBLIC.BANKS
(
    id,          -- 1
    bank_name,   -- 2
    bank_code,   -- 3
    short_name,  -- 4
    logo,        -- 5
    created_at,  -- 6
    updated_at   -- 7
) VALUES
(
    $1, $2, $3, $4, $5, now(), now()
)
ON CONFLICT (id) DO UPDATE
SET 
    bank_name = EXCLUDED.bank_name,
    bank_code = EXCLUDED.bank_code,
    short_name = EXCLUDED.short_name,
    logo = EXCLUDED.logo,
    updated_at = now()
`

type CreateBankParams struct {
	ID        int32   `json:"id"`
	BankName  string  `json:"bank_name"`
	BankCode  string  `json:"bank_code"`
	ShortName *string `json:"short_name"`
	Logo      *string `json:"logo"`
}

func (q *Queries) CreateBank(ctx context.Context, arg CreateBankParams) error {
	_, err := q.db.Exec(ctx, createBank,
		arg.ID,
		arg.BankName,
		arg.BankCode,
		arg.ShortName,
		arg.Logo,
	)
	return err
}

const getAllBanks = `-- name: GetAllBanks :many
SELECT id,
       bank_name,
       bank_code,
       short_name,
       logo,
       created_at,
       updated_at
FROM PUBLIC.BANKS
`

type GetAllBanksRow struct {
	ID        int32            `json:"id"`
	BankName  string           `json:"bank_name"`
	BankCode  string           `json:"bank_code"`
	ShortName *string          `json:"short_name"`
	Logo      *string          `json:"logo"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) GetAllBanks(ctx context.Context) ([]GetAllBanksRow, error) {
	rows, err := q.db.Query(ctx, getAllBanks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetAllBanksRow
	for rows.Next() {
		var i GetAllBanksRow
		if err := rows.Scan(
			&i.ID,
			&i.BankName,
			&i.BankCode,
			&i.ShortName,
			&i.Logo,
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

const getBankByID = `-- name: GetBankByID :one
SELECT 
    id,
    bank_name,
    bank_code,
    short_name,
    logo,
    created_at,
    updated_at
FROM 
    PUBLIC.BANKS
WHERE 
    id = $1
`

type GetBankByIDRow struct {
	ID        int32            `json:"id"`
	BankName  string           `json:"bank_name"`
	BankCode  string           `json:"bank_code"`
	ShortName *string          `json:"short_name"`
	Logo      *string          `json:"logo"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
	UpdatedAt pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) GetBankByID(ctx context.Context, id int32) (GetBankByIDRow, error) {
	row := q.db.QueryRow(ctx, getBankByID, id)
	var i GetBankByIDRow
	err := row.Scan(
		&i.ID,
		&i.BankName,
		&i.BankCode,
		&i.ShortName,
		&i.Logo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
