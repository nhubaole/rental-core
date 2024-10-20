package requests

type CreateConversationReq struct {
	UserB         int32  `json:"user_b"`
	LastMessageID *int32 `json:"last_message_id"`
}