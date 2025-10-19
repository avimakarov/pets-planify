package storage_tasks

import (
	"context"
	"database/sql"
	"emperror.dev/errors"
	"github.com/google/uuid"
	"pets-planify/internal/models"
)

const (
	sqlGetByID = `
	select id, user_id, is_done, created_at, name, planed_to, planed_from, description
	from tasks
	where id = $1
	`
)

func (s *Storage) GetByID(ctx context.Context, id uuid.UUID) (*models.Task, error) {
	var res models.Task

	err := s.db.QueryRowContext(ctx, sqlGetByID, id).Scan(
		&res.UUID, &res.UserID, &res.IsDone, &res.CreatedAt,
		&res.Name, &res.PlanedTo, &res.PlanedFrom, &res.Description,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "s.db.QueryRowContext")
	}

	return &res, nil
}
