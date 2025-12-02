package usecase_login

import (
	"context"
	"github.com/google/uuid"
)

//go:generate mockgen --source=deps.go --destination=deps_test.go --package=${GOPACKAGE}_test

type ServiceUsers interface {
	CreateWithEmail(ctx context.Context, email string) (*uuid.UUID, error)
}

type ServiceUserMails interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	GetUserIdByEmail(ctx context.Context, email string) (*uuid.UUID, error)
}

type QueueConfirmationCodeRequested interface {
	Produce(ctx context.Context, userID, confirmationID, confirmationToken uuid.UUID) error
}
