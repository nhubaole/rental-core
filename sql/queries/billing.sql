-- name: CreateBill :exec
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
    updated_at --15
) VALUES
(
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,  now(), now()
);


-- name: GetBillByMonth :many
SELECT b.code,
        b.contract_id,
        b.addition_fee,
        b.addition_note,
        b.total_amount,
        b.month,
        b.year,
        b.created_at,
        b.updated_at
FROM PUBLIC.BILLING as b
WHERE b.year = $1 
    AND b.month=$2;

-- name: GetBillByID :one
SELECT  code,
        contract_id,
        addition_fee,
        addition_note,
        total_amount,
        month,
        year,
        status,
        created_at,
        updated_at
FROM PUBLIC.BILLING
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
SELECT  code,
        contract_id,
        addition_fee,
        addition_note,
        total_amount,
        month,
        year,
        status,
        created_at,
        updated_at
FROM PUBLIC.BILLING
WHERE deleted_at IS NULL 
      AND status = $1;

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