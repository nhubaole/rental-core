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
SELECT status , deleted_at
FROM PUBLIC.RENTAL_REQUESTS 
WHERE room_id = $1 and sender_id = $2;


-- name: DeleteRequest :exec
UPDATE public.rental_requests
SET deleted_at = now()
WHERE id = $1 ;

-- name: GetRequestByID :one
SELECT *
FROM PUBLIC.RENTAL_REQUESTS 
WHERE room_id = $1;


-- name: GetRequestBySenderID :one
SELECT *
FROM PUBLIC.RENTAL_REQUESTS 
WHERE sender_id = $1;

-- name: GetRequestByOwnerID :one
SELECT     
    RENTAL_REQUESTS.id,
    RENTAL_REQUESTS.code,
    RENTAL_REQUESTS.sender_id,
    RENTAL_REQUESTS.room_id,
    RENTAL_REQUESTS.suggested_price,
    RENTAL_REQUESTS.num_of_person,
    RENTAL_REQUESTS.begin_date,
    RENTAL_REQUESTS.end_date,
    RENTAL_REQUESTS.addition_request,
    RENTAL_REQUESTS.status,
    RENTAL_REQUESTS.created_at,
    RENTAL_REQUESTS.updated_at
FROM PUBLIC.RENTAL_REQUESTS ,PUBLIC.ROOMS 
WHERE owner = $1 and RENTAL_REQUESTS.room_id =ROOMS.id and RENTAL_REQUESTS.deleted_at != NULL ;