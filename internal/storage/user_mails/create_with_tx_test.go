//go:build integration

package storage_user_mails_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	config_postgres "pets-planify/internal/config/postgres"
	storage_user_mails "pets-planify/internal/storage/user_mails"
	"testing"
)

func TestStorage_CreateWithTx(t *testing.T) {
	cf := config_postgres.New()

	db, err := cf.GetConnection()
	assert.NoError(t, err)

	testUserID := uuid.New()
	testUserEmail1 := uuid.NewString()

	_, execErr1 := db.Exec(`insert into users (id) values ($1)`, testUserID)
	assert.NoError(t, execErr1)

	st := storage_user_mails.New(db)

	tx, txErr := db.Begin()
	assert.NoError(t, txErr)

	createErr := st.CreateWithTx(context.Background(), testUserID, testUserEmail1, tx)
	assert.NoError(t, createErr)

	assert.NoError(t, tx.Commit())
}
