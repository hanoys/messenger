package domain

import "time"

type Message struct {
	SenderID    int       `json:"sender_id"`
	RecipientID int       `json:"recipient_id"`
	TimeSending time.Time `json:"time_sending"`
	Body        string    `json:"body"`
}

