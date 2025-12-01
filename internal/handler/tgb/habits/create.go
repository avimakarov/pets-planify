package handler_tgb_habits

import (
	"fmt"
	"pets-planify/internal/models"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/mymmrac/telego/telegoutil"
)

const (
	CommandSlugCreate      = "habit"
	TemplateResponseCreate = "Привычка создана\nОткрыть: %s"
)

func (h *Handler) Create(ctx *th.Context, msg telego.Update) error {
	chatID := msg.Message.Chat.ID
	userID := msg.Message.From.ID

	id, err := h.serviceHabits.Create(
		ctx.Context(), models.Habit{
			UserID: userID, ChatID: chatID,
		},
	)
	if err != nil {
		return fmt.Errorf("h.serviceHabits.Create: %w", err)
	}

	outMsg := telegoutil.Messagef(
		telegoutil.ID(chatID), TemplateResponseCreate, id.String(),
	)

	if _, sendErr := h.serviceBot.SendMessage(ctx, outMsg); sendErr != nil {
		return fmt.Errorf("h.serviceBot.SendMessage: %w", sendErr)
	}

	return nil
}
