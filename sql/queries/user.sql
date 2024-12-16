-- name: GetUsers :many
SELECT id, phone_number,avatar_url,role, full_name, address, created_at 
FROM PUBLIC.USERS
WHERE deleted_at IS NULL;

-- name: CreateUser :exec
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
);

-- name: GetUserByPhone :one
SELECT id, phone_number, avatar_url, role, full_name, password, address, wallet_address, private_key_hex, otp, created_at
FROM PUBLIC.USERS
WHERE deleted_at IS NULL 
    AND phone_number = $1;

-- name: GetUserByID :one
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
    AND u.deleted_at IS NULL;


-- name: UpdateUser :one
UPDATE users
SET
    phone_number = COALESCE($2, phone_number),
    full_name = COALESCE($3, full_name),
    address = COALESCE($4, address),
    role = COALESCE($5, role),
    otp = $6
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, phone_number, full_name, address, role::text, created_at;

-- name: UpdateDeviceToken :exec
UPDATE PUBLIC.USERS
SET
    device_token = $2
WHERE
    id = $1
    AND deleted_at IS NULL;

-- name: GetDeviceTokenByUserID :one
SELECT
    device_token
FROM
    PUBLIC.USERS
WHERE
    id = $1
    AND deleted_at IS NULL;