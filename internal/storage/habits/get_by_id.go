package storage_habits

import (
	"context"
	"fmt"
	"pets-planify/internal/models"

	"github.com/google/uuid"
)

const (
	sqlGetByID = `
	select id, name, user_id, chat_id, canceled, created_at from habits where id = $1
	`
)

func (s *Storage) GetByID(ctx context.Context, id uuid.UUID) (*models.Habit, error) {
	var habit models.Habit

	err := s.db.QueryRowContext(
		ctx, sqlGetByID, id,
	).Scan(
		&habit.ID, &habit.Name, &habit.UserID,
		&habit.ChatID, &habit.Canceled, &habit.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("s.db.QueryRowContext: %w", err)
	}

	return &habit, nil

}
