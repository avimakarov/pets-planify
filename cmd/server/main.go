package main

import (
	"emperror.dev/errors"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"pets-planify/internal/api"
	config_postgres "pets-planify/internal/config/postgres"
	server_gen "pets-planify/internal/generated/openapi/server"
	handler_api_get_task_by_id "pets-planify/internal/handler/api/get_task_by_id"
	handler_api_info "pets-planify/internal/handler/api/info"
	service_tasks "pets-planify/internal/service/tasks"
	storage_tasks "pets-planify/internal/storage/tasks"
)

func main() {
	cfgPostgres := config_postgres.New()

	dbPostgres, err := cfgPostgres.GetConnection()
	if err != nil {
		log.Fatalln(errors.Wrap(err, "configPg.GetConnection"))
	}

	storageTasks := storage_tasks.New(dbPostgres)
	serviceTasks := service_tasks.New(storageTasks)

	handlerApiInfo := handler_api_info.New()
	handlerApiGetTaskByID := handler_api_get_task_by_id.New(serviceTasks)

	server := api.New(
		handlerApiInfo,
		handlerApiGetTaskByID,
	)
	router := chi.NewRouter()

	handler := server_gen.HandlerFromMux(server, router)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Println(errors.Wrap(err, "http.ListenAndServe").Error())
	}
}
