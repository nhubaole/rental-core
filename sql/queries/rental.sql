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


-- name: GetRequestBySenderID :many
SELECT *
FROM PUBLIC.RENTAL_REQUESTS 
WHERE sender_id = $1;

-- name: GetRequestByUserID :many
SELECT     
    RR.id,
    RR.code,
    RR.sender_id,
    RR.room_id,
    RR.suggested_price,
    RR.num_of_person,
    RR.begin_date,
    RR.end_date,
    RR.addition_request,
    RR.status,
    RR.created_at,
    RR.updated_at
FROM PUBLIC.RENTAL_REQUESTS  RR left join PUBLIC.ROOMS
	on RR.room_id = ROOMS.id
WHERE (owner = $1   or sender_id = $1) 
	and RR.deleted_at is NULL ;

-- name: UpdateRequestStatusById :exec
update PUBLIC.RENTAL_REQUESTS
set status = $1 WHERE id = $2 and status = 1;
