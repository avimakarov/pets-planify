package storage_users

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
	"pets-planify/internal"
)

const (
	sqlCreateWithTx = `
	insert into users (id) values ($1)
	`
)

func (s *Storage) CreateWithTx(ctx context.Context, id uuid.UUID, tx internal.Tx) error {
	_, err := tx.ExecContext(ctx, sqlCreateWithTx, id)
	if err != nil {
		return errors.Wrap(err, "tx.ExecContext")
	}

	return nil
}
