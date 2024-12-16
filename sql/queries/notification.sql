-- name: GetNotificationsByUserID :many
SELECT id, user_id, reference_id, reference_type, title, is_read, created_at, updated_at
FROM public.notifications
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: CreateNotification :exec
INSERT INTO public.notifications (
    user_id,
    reference_id,
    reference_type,
    title,
    is_read,
    created_at,
    updated_at
) VALUES (
    $1, -- user_id
    $2, -- reference_id
    $3, -- reference_type
    $4, -- title
    false, -- is_read (mặc định chưa đọc)
    now(), -- created_at
    now()  -- updated_at
);