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
	handlerHabitGet    HandlerHabitGet
	handlerHabitUpdate HandlerHabitUpdate
}

func New(
	handlerHabitGet HandlerHabitGet,
	handlerHabitUpdate HandlerHabitUpdate,
) *API {
	return &API{
		handlerHabitGet:    handlerHabitGet,
		handlerHabitUpdate: handlerHabitUpdate,
	}
}
