package handler_server_get

import (
	"context"
	"fmt"
	schema "pets-planify/internal/generated/openapi/server"

	"github.com/google/uuid"
)

type Handler struct {
	serviceHabits ServiceHabits
}

func New(
	serviceHabits ServiceHabits,
) *Handler {
	return &Handler{
		serviceHabits: serviceHabits,
	}
}

func (h *Handler) HabitGet(ctx context.Context, id uuid.UUID, out *schema.HabitGetOut) error {
	res, err := h.serviceHabits.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("h.serviceHabits.GetByID: %w", err)
	}

	out.Habit = &schema.Habit{
		Id:        res.ID,
		Name:      res.Name,
		UserId:    res.UserID,
		ChatId:    res.ChatID,
		Canceled:  res.Canceled,
		CreatedAt: res.CreatedAt,
	}

	return nil
}
