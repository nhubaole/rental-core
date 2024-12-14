-- name: GetNotificationsByUserID :many
SELECT id, user_id, reference_id, reference_type, title, is_read, created_at, updated_at
FROM public.notifications
WHERE user_id = $1
ORDER BY created_at DESC;