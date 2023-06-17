package postgres

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type usersRepository struct {
	db *pgxpool.Pool
}

func NewUsersRepository(db *pgxpool.Pool) *usersRepository {
	return &usersRepository{db}
}

func (repo *usersRepository) Create(ctx context.Context, firstName string, lastName string, email string, nickname string, password string) (domain.User, error) {
	var createdUser domain.User
	err := repo.db.QueryRow(ctx,
		"INSERT INTO users(first_name, last_name, email, nickname, password, created_at) values($1, $2, $3, $4, $5, now()) RETURNING *",
		firstName,
		lastName,
		email,
		nickname,
		password).Scan(&createdUser.ID,
		&createdUser.FirstName,
		&createdUser.LastName,
		&createdUser.Email,
		&createdUser.Nickname,
		&createdUser.Password,
		&createdUser.CreatedAt)
	if err != nil {
		return domain.User{}, err
	}

	return createdUser, nil
}

func (repo *usersRepository) FindAll(ctx context.Context) ([]domain.User, error) {
	rows, err := repo.db.Query(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Nickname,
			&user.Password,
			&user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// TODO: make id uppercase
func (repo *usersRepository) FindById(ctx context.Context, id int) (domain.User, error) {
	row := repo.db.QueryRow(ctx, "SELECT * FROM users WHERE id = $1", id)
	var user domain.User

	if err := row.Scan(&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Nickname,
		&user.Password,
		&user.CreatedAt); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo *usersRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	row := repo.db.QueryRow(ctx, "SELECT * FROM users WHERE email = $1", email)
	var user domain.User

	if err := row.Scan(&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Nickname,
		&user.Password,
		&user.CreatedAt); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo *usersRepository) FindByCredentials(ctx context.Context, email string, password string) (domain.User, error) {
	row := repo.db.QueryRow(ctx,
		"SELECT * FROM users WHERE email = $1 and password = $2",
		email, password)
	var user domain.User

	if err := row.Scan(&user.ID,
        &user.FirstName,
        &user.LastName,
		&user.Email,
        &user.Nickname,
		&user.Password,
		&user.CreatedAt); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo *usersRepository) Delete(ctx context.Context, id int) (domain.User, error) {
	row := repo.db.QueryRow(ctx, "DELETE FROM users WHERE id = $1 RETURNING *", id)

	var deletedUser domain.User
	if err := row.Scan(&deletedUser.ID,
		&deletedUser.FirstName,
		&deletedUser.LastName,
		&deletedUser.Email,
		&deletedUser.Nickname,
		&deletedUser.Password,
		&deletedUser.CreatedAt); err != nil {
		return domain.User{}, err
	}

	return deletedUser, nil
}

func (repo *usersRepository) Update(ctx context.Context, id int, firstName string, lastName string, email string, nickname string, password string) (domain.User, error) {
	row := repo.db.QueryRow(ctx,
		"UPDATE users SET first_name = $2, last_name = $3, email = $4, nickname = $5, password = $6 WHERE id = $1",
		id, firstName, lastName, email, nickname, password)

	var updatedUser domain.User
	if err := row.Scan(&updatedUser.ID,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.Email,
		&updatedUser.Nickname,
		&updatedUser.Password,
		&updatedUser.CreatedAt); err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}
