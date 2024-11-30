-- name: CreateBank :exec
INSERT INTO PUBLIC.BANKS
(
    id,          -- 1
    bank_name,   -- 2
    bank_code,   -- 3
    short_name,  -- 4
    logo,        -- 5
    created_at,  -- 6
    updated_at   -- 7
) VALUES
(
    $1, $2, $3, $4, $5, now(), now()
)
ON CONFLICT (id) DO UPDATE
SET 
    bank_name = EXCLUDED.bank_name,
    bank_code = EXCLUDED.bank_code,
    short_name = EXCLUDED.short_name,
    logo = EXCLUDED.logo,
    updated_at = now();

-- name: GetAllBanks :many
SELECT id,
       bank_name,
       bank_code,
       short_name,
       logo,
       created_at,
       updated_at
FROM PUBLIC.BANKS;
