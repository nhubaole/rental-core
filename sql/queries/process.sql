-- name: CreateProcessTracking :one
insert INTO PUBLIC.PROCESS_TRACKING(
  actor, action, issued_at, request_id
)
VALUES ($1, $2 , now(), $3)
RETURNING id, actor, action, issued_at, request_id;


-- name: GetProcessTrackingByRentalId :many
select * from PUBLIC.PROCESS_TRACKING
where request_id = $1;


-- name: GetAllProcessTracking :many
select * from PUBLIC.PROCESS_TRACKING
where actor = $1;