package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vanya/backend/internal/domain"
)

func (h *Handler) initTelemetryRoutes(api *gin.RouterGroup) {
	api.Use(ipIdentity(h.Geoip))
	api.POST("/track", h.track)
}

type telemetryRequest struct {
	UserID    string    `json:"user_id" binding:"required"`
	ScreenID  uint32    `json:"screen_id" binding:"required"`
	Action    string    `json:"action" binding:"required"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
}

// @Summary		Track telemetry
// @Tags		Telemetry
// @Description	track telemetry
// @ModuleID	track
// @Accept		json
// @Produce		json
// @Param		input	body		telemetryRequest	true	"track info"
// @Success		200		{string}	string				"ok"
// @Failure		400		{object}	response
// @Failure		500		{object}	response
// @Failure		default	{object}	response
// @Router		/track [post]
func (h *Handler) track(c *gin.Context) {
	payload, err := getPayloadByContext(c)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	var req telemetryRequest
	if err := c.BindJSON(&req); err != nil {
		newResponse(c, http.StatusBadRequest, domain.ErrInvalidInput.Error())

		return
	}

	inp, err := NewTelemetryInput(req, payload)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())

		return
	}

	err = h.Services.Telemetry.Append(c, inp)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())

		return
	}

	c.Status(http.StatusOK)
}
