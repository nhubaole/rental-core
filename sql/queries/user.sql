-- name: GetUsers :many
SELECT phone_number, full_name, address, created_at 
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
    created_at
) VALUES
(
    $1,$2,$3,$4,$5,now()
);

-- name: GetUserByPhone :one
SELECT phone_number, full_name, address, created_at
FROM PUBLIC.USERS
WHERE deleted_at IS NULL 
    AND phone_number = $1;

-- name: GetUserByID :one
SELECT id, phone_number, full_name, address, created_at
FROM PUBLIC.USERS
WHERE id = $1 AND deleted_at IS NULL;

-- name: UpdateUser :one
UPDATE users
SET
    phone_number = COALESCE($2, phone_number),
    full_name = COALESCE($3, full_name),
    address = COALESCE($4, address),
    role = COALESCE($5, role)
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, phone_number, full_name, address, role::text, created_at;
