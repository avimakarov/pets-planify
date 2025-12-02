package storage_user_mails

import (
	"context"
	"emperror.dev/errors"
)

const (
	sqlExistByEmail = `
	select exists (
		select id from user_mails where email = $1
	)
`
)

func (s *Storage) ExistByEmail(ctx context.Context, email string) (bool, error) {
	var e bool

	if err := s.db.QueryRowContext(ctx, sqlExistByEmail, email).Scan(&e); err != nil {
		return false, errors.Wrap(err, "s.db.QueryRowContext")
	}

	return e, nil
}
