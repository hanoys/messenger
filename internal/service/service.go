package service

import (
	"context"
	"errors"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/repository"
)

var RowsNotFoundErr = errors.New("rows not found")

type Users interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, id int) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, id int) (domain.User, error)
	Update(ctx context.Context, user domain.User) (domain.User, error)
}

type Chats interface {
	FindAll(ctx context.Context) ([]domain.Chat, error)
	FindByID(ctx context.Context, id int) (domain.Chat, error)
	Create(ctx context.Context, chat domain.Chat) (domain.Chat, error)
	Delete(ctx context.Context, id int) (domain.Chat, error)
	Update(ctx context.Context, chat domain.Chat) (domain.Chat, error)
}

type Messages interface {
	Add(ctx context.Context, msg domain.Message) (domain.Message, error)
	FindAll(ctx context.Context) ([]domain.Message, error)
	FindByID(ctx context.Context, id int) (domain.Message, error)
	FindBySenderID(ctx context.Context, id int) ([]domain.Message, error)
	FindByRecipientID(ctx context.Context, id int) ([]domain.Message, error)
	Delete(ctx context.Context, id int) (domain.Message, error)
}

type Services struct {
	Users
	Chats
	Messages
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{Users: newUsersService(repositories),
		Chats:    newChatService(repositories),
		Messages: newMessagesService(repositories)}
}
