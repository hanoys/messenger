package simplerepository

import (
	"context"
	"errors"
	"fmt"

	"github.com/hanoy/messenger/internal/domain"
)

type UsersRepository struct {
	Users map[int]domain.User
}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{Users: make(map[int]domain.User, 128)}
}

func (repo *UsersRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	users := make([]domain.User, 0)
	for _, user := range repo.Users {
		users = append(users, user)
	}

	return users, nil
}

func (repo *UsersRepository) FindById(ctx context.Context, id int) (domain.User, error) {
	user, ok := repo.Users[id]
	if !ok {
		return user, errors.New(fmt.Sprintf("user with id %v not found", id))
	}

	return user, nil
}

func (repo *UsersRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	for _, user := range repo.Users {
		if user.Email == email {
			return user, nil
		}
	}

	return domain.User{}, errors.New(fmt.Sprintf("user with email %v not found", email))
}

func (repo *UsersRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	if _, ok := repo.Users[user.ID]; ok {
		return user, errors.New(fmt.Sprintf("user with id %v already exists", user.ID))
	}

	repo.Users[user.ID] = user
	return user, nil
}

func (repo *UsersRepository) Delete(ctx context.Context, id int) (domain.User, error) {
	if user, ok := repo.Users[id]; ok {
		delete(repo.Users, id)
		return user, nil
	}

	return domain.User{}, errors.New(fmt.Sprintf("user with id %v not found", id))
}

func (repo *UsersRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	if _, ok := repo.Users[user.ID]; ok {
		repo.Users[user.ID] = user
		return user, nil
	}

	return user, errors.New(fmt.Sprintf("user with id %v not found", user.ID))
}
