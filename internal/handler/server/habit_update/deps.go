package handler_server_habit_update

import (
	"context"
	"pets-planify/internal/models"

	"github.com/google/uuid"
)

type ServiceHabits interface {
	Update(ctx context.Context, habit models.Habit) error
	GetByID(ctx context.Context, id uuid.UUID) (*models.Habit, error)
}
