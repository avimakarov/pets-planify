package handler_server_habit_update

import (
	"context"
	"fmt"
	schema "pets-planify/internal/generated/openapi/server"
	"pets-planify/internal/models"
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

func (h *Handler) HabitUpdate(ctx context.Context, in schema.Habit, out *schema.HabitUpdateOut) error {
	req := models.Habit{
		ID:        in.Id,
		Name:      in.Name,
		UserID:    in.UserId,
		ChatID:    in.ChatId,
		Canceled:  in.Canceled,
		CreatedAt: in.CreatedAt,
	}

	err := h.serviceHabits.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("h.serviceHabits.Update: %w", err)
	}

	res, err := h.serviceHabits.GetByID(ctx, req.ID)
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
