package api

import (
	"context"
	schema "pets-planify/internal/generated/openapi/server"
)

type HandlerLogin interface {
	Login(ctx context.Context, in *schema.LoginIn, out *schema.LoginOut) error
}
