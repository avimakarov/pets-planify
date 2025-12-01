package api

import (
	"context"
	"errors"
	"net/http"
	schema "pets-planify/internal/generated/openapi/server"

	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (a *API) HabitGet(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	chn := make(chan error, 1)
	out := new(schema.HabitGetOut)

	ctx, cancel := context.WithTimeout(r.Context(), defaultHandlerDeadline)
	defer cancel()

	go func() {
		chn <- a.handlerHabitGet.HabitGet(ctx, uuid.UUID(id), out)
	}()

	select {
	case <-ctx.Done():
		{
			msg := schema.HabitGetOut{
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
					msg := schema.HabitGetOut{
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
