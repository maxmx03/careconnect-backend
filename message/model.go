package message

type MessageModel struct {
	MessageID        int `json:"message_id"`
	SenderID         int
	RecipientID      int
	DateTime         string `json:"date"`
	Content          string `json:"content"`
	SenderName       string `json:"sender_name"`
	SenderSurname    string `json:"sender_surname"`
	SenderType       string `json:"sender_type"`
	RecipientName    string `json:"recipient_name"`
	RecipientSurname string `json:"recipient_surname"`
	RecipientType    string `json:"recipient_type"`
}
