-- name: CreateContractTemplate :exec
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
RETURNING id;

-- name: CreateContract :one
INSERT INTO PUBLIC.contracts
(
    room_id,
    signature_a
) VALUES
(
    $1, $2
) RETURNING id;

-- name: GetContractTemplateByAddress :one
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
    AND deleted_at IS NULL;

-- name: ListContractIds :many
SELECT id
FROM public.contracts;

-- name: ListContractByRoomId :many
SELECT id
FROM public.contracts
WHERE room_id = $1;