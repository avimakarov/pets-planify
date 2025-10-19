package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	UUID        uuid.UUID  `json:"uuid"`
	Name        *string    `json:"name"`
	UserID      int64      `json:"user_id"`
	IsDone      bool       `json:"is_done"`
	PlanedTo    *time.Time `json:"planed_to"`
	CreatedAt   *time.Time `json:"created_at"`
	PlanedFrom  *time.Time `json:"planed_from"`
	Description *string    `json:"description"`
}
