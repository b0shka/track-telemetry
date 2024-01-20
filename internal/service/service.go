package service

import (
	"github.com/gin-gonic/gin"
	"github.com/vanya/backend/internal/domain/telemetry"
	repository "github.com/vanya/backend/internal/repository/clickhouse"
)

type Telemetry interface {
	Append(ctx *gin.Context, inp telemetry.Input) error
}

type Services struct {
	Telemetry
}

type Deps struct {
	Repos *repository.Repositories
}

func NewServices(deps Deps) *Services {
	return &Services{
		Telemetry: NewTelemetryService(deps.Repos.Telemetry),
	}
}
