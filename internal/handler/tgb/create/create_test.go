package handler_tgb_create_test

import (
	"emperror.dev/errors"
	"github.com/AlekSi/pointer"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/stretchr/testify/assert"
	"math/rand"
	handler_tgb_create "pets-planify/internal/handler/tgb/create"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	testErr := errors.New("oops")

	testTaskID := uuid.New()

	testChatID := rand.Int63()
	testUserID := rand.Int63()

	tests := []struct {
		name             string
		args             telego.Update
		wantErr          error
		mockServiceBot   func(m *MockServiceBot)
		mockServiceTasks func(m *MockServiceTasks)
	}{
		{
			name: "success",
			args: telego.Update{
				Message: &telego.Message{
					Chat: telego.Chat{
						ID: testChatID,
					},
					From: &telego.User{
						ID: testUserID,
					},
				},
			},
			wantErr: nil,
			mockServiceBot: func(m *MockServiceBot) {
				m.EXPECT().SendMessage(gomock.Any(), gomock.Any()).Return(nil, nil)
			},
			mockServiceTasks: func(m *MockServiceTasks) {
				m.EXPECT().Create(gomock.Any(), testUserID).Return(
					&testTaskID, nil,
				)
			},
		},
		{
			name: "got_error_from_task_service",
			args: telego.Update{
				Message: &telego.Message{
					Chat: telego.Chat{
						ID: testChatID,
					},
					From: &telego.User{
						ID: testUserID,
					},
				},
			},
			wantErr: testErr,
			mockServiceBot: func(m *MockServiceBot) {
			},
			mockServiceTasks: func(m *MockServiceTasks) {
				m.EXPECT().Create(gomock.Any(), testUserID).Return(
					nil, testErr,
				)
			},
		},
		{
			name: "got_error_from_tg_bot_service",
			args: telego.Update{
				Message: &telego.Message{
					Chat: telego.Chat{
						ID: testChatID,
					},
					From: &telego.User{
						ID: testUserID,
					},
				},
			},
			wantErr: testErr,
			mockServiceBot: func(m *MockServiceBot) {
				m.EXPECT().SendMessage(gomock.Any(), gomock.Any()).Return(nil, testErr)
			},
			mockServiceTasks: func(m *MockServiceTasks) {
				m.EXPECT().Create(gomock.Any(), testUserID).Return(
					&testTaskID, nil,
				)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			mockServiceBot := NewMockServiceBot(ctrl)
			if tt.mockServiceBot != nil {
				tt.mockServiceBot(mockServiceBot)
			}

			mockServiceTasks := NewMockServiceTasks(ctrl)
			if tt.mockServiceTasks != nil {
				tt.mockServiceTasks(mockServiceTasks)
			}

			got := handler_tgb_create.New(mockServiceBot, mockServiceTasks).Create(
				pointer.To(th.Context{}), tt.args,
			)
			assert.ErrorIs(t, got, tt.wantErr)
		})
	}
}
