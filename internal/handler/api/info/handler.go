package handler_api_info

import (
	"context"
	schema "pets-planify/internal/generated/openapi/server"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Info(_ context.Context, out *schema.InfoOut) error {
	out.Status = "ok"
	return nil
}
