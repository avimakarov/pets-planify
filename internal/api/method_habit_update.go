package api

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	schema "pets-planify/internal/generated/openapi/server"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (a *API) HabitUpdate(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	var in schema.Habit

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		msg := schema.HabitUpdateOut{
			Err: &schema.Error{
				Message: ErrInvalidRequestBody.Error(),
			},
		}
		a.writeResponse(w, msg, http.StatusBadRequest)
		return
	}

	chn := make(chan error, 1)
	out := new(schema.HabitUpdateOut)

	ctx, cancel := context.WithTimeout(r.Context(), defaultHandlerDeadline)
	defer cancel()

	go func() {
		chn <- a.handlerHabitUpdate.HabitUpdate(ctx, in, out)
	}()

	select {
	case <-ctx.Done():
		{
			msg := schema.HabitUpdateOut{
				Err: &schema.Error{
					Message: ErrRequestDeadline.Error(),
				},
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
					msg := schema.HabitUpdateOut{
						Err: &schema.Error{
							Message: ErrInternal.Error(),
						},
					}
					a.writeResponse(w, msg, http.StatusInternalServerError)
				}
			}
		}
	}
}
