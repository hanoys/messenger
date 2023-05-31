package postgres

import (
	"context"

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
		"INSERT INTO chats(name, type) VALUES ($1, $2) RETURNING id",
		chat.Name, chat.Type).Scan(&id)
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
		if err := rows.Scan(&chat.ID, &chat.Name, &chat.Type); err != nil {
			return nil, err
		}

		chats = append(chats, chat)
	}

	return chats, nil
}

func (repo *chatRepository) FindByID(ctx context.Context, id int) (domain.Chat, error) {
	row := repo.db.QueryRow(ctx, "SELECT * FROM chats WHERE id = $1", id)
	var chat domain.Chat

	if err := row.Scan(&chat.ID, &chat.Name, &chat.Type); err != nil {
		return domain.Chat{}, err
	}

	return chat, nil
}

func (repo *chatRepository) Delete(ctx context.Context, id int) (domain.Chat, error) {
	row := repo.db.QueryRow(ctx, "DELETE FROM chats WHERE id = $1 RETURNING *", id)

	var deletedChat domain.Chat
	if err := row.Scan(&deletedChat.ID, &deletedChat.Name, &deletedChat.Type); err != nil {
		return domain.Chat{}, err
	}

	return deletedChat, nil
}

func (repo *chatRepository) Update(ctx context.Context, chat domain.Chat) (domain.Chat, error) {
	row := repo.db.QueryRow(ctx,
		"UPDATE chats SET name = $2, type = $3 WHERE id = $1", chat.ID, chat.Name, chat.Type)

	var updatedChat domain.Chat
	if err := row.Scan(&updatedChat.ID, &updatedChat.Name, &updatedChat.Type); err != nil {
		return domain.Chat{}, err
	}

	return updatedChat, nil
}
