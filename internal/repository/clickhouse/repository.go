package repository

import (
	"context"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type Telemetry interface {
	Append(ctx context.Context, arg AppendTelemetryParams) error
}

type Repositories struct {
	Telemetry
}

func NewRepositories(db driver.Conn) *Repositories {
	return &Repositories{
		Telemetry: NewTelemetryRepo(db),
	}
}
