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

-- name: CreateContract :exec
INSERT INTO PUBLIC.contracts
(
    code,                   -- Mã hợp đồng
    party_a,                -- ID của bên A
    party_b,                -- ID của bên B
    request_id,             -- ID của yêu cầu liên quan đến hợp đồng
    room_id,                -- ID của phòng
    actual_price,           -- Giá thực tế
    payment_method,         -- Phương thức thanh toán
    electricity_method,     -- Phương thức tính tiền điện
    electricity_cost,       -- Chi phí điện
    water_method,           -- Phương thức tính tiền nước
    water_cost,             -- Chi phí nước
    internet_cost,          -- Chi phí Internet
    parking_fee,            -- Phí giữ xe
    deposit,                -- Tiền cọc
    begin_date,             -- Ngày bắt đầu
    end_date,               -- Ngày kết thúc
    responsibility_a,       -- Trách nhiệm của bên A
    responsibility_b,       -- Trách nhiệm của bên B
    general_responsibility, -- Trách nhiệm chung
    signature_a,            -- Chữ ký của bên A
    signed_time_a,          -- Thời gian ký của bên A
    contract_template_id,   -- ID mẫu hợp đồng
    created_at,             -- Thời gian tạo
    updated_at              -- Thời gian cập nhật
) VALUES
(
    $1,  -- Mã hợp đồng
    $2,  -- ID của bên A
    $3,  -- ID của bên B
    $4,  -- ID của yêu cầu
    $5,  -- ID của phòng
    $6,  -- Giá thực tế
    $7,  -- Phương thức thanh toán
    $8,  -- Phương thức tính tiền điện
    $9,  -- Chi phí điện
    $10, -- Phương thức tính tiền nước
    $11, -- Chi phí nước
    $12, -- Chi phí Internet
    $13, -- Phí giữ xe
    $14, -- Tiền cọc
    $15, -- Ngày bắt đầu
    $16, -- Ngày kết thúc
    $17, -- Trách nhiệm bên A
    $18, -- Trách nhiệm bên B
    $19, -- Trách nhiệm chung
    $20, -- Chữ ký bên A
    $21, -- Thời gian ký của bên A
    $22, -- ID mẫu hợp đồng
    NOW(), -- Thời gian tạo
    NOW()  -- Thời gian cập nhật
);

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


-- name: GetContractByID :one
SELECT id, code, party_a, party_b, request_id, room_id, actual_price, payment_method, electricity_method, electricity_cost, water_method, water_cost, internet_cost, parking_fee, deposit, begin_date, end_date, responsibility_a, responsibility_b, general_responsibility, signature_a, signed_time_a, signature_b, signed_time_b, created_at, updated_at, contract_template_id
FROM public.contracts
WHERE id = $1 AND deleted_at IS NULL;

-- name: ListContractByStatus :many
SELECT id, code, party_a, party_b, request_id, room_id, actual_price, payment_method, electricity_method, electricity_cost, water_method, water_cost, internet_cost, parking_fee, deposit, begin_date, end_date, responsibility_a, responsibility_b, general_responsibility, signature_a, signed_time_a, signature_b, signed_time_b, created_at, updated_at, contract_template_id
FROM public.contracts
WHERE status = $1 AND deleted_at IS NULL;

-- name: SignContract :exec
UPDATE public.contracts
SET signature_b = $2,
    signed_time_b = NOW(),
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL;

-- name: DeclineContract :exec
UPDATE public.contracts
SET status = 3,
    updated_at = NOW()
WHERE id = $1 AND deleted_at IS NULL;
