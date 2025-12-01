package storage_habits

import (
	"context"
	"fmt"
	"pets-planify/internal/models"
)

const (
	sqlUpdate = `
	update habits set name = $1 where id = $2
	`
)

func (s *Storage) Update(ctx context.Context, habit models.Habit) error {
	_, err := s.db.ExecContext(ctx, sqlUpdate, habit.Name, habit.ID)
	if err != nil {
		return fmt.Errorf("s.db.ExecContext: %w", err)
	}

	return nil
}
