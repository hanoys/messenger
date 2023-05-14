package domain

type Chat struct {
	ID      int   `json:"chat_id"`
	UsersID []int `json:"users_id"`
}
