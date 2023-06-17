package repository

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/repository/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Users interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindById(ctx context.Context, id int) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindByCredentials(ctx context.Context, email string, password string) (domain.User, error)
	Create(ctx context.Context, firstName string, lastName string, email string, nickname string, password string) (domain.User, error)
	Delete(ctx context.Context, id int) (domain.User, error)
	Update(ctx context.Context, id int, firstName string, lastName string, email string, nickname string, password string) (domain.User, error)
}

type Admins interface {
	FindByCredentials(ctx context.Context, email string, password string) (domain.Admin, error)
}

type Chats interface {
	FindAll(ctx context.Context) ([]domain.Chat, error)
	FindByID(ctx context.Context, id int) (domain.Chat, error)
	Create(ctx context.Context, name string, chat_type string) (domain.Chat, error)
	Delete(ctx context.Context, id int) (domain.Chat, error)
	Update(ctx context.Context, id int, name string, chat_type string) (domain.Chat, error)
}

type Messages interface {
	Add(ctx context.Context, sender_id int, recipient_id int, chat_id int, body string) (domain.Message, error)
	FindAll(ctx context.Context) ([]domain.Message, error)
	FindByID(ctx context.Context, id int) (domain.Message, error)
	FindBySenderID(ctx context.Context, id int) ([]domain.Message, error)
	FindByRecipientID(ctx context.Context, id int) ([]domain.Message, error)
	Delete(ctx context.Context, id int) (domain.Message, error)
}

type Repositories struct {
	Users
	Admins
	Chats
	Messages
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{Users: postgres.NewUsersRepository(db),
		Admins:   postgres.NewAdminsRepository(db),
		Chats:    postgres.NewChatRepository(db),
		Messages: postgres.NewMessageRepository(db)}
}
