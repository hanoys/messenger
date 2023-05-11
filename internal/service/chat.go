package service

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/repository"
)

type chatService struct {
    repositories repository.Repositories
}

func newChatService(repositories *repository.Repositories) *chatService {
    return &chatService{repositories: *repositories}
}

func (s *chatService) FindAll() ([]domain.Chat, error) {
    return s.repositories.Chats.FindAll(context.TODO())
}

func (s *chatService) FindByID(id int) (domain.Chat, error) {
    return s.repositories.Chats.FindByID(context.TODO(), id)
}

func (s *chatService) Create(chat domain.Chat) (domain.Chat, error) {
    return s.repositories.Chats.Create(context.TODO(), chat)
}

func (s *chatService) Delete(id int) (domain.Chat, error) {
    return s.repositories.Chats.Delete(context.TODO(), id)
}

func (s *chatService) Update(chat domain.Chat) (domain.Chat, error) {
    return s.repositories.Chats.Update(context.TODO(), chat)
}
