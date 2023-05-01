package service

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/repository"
)

type UsersService struct {
    repository *repository.UsersRepository
}

func NewUsersService(repository *repository.UsersRepository) *UsersService {
    return &UsersService{repository: repository}
}

func (s *UsersService) FindAll() ([]domain.User, error) {
    return s.repository.FindAll(context.TODO())
}

func (s *UsersService) FindByID(id int) (domain.User, error) {
    return s.repository.FindById(context.TODO(), id)
}

func (s *UsersService) FindByEmail(email string) (domain.User, error) {
    return s.repository.FindByEmail(context.TODO(), email)
}

func (s *UsersService) Create(user domain.User) (domain.User, error) {
    return s.repository.Create(context.TODO(), user)
}

func (s *UsersService) Delete(id int) (domain.User, error) {
    return s.repository.Delete(context.TODO(), id)
}

func (s *UsersService) Update(user domain.User) (domain.User, error) {
    return s.repository.Update(context.TODO(), user)
}
