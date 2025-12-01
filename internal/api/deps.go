package api

import (
	"context"

	schema "pets-planify/internal/generated/openapi/server"

	"github.com/google/uuid"
)

type HandlerHabitGet interface {
	HabitGet(ctx context.Context, id uuid.UUID, out *schema.HabitGetOut) error
}

type HandlerHabitUpdate interface {
	HabitUpdate(ctx context.Context, in schema.Habit, out *schema.HabitUpdateOut) error
}
