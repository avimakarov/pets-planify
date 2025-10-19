package api

import (
	"context"
	"emperror.dev/errors"
	"log"
	"net/http"
	schema "pets-planify/internal/generated/openapi/server"
)

func (a *API) Info(w http.ResponseWriter, r *http.Request) {
	chn := make(chan error, 1)
	out := new(schema.InfoOut)

	ctx, cancel := context.WithTimeout(r.Context(), defaultHandlerDeadline)
	defer cancel()

	go func() {
		chn <- a.handlerInfo.Info(ctx, out)
	}()

	select {
	case <-ctx.Done():
		{
			msg := schema.ErrorMessage{
				Message: ErrRequestDeadline.Error(),
			}
			a.writeResponse(w, msg, http.StatusRequestTimeout)
		}
	case err := <-chn:
		{
			switch {
			case errors.Is(err, nil):
				{
					a.writeResponse(w, out, http.StatusOK)
				}
			default:
				{
					msg := schema.ErrorMessage{
						Message: ErrInternal.Error(),
					}
					a.writeResponse(w, msg, http.StatusInternalServerError)
					log.Println(errors.Wrap(err, "a.handlerInfo.Info").Error())
				}
			}
		}
	}
}
