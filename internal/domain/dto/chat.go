package dto

type CreateChatDTO struct {
	Name string `json:"chat_name"`
	Type string `json:"chat_type"`
}

type UpdateChatDTO struct {
	ID   int    `json:"chat_id"`
	Name string `json:"chat_name"`
	Type string `json:"chat_type"`
}
