package service

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/domain/dto"
	"github.com/hanoy/messenger/internal/repository"
)

type adminsService struct {
	repositories *repository.Repositories
}

func newAdminsService(repositories *repository.Repositories) *adminsService {
	return &adminsService{repositories: repositories}
}

func (s *adminsService) FindByCredentials(ctx context.Context, adminDTO dto.LogInAdminDTO) (domain.Admin, error) {
    return s.repositories.Admins.FindByCredentials(ctx, adminDTO.Email, adminDTO.Password) 
}
