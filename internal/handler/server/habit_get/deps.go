package handler_server_get

import (
	"context"
	"pets-planify/internal/models"

	"github.com/google/uuid"
)

type ServiceHabits interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Habit, error)
}
