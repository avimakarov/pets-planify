package models_queues

import "github.com/google/uuid"

type MsgConfirmationCodeRequested struct {
	UserID            uuid.UUID
	ConfirmationID    uuid.UUID
	ConfirmationToken uuid.UUID
}
