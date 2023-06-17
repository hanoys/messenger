package domain

type Chat struct {
	ID   int    `json:"chat_id"`
	Name string `json:"chat_name"`
	Type string `json:"chat_type"`
}
