package domain

import "time"

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Nickname  string    `json:"nickname"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
