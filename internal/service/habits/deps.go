package service_habits

import (
	"context"
	"pets-planify/internal/models"

	"github.com/google/uuid"
)

type StorageHabits interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Habit, error)
	Update(ctx context.Context, habit models.Habit) error
	Create(ctx context.Context, habit models.Habit) (*uuid.UUID, error)
}
