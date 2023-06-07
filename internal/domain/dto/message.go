package dto

type AddMessageDTO struct {
    SenderID    int       `json:"sender_id"`
	RecipientID int       `json:"recipient_id"`
	ChatID      int       `json:"chat_id"`
	Body        string    `json:"body"`
}

