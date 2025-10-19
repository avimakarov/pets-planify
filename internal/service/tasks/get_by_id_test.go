package service_tasks_test

import (
	"context"
	"emperror.dev/errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"pets-planify/internal/models"
	service_tasks "pets-planify/internal/service/tasks"
	"testing"
)

func TestService_GetByID(t *testing.T) {
	testErr := errors.New("oops")

	testTaskID := uuid.New()

	type testCaseArgs struct {
		ID uuid.UUID
	}

	tests := []struct {
		name             string
		args             testCaseArgs
		want             *models.Task
		wantErr          error
		mockStorageTasks func(m *MockStorageTasks)
	}{
		{
			name: "success",
			args: testCaseArgs{
				ID: testTaskID,
			},
			want: &models.Task{
				UUID: testTaskID,
			},
			wantErr: nil,
			mockStorageTasks: func(m *MockStorageTasks) {
				m.EXPECT().GetByID(gomock.Any(), testTaskID).Return(
					&models.Task{
						UUID: testTaskID,
					},
					nil,
				)
			},
		},
		{
			name: "got_error_from_storage_method",
			args: testCaseArgs{
				ID: testTaskID,
			},
			want:    nil,
			wantErr: testErr,
			mockStorageTasks: func(m *MockStorageTasks) {
				m.EXPECT().GetByID(gomock.Any(), testTaskID).Return(
					nil,
					testErr,
				)
			},
		},
		{
			name: "got_nil_response_from_storage_method",
			args: testCaseArgs{
				ID: testTaskID,
			},
			want:    nil,
			wantErr: nil,
			mockStorageTasks: func(m *MockStorageTasks) {
				m.EXPECT().GetByID(gomock.Any(), testTaskID).Return(
					nil, nil,
				)
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
			got, err := service.GetByID(context.Background(), tt.args.ID)
			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			} else {
				assert.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}
