package service

import (
	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/repository"
)

type Users interface {
    FindAll() ([]domain.User, error)
    FindByID(id int) (domain.User, error)
    FindByEmail(email string) (domain.User, error)
    Create(user domain.User) (domain.User, error)
    Delete(id int) (domain.User, error)
    Update(user domain.User) (domain.User, error)
}

type Services struct {
    Users
}

func NewServices(repository *repository.UsersRepository) *Services {
    usersService := newUsersService(repository)
    return &Services{Users: usersService}
}
