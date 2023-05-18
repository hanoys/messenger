package service

import (
	"context"
	"errors"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/repository"
)

type chatService struct {
	repositories repository.Repositories
}

func newChatService(repositories *repository.Repositories) *chatService {
	return &chatService{repositories: *repositories}
}

func (s *chatService) FindAll(ctx context.Context) ([]domain.Chat, error) {
	chats, err := s.repositories.Chats.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	if len(chats) == 0 {
		return nil, errors.New("chats not found")
	}

	return chats, nil
}

func (s *chatService) FindByID(ctx context.Context, id int) (domain.Chat, error) {
	return s.repositories.Chats.FindByID(ctx, id)
}

func (s *chatService) Create(ctx context.Context, chat domain.Chat) (domain.Chat, error) {
	return s.repositories.Chats.Create(ctx, chat)
}

func (s *chatService) Delete(ctx context.Context, id int) (domain.Chat, error) {
	return s.repositories.Chats.Delete(ctx, id)
}

func (s *chatService) Update(ctx context.Context, chat domain.Chat) (domain.Chat, error) {
	return s.repositories.Chats.Update(ctx, chat)
}
