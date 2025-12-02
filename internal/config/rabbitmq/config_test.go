//go:build integration

package config_rabbitmq_test

import (
	"github.com/stretchr/testify/assert"
	config_rabbitmq "pets-planify/internal/config/rabbitmq"
	"testing"
)

func TestConfig_GetChannel(t *testing.T) {
	_, err := config_rabbitmq.New().GetChannel()
	assert.NoError(t, err)
}
