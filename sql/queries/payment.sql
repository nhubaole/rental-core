-- name: GetPaymentByID :one
SELECT 
    p.id, 
    p.code, 
    p.sender_id, 
    u.full_name AS sender_name, 
    p.bill_id, 
    p.contract_id, 
    p.amount, 
    p.status, 
    p.return_request_id, 
    p.paid_time, 
    p.transfer_content, 
    p.evidence_image
FROM 
    public.payments p
JOIN 
    public.users u ON p.sender_id = u.id
WHERE 
    p.id = $1;

-- name: CreatePayment :one
INSERT INTO public.payments(
    code, --1
    sender_id, --2
    bill_id, --3
    contract_id, --4
    amount, 
    status, 
    return_request_id, 
    transfer_content, 
    evidence_image, 
    paid_time
    )VALUES(
        $1, $2, $3, $4, $5, $6, $7, $8, $9, now())
    RETURNING id;

-- name: GetAllPayments :many
SELECT id, code, sender_id, bill_id, contract_id, amount, status, return_request_id, transfer_content, evidence_image, paid_time
FROM public.payments;

-- name: ConfirmPayment :one
UPDATE public.payments
SET status = 1
WHERE id = $1
RETURNING id;

