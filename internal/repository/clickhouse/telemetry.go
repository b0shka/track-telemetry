package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/google/uuid"
)

type TelemetryRepo struct {
	db driver.Conn
}

func NewTelemetryRepo(db driver.Conn) *TelemetryRepo {
	return &TelemetryRepo{
		db: db,
	}
}

type AppendTelemetryParams struct {
	UserID    uuid.UUID `json:"user_id"`
	ScreenID  uint32    `json:"screen_id"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
	Country   string    `json:"country"`
	OS        string    `json:"os"`
}

func (r *TelemetryRepo) Append(ctx context.Context, row AppendTelemetryParams) error {
	q := `INSERT INTO %s (ts, user_id, screen_id, action, country, os)`

	batch, err := r.db.PrepareBatch(ctx, fmt.Sprintf(q, ActionsTable))
	if err != nil {
		return err
	}

	err = batch.Append(
		row.Timestamp,
		row.UserID,
		row.ScreenID,
		row.Action,
		row.Country,
		row.OS,
	)

	if err != nil {
		return err
	}

	return batch.Send()
}
