package main

import (
	"context"
	"log"
	"os"
	config_postgres "pets-planify/internal/config/postgres"
	handler_tgb_habits "pets-planify/internal/handler/tgb/habits"
	service_habits "pets-planify/internal/service/habits"
	storage_habits "pets-planify/internal/storage/habits"

	"emperror.dev/errors"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

var (
	token = os.Getenv("BOT_TOKEN")
)

func main() {
	ctx := context.Background()

	bot, err := telego.NewBot(token)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "telego.NewBot"))
	}

	upd, err := bot.UpdatesViaLongPolling(ctx, nil)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "bot.UpdatesViaLongPolling"))
	}

	handler, err := th.NewBotHandler(bot, upd)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "th.NewBotHandler"))
	}

	pgConn, err := config_postgres.New().GetConnection()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "config_postgres.New.GetConnection"))
	}

	storageHabits := storage_habits.New(pgConn)
	serviceHabits := service_habits.New(storageHabits)

	handlerHabits := handler_tgb_habits.New(bot, serviceHabits)

	handler.Handle(handlerHabits.Create, th.CommandEqual(handler_tgb_habits.CommandSlugCreate))

	if startErr := handler.Start(); startErr != nil {
		log.Fatalln(errors.Wrap(startErr, "handler.Start"))
	}
}
