-- name: CreateBill :one
INSERT INTO PUBLIC.BILLING
(
    code,
    contract_id,
    index_id,
    addition_fee,
    addition_note,
    total_amount,
    month,
    year,
    paid_time,
    created_at,
    updated_at
) VALUES
(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, now(), now()
)
RETURNING id, code, contract_id, index_id, addition_fee, addition_note, total_amount, month, year, paid_time, created_at;


-- name: GetBillByMonth :many
SELECT *
FROM PUBLIC.BILLING as bi
left join public.contracts as ct on ct.id = bi.contract_id
WHERE year = $1 and month=$2 and (party_a = $3 or party_b = $3);

