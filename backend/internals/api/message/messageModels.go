package message

type MessageBody struct {
	SenderID   uint           `json:"sender_id"`
	ReceiverID uint           `json:"receiver_id"`
	Message    string         `json:"message"`
}
