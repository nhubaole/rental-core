-- name: UpsertIndex :one
INSERT INTO PUBLIC.INDEX(
  water_index,
  electricity_index,
  room_id,
  month,
  year
)
VALUES(
    $1, $2, $3, $4, $5
  )
ON CONFLICT (room_id, month, year)
DO UPDATE
SET
  water_index = COALESCE(EXCLUDED.water_index, PUBLIC.INDEX.water_index),
  electricity_index = COALESCE(EXCLUDED.electricity_index, PUBLIC.INDEX.electricity_index)
RETURNING id, water_index, electricity_index, room_id, month, year;

-- name: GetIndexById :one
SELECT * 
FROM PUBLIC.INDEX
WHERE id = $1;

-- name: GetIndexByRoomId :many
SELECT *
FROM PUBLIC.INDEX
WHERE room_id = $1;

-- name: GetIndexByOwnerId :many
SELECT ro.id as room_id, t.id, t.prev_month, t.curr_month, t.prev_water, t.curr_water,t.prev_electricity,t.curr_electricity, t.year
FROM (
	SELECT id , LAG(i.month) OVER(ORDER BY month) AS prev_month , month as curr_month,
	LAG(i.water_index) OVER(ORDER BY month) as prev_water , water_index as curr_water, 
	LAG(i.electricity_index) OVER(ORDER BY month) as prev_electricity , electricity_index as curr_electricity, 
	room_id, year
	FROM public.index as i
) AS t
LEFT JOIN PUBLIC.ROOMS AS ro ON t.room_id = ro.id
LEFT JOIN public.INDEX idx ON t.id = idx.id
WHERE  ro.owner = $1
AND idx.month = $2
AND idx.year = $3;


-- name: GetIndexByOwnerIdShort :many
SELECT idx.id, idx.room_id, idx.water_index, idx.electricity_index, idx.month, idx.year
FROM  PUBLIC.INDEX AS idx 
LEFT JOIN public.ROOMS ro ON ro.id = idx.room_id
WHERE  ro.owner = $1;
