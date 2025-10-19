//go:build integration

package storage_tasks_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"math/rand"
	config_postgres "pets-planify/internal/config/postgres"
	storage_tasks "pets-planify/internal/storage/tasks"
	"testing"
)

func TestStorage_GetByID(t *testing.T) {
	cfg := config_postgres.New()

	db, err := cfg.GetConnection()
	assert.NoError(t, err)

	storage := storage_tasks.New(db)

	testUserID := rand.Int63()

	res, err := storage.GetByID(context.Background(), uuid.New())
	assert.Nil(t, res)
	assert.NoError(t, err)

	taskID, err := storage.Create(context.Background(), testUserID)
	assert.NotNil(t, taskID)
	assert.NoError(t, err)

	res2, err := storage.GetByID(context.Background(), *taskID)
	assert.NotNil(t, res2)
	assert.NoError(t, err)
}
