//go:build integration

package storage_tasks_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"math/rand"
	config_postgres "pets-planify/internal/config/postgres"
	storage_tasks "pets-planify/internal/storage/tasks"
	"testing"
)

func TestStorage_Create(t *testing.T) {
	cfg := config_postgres.New()

	db, err := cfg.GetConnection()
	assert.NoError(t, err)

	storage := storage_tasks.New(db)

	testUserID := rand.Int63()

	taskID, createErr := storage.Create(context.Background(), testUserID)
	assert.NotNil(t, taskID)
	assert.NoError(t, createErr)
}
