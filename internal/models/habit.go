package models

import (
	"time"

	"github.com/google/uuid"
)

type Habit struct {
	ID        uuid.UUID
	Name      *string
	UserID    int64
	ChatID    int64
	Canceled  bool
	CreatedAt time.Time
}
