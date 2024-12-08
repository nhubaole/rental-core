-- name: CreateReturnRequest :exec
INSERT INTO public.return_requests
    (contract_id, --1
     reason, --2
     return_date, --3 
     status, --4
     deduct_amount, --5 
     total_return_deposit, --6 
     created_user, --7
     created_at,
     updated_at)
VALUES(
    $1, $2, $3, $4, $5, $6, $7, now(), now()
    );

-- name: GetReturnRequestByID :one
SELECT id, contract_id, reason, return_date, status, deduct_amount, total_return_deposit, created_user, created_at, updated_at
FROM public.return_requests
WHERE deleted_at IS NULL
    AND id = $1;

-- name: ApproveReturnRequest :exec
UPDATE public.return_requests
SET status = 2,
    updated_at= now()
WHERE deleted_at IS NULL
    AND id = $1;

-- name: GetReturnRequestByLandlordID :many
SELECT rr.id, rr.contract_id, rr.reason, rr.return_date, rr.status, rr.deduct_amount, rr.total_return_deposit, rr.created_user, rr.created_at, rr.updated_at
FROM public.return_requests rr LEFT JOIN public.contracts c
ON rr.contract_id = c.id
LEFT JOIN public.rooms r ON c.room_id = r.id
WHERE r.owner = $1;