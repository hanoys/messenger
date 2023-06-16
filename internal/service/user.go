package service

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/domain/dto"
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

func (s *usersService) FindByEmail(ctx context.Context, userDTO dto.FindByEmailUserDTO) (domain.User, error) {
	return s.repositories.Users.FindByEmail(ctx, userDTO.Email)
}

func (s *usersService) FindByCredentials(ctx context.Context, userDTO dto.LogInUserDTO) (domain.User, error) {
    return s.repositories.Users.FindByCredentials(ctx, userDTO.Email, userDTO.Password)
}

func (s *usersService) Create(ctx context.Context, userDTO dto.CreateUserDTO) (domain.User, error) {
	return s.repositories.Users.Create(ctx, userDTO.FirstName,
		userDTO.LastName, userDTO.Email, userDTO.Login, userDTO.Password)
}

func (s *usersService) Delete(ctx context.Context, id int) (domain.User, error) {
	return s.repositories.Users.Delete(ctx, id)
}

func (s *usersService) Update(ctx context.Context, userDTO dto.UpdateUserDTO) (domain.User, error) {
	return s.repositories.Users.Update(ctx, userDTO.ID, userDTO.FirstName,
		userDTO.LastName, userDTO.Email, userDTO.Login, userDTO.Password)
}
