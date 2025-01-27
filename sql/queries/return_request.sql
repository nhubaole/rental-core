-- name: CreateReturnRequest :one
INSERT INTO public.return_requests
    (contract_id, --1
     reason, --2
     return_date, --3 
     deduct_amount, --5 
     total_return_deposit, --6 
     created_user, --7
     status, --4
     created_at,
     updated_at)
VALUES(
    $1, $2, $3, $4, $5, $6, 0, now(), now()
    ) RETURNING id;

-- name: GetReturnRequestByID :one
SELECT rr.id, rr.reason,c.id as contract_id, c.room_id, rr.return_date, rr.status, rr.deduct_amount, rr.total_return_deposit, rr.created_user, rr.created_at, rr.updated_at
FROM public.return_requests rr 
LEFT JOIN public.contracts c
ON rr.contract_id = c.id
WHERE deleted_at IS NULL
    AND rr.id = $1;

-- name: ApproveReturnRequest :exec
UPDATE public.return_requests
SET status = $1,
    updated_at= now()
WHERE deleted_at IS NULL
    AND id = $2;

-- name: GetReturnRequestByLandlordID :many
SELECT rr.id, rr.contract_id, r.id as room_id, rr.reason, rr.return_date, rr.status, rr.deduct_amount, rr.total_return_deposit, rr.created_user, rr.created_at, rr.updated_at
FROM public.return_requests rr LEFT JOIN public.contracts c
ON rr.contract_id = c.id
LEFT JOIN public.rooms r ON c.room_id = r.id
WHERE r.owner = $1;

-- name: GetReturnRequestByTenantID :many
SELECT rr.id, rr.contract_id, r.id as room_id, rr.reason, rr.return_date, rr.status, rr.deduct_amount, rr.total_return_deposit, rr.created_user, rr.created_at, rr.updated_at
FROM public.return_requests rr LEFT JOIN public.contracts c
ON rr.contract_id = c.id
LEFT JOIN public.rooms r ON c.room_id = r.id
WHERE rr.created_user = $1;