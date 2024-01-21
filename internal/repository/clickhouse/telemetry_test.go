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
		UserID:     uuid.New(),
		Screen:     "home",
		Action:     "click button 'Profile'",
		Timestamp:  time.Now().UTC(),
		AppVersion: "v1.1.18",
		Country:    "RU",
		OS:         "Android 11",
	}

	err := testRepos.Telemetry.Append(context.Background(), arg)
	require.NoError(t, err)
}
