package postgres

import (
	"context"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type messageRepository struct {
	db *pgxpool.Pool
}

func NewMessageRepository(db *pgxpool.Pool) *messageRepository {
	return &messageRepository{db}
}

func (repo *messageRepository) Add(ctx context.Context, sender_id int, recipient_id int, chat_id int, body string) (domain.Message, error) {
	var addedMessage domain.Message
	err := repo.db.QueryRow(ctx,
		"INSERT INTO messages(sender_id, recipient_id, chat_id, time, body) VALUES ($1, $2, $3, now(), $4) RETURNING *",
		sender_id,
		recipient_id,
		chat_id,
		body).Scan(&addedMessage.ID,
		&addedMessage.SenderID,
		&addedMessage.RecipientID,
		&addedMessage.ChatID,
		&addedMessage.Time,
		&addedMessage.Body)

	if err != nil {
		return domain.Message{}, err
	}

	return addedMessage, nil
}

func (repo *messageRepository) FindAll(ctx context.Context) ([]domain.Message, error) {
	rows, err := repo.db.Query(ctx, "SELECT * FROM messages")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []domain.Message
	for rows.Next() {
		var msg domain.Message
		if err := rows.Scan(&msg.ID,
			&msg.SenderID,
			&msg.RecipientID,
			&msg.ChatID,
			&msg.Time,
			&msg.Body); err != nil {
			return nil, err
		}

		messages = append(messages, msg)
	}

	return messages, nil
}

func (repo *messageRepository) FindByID(ctx context.Context, id int) (domain.Message, error) {
	row := repo.db.QueryRow(ctx, "SELECT * FROM messages WHERE id = $1", id)
	var msg domain.Message

	if err := row.Scan(&msg.ID,
		&msg.SenderID,
		&msg.RecipientID,
		&msg.ChatID,
		&msg.Time,
		&msg.Body); err != nil {
		return domain.Message{}, err
	}

	return msg, nil
}

func (repo *messageRepository) FindBySenderID(ctx context.Context, sender_id int) ([]domain.Message, error) {
	rows, err := repo.db.Query(ctx, "SELECT * FROM messages WHERE sender_id = $1", sender_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []domain.Message
	for rows.Next() {
		var msg domain.Message
		if err := rows.Scan(&msg.ID,
			&msg.SenderID,
			&msg.RecipientID,
			&msg.ChatID,
			&msg.Time,
			&msg.Body); err != nil {
			return nil, err
		}

		messages = append(messages, msg)
	}

	return messages, nil
}

func (repo *messageRepository) FindByRecipientID(ctx context.Context, recipient_id int) ([]domain.Message, error) {
	rows, err := repo.db.Query(ctx, "SELECT * FROM messages WHERE sender_id = $1", recipient_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var messages []domain.Message
	for rows.Next() {
		var msg domain.Message
		if err := rows.Scan(&msg.ID,
			&msg.SenderID,
			&msg.RecipientID,
			&msg.ChatID,
			&msg.Time,
			&msg.Body); err != nil {
			return nil, err
		}

		messages = append(messages, msg)
	}

	return messages, nil
}

func (repo *messageRepository) Delete(ctx context.Context, id int) (domain.Message, error) {
	row := repo.db.QueryRow(ctx, "DELETE FROM messages WHERE id = $1 RETURNING *", id)

	var deletedMsg domain.Message
	if err := row.Scan(&deletedMsg.ID,
		&deletedMsg.SenderID,
		&deletedMsg.RecipientID,
		&deletedMsg.ChatID,
		&deletedMsg.Time,
		&deletedMsg.Body); err != nil {
		return domain.Message{}, err
	}

	return deletedMsg, nil
}
