// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO PUBLIC.USERS
(
    phone_number,
    full_name,
    address,
    password,
    role,
    otp,
    wallet_address,
    private_key_hex,
    created_at
    
) VALUES
(
    $1,$2,$3,$4,$5,$6, $7, $8, now()
)
`

type CreateUserParams struct {
	PhoneNumber   string  `json:"phone_number"`
	FullName      string  `json:"full_name"`
	Address       *string `json:"address"`
	Password      string  `json:"password"`
	Role          int32   `json:"role"`
	Otp           *int32  `json:"otp"`
	WalletAddress *string `json:"wallet_address"`
	PrivateKeyHex *string `json:"private_key_hex"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.PhoneNumber,
		arg.FullName,
		arg.Address,
		arg.Password,
		arg.Role,
		arg.Otp,
		arg.WalletAddress,
		arg.PrivateKeyHex,
	)
	return err
}

const getUserByID = `-- name: GetUserByID :one
SELECT 
    u.id,
    u.phone_number,
    u.avatar_url,
    u.role,
    u.full_name,
    u.address,
    u.wallet_address,
    u.private_key_hex,
    u.created_at
FROM PUBLIC.USERS u
WHERE 
    u.id = $1
    AND u.deleted_at IS NULL
`

type GetUserByIDRow struct {
	ID            int32              `json:"id"`
	PhoneNumber   string             `json:"phone_number"`
	AvatarUrl     *string            `json:"avatar_url"`
	Role          int32              `json:"role"`
	FullName      string             `json:"full_name"`
	Address       *string            `json:"address"`
	WalletAddress *string            `json:"wallet_address"`
	PrivateKeyHex *string            `json:"private_key_hex"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) GetUserByID(ctx context.Context, id int32) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i GetUserByIDRow
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.AvatarUrl,
		&i.Role,
		&i.FullName,
		&i.Address,
		&i.WalletAddress,
		&i.PrivateKeyHex,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByPhone = `-- name: GetUserByPhone :one
SELECT id, phone_number, avatar_url, role, full_name, password, address, wallet_address, private_key_hex, otp, created_at
FROM PUBLIC.USERS
WHERE deleted_at IS NULL 
    AND phone_number = $1
`

type GetUserByPhoneRow struct {
	ID            int32              `json:"id"`
	PhoneNumber   string             `json:"phone_number"`
	AvatarUrl     *string            `json:"avatar_url"`
	Role          int32              `json:"role"`
	FullName      string             `json:"full_name"`
	Password      string             `json:"password"`
	Address       *string            `json:"address"`
	WalletAddress *string            `json:"wallet_address"`
	PrivateKeyHex *string            `json:"private_key_hex"`
	Otp           *int32             `json:"otp"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) GetUserByPhone(ctx context.Context, phoneNumber string) (GetUserByPhoneRow, error) {
	row := q.db.QueryRow(ctx, getUserByPhone, phoneNumber)
	var i GetUserByPhoneRow
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.AvatarUrl,
		&i.Role,
		&i.FullName,
		&i.Password,
		&i.Address,
		&i.WalletAddress,
		&i.PrivateKeyHex,
		&i.Otp,
		&i.CreatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT id, phone_number,avatar_url,role, full_name, address, created_at 
FROM PUBLIC.USERS
WHERE deleted_at IS NULL
`

type GetUsersRow struct {
	ID          int32              `json:"id"`
	PhoneNumber string             `json:"phone_number"`
	AvatarUrl   *string            `json:"avatar_url"`
	Role        int32              `json:"role"`
	FullName    string             `json:"full_name"`
	Address     *string            `json:"address"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) GetUsers(ctx context.Context) ([]GetUsersRow, error) {
	rows, err := q.db.Query(ctx, getUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersRow
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.PhoneNumber,
			&i.AvatarUrl,
			&i.Role,
			&i.FullName,
			&i.Address,
			&i.CreatedAt,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    phone_number = COALESCE($2, phone_number),
    full_name = COALESCE($3, full_name),
    address = COALESCE($4, address),
    role = COALESCE($5, role),
    otp = $6
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, phone_number, full_name, address, role::text, created_at
`

type UpdateUserParams struct {
	ID          int32   `json:"id"`
	PhoneNumber string  `json:"phone_number"`
	FullName    string  `json:"full_name"`
	Address     *string `json:"address"`
	Role        int32   `json:"role"`
	Otp         *int32  `json:"otp"`
}

type UpdateUserRow struct {
	ID          int32              `json:"id"`
	PhoneNumber string             `json:"phone_number"`
	FullName    string             `json:"full_name"`
	Address     *string            `json:"address"`
	Role        string             `json:"role"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.PhoneNumber,
		arg.FullName,
		arg.Address,
		arg.Role,
		arg.Otp,
	)
	var i UpdateUserRow
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.FullName,
		&i.Address,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}
