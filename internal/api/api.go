package api

import (
	"emperror.dev/errors"
	"time"
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
	handlerInfo        HandlerInfo
	handlerGetTaskByID HandlerGetTaskByID
}

func New(
	handlerInfo HandlerInfo,
	handlerGetTaskByID HandlerGetTaskByID,
) *API {
	return &API{
		handlerInfo:        handlerInfo,
		handlerGetTaskByID: handlerGetTaskByID,
	}
}
