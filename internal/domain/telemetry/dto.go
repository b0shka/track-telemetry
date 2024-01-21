package telemetry

import (
	"time"

	"github.com/google/uuid"
)

type Input struct {
	UserID     uuid.UUID `json:"user_id"`
	Screen     string    `json:"screen"`
	Action     string    `json:"action"`
	Timestamp  time.Time `json:"timestamp"`
	AppVersion string    `json:"app_version"`
	Country    string    `json:"country"`
	OS         string    `json:"os"`
}

func NewTelemetryInput(
	userID uuid.UUID,
	screen string,
	action string,
	timestamp time.Time,
	appVersion string,
	country string,
	os string,
) Input {
	return Input{
		UserID:     userID,
		Screen:     screen,
		Action:     action,
		Timestamp:  timestamp,
		AppVersion: appVersion,
		Country:    country,
		OS:         os,
	}
}
