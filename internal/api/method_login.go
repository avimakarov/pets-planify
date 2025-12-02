package api

import (
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"log"
	"net/http"
	schema "pets-planify/internal/generated/openapi/server"
)

func (a *API) Login(w http.ResponseWriter, r *http.Request) {
	in := new(schema.LoginIn)
	out := new(schema.LoginOut)

	if err := json.NewDecoder(r.Body).Decode(in); err != nil {
		a.writeResponse(w, out, http.StatusBadRequest)
		return
	}

	chn := make(chan error, 1)
	ctx, cancel := context.WithTimeout(r.Context(), defaultHandlerDeadline)
	defer cancel()

	go func() {
		chn <- a.handlerLogin.Login(ctx, in, out)
	}()

	select {
	case <-ctx.Done():
		{
			a.writeResponse(w, out, http.StatusRequestTimeout)
			return
		}
	case err := <-chn:
		{
			switch {
			case errors.Is(err, nil):
				{
					a.writeResponse(w, out, http.StatusOK)
					return
				}
			default:
				{
					a.writeResponse(w, out, http.StatusInternalServerError)
					log.Println(errors.Wrap(err, "a.handlerGetSuitsByInn.GetSuitsByInn").Error())

					return
				}
			}
		}
	}
}
