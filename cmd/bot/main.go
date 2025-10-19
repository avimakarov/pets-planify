package main

import (
	"context"
	"emperror.dev/errors"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"log"
	"os"
	config_postgres "pets-planify/internal/config/postgres"
	handler_tgb_create "pets-planify/internal/handler/tgb/create"
	service_tasks "pets-planify/internal/service/tasks"
	storage_tasks "pets-planify/internal/storage/tasks"
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

	storageTasks := storage_tasks.New(pgConn)
	serviceTasks := service_tasks.New(storageTasks)

	handlerTaskCreate := handler_tgb_create.New(bot, serviceTasks)

	handler.Handle(handlerTaskCreate.Create, th.CommandEqual(handler_tgb_create.SlugCommandHandler))

	if startErr := handler.Start(); startErr != nil {
		log.Fatalln(errors.Wrap(startErr, "handler.Start"))
	}
}
