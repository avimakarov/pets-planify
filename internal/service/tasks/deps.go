package service_tasks

import (
	"context"
	"github.com/google/uuid"
	"pets-planify/internal/models"
)

//go:generate mockgen --source=deps.go --destination=deps_test.go --package=${GOPACKAGE}_test

type StorageTasks interface {
	Create(ctx context.Context, userID int64) (*uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (*models.Task, error)
}
