package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/corpix/uarand"
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vanya/backend/internal/domain/telemetry"
	"github.com/vanya/backend/internal/service"
	"github.com/vanya/backend/pkg/utils"
)

func TestHandler_ipIdentity(t *testing.T) {
	testTable := []struct {
		name      string
		setupAuth func(
			t *testing.T,
			request *http.Request,
			geoip *geoip2.Reader,
		)
		statusCode   int
		responseBody string
	}{
		{
			name: "ok",
			setupAuth: func(t *testing.T, request *http.Request, geoip *geoip2.Reader) {
				request.Header.Set("User-Agent", uarand.GetRandom())
			},
			statusCode:   200,
			responseBody: "",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			reader, err := geoip2.Open("../../../GeoLite2-Country-Test.mmdb")
			require.NoError(t, err)

			handler := Handler{
				Services: &service.Services{},
				Geoip:    reader,
				Logger:   &logrus.Logger{},
			}

			router := gin.Default()
			router.GET(
				"/identity",
				handler.ipIdentity(reader),
				func(c *gin.Context) {
					c.Status(http.StatusOK)
				},
			)

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/identity", nil)

			testCase.setupAuth(t, req, reader)
			router.ServeHTTP(recorder, req)

			assert.Equal(t, testCase.statusCode, recorder.Code)
			assert.Equal(t, testCase.responseBody, recorder.Body.String())
		})
	}
}

func TestGetPayload(t *testing.T) {
	payloadContext := telemetry.NewPayloadContext("RU", "Linux x86_64")

	normalContext := &gin.Context{}
	normalContext.Set(payloadCtx, payloadContext)

	key, err := utils.RandomString(10)
	require.NoError(t, err)

	invalidContext := &gin.Context{}
	invalidContext.Set(payloadCtx, key)

	tests := []struct {
		name      string
		ctx       *gin.Context
		payload   telemetry.PayloadContext
		shouldErr bool
	}{
		{
			name:      "ok",
			ctx:       normalContext,
			payload:   payloadContext,
			shouldErr: false,
		},
		{
			name:      "empty context",
			ctx:       &gin.Context{},
			shouldErr: true,
		},
		{
			name:      "invalid payload",
			ctx:       invalidContext,
			shouldErr: true,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			payload, err := getPayloadByContext(testCase.ctx)

			if testCase.shouldErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, testCase.payload, payload)
		})
	}
}
