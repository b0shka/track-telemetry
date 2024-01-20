package http

import (
	"github.com/google/uuid"
	"github.com/vanya/backend/internal/domain/telemetry"
)

func NewTelemetryInput(
	req telemetryRequest,
	payload telemetry.PayloadContext,
) (telemetry.Input, error) {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return telemetry.Input{}, err
	}

	return telemetry.NewTelemetryInput(
		userID,
		req.ScreenID,
		req.Action,
		req.Timestamp,
		payload.Country,
		payload.OS,
	), nil
}
