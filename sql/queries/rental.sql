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
RETURNING id, code, sender_id, room_id, suggested_price, num_of_person, begin_date, end_date,     addition_request, status, created_at;

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
WHERE id = $1 and deleted_at is null;


-- name: GetRequestBySenderID :many
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
FROM PUBLIC.RENTAL_REQUESTS  RR
WHERE sender_id = $1
    AND RR.deleted_at is NULL;

-- name: GetRequestByOwnerID :many
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
FROM PUBLIC.RENTAL_REQUESTS  RR left join PUBLIC.ROOMS r
	on RR.room_id = r.id
WHERE r.owner = $1 
	and RR.deleted_at is NULL ;

-- name: UpdateRequestStatusById :exec
update PUBLIC.RENTAL_REQUESTS
set status = $1 WHERE id = $2 and status = 1;


-- name: GetRentalRequestSuccessByRoomId :one
SELECT *
FROM public.rental_requests
WHERE room_id = $1
  AND status = 2
  AND begin_date <= now()
  AND end_date >= now();

-- name: GetRequestByRoomID :many
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
FROM PUBLIC.RENTAL_REQUESTS  RR 
WHERE RR.room_id = $1
	and RR.deleted_at is NULL ;