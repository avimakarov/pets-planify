package handler_tgb_create

import (
	"context"
	"github.com/google/uuid"
	"github.com/mymmrac/telego"
)

//go:generate mockgen --source=deps.go --destination=deps_test.go --package=${GOPACKAGE}_test

type ServiceBot interface {
	SendMessage(ctx context.Context, params *telego.SendMessageParams) (*telego.Message, error)
}

type ServiceTasks interface {
	Create(ctx context.Context, userID int64) (*uuid.UUID, error)
}
