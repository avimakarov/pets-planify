package config_postgres

import (
	"database/sql"
	"emperror.dev/errors"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var (
	host = os.Getenv("POSTGRES_HOST_APP")
	port = os.Getenv("POSTGRES_PORT_APP")
	user = os.Getenv("POSTGRES_USER_APP")
	pasw = os.Getenv("POSTGRES_PASW_APP")
	name = os.Getenv("POSTGRES_NAME_APP")
)

type Config struct{}

func New() *Config {
	return &Config{}
}

func (c *Config) dsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		pasw,
		name,
	)
}

func (c *Config) GetConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", c.dsn())
	if err != nil {
		return nil, errors.Wrap(err, "sql.Open")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "db..Ping")
	}

	return db, nil
}
