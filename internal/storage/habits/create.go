package storage_habits

import (
	"context"
	"fmt"
	"pets-planify/internal/models"

	"github.com/google/uuid"
)

const (
	sqlCreate = `
	insert into habits (name, user_id, chat_id) values ($1, $2, $3) returning id
	`
)

func (s *Storage) Create(ctx context.Context, habit models.Habit) (*uuid.UUID, error) {
	var habitID uuid.UUID

	res := s.db.QueryRowContext(
		ctx, sqlCreate,
		habit.Name, habit.UserID, habit.ChatID,
	)

	if err := res.Scan(&habitID); err != nil {
		return nil, fmt.Errorf("res.Scan: %w", err)
	}

	return &habitID, nil
}
