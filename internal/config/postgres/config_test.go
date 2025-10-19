//go:build integration

package config_postgres_test

import (
	"github.com/stretchr/testify/assert"
	config_postgres "pets-planify/internal/config/postgres"
	"testing"
)

func TestConfig_GetConnection(t *testing.T) {
	cfg := config_postgres.New()

	_, err := cfg.GetConnection()
	assert.NoError(t, err)
}
