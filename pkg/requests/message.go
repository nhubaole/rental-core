package requests

type MessageReq struct {
	Message string `json:"message"`
	SocketID string `json:"socket_id"`
}