package service

import (
	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/repository"
)

type Users interface {
	FindAll() ([]domain.User, error)
	FindByID(id int) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Create(user domain.User) (domain.User, error)
	Delete(id int) (domain.User, error)
	Update(user domain.User) (domain.User, error)
}

type Chats interface {
	FindAll() ([]domain.Chat, error)
	FindByID(id int) (domain.Chat, error)
	Create(chat domain.Chat) (domain.Chat, error)
	Delete(id int) (domain.Chat, error)
	Update(chat domain.Chat) (domain.Chat, error)
}

type Services struct {
	Users
	Chats
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{Users: newUsersService(repositories),
		Chats: newChatService(repositories)}
}
