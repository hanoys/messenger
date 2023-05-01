package repository

import (
	"context"
	"log"

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

type UsersRepository struct {
	Users
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{Users: simplerepository.NewUsersRepository()}
}

func NewUsersRepositoryPostgres() *UsersRepository {
    dbpool, err := pgxpool.New(context.TODO(), "postgres://messenger:messenger@localhost:5432/messenger")
    if err != nil {
        log.Fatalf("unable to establish connection with database: %v", err)
    }
	userRepo, err := postgres.NewUsersRepository(dbpool)
	if err != nil {
        log.Fatalf("unable to create repository: %v", err)
	}
	return &UsersRepository{Users: userRepo}
}
