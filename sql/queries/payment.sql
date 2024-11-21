-- name: GetPaymentByID :one
SELECT id, code, sender_id, bill_id, contract_id, amount, status,return_request_id
FROM public.payments
WHERE id = $1;