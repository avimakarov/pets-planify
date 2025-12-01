package handler_tgb_habits

import (
	"context"
	"pets-planify/internal/models"

	"github.com/google/uuid"
	"github.com/mymmrac/telego"
)

//go:generate mockgen --source=deps.go --destination=deps_test.go --package=${GOPACKAGE}_test

type ServiceBot interface {
	SendMessage(ctx context.Context, params *telego.SendMessageParams) (*telego.Message, error)
}

type ServiceHabits interface {
	Create(ctx context.Context, habit models.Habit) (*uuid.UUID, error)
}
