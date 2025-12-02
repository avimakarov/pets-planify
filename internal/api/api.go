package api

import (
	"time"

	"emperror.dev/errors"
)

var (
	ErrInternal           = errors.New("internal error")
	ErrRequestDeadline    = errors.New("deadline exceeded")
	ErrInvalidRequestBody = errors.New("invalid request body")
)

const (
	defaultHandlerDeadline = time.Second
)

type API struct {
	handlerLogin HandlerLogin
}

func New(
	handlerLogin HandlerLogin,
) *API {
	return &API{
		handlerLogin: handlerLogin,
	}
}
