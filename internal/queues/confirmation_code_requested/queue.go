package queue_email_confirmation_code

import (
	"context"
	"emperror.dev/errors"
	"encoding/json"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	models_queues "pets-planify/internal/models/queues"
)

const (
	contentTypeJson = "application/json"
)

const (
	queueName = "confirmation_code_requested"
)

type Queue struct {
	channel *amqp.Channel
}

func New(channel *amqp.Channel) *Queue {
	return &Queue{
		channel: channel,
	}
}

func (q *Queue) Produce(ctx context.Context, userID, confirmationID, confirmationToken uuid.UUID) error {
	msg := models_queues.MsgConfirmationCodeRequested{
		UserID:            userID,
		ConfirmationID:    confirmationID,
		ConfirmationToken: confirmationToken,
	}

	enc, err := json.Marshal(msg)
	if err != nil {
		return errors.Wrap(err, "json.Marshal")
	}

	queue, err := q.channel.QueueDeclare(
		queueName, false, false, false, false, nil,
	)
	if err != nil {
		return errors.Wrap(err, "q.channel.QueueDeclare")
	}

	publishErr := q.channel.PublishWithContext(
		ctx, "", queue.Name, false, false,
		amqp.Publishing{Body: enc, ContentType: contentTypeJson},
	)
	if publishErr != nil {
		return errors.Wrap(publishErr, "q.channel.PublishWithContext")
	}

	return nil
}
