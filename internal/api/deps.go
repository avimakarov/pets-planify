package api

import (
	"context"
	schema "pets-planify/internal/generated/openapi/server"
)

//go:generate mockgen --source=deps.go --destination=deps_test.go --package=${GOPACKAGE}_test

type HandlerInfo interface {
	Info(ctx context.Context, out *schema.InfoOut) error
}

type HandlerGetTaskByID interface {
	GetTaskByID(ctx context.Context, in *schema.GetTaskByIdIn, out *schema.GetTaskByIdOut) error
}
