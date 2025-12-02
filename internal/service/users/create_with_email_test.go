package service_users_test

import (
	"context"
	"database/sql"
	"emperror.dev/errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	service_users "pets-planify/internal/service/users"
	"testing"
)

func TestService_CreateWithEmail(t *testing.T) {
	testErr := errors.New("ops")
	testMail := uuid.NewString()

	tests := []struct {
		name                 string
		mockTx               func(m *MockTx)
		wantErr              error
		mockStorage          func(m *MockStorage, tx *MockTx)
		mockStorageUsers     func(m *MockStorageUsers)
		mockStorageUserMails func(m *MockStorageUserMails)
	}{
		{
			name: "success",
			mockTx: func(m *MockTx) {
				m.EXPECT().Commit().Return(nil)
				m.EXPECT().Rollback().Return(sql.ErrTxDone)
			},
			mockStorage: func(m *MockStorage, tx *MockTx) {
				m.EXPECT().Tx().Return(tx, nil)
			},
			mockStorageUsers: func(m *MockStorageUsers) {
				m.EXPECT().CreateWithTx(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			mockStorageUserMails: func(m *MockStorageUserMails) {
				m.EXPECT().CreateWithTx(gomock.Any(), gomock.Any(), testMail, gomock.Any()).Return(nil)
			},
		},
		{
			name:    "error_got_error_from_tx_begin",
			wantErr: testErr,
			mockTx: func(m *MockTx) {
			},
			mockStorage: func(m *MockStorage, tx *MockTx) {
				m.EXPECT().Tx().Return(tx, testErr)
			},
			mockStorageUsers: func(m *MockStorageUsers) {
			},
			mockStorageUserMails: func(m *MockStorageUserMails) {
			},
		},
		{
			name:    "error_got_error_from_tx_commit",
			wantErr: testErr,
			mockTx: func(m *MockTx) {
				m.EXPECT().Commit().Return(testErr)
				m.EXPECT().Rollback().Return(nil)
			},
			mockStorage: func(m *MockStorage, tx *MockTx) {
				m.EXPECT().Tx().Return(tx, nil)
			},
			mockStorageUsers: func(m *MockStorageUsers) {
				m.EXPECT().CreateWithTx(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			mockStorageUserMails: func(m *MockStorageUserMails) {
				m.EXPECT().CreateWithTx(gomock.Any(), gomock.Any(), testMail, gomock.Any()).Return(nil)
			},
		},
		{
			name:    "error_got_error_from_storage_users",
			wantErr: testErr,
			mockTx: func(m *MockTx) {
				m.EXPECT().Rollback().Return(nil)
			},
			mockStorage: func(m *MockStorage, tx *MockTx) {
				m.EXPECT().Tx().Return(tx, nil)
			},
			mockStorageUsers: func(m *MockStorageUsers) {
				m.EXPECT().CreateWithTx(gomock.Any(), gomock.Any(), gomock.Any()).Return(testErr)
			},
			mockStorageUserMails: func(m *MockStorageUserMails) {
			},
		},
		{
			name:    "error_got_error_from_storage_user_mails",
			wantErr: testErr,
			mockTx: func(m *MockTx) {
				m.EXPECT().Rollback().Return(nil)
			},
			mockStorage: func(m *MockStorage, tx *MockTx) {
				m.EXPECT().Tx().Return(tx, nil)
			},
			mockStorageUsers: func(m *MockStorageUsers) {
				m.EXPECT().CreateWithTx(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
			mockStorageUserMails: func(m *MockStorageUserMails) {
				m.EXPECT().CreateWithTx(gomock.Any(), gomock.Any(), testMail, gomock.Any()).Return(testErr)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockTx := NewMockTx(ctrl)
			if tt.mockTx != nil {
				tt.mockTx(mockTx)
			}

			mockStorage := NewMockStorage(ctrl)
			if tt.mockStorage != nil {
				tt.mockStorage(mockStorage, mockTx)
			}

			mockStorageUsers := NewMockStorageUsers(ctrl)
			if tt.mockStorageUsers != nil {
				tt.mockStorageUsers(mockStorageUsers)
			}

			mockStorageUserMails := NewMockStorageUserMails(ctrl)
			if tt.mockStorageUserMails != nil {
				tt.mockStorageUserMails(mockStorageUserMails)
			}

			service := service_users.New(mockStorage, mockStorageUsers, mockStorageUserMails)
			_, err := service.CreateWithEmail(context.Background(), testMail)
			if tt.wantErr != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
