package service_habits

import (
	"context"
	"fmt"
	"pets-planify/internal/models"

	"github.com/google/uuid"
)

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*models.Habit, error) {
	res, err := s.storageHabits.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("s.storageHabits.GetByID: %w", err)
	}

	return res, nil
}
