package handler_tgb_create

import (
	"emperror.dev/errors"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

const (
	ResponseTemplateSuccess = "Задача создана\nОткрыть задачу: %s"
)

func (h *Handler) Create(ctx *th.Context, msg telego.Update) error {
	chatID := msg.Message.Chat.ID
	userID := msg.Message.From.ID

	id, err := h.serviceTasks.Create(ctx.Context(), userID)
	if err != nil {
		return errors.Wrap(err, "h.serviceTasks.Create")
	}

	result := telegoutil.Messagef(
		telegoutil.ID(chatID), ResponseTemplateSuccess, id.String(),
	)

	if _, sendErr := h.serviceBot.SendMessage(ctx, result); sendErr != nil {
		return errors.Wrap(sendErr, "h.serviceBot.SendMessage")
	}

	return nil
}
