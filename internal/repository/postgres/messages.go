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

func (repo *messageRepository) Add(ctx context.Context, msg domain.Message) (domain.Message, error) {
	var id int
	err := repo.db.QueryRow(ctx,
		"INSERT INTO messages(sender_id, recipient_id, time, body) VALUES ($1, $2, now(), $3) RETURNING id",
		msg.SenderID,
		msg.RecipientID,
		msg.Body).Scan(&id)

	if err != nil {
		return domain.Message{}, err
	}

	msg.ID = id
	return msg, nil
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
		&deletedMsg.Time,
		&deletedMsg.Body); err != nil {
		return domain.Message{}, err
	}

	return deletedMsg, nil
}
