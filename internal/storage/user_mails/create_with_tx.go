package storage_user_mails

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
	"pets-planify/internal"
)

const (
	sqlCreateWithTx = `
	insert into user_mails (email, user_id) values ($1, $2)
	`
)

func (s *Storage) CreateWithTx(ctx context.Context, userID uuid.UUID, email string, tx internal.Tx) error {
	_, err := tx.ExecContext(ctx, sqlCreateWithTx, email, userID)
	if err != nil {
		return errors.Wrap(err, "tx.ExecContext")
	}

	return nil
}
