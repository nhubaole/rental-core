-- name: CreateTenant :exec
INSERT INTO "tenants" 
(
  "room_id", 
  "tenant_id", 
  "begin_date", 
  "end_date", 
  "created_at", 
  "updated_at", 
  "deleted_at"
) 
VALUES 
(
  $1, $2, $3, $4, now(), now(), $5
);

-- name: DeleteTenantByRoomID :exec
UPDATE public.tenants
SET deleted_at = now()
WHERE deleted_at IS NULL
  AND room_id = $1;

-- name: GetRoomByTenantID :many
SELECT 
    t.id, 
    r.id AS room_id, 
    r.title, 
    r.address, 
    r.room_number, 
    t.tenant_id, 
    t.begin_date, 
    t.end_date, 
    t.created_at, 
    t.updated_at
FROM 
    public.tenants t
LEFT JOIN 
    public.rooms r 
ON 
    t.room_id = r.id
WHERE 
    t.tenant_id = $1 
    AND t.deleted_at IS NULL;