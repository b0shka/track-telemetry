package telemetry

import (
	"time"

	"github.com/google/uuid"
)

type Input struct {
	UserID    uuid.UUID `json:"user_id"`
	ScreenID  uint32    `json:"screen_id"`
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
	Country   string    `json:"country"`
	OS        string    `json:"os"`
}

func NewTelemetryInput(
	userID uuid.UUID,
	screenID uint32,
	action string,
	timestamp time.Time,
	country string,
	os string,
) Input {
	return Input{
		UserID:    userID,
		ScreenID:  screenID,
		Action:    action,
		Timestamp: timestamp,
		Country:   country,
		OS:        os,
	}
}
