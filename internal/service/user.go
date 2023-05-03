package service

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/repository"
)

type usersService struct {
    repository *repository.UsersRepository
}

func newUsersService(repository *repository.UsersRepository) *usersService {
    return &usersService{repository: repository}
}

func (s *usersService) FindAll() ([]domain.User, error) {
    return s.repository.FindAll(context.TODO())
}

func (s *usersService) FindByID(id int) (domain.User, error) {
    return s.repository.FindById(context.TODO(), id)
}

func (s *usersService) FindByEmail(email string) (domain.User, error) {
    return s.repository.FindByEmail(context.TODO(), email)
}

func (s *usersService) Create(user domain.User) (domain.User, error) {
    return s.repository.Create(context.TODO(), user)
}

func (s *usersService) Delete(id int) (domain.User, error) {
    return s.repository.Delete(context.TODO(), id)
}

func (s *usersService) Update(user domain.User) (domain.User, error) {
    return s.repository.Update(context.TODO(), user)
}
