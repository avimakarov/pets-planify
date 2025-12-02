package main

import (
	"log"
	"net/http"
	"pets-planify/internal/api"
	config_postgres "pets-planify/internal/config/postgres"
	config_rabbitmq "pets-planify/internal/config/rabbitmq"
	server_api_gen "pets-planify/internal/generated/openapi/server"
	handler_login "pets-planify/internal/handler/server/login"
	queue_email_confirmation_code "pets-planify/internal/queues/confirmation_code_requested"
	service_user_mails "pets-planify/internal/service/user_mails"
	service_users "pets-planify/internal/service/users"
	storage "pets-planify/internal/storage"
	storage_user_mails "pets-planify/internal/storage/user_mails"
	storage_users "pets-planify/internal/storage/users"
	usecase_login "pets-planify/internal/usecase/login"
	"time"

	"emperror.dev/errors"
	"github.com/go-chi/chi/v5"
)

func main() {
	time.Sleep(10 * time.Second)

	configPostgres := config_postgres.New()
	configRabbitmq := config_rabbitmq.New()

	ch, err := configRabbitmq.GetChannel()
	if err != nil {
		log.Panicln(errors.Wrap(err, "configRabbitmq.GetChannel"))
	}

	db, err := configPostgres.GetConnection()
	if err != nil {
		log.Panicln(errors.Wrap(err, "configPostgres.GetConnection"))
	}

	storage := storage.New(db)
	storageUsers := storage_users.New(db)
	storageUserMails := storage_user_mails.New(db)

	queueConfirmationCodeRequested := queue_email_confirmation_code.New(ch)

	serviceUsers := service_users.New(storage, storageUsers, storageUserMails)
	serviceUserMails := service_user_mails.New(storageUserMails)

	useCaseLogin := usecase_login.New(serviceUsers, serviceUserMails, queueConfirmationCodeRequested)

	handlerLogin := handler_login.New(useCaseLogin)

	server := api.New(
		handlerLogin,
	)

	router := chi.NewRouter()
	handler := server_api_gen.HandlerFromMux(server, router)

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Println(errors.Wrap(err, "http.ListenAndServe").Error())
	}
}
