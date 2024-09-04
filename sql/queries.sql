-- name: GetUsers :many
SELECT * 
FROM PUBLIC.USERS;

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