package config_rabbitmq

import (
	"emperror.dev/errors"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"os"
)

var (
	host = os.Getenv("RABBITMQ_HOST_APP")
	port = os.Getenv("RABBITMQ_PORT_APP")
	user = os.Getenv("RABBITMQ_USER_APP")
	pasw = os.Getenv("RABBITMQ_PASW_APP")
)

type Config struct{}

func New() *Config {
	return &Config{}
}

func (c *Config) dsn() string {
	return fmt.Sprintf(
		"amqp://%s:%s@%s:%s/", user, pasw, host, port,
	)
}

func (c *Config) GetChannel() (*amqp.Channel, error) {
	con, err := amqp.Dial(c.dsn())
	if err != nil {
		return nil, errors.Wrap(err, "amqp.Dial")
	}

	chn, err := con.Channel()
	if err != nil {
		return nil, errors.Wrap(err, "con.Channel")
	}

	return chn, nil
}
