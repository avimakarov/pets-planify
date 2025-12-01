//go:build integration

package config_postgres_test

import (
	config_postgres "pets-planify/internal/config/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig_GetConnection(t *testing.T) {
	_, err := config_postgres.New().GetConnection()
	assert.NoError(t, err)
}
