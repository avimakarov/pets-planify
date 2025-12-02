package service_user_mails

import (
	"context"
	"github.com/google/uuid"
)

//go:generate mockgen --source=deps.go --destination=deps_test.go --package=${GOPACKAGE}_test

type StorageUserMails interface {
	ExistByEmail(ctx context.Context, email string) (bool, error)
	GetUserIdByEmail(ctx context.Context, email string) (*uuid.UUID, error)
}
