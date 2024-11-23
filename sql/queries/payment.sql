-- name: GetPaymentByID :one
SELECT id, code, sender_id, bill_id, contract_id, amount, status,return_request_id
FROM public.payments
WHERE id = $1;

-- name: CreatePayment :exec
INSERT INTO public.payments(
    code, 
    sender_id,
    bill_id,
    contract_id,
    amount, 
    status, 
    return_request_id, 
    transfer_content, 
    evidence_image, 
    paid_time
    )VALUES(
        $1, $2, $3, $4, $5, $6, $7, $8, $9, now());
