package service_user_mails

import (
	"context"
	"emperror.dev/errors"
)

func (s *Service) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	exists, err := s.storageUserMails.ExistByEmail(ctx, email)
	if err != nil {
		return false, errors.Wrap(err, "s.storageUserMails.ExistByEmail")
	}

	return exists, nil
}
