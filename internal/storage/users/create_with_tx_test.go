//go:build integration

package storage_users_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	config_postgres "pets-planify/internal/config/postgres"
	storage_users "pets-planify/internal/storage/users"
	"testing"
)

func TestStorage_CreateWithTx(t *testing.T) {
	cf := config_postgres.New()

	db, err := cf.GetConnection()
	assert.NoError(t, err)

	testUserID := uuid.New()

	st := storage_users.New(db)

	tx, err := db.Begin()
	assert.NoError(t, err)

	createErr := st.CreateWithTx(context.Background(), testUserID, tx)
	assert.NoError(t, createErr)

	assert.NoError(t, tx.Commit())
}
