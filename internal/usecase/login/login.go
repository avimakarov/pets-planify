package usecase_login

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
)

type LoginIn struct {
	Email string
}

type LoginOut struct {
	ConfirmationID uuid.UUID
}

func (u *Usecase) Login(ctx context.Context, in LoginIn) (*LoginOut, error) {
	confirmationID := uuid.New()
	confirmationToken := uuid.New()

	mailExists, err := u.serviceUserMails.ExistsByEmail(ctx, in.Email)
	if err != nil {
		return nil, errors.Wrap(err, "u.serviceUserMails.ExistsByEmail")
	}

	if !mailExists {
		_, err := u.serviceUsers.CreateWithEmail(ctx, in.Email)
		if err != nil {
			return nil, errors.Wrap(err, "u.serviceUsers.CreateWithEmail")
		}
	}

	userID, err := u.serviceUserMails.GetUserIdByEmail(ctx, in.Email)
	if err != nil {
		return nil, errors.Wrap(err, "u.serviceUserMails.GetUserIdByEmail")
	}

	if err := u.queueConfirmationCodeRequested.Produce(ctx, *userID, confirmationID, confirmationToken); err != nil {
		return nil, errors.Wrap(err, "u.queueConfirmationCodeRequested.Produce")
	}

	return &LoginOut{
		ConfirmationID: confirmationID,
	}, nil
}
