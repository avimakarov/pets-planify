package main

import (
	"log"
	"net/http"
	"pets-planify/internal/api"
	config_postgres "pets-planify/internal/config/postgres"
	server_api_gen "pets-planify/internal/generated/openapi/server"
	handler_server_get "pets-planify/internal/handler/server/habit_get"
	handler_server_habit_update "pets-planify/internal/handler/server/habit_update"
	service_habits "pets-planify/internal/service/habits"
	storage_habits "pets-planify/internal/storage/habits"

	"emperror.dev/errors"
	"github.com/go-chi/chi/v5"
)

func main() {
	configPostgres := config_postgres.New()

	db, err := configPostgres.GetConnection()
	if err != nil {
		log.Panicln(errors.Wrap(err, "configPostgres.GetConnection"))
	}

	storageHabits := storage_habits.New(db)
	serviceHabits := service_habits.New(storageHabits)

	handlerHabitGet := handler_server_get.New(serviceHabits)
	handlerHabitUpdate := handler_server_habit_update.New(serviceHabits)

	server := api.New(
		handlerHabitGet,
		handlerHabitUpdate,
	)

	router := chi.NewRouter()
	handler := server_api_gen.HandlerFromMux(server, router)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Println(errors.Wrap(err, "http.ListenAndServe").Error())
	}
}
