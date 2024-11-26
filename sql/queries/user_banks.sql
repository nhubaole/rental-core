-- name: CreateUserBank :exec
INSERT INTO user_banks (
    user_id, 
    bank_id, 
    account_number, 
    account_name, 
    card_number, 
    currency, 
    created_at
) VALUES (
    $1, $2, $3, $4, $5, $6, now()
);

-- name: UpdateUserBank :one
UPDATE user_banks
SET
    account_number = COALESCE($2, account_number),
    account_name = COALESCE($3, account_name),
    card_number = COALESCE($4, card_number),
    currency = COALESCE($5, currency),
    bank_id = COALESCE($6, bank_id),
    updated_at = now()
WHERE 
    user_id = $1
RETURNING 
    user_id, 
    bank_id, 
    account_number, 
    account_name, 
    card_number, 
    currency, 
    created_at, 
    updated_at;

-- name: GetBankInfoByUserID :one
SELECT