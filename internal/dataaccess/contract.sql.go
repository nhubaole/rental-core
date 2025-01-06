// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: contract.sql

package dataaccess

import (
	"context"
)

const createContract = `-- name: CreateContract :one
INSERT INTO PUBLIC.contracts
(
    room_id,
    signature_a
) VALUES
(
    $1, $2
) RETURNING id
`

type CreateContractParams struct {
	RoomID     *int32  `json:"room_id"`
	SignatureA *string `json:"signature_a"`
}

func (q *Queries) CreateContract(ctx context.Context, arg CreateContractParams) (int32, error) {
	row := q.db.QueryRow(ctx, createContract, arg.RoomID, arg.SignatureA)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createContractTemplate = `-- name: CreateContractTemplate :exec
INSERT INTO PUBLIC.contract_templates
(
    party_a,
    address,
    electricity_method,
    electricity_cost,
    water_method,
    water_cost,
    internet_cost,
    parking_fee,
    responsibility_a,
    responsibility_b,
    general_responsibility,
    created_at,
    updated_at
) VALUES
(
    $1, -- ID của bên A
    $2, -- Địa chỉ (mảng chuỗi)
    $3, -- Phương thức tính tiền điện
    $4, -- Chi phí điện
    $5, -- Phương thức tính tiền nước
    $6, -- Chi phí nước
    $7, -- Chi phí Internet
    $8, -- Phí giữ xe
    $9, -- Trách nhiệm bên A
    $10, -- Trách nhiệm bên B
    $11, -- Trách nhiệm chung
    NOW(), -- Thời gian tạo
    NOW()  -- Thời gian cập nhật
)
RETURNING id
`

type CreateContractTemplateParams struct {
	PartyA                int32    `json:"party_a"`
	Address               []string `json:"address"`
	ElectricityMethod     string   `json:"electricity_method"`
	ElectricityCost       float64  `json:"electricity_cost"`
	WaterMethod           string   `json:"water_method"`
	WaterCost             float64  `json:"water_cost"`
	InternetCost          float64  `json:"internet_cost"`
	ParkingFee            float64  `json:"parking_fee"`
	ResponsibilityA       string   `json:"responsibility_a"`
	ResponsibilityB       string   `json:"responsibility_b"`
	GeneralResponsibility string   `json:"general_responsibility"`
}

func (q *Queries) CreateContractTemplate(ctx context.Context, arg CreateContractTemplateParams) error {
	_, err := q.db.Exec(ctx, createContractTemplate,
		arg.PartyA,
		arg.Address,
		arg.ElectricityMethod,
		arg.ElectricityCost,
		arg.WaterMethod,
		arg.WaterCost,
		arg.InternetCost,
		arg.ParkingFee,
		arg.ResponsibilityA,
		arg.ResponsibilityB,
		arg.GeneralResponsibility,
	)
	return err
}

const getContractById = `-- name: GetContractById :one
SELECT id, room_id, signature_a, signature_b from public.contracts
WHERE id = $1
`

func (q *Queries) GetContractById(ctx context.Context, id int32) (Contract, error) {
	row := q.db.QueryRow(ctx, getContractById, id)
	var i Contract
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.SignatureA,
		&i.SignatureB,
	)
	return i, err
}

const getContractTemplateByAddress = `-- name: GetContractTemplateByAddress :one
SELECT 
    id,
    party_a,
    address,
    electricity_method,
    electricity_cost,
    water_method,
    water_cost,
    internet_cost,
    parking_fee,
    responsibility_a,
    responsibility_b,
    general_responsibility,
    created_at,
    updated_at,
    deleted_at
FROM 
    public.contract_templates
WHERE 
    address = $1::varchar[]
    AND deleted_at IS NULL
`

func (q *Queries) GetContractTemplateByAddress(ctx context.Context, dollar_1 []string) (ContractTemplate, error) {
	row := q.db.QueryRow(ctx, getContractTemplateByAddress, dollar_1)
	var i ContractTemplate
	err := row.Scan(
		&i.ID,
		&i.PartyA,
		&i.Address,
		&i.ElectricityMethod,
		&i.ElectricityCost,
		&i.WaterMethod,
		&i.WaterCost,
		&i.InternetCost,
		&i.ParkingFee,
		&i.ResponsibilityA,
		&i.ResponsibilityB,
		&i.GeneralResponsibility,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const listContractByRoomId = `-- name: ListContractByRoomId :many
SELECT id
FROM public.contracts
WHERE room_id = $1
`

func (q *Queries) ListContractByRoomId(ctx context.Context, roomID *int32) ([]int32, error) {
	rows, err := q.db.Query(ctx, listContractByRoomId, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listContractIds = `-- name: ListContractIds :many
SELECT id
FROM public.contracts
`

func (q *Queries) ListContractIds(ctx context.Context) ([]int32, error) {
	rows, err := q.db.Query(ctx, listContractIds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateSignatureB = `-- name: UpdateSignatureB :exec
UPDATE public.contracts
SET signature_b = $2
WHERE id = $1
`

type UpdateSignatureBParams struct {
	ID         int32   `json:"id"`
	SignatureB *string `json:"signature_b"`
}

func (q *Queries) UpdateSignatureB(ctx context.Context, arg UpdateSignatureBParams) error {
	_, err := q.db.Exec(ctx, updateSignatureB, arg.ID, arg.SignatureB)
	return err
}
