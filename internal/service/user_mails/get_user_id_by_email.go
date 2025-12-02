package service_user_mails

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
)

func (s *Service) GetUserIdByEmail(ctx context.Context, email string) (*uuid.UUID, error) {
	res, err := s.storageUserMails.GetUserIdByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err, "s.storageUserMails.GetUserIdByEmail")
	}

	return res, nil
}
