package service_users

import (
	"context"
	"database/sql"
	"emperror.dev/errors"
	"github.com/google/uuid"
	"log"
)

func (s *Service) CreateWithEmail(ctx context.Context, email string) (*uuid.UUID, error) {
	// todo сделать генератор uuid как отдельный сервис / утилиту
	userID := uuid.New()

	tx, err := s.storage.Tx()
	if err != nil {
		return nil, errors.Wrap(err, "s.storage.Tx")
	}
	defer func() {
		if txErr := tx.Rollback(); txErr != nil && !errors.Is(txErr, sql.ErrTxDone) {
			log.Println(errors.Wrap(txErr, "tx.Rollback"))
		}
	}()

	if userErr := s.storageUsers.CreateWithTx(ctx, userID, tx); userErr != nil {
		return nil, errors.Wrap(userErr, "s.storageUsers.CreateWithTx")
	}

	if mailErr := s.storageUserMails.CreateWithTx(ctx, userID, email, tx); mailErr != nil {
		return nil, errors.Wrap(mailErr, "s.storageUserMails.CreateWithTx")
	}

	if commitErr := tx.Commit(); commitErr != nil {
		return nil, errors.Wrap(commitErr, "tx.Commit")
	}

	return &userID, nil
}
