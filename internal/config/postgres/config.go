package config_postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/lib/pq"
)

var (
	host = os.Getenv("POSTGRES_HOST_APP")
	port = os.Getenv("POSTGRES_PORT_APP")
	user = os.Getenv("POSTGRES_USER_APP")
	pasw = os.Getenv("POSTGRES_PASW_APP")
	name = os.Getenv("POSTGRES_NAME_APP")
)

func init() {
	v := map[string]string{
		"POSTGRES_HOST_APP": host,
		"POSTGRES_PORT_APP": port,
		"POSTGRES_USER_APP": user,
		"POSTGRES_PASW_APP": pasw,
		"POSTGRES_NAME_APP": name,
	}

	for k, v := range v {
		if len(strings.TrimSpace(v)) == 0 {
			log.Fatalln(fmt.Errorf("env is empty: %s", k))
		}
	}
}

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
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("db..Ping: %w", err)
	}

	return db, nil
}
