package postgres

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type adminsRepository struct {
	db *pgxpool.Pool
}

func NewAdminsRepository(db *pgxpool.Pool) *adminsRepository {
	return &adminsRepository{db}
}

func (repo *adminsRepository) FindByCredentials(ctx context.Context, email string, password string) (domain.Admin, error) {
	row := repo.db.QueryRow(ctx,
		"SELECT * FROM admins WHERE email = $1 and password = $2",
		email, password)
	var admin domain.Admin

	if err := row.Scan(&admin.ID,
		&admin.Email,
		&admin.Password,
		&admin.CreatedAt); err != nil {
		return domain.Admin{}, err
	}

	return admin, nil
}
