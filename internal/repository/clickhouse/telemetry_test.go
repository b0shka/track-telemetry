package repository

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestTelemetryRepository_Append(t *testing.T) {
	arg := AppendTelemetryParams{
		UserID:    uuid.New(),
		ScreenID:  1,
		Action:    "click",
		Timestamp: time.Now().UTC(),
	}

	err := testRepos.Telemetry.Append(context.Background(), arg)
	require.NoError(t, err)
}
