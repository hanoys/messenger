package service

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/repository"
)

type usersService struct {
    repositories *repository.Repositories
}

func newUsersService(repositories *repository.Repositories) *usersService {
    return &usersService{repositories: repositories}
}

func (s *usersService) FindAll() ([]domain.User, error) {
    return s.repositories.Users.FindAll(context.TODO())
}

func (s *usersService) FindByID(id int) (domain.User, error) {
    return s.repositories.Users.FindById(context.TODO(), id)
}

func (s *usersService) FindByEmail(email string) (domain.User, error) {
    return s.repositories.Users.FindByEmail(context.TODO(), email)
}

func (s *usersService) Create(user domain.User) (domain.User, error) {
    return s.repositories.Users.Create(context.TODO(), user)
}

func (s *usersService) Delete(id int) (domain.User, error) {
    return s.repositories.Users.Delete(context.TODO(), id)
}

func (s *usersService) Update(user domain.User) (domain.User, error) {
    return s.repositories.Users.Update(context.TODO(), user)
}
