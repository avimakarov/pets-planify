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

func TestStorage_ExistByEmail(t *testing.T) {
	cf := config_postgres.New()

	db, err := cf.GetConnection()
	assert.NoError(t, err)

	testMailID := uuid.NewString()
	testUserID := uuid.NewString()
	testUserEmail1 := uuid.NewString()
	testUserEmail2 := uuid.NewString()

	_, execErr1 := db.Exec(`insert into users (id) values ($1)`, testUserID)
	assert.NoError(t, execErr1)

	_, execErr2 := db.Exec(`insert into user_mails (id, email, user_id) values ($1, $2, $3)`, testMailID, testUserEmail1, testUserID)
	assert.NoError(t, execErr2)

	st := storage_user_mails.New(db)

	resp1, err := st.ExistByEmail(context.Background(), testUserEmail1)
	assert.True(t, resp1)
	assert.NoError(t, err)

	resp2, err := st.ExistByEmail(context.Background(), testUserEmail2)
	assert.False(t, resp2)
	assert.NoError(t, err)
}
