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
    otp,
    created_at
) VALUES
(
    $1,$2,$3,$4,$5,$6,now()
);

-- name: GetUserByPhone :one
SELECT id, phone_number, password, role, full_name, address, otp, created_at
FROM PUBLIC.USERS
WHERE deleted_at IS NULL 
    AND phone_number = $1;
    