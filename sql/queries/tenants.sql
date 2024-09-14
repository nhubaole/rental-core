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