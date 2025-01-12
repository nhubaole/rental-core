// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: message.sql

package dataaccess

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getMessageByConversationID = `-- name: GetMessageByConversationID :many
SELECT id, 
       conversation_id, 
       sender_id, 
       type, 
       content, 
       rent_auto_content::text, 
       created_at
FROM public.messages
WHERE conversation_id = $1
ORDER BY created_at DESC
`

type GetMessageByConversationIDRow struct {
	ID              int32            `json:"id"`
	ConversationID  int32            `json:"conversation_id"`
	SenderID        int32            `json:"sender_id"`
	Type            int32            `json:"type"`
	Content         *string          `json:"content"`
	RentAutoContent string           `json:"rent_auto_content"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) GetMessageByConversationID(ctx context.Context, conversationID int32) ([]GetMessageByConversationIDRow, error) {
	rows, err := q.db.Query(ctx, getMessageByConversationID, conversationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMessageByConversationIDRow
	for rows.Next() {
		var i GetMessageByConversationIDRow
		if err := rows.Scan(
			&i.ID,
			&i.ConversationID,
			&i.SenderID,
			&i.Type,
			&i.Content,
			&i.RentAutoContent,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMessageByID = `-- name: GetMessageByID :one
SELECT id, conversation_id, sender_id, type, content, rent_auto_content::text,  created_at
FROM public.messages
WHERE id = $1
`

type GetMessageByIDRow struct {
	ID              int32            `json:"id"`
	ConversationID  int32            `json:"conversation_id"`
	SenderID        int32            `json:"sender_id"`
	Type            int32            `json:"type"`
	Content         *string          `json:"content"`
	RentAutoContent string           `json:"rent_auto_content"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) GetMessageByID(ctx context.Context, id int32) (GetMessageByIDRow, error) {
	row := q.db.QueryRow(ctx, getMessageByID, id)
	var i GetMessageByIDRow
	err := row.Scan(
		&i.ID,
		&i.ConversationID,
		&i.SenderID,
		&i.Type,
		&i.Content,
		&i.RentAutoContent,
		&i.CreatedAt,
	)
	return i, err
}

const getMessages = `-- name: GetMessages :many
SELECT id, conversation_id, sender_id, type,  content, rent_auto_content::text, created_at
FROM public.messages
`

type GetMessagesRow struct {
	ID              int32            `json:"id"`
	ConversationID  int32            `json:"conversation_id"`
	SenderID        int32            `json:"sender_id"`
	Type            int32            `json:"type"`
	Content         *string          `json:"content"`
	RentAutoContent string           `json:"rent_auto_content"`
	CreatedAt       pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) GetMessages(ctx context.Context) ([]GetMessagesRow, error) {
	rows, err := q.db.Query(ctx, getMessages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetMessagesRow
	for rows.Next() {
		var i GetMessagesRow
		if err := rows.Scan(
			&i.ID,
			&i.ConversationID,
			&i.SenderID,
			&i.Type,
			&i.Content,
			&i.RentAutoContent,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
