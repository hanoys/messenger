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

func (repo *usersRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
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
			&user.Login,
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
		&user.Login,
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
		&user.Login,
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
		&deletedUser.Login,
		&deletedUser.Password,
		&deletedUser.CreatedAt); err != nil {
		return domain.User{}, err
    }

    return deletedUser, nil
}

func (repo *usersRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
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
