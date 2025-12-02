package handler_login

import (
	"context"
	"emperror.dev/errors"
	schema "pets-planify/internal/generated/openapi/server"
	usecase_login "pets-planify/internal/usecase/login"
)

type Handler struct {
	useCaseLogin UseCaseLogin
}

func New(
	useCaseLogin UseCaseLogin,
) *Handler {
	return &Handler{
		useCaseLogin: useCaseLogin,
	}
}

func (h *Handler) Login(ctx context.Context, in *schema.LoginIn, out *schema.LoginOut) error {
	req := usecase_login.LoginIn{
		Email: string(in.Email),
	}

	res, err := h.useCaseLogin.Login(ctx, req)
	if err != nil {
		return errors.Wrap(err, "h.useCaseLogin.Login")
	}

	out.ConfirmationToken = res.ConfirmationID

	return nil
}
