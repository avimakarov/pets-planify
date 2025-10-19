package service_tasks

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
	"pets-planify/internal/models"
)

func (s *Service) GetByID(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	task, err := s.storageTasks.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "s.storageTasks.GetByID")
	}

	return task, nil
}
