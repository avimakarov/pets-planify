package service_habits

import (
	"context"
	"fmt"
	"pets-planify/internal/models"

	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, habit models.Habit) (*uuid.UUID, error) {
	habitID, err := s.storageHabits.Create(ctx, habit)
	if err != nil {
		return nil, fmt.Errorf("s.storageHabits.Create: %w", err)
	}

	return habitID, nil
}
