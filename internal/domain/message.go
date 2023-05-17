package domain

import "time"

type Message struct {
	ID          int       `json:"chat_id"`
	SenderID    int       `json:"sender_id"`
	RecipientID int       `json:"recipient_id"`
	Time        time.Time `json:"time_sending"`
	Body        string    `json:"body"`
}
