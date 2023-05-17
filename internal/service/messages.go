package service

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/repository"
)

type messagesService struct {
	repositories *repository.Repositories
}

func newMessagesService(repositories *repository.Repositories) *messagesService {
	return &messagesService{repositories: repositories}
}

func (s *messagesService) Add(msg domain.Message) (domain.Message, error) {
	return s.repositories.Messages.Add(context.TODO(), msg)
}

func (s *messagesService) FindAll() ([]domain.Message, error) {
	messages, err := s.repositories.Messages.FindAll(context.TODO())
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, RowsNotFoundErr
	}

	return messages, nil
}

func (s *messagesService) FindByID(id int) (domain.Message, error) {
	return s.repositories.Messages.FindByID(context.TODO(), id)
}

func (s *messagesService) FindBySenderID(id int) ([]domain.Message, error) {
	messages, err := s.repositories.Messages.FindBySenderID(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, RowsNotFoundErr
	}

	return messages, nil
}

func (s *messagesService) FindByRecipientID(id int) ([]domain.Message, error) {
	messages, err := s.repositories.Messages.FindByRecipientID(context.TODO(), id)
	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, RowsNotFoundErr
	}

	return messages, nil
}

func (s *messagesService) Delete(id int) (domain.Message, error) {
	return s.repositories.Messages.Delete(context.TODO(), id)
}
