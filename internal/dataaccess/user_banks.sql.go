// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user_banks.sql

package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUserBank = `-- name: CreateUserBank :exec
INSERT INTO user_banks (
    user_id, 
    bank_id, 
    account_number, 
    account_name, 
    card_number, 
    currency, 
    created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, now()
)
`

type CreateUserBankParams struct {
	UserID        int32   `json:"user_id"`
	BankID        int32   `json:"bank_id"`
	AccountNumber string  `json:"account_number"`
	AccountName   string  `json:"account_name"`
	CardNumber    *string `json:"card_number"`
	Currency      *string `json:"currency"`
}

func (q *Queries) CreateUserBank(ctx context.Context, arg CreateUserBankParams) error {
	_, err := q.db.Exec(ctx, createUserBank,
		arg.UserID,
		arg.BankID,
		arg.AccountNumber,
		arg.AccountName,
		arg.CardNumber,
		arg.Currency,
	)
	return err
}

const getBankInfoByUserID = `-- name: GetBankInfoByUserID :one
SELECT 
    user_id, 
    bank_id,
    b.bank_name,
    b.short_name, 
    account_number, 
    account_name, 
    card_number, 
    currency, 
    ub.created_at, 
    ub.updated_at
FROM 
    user_banks ub
LEFT JOIN banks b ON ub.bank_id = b.id
WHERE 
    user_id = $1
`

type GetBankInfoByUserIDRow struct {
	UserID        int32            `json:"user_id"`
	BankID        int32            `json:"bank_id"`
	BankName      *string          `json:"bank_name"`
	ShortName     *string          `json:"short_name"`
	AccountNumber string           `json:"account_number"`
	AccountName   string           `json:"account_name"`
	CardNumber    *string          `json:"card_number"`
	Currency      *string          `json:"currency"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
	UpdatedAt     pgtype.Timestamp `json:"updated_at"`
}

func (q *Queries) GetBankInfoByUserID(ctx context.Context, userID int32) (GetBankInfoByUserIDRow, error) {
	row := q.db.QueryRow(ctx, getBankInfoByUserID, userID)
	var i GetBankInfoByUserIDRow
	err := row.Scan(
		&i.UserID,
		&i.BankID,
		&i.BankName,
		&i.ShortName,
		&i.AccountNumber,
		&i.AccountName,
		&i.CardNumber,
		&i.Currency,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserBank = `-- name: UpdateUserBank :one
UPDATE user_banks
SET
    account_number = COALESCE($2, account_number),
    account_name = COALESCE($3, account_name),
    card_number = COALESCE($4, card_number),
    currency = COALESCE($5, currency),
    bank_id = COALESCE($6, bank_id),
    updated_at = now()
WHERE 
    user_id = $1
RETURNING 
    user_id, 
    bank_id, 
    account_number, 
    account_name, 
    card_number, 
    currency, 
    created_at, 
    updated_at
`

type UpdateUserBankParams struct {
	UserID        int32   `json:"user_id"`
	AccountNumber string  `json:"account_number"`
	AccountName   string  `json:"account_name"`
	CardNumber    *string `json:"card_number"`
	Currency      *string `json:"currency"`
	BankID        int32   `json:"bank_id"`
}

func (q *Queries) UpdateUserBank(ctx context.Context, arg UpdateUserBankParams) (UserBank, error) {
	row := q.db.QueryRow(ctx, updateUserBank,
		arg.UserID,
		arg.AccountNumber,
		arg.AccountName,
		arg.CardNumber,
		arg.Currency,
		arg.BankID,
	)
	var i UserBank
	err := row.Scan(
		&i.UserID,
		&i.BankID,
		&i.AccountNumber,
		&i.AccountName,
		&i.CardNumber,
		&i.Currency,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
