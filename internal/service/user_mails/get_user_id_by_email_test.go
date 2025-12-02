package service_user_mails_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	service_user_mails "pets-planify/internal/service/user_mails"
	"testing"
)

func TestService_GetUserIdByEmail(t *testing.T) {
	testErr := errors.New("oops")

	testEmail, testUserID := uuid.NewString(), uuid.New()

	tests := []struct {
		name                 string
		want                 *uuid.UUID
		wantErr              error
		mockStorageUserMails func(m *MockStorageUserMails)
	}{
		{
			name: "success",
			want: &testUserID,
			mockStorageUserMails: func(m *MockStorageUserMails) {
				m.EXPECT().GetUserIdByEmail(gomock.Any(), testEmail).Return(&testUserID, nil)
			},
		},
		{
			name:    "error_got_error_from_storage_user_mail",
			wantErr: testErr,
			mockStorageUserMails: func(m *MockStorageUserMails) {
				m.EXPECT().GetUserIdByEmail(gomock.Any(), testEmail).Return(nil, testErr)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockStorageUserMails := NewMockStorageUserMails(ctrl)
			if tt.mockStorageUserMails != nil {
				tt.mockStorageUserMails(mockStorageUserMails)
			}

			service := service_user_mails.New(mockStorageUserMails)
			got, err := service.GetUserIdByEmail(context.Background(), testEmail)
			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr.Error())
			}
		})
	}
}
