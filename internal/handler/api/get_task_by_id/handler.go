package handler_api_get_task_by_id

import (
	"context"
	"emperror.dev/errors"
	"github.com/google/uuid"
	schema "pets-planify/internal/generated/openapi/server"
	"pets-planify/internal/models"
)

type Handler struct {
	serviceTasks ServiceTasks
}

func New(
	serviceTasks ServiceTasks,
) *Handler {
	return &Handler{
		serviceTasks: serviceTasks,
	}
}

var (
	ErrTaskNotFound     = errors.New("task not found")
	ErrTaskIdIsNotValid = errors.New("task id is not valid")
)

func (h *Handler) GetTaskByID(ctx context.Context, in *schema.GetTaskByIdIn, out *schema.GetTaskByIdOut) error {
	if err := h.validate(in); err != nil {
		return err
	}

	task, err := h.serviceTasks.GetByID(ctx, in.TaskId)
	if err != nil {
		return errors.Wrap(err, "h.serviceTasks.GetByID")
	}

	return h.buildResponse(task, out)
}

func (h *Handler) validate(in *schema.GetTaskByIdIn) error {
	if in.TaskId == uuid.Nil {
		return ErrTaskIdIsNotValid
	}

	err := uuid.Validate(in.TaskId.String())
	if err != nil {
		return ErrTaskIdIsNotValid
	}

	return nil
}

func (h *Handler) buildResponse(task *models.Task, out *schema.GetTaskByIdOut) error {
	if task == nil {
		return ErrTaskNotFound
	}

	res := new(schema.Task)

	res.Id = task.UUID
	res.Name = task.Name
	res.UserId = task.UserID
	res.IsDone = task.IsDone
	res.PlanedTo = task.PlanedTo
	res.PlanedFrom = task.PlanedFrom
	res.Description = task.Description

	out.Task = res

	return nil
}
