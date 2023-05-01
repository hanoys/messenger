package postgres

import (
	"context"
	"errors"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsersRepository struct {
	db *pgxpool.Pool
}

func NewUsersRepository(db *pgxpool.Pool) (*UsersRepository, error) {
	repo := &UsersRepository{db}
	err := repo.Init(context.Background())
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (repo *UsersRepository) Init(ctx context.Context) error {
	query := `
        CREATE TABLE IF NOT EXISTS users(
            id SERIAL PRIMARY KEY,
            first_name TEXT NOT NULL,
            last_name TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE,
            login TEXT NOT NULL UNIQUE,
            password TEXT NOT NULL UNIQUE,
            created_at TIMESTAMP
        );
    `

	_, err := repo.db.Exec(ctx, query)
	return err
}

func (repo *UsersRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	var id int
    err := repo.db.QueryRow(ctx,
		"INSERT INTO users(first_name, last_name, email, login, password, created_at) values($1, $2, $3, $4, $5, now()) RETURNING id",
		user.FirstName,
		user.LastName,
		user.Email,
		user.Login,
		user.Password).Scan(&id)
	if err != nil {
		return domain.User{}, err
	}

	user.ID = id
	return user, nil
}

func (repo *UsersRepository) FindAll(ctx context.Context) ([]domain.User, error) {
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
			&user.Login,
			&user.Password,
			&user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *UsersRepository) FindById(ctx context.Context, id int) (domain.User, error) {
	row := repo.db.QueryRow(ctx, "SELECT * FROM users WHERE id = $1", id)
	var user domain.User

	if err := row.Scan(&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Login,
		&user.Password,
		&user.CreatedAt); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repo *UsersRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	row := repo.db.QueryRow(ctx, "SELECT * FROM users WHERE email = $1", email)
	var user domain.User

	if err := row.Scan(&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Login,
		&user.Password,
		&user.CreatedAt); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

// TODO: return user
func (repo *UsersRepository) Delete(ctx context.Context, id int) (domain.User, error) {
	res, err := repo.db.Exec(ctx, "DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return domain.User{}, err
	}

	if res.RowsAffected() == 0 {
		return domain.User{}, errors.New("user not found")
	}

	return domain.User{}, nil
}

func (repo *UsersRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	row := repo.db.QueryRow(ctx,
		"UPDATE users SET first_name = $2, last_name = $3, email = $4, login = $5, password = $6 WHERE id = $1",
		user.ID, user.FirstName, user.LastName, user.Email, user.Login, user.Password)

	var updatedUser domain.User
	if err := row.Scan(&updatedUser.ID,
		&updatedUser.FirstName,
		&updatedUser.LastName,
		&updatedUser.Email,
		&updatedUser.Login,
		&updatedUser.Password,
		&updatedUser.CreatedAt); err != nil {
		return domain.User{}, err
	}

	return updatedUser, nil
}
