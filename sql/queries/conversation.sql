-- name: CreateConversation :one
INSERT INTO public.conversations
(user_a, user_b, last_message_id, created_at)
VALUES(
    $1, $2, $3, now()
)
RETURNING id;

-- name: GetConversationByUserID :many
SELECT id, user_a, user_b, last_message_id, created_at
FROM public.conversations
WHERE user_a = $1 OR user_b = $1;
