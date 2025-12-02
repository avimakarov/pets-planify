package storage_user_mails

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
)

const (
	sqlGetUserIdByEmail = `
	select user_id from user_mails where email = $1
	`
)

func (s *Storage) GetUserIdByEmail(ctx context.Context, email string) (*uuid.UUID, error) {
	var userID uuid.UUID

	if err := s.db.QueryRowContext(ctx, sqlGetUserIdByEmail, email).Scan(&userID); err != nil {
		return nil, errors.Wrap(err, "s.db.QueryRowContext")
	}

	return &userID, nil
}
