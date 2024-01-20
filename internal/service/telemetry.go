package service

import (
	"github.com/gin-gonic/gin"
	"github.com/vanya/backend/internal/domain/telemetry"
	repository "github.com/vanya/backend/internal/repository/clickhouse"
)

type TelemetryService struct {
	repo repository.Telemetry
}

func NewTelemetryService(repo repository.Telemetry) *TelemetryService {
	return &TelemetryService{
		repo: repo,
	}
}

func (s *TelemetryService) Append(ctx *gin.Context, inp telemetry.Input) error {
	return s.repo.Append(ctx, repository.AppendTelemetryParams{
		UserID:    inp.UserID,
		ScreenID:  inp.ScreenID,
		Action:    inp.Action,
		Timestamp: inp.Timestamp,
		Country:   inp.Country,
		OS:        inp.OS,
	})
}
