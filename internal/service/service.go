package service

import (
	"context"
	"errors"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/domain/dto"
	"github.com/hanoy/messenger/internal/repository"
)

var RowsNotFoundErr = errors.New("rows not found")

type Users interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, id int) (domain.User, error)
	FindByEmail(ctx context.Context, userDTO dto.FindByEmailUserDTO) (domain.User, error)
	FindByCredentials(ctx context.Context, userDTO dto.LogInUserDTO) (domain.User, error)
	Create(ctx context.Context, userDTO dto.CreateUserDTO) (domain.User, error)
	Delete(ctx context.Context, id int) (domain.User, error)
	Update(ctx context.Context, userDTO dto.UpdateUserDTO) (domain.User, error)
}

type Admins interface {
	FindByCredentials(ctx context.Context, adminDTO dto.LogInAdminDTO) (domain.Admin, error)
}

type Chats interface {
	FindAll(ctx context.Context) ([]domain.Chat, error)
	FindByID(ctx context.Context, id int) (domain.Chat, error)
	Create(ctx context.Context, chatDTO dto.CreateChatDTO) (domain.Chat, error)
	Delete(ctx context.Context, id int) (domain.Chat, error)
	Update(ctx context.Context, chatDTO dto.UpdateChatDTO) (domain.Chat, error)
}

type Messages interface {
	Add(ctx context.Context, msgDTO dto.AddMessageDTO) (domain.Message, error)
	FindAll(ctx context.Context) ([]domain.Message, error)
	FindByID(ctx context.Context, id int) (domain.Message, error)
	FindBySenderID(ctx context.Context, id int) ([]domain.Message, error)
	FindByRecipientID(ctx context.Context, id int) ([]domain.Message, error)
	Delete(ctx context.Context, id int) (domain.Message, error)
}

type Services struct {
	Users
	Admins
	Chats
	Messages
}

func NewServices(repositories *repository.Repositories) *Services {
	return &Services{Users: newUsersService(repositories),
		Admins:   newAdminsService(repositories),
		Chats:    newChatService(repositories),
		Messages: newMessagesService(repositories)}
}
