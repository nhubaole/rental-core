-- name: GetMessages :many
SELECT id, conversation_id, sender_id, type,  content, rent_auto_content::text, created_at
FROM public.messages;

-- name: GetMessageByID :one
SELECT id, conversation_id, sender_id, type, content, rent_auto_content::text,  created_at
FROM public.messages
WHERE id = $1;

-- name: GetMessageByConversationID :many
SELECT id, conversation_id, sender_id, type, content, rent_auto_content::text,  created_at
FROM public.messages
WHERE conversation_id = $1;
