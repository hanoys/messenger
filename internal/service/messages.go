package service

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/domain/dto"
	"github.com/hanoy/messenger/internal/repository"
)

type messagesService struct {
	repositories *repository.Repositories
}

func newMessagesService(repositories *repository.Repositories) *messagesService {
	return &messagesService{repositories: repositories}
}

func (s *messagesService) Add(ctx context.Context, msgDTO dto.AddMessageDTO) (domain.Message, error) {
	return s.repositories.Messages.Add(ctx, msgDTO.SenderID, msgDTO.RecipientID, msgDTO.ChatID, msgDTO.Body)
}

func (s *messagesService) FindAll(ctx context.Context) ([]domain.Message, error) {
	messages, err := s.repositories.Messages.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, RowsNotFoundErr
	}

	return messages, nil
}

func (s *messagesService) FindByID(ctx context.Context, id int) (domain.Message, error) {
	return s.repositories.Messages.FindByID(ctx, id)
}

func (s *messagesService) FindBySenderID(ctx context.Context, id int) ([]domain.Message, error) {
	messages, err := s.repositories.Messages.FindBySenderID(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, RowsNotFoundErr
	}

	return messages, nil
}

func (s *messagesService) FindByRecipientID(ctx context.Context, id int) ([]domain.Message, error) {
	messages, err := s.repositories.Messages.FindByRecipientID(ctx, id)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, RowsNotFoundErr
	}

	return messages, nil
}

func (s *messagesService) Delete(ctx context.Context, id int) (domain.Message, error) {
	return s.repositories.Messages.Delete(ctx, id)
}
