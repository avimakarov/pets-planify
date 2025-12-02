package service_users

import (
	"context"
	"github.com/google/uuid"
	"pets-planify/internal"
)

//go:generate mockgen --source=deps.go --destination=deps_test.go --package=${GOPACKAGE}_test

type Tx interface {
	internal.Tx
}

type Storage interface {
	Tx() (internal.Tx, error)
}

type StorageUsers interface {
	CreateWithTx(ctx context.Context, id uuid.UUID, tx internal.Tx) error
}

type StorageUserMails interface {
	CreateWithTx(ctx context.Context, userID uuid.UUID, email string, tx internal.Tx) error
}
