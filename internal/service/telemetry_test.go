package service_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/vanya/backend/internal/domain/telemetry"
	mock_repository "github.com/vanya/backend/internal/repository/clickhouse/mocks"
	"github.com/vanya/backend/internal/service"
)

func mockTelemetryService(t *testing.T) (
	*service.TelemetryService,
	*mock_repository.MockTelemetry,
) {
	repoCtl := gomock.NewController(t)
	defer repoCtl.Finish()

	telemetryRepo := mock_repository.NewMockTelemetry(repoCtl)
	telemetryService := service.NewTelemetryService(telemetryRepo)

	return telemetryService, telemetryRepo
}

func TestTelemetryService_Append(t *testing.T) {
	telemetryService, telemetryRepo := mockTelemetryService(t)

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	telemetryRepo.EXPECT().Append(ctx, gomock.Any())

	err := telemetryService.Append(ctx, telemetry.Input{})
	require.NoError(t, err)
}
