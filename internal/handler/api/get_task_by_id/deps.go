package handler_api_get_task_by_id

import (
	"context"
	"github.com/google/uuid"
	"pets-planify/internal/models"
)

//go:generate mockgen --source=deps.go --destination=deps_test.go --package=${GOPACKAGE}_test

type ServiceTasks interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Task, error)
}
