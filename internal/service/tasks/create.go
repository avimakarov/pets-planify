package service_tasks

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
)

func (s *Service) Create(ctx context.Context, userID int64) (*uuid.UUID, error) {
	taskID, err := s.storageTasks.Create(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "s.storageTasks.Create")
	}

	return taskID, nil
}
