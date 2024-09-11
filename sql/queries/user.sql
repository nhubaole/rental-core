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
