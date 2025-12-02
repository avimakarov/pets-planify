package service_user_mails_test

import (
	"context"
	"emperror.dev/errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	service_user_mails "pets-planify/internal/service/user_mails"
	"testing"
)

func TestService_ExistsByEmail(t *testing.T) {
	testErr := errors.New("oops")

	testEmail := uuid.NewString()

	tests := []struct {
		name                 string
		want                 bool
		wantErr              error
		mockStorageUserMails func(m *MockStorageUserMails)
	}{
		{
			name: "success",
			want: true,
			mockStorageUserMails: func(m *MockStorageUserMails) {
				m.EXPECT().ExistByEmail(gomock.Any(), testEmail).Return(true, nil)
			},
		},
		{
			name:    "error_got_error_from_storage_method",
			wantErr: testErr,
			mockStorageUserMails: func(m *MockStorageUserMails) {
				m.EXPECT().ExistByEmail(gomock.Any(), testEmail).Return(true, testErr)
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
			got, err := service.ExistsByEmail(context.Background(), testEmail)
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
