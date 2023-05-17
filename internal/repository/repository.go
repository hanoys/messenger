package repository

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/repository/postgres"
	simplerepository "github.com/hanoy/messenger/internal/repository/simple_repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Users interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindById(ctx context.Context, id int) (domain.User, error)
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

type UsersRepository struct {
	Users
}

type Repositories struct {
	Users
	Chats
	Messages
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{Users: simplerepository.NewUsersRepository()}
}

func NewUsersRepositoryPostgres(db *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{Users: postgres.NewUsersRepository(db)}
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{Users: postgres.NewUsersRepository(db),
		Chats:    postgres.NewChatRepository(db),
		Messages: postgres.NewMessageRepository(db)}
}
