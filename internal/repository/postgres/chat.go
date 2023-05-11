package postgres

import (
	"context"
	"errors"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type chatRepository struct {
	db *pgxpool.Pool
}

func NewChatRepository(db *pgxpool.Pool) *chatRepository {
	return &chatRepository{db}
}

func (repo *chatRepository) Create(ctx context.Context, chat domain.Chat) (domain.Chat, error) {
	var id int
	err := repo.db.QueryRow(ctx,
		"INSERT INTO chats(id, users_id) values ($1, $2) RETURNING id",
		chat.ID,
		chat.UsersID).Scan(&id)
	if err != nil {
		return domain.Chat{}, err
	}

	chat.ID = id
	return chat, nil
}

func (repo *chatRepository) FindAll(ctx context.Context) ([]domain.Chat, error) {
	rows, err := repo.db.Query(ctx, "SELECT * FROM chats")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var chats []domain.Chat
	for rows.Next() {
		var chat domain.Chat
		if err := rows.Scan(&chat.ID, &chat.UsersID); err != nil {
			return nil, err
		}

		chats = append(chats, chat)
	}

	return chats, nil
}

func (repo *chatRepository) FindByID(ctx context.Context, id int) (domain.Chat, error) {
	row := repo.db.QueryRow(ctx, "SELECT * FROM chats WHERE id = $1", id)
	var chat domain.Chat

	if err := row.Scan(&chat.ID, &chat.UsersID); err != nil {
		return domain.Chat{}, err
	}

	return chat, nil
}

// TODO: return chat
func (repo *chatRepository) Delete(ctx context.Context, id int) (domain.Chat, error) {
	res, err := repo.db.Exec(ctx, "DELETE FROM chats WHERE id = $1", id)
	if err != nil {
		return domain.Chat{}, err
	}

	if res.RowsAffected() == 0 {
		return domain.Chat{}, errors.New("chat not found")
	}

	return domain.Chat{}, nil
}

func (repo *chatRepository) Update(ctx context.Context, chat domain.Chat) (domain.Chat, error) {
	row := repo.db.QueryRow(ctx,
		"UPDATE chats SET users_id = $2 WHERE id = $1", chat.ID, chat.UsersID)

	var updatedChat domain.Chat
	if err := row.Scan(&updatedChat.ID, &updatedChat.UsersID); err != nil {
		return domain.Chat{}, err
	}

	return updatedChat, nil
}
