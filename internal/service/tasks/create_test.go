package service_tasks_test

import (
	"context"
	"emperror.dev/errors"
	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"pets-planify/internal/service/tasks"
	"testing"
)

func TestService_Create(t *testing.T) {
	testID := uuid.New()
	testErr := errors.New("oops")

	testUserID := rand.Int63()

	type testCaseArgs struct {
		userID int64
	}

	tests := []struct {
		name             string
		args             testCaseArgs
		want             *uuid.UUID
		wantErr          error
		mockStorageTasks func(m *MockStorageTasks)
	}{
		{
			name: "success",
			args: testCaseArgs{
				userID: testUserID,
			},
			want:    pointer.To(testID),
			wantErr: nil,
			mockStorageTasks: func(m *MockStorageTasks) {
				m.EXPECT().Create(gomock.Any(), testUserID).Return(pointer.To(testID), nil)
			},
		},
		{
			name: "got_error_from_storage_method",
			args: testCaseArgs{
				userID: testUserID,
			},
			want:    nil,
			wantErr: testErr,
			mockStorageTasks: func(m *MockStorageTasks) {
				m.EXPECT().Create(gomock.Any(), testUserID).Return(nil, testErr)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockStorageTasks := NewMockStorageTasks(ctrl)
			if tt.mockStorageTasks != nil {
				tt.mockStorageTasks(mockStorageTasks)
			}

			service := service_tasks.New(mockStorageTasks)
			got, err := service.Create(context.Background(), tt.args.userID)
			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}
