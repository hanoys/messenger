package dto

type CreateUserDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

type UpdateUserDTO struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

type FindByEmailUserDTO struct {
	Email string `json:"email"`
}

type SignUpUserDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	Password  string `json:"password"`
}

type LogInUserDTO struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}
