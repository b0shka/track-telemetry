package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/oschwald/geoip2-golang"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/vanya/backend/internal/domain/telemetry"
	"github.com/vanya/backend/internal/service"
	mock_service "github.com/vanya/backend/internal/service/mocks"
)

var ErrInternalServerError = errors.New("test: internal server error")

func TestHandler_track(t *testing.T) {
	type mockBehavior func(s *mock_service.MockTelemetry, input telemetry.Input)

	userID := uuid.New()
	timestamp := time.Now().UTC()

	tests := []struct {
		name         string
		body         gin.H
		userInput    telemetry.Input
		mockBehavior mockBehavior
		statusCode   int
		responseBody string
	}{
		{
			name: "ok",
			body: gin.H{
				"user_id":   userID,
				"screen_id": 1,
				"action":    "click",
				"timestamp": timestamp,
			},
			userInput: telemetry.Input{
				UserID:    userID,
				ScreenID:  1,
				Action:    "click",
				Timestamp: timestamp,
			},
			mockBehavior: func(s *mock_service.MockTelemetry, input telemetry.Input) {
				s.EXPECT().Append(gomock.Any(), input).Return(nil)
			},
			statusCode:   200,
			responseBody: "",
		},
		{
			name: "error telemetry append",
			body: gin.H{
				"user_id":   userID,
				"screen_id": 1,
				"action":    "click",
				"timestamp": timestamp,
			},
			mockBehavior: func(s *mock_service.MockTelemetry, input telemetry.Input) {
				s.EXPECT().Append(gomock.Any(), gomock.Any()).
					Return(ErrInternalServerError)
			},
			statusCode:   500,
			responseBody: fmt.Sprintf(`{"message":"%s"}`, ErrInternalServerError),
		},
		{
			name:         "empty fields",
			body:         gin.H{},
			mockBehavior: func(s *mock_service.MockTelemetry, input telemetry.Input) {},
			statusCode:   400,
			responseBody: `{"message":"invalid input body"}`,
		},
		{
			name: "invalid uuid",
			body: gin.H{
				"user_id":   "944c1052-a14b-47e7-9efa-31b7900bda3",
				"screen_id": 1,
				"action":    "click",
				"timestamp": timestamp,
			},
			mockBehavior: func(s *mock_service.MockTelemetry, input telemetry.Input) {},
			statusCode:   400,
			responseBody: `{"message":"invalid UUID length: 35"}`,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			mockCtl := gomock.NewController(t)
			defer mockCtl.Finish()

			telemetryService := mock_service.NewMockTelemetry(mockCtl)
			testCase.mockBehavior(telemetryService, testCase.userInput)

			services := &service.Services{Telemetry: telemetryService}
			handler := Handler{
				Services: services,
				Geoip:    &geoip2.Reader{},
				Logger:   &logrus.Logger{},
			}

			reader, err := geoip2.Open("../../../GeoLite2-Country-Test.mmdb")
			require.NoError(t, err)

			router := gin.Default()
			router.POST(
				"/track",
				handler.ipIdentity(reader),
				handler.track,
			)

			recorder := httptest.NewRecorder()

			data, err := json.Marshal(testCase.body)
			require.NoError(t, err)

			req := httptest.NewRequest(
				http.MethodPost,
				"/track",
				bytes.NewReader(data),
			)

			router.ServeHTTP(recorder, req)

			require.Equal(t, testCase.statusCode, recorder.Code)
			require.Equal(t, testCase.responseBody, recorder.Body.String())
		})
	}
}
