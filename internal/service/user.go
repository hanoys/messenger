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

func (s *usersService) FindAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.repositories.Users.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, RowsNotFoundErr
	}

	return users, nil
}

func (s *usersService) FindByID(ctx context.Context, id int) (domain.User, error) {
	return s.repositories.Users.FindById(ctx, id)
}

func (s *usersService) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.repositories.Users.FindByEmail(ctx, email)
}

func (s *usersService) Create(ctx context.Context, user domain.User) (domain.User, error) {
	return s.repositories.Users.Create(ctx, user)
}

func (s *usersService) Delete(ctx context.Context, id int) (domain.User, error) {
	return s.repositories.Users.Delete(ctx, id)
}

func (s *usersService) Update(ctx context.Context, user domain.User) (domain.User, error) {
	return s.repositories.Users.Update(ctx, user)
}
