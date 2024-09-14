-- name: CreateRentalRequest :one
INSERT INTO PUBLIC.RENTAL_REQUESTS
(
    code,
    sender_id,
    room_id,
    suggested_price,
    num_of_person,
    begin_date,
    end_date,
    addition_request,
    status,
    created_at,
    updated_at
) VALUES
(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, now(), now()
)
RETURNING code, sender_id, room_id, suggested_price, num_of_person, begin_date, end_date,     addition_request, status, created_at;

-- name: CheckRoomExisted :one
SELECT id 
FROM PUBLIC.ROOMS 
WHERE id = $1 ;

-- name: CheckRequestExisted :one
SELECT status 
FROM PUBLIC.RENTAL_REQUESTS 
WHERE room_id = $1 and sender_id = $2;