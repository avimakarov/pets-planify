package service_habits

import (
	"context"
	"fmt"
	"pets-planify/internal/models"
)

func (s *Service) Update(ctx context.Context, habit models.Habit) error {
	err := s.storageHabits.Update(ctx, habit)
	if err != nil {
		return fmt.Errorf("s.storageHabits.Update: %w", err)
	}

	return nil
}
