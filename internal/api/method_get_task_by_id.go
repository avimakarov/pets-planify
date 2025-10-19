package api

import (
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"log"
	"net/http"
	schema "pets-planify/internal/generated/openapi/server"
	handler_api_get_task_by_id "pets-planify/internal/handler/api/get_task_by_id"
)

func (a *API) GetTaskById(w http.ResponseWriter, r *http.Request) {
	var in schema.GetTaskByIdIn

	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		msg := schema.ErrorMessage{
			Message: ErrInvalidRequestBody.Error(),
		}
		a.writeResponse(w, msg, http.StatusBadRequest) // todo в схеиу
		return
	}

	chn := make(chan error, 1)
	out := new(schema.GetTaskByIdOut)

	ctx, cancel := context.WithTimeout(r.Context(), defaultHandlerDeadline)
	defer cancel()

	go func() {
		chn <- a.handlerGetTaskByID.GetTaskByID(ctx, &in, out)
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
					return
				}
			case errors.Is(err, handler_api_get_task_by_id.ErrTaskNotFound):
				{
					msg := schema.ErrorMessage{
						Message: err.Error(),
					}
					a.writeResponse(w, msg, http.StatusNotFound)
					return
				}
			case errors.Is(err, handler_api_get_task_by_id.ErrTaskIdIsNotValid):
				{
					msg := schema.ErrorMessage{
						Message: err.Error(),
					}
					a.writeResponse(w, msg, http.StatusBadRequest)
					return
				}
			default:
				{
					msg := schema.ErrorMessage{
						Message: ErrInternal.Error(),
					}
					a.writeResponse(w, msg, http.StatusInternalServerError)
					log.Println(errors.Wrap(err, "a.handlerGetTaskByID.GetTaskByID").Error())

					return
				}
			}
		}
	}
}
