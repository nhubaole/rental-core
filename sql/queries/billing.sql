-- name: CreateBill :one
INSERT INTO PUBLIC.BILLING
(
    code,
    contract_id, --1
    old_water_index, --2
    old_electricity_index, --3
    new_water_index, --4
    new_electricity_index, --5
    total_water_cost, --6
    total_electricity_cost, --7
    addition_fee,  --8
    addition_note, --9
    total_amount, --10
    month, --11
    year, --12
    created_at,  --13
    updated_at, --15
    status
) VALUES
(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,  now(), now(), 0
) RETURNING id;


-- name: GetBillByMonth :many
SELECT 
    r.address,
    b.id AS bill_id,
    COALESCE(b.status, -1) AS bill_status,
    r.room_number,
    c.id as contract_id,
    b.payment_id,
    b.total_amount,
    b.created_at AS bill_created_at
FROM 
    public.rooms r
LEFT JOIN 
    public.contracts c ON c.room_id = r.id
LEFT JOIN 
    public.billing b ON c.id = b.contract_id AND b.year = $1 AND b.month = $2
WHERE 
    r.owner = $3;


-- name: GetBillByID :one
SELECT  b.id,
        b.code,
        b.contract_id,
        b.addition_fee,
        b.addition_note,
        b.total_amount,
        b.payment_id,
        b.month,
        b.year,
        b.old_water_index, 
        b.old_electricity_index, 
        b.new_water_index, 
        b.new_electricity_index, 
        b.total_water_cost, 
        b.total_electricity_cost,
        b.status,
        b.created_at,
        b.updated_at
FROM PUBLIC.BILLING b
WHERE deleted_at IS NULL 
      AND id = $1;

-- name: GetAllMetric4BillByRoomID :one
SELECT t.room_id,
       t.prev_month,
       t.curr_month,
       t.prev_water,
       t.curr_water,
       t.prev_electricity,
       t.curr_electricity, 
       t.year
FROM (
	SELECT id , LAG(i.month) OVER(ORDER BY month, year) AS prev_month , month as curr_month,
	LAG(i.water_index) OVER(ORDER BY month, year) as prev_water , water_index as curr_water, 
	LAG(i.electricity_index) OVER(ORDER BY month,year) as prev_electricity, electricity_index as curr_electricity, 
	room_id, year
	FROM public.index as i
) AS t
LEFT JOIN public.INDEX idx ON t.id = idx.id
WHERE idx.room_id = $1
AND idx.month = $2
AND idx.year = $3;

-- name: GetBillByStatus :many
SELECT  b.id,
        b.code,
        r.address,
        r.room_number,
        b.contract_id,
        b.addition_fee,
        b.addition_note,
        b.total_amount,
        b.month,
        b.year,
        b.old_water_index, 
        b.old_electricity_index, 
        b.new_water_index, 
        b.new_electricity_index, 
        b.total_water_cost, 
        b.total_electricity_cost,
        b.status,
        b.created_at,
        b.updated_at,
        (b.updated_at + interval '10 days')::timestamp AS deadline
FROM PUBLIC.BILLING b
LEFT JOIN public.contracts c ON b.contract_id = c.id
LEFT JOIN public.rooms r ON c.room_id = r.id
WHERE b.deleted_at IS NULL 
      AND b.status = $1;

-- name: GetBillOfRentedRoomByOwnerID :many

SELECT r.id AS room_id,
		r.room_number,
		b.id,
		b.status,
		b.total_amount,
		u.full_name
FROM rooms r 
LEFT JOIN contracts c ON r.id = c.room_id 
LEFT JOIN billing b ON c.id = b.contract_id 
RIGHT JOIN users u ON c.party_b = u.id 
WHERE r."owner" = $1
	AND r.is_rent = true;

-- name: UpdatePaymentIDByBillID :exec
UPDATE public.billing
SET payment_id = $1,
    updated_at = now()
WHERE id = $2
  AND deleted_at IS NULL;

-- name: UpdateBillStatus :exec
UPDATE public.billing
SET status = $1,
    updated_at = now()
WHERE id = $2
  AND deleted_at IS NULL;