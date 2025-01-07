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

-- name: GetAllIndex :many
SELECT 
    ro.id AS room_id, 
    t.id, 
    t.prev_month, 
    t.curr_month, 
    t.prev_water, 
    t.curr_water, 
    t.prev_electricity, 
    t.curr_electricity, 
    t.year
FROM (
    SELECT 
        id, 
        LAG(i.month) OVER (PARTITION BY i.room_id ORDER BY i.year, i.month) AS prev_month, 
        i.month AS curr_month,
        LAG(i.water_index) OVER (PARTITION BY i.room_id ORDER BY i.year, i.month) AS prev_water, 
        i.water_index AS curr_water, 
        LAG(i.electricity_index) OVER (PARTITION BY i.room_id ORDER BY i.year, i.month) AS prev_electricity, 
        i.electricity_index AS curr_electricity, 
        i.room_id, 
        i.year
    FROM public.index AS i
) AS t
LEFT JOIN public.rooms AS ro ON t.room_id = ro.id
LEFT JOIN public.index AS idx ON t.id = idx.id
WHERE ro.id = $1
  AND idx.month = $2
  AND idx.year = $3
ORDER BY t.year, t.curr_month;

-- name: GetIndexByOwnerIdShort :many
SELECT idx.id, idx.room_id, idx.water_index, idx.electricity_index, idx.month, idx.year
FROM  PUBLIC.INDEX AS idx 
LEFT JOIN public.ROOMS ro ON ro.id = idx.room_id
WHERE  ro.owner = $1;
