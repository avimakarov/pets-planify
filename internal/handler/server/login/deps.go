package handler_login

import (
	"context"
	usecase_login "pets-planify/internal/usecase/login"
)

type UseCaseLogin interface {
	Login(ctx context.Context, in usecase_login.LoginIn) (*usecase_login.LoginOut, error)
}
