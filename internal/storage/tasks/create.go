package storage_tasks

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
)

const (
	sqlCreate = `
	insert into tasks (user_id) values ($1) returning id
	`
)

func (s *Storage) Create(ctx context.Context, userID int64) (*uuid.UUID, error) {
	var taskID uuid.UUID

	if err := s.db.QueryRowContext(ctx, sqlCreate, userID).Scan(&taskID); err != nil {
		return nil, errors.Wrap(err, "s.db.QueryRowContext")
	}

	return &taskID, nil
}
