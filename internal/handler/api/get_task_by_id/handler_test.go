package handler_api_get_task_by_id_test

import (
	"context"
	"emperror.dev/errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	schema "pets-planify/internal/generated/openapi/server"
	handler_api_get_task_by_id "pets-planify/internal/handler/api/get_task_by_id"
	"pets-planify/internal/models"
	"testing"
)

func TestHandler_GetTaskByID(t *testing.T) {
	testErr := errors.New("oops")

	testTaskID := uuid.New()

	tests := []struct {
		name             string
		args             *schema.GetTaskByIdIn
		want             *schema.GetTaskByIdOut
		wantErr          error
		mockServiceTasks func(m *MockServiceTasks)
	}{
		{
			name: "success",
			args: &schema.GetTaskByIdIn{
				TaskId: testTaskID,
			},
			want: &schema.GetTaskByIdOut{
				Task: &schema.Task{
					Id: testTaskID,
				},
			},
			wantErr: nil,
			mockServiceTasks: func(m *MockServiceTasks) {
				m.EXPECT().GetByID(gomock.Any(), testTaskID).Return(
					&models.Task{
						UUID: testTaskID,
					},
					nil,
				)
			},
		},
		{
			name: "success_task_not_found",
			args: &schema.GetTaskByIdIn{
				TaskId: testTaskID,
			},
			want: &schema.GetTaskByIdOut{
				Task: &schema.Task{
					Id: testTaskID,
				},
			},
			wantErr: handler_api_get_task_by_id.ErrTaskNotFound,
			mockServiceTasks: func(m *MockServiceTasks) {
				m.EXPECT().GetByID(gomock.Any(), testTaskID).Return(
					nil,
					nil,
				)
			},
		},
		{
			name: "got_invalid_request",
			args: &schema.GetTaskByIdIn{
				TaskId: uuid.UUID{},
			},
			want:    new(schema.GetTaskByIdOut),
			wantErr: handler_api_get_task_by_id.ErrTaskIdIsNotValid,
			mockServiceTasks: func(m *MockServiceTasks) {
			},
		},
		{
			name: "got_error_from_service",
			args: &schema.GetTaskByIdIn{
				TaskId: testTaskID,
			},
			want:    new(schema.GetTaskByIdOut),
			wantErr: testErr,
			mockServiceTasks: func(m *MockServiceTasks) {
				m.EXPECT().GetByID(gomock.Any(), testTaskID).Return(
					nil, testErr,
				)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockServiceTasks := NewMockServiceTasks(ctrl)
			if tt.mockServiceTasks != nil {
				tt.mockServiceTasks(mockServiceTasks)
			}

			handler := handler_api_get_task_by_id.New(mockServiceTasks)

			out := new(schema.GetTaskByIdOut)
			err := handler.GetTaskByID(context.Background(), tt.args, out)
			if tt.wantErr == nil {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, out)
			} else {
				assert.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}
