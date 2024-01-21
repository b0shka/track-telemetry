package http

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mssola/useragent"
	"github.com/oschwald/geoip2-golang"
	"github.com/vanya/backend/internal/domain/telemetry"
)

const (
	payloadCtx = "payload"
)

func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != http.MethodOptions {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}

func (h *Handler) ipIdentity(geoip *geoip2.Reader) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		country, err := geoip.Country(net.ParseIP(ip))
		if err != nil {
			h.Logger.Errorf("Error parsing geoip: %s", err.Error())
		}

		ua := useragent.New(c.Request.Header.Get("User-Agent"))

		payload := telemetry.NewPayloadContext(country.Country.IsoCode, ua.OS())
		c.Set(payloadCtx, payload)
	}
}

func getPayloadByContext(c *gin.Context) (telemetry.PayloadContext, error) {
	payloadFromCtx, ok := c.Get(payloadCtx)
	if !ok {
		return telemetry.PayloadContext{}, fmt.Errorf("%s not found", payloadCtx)
	}

	payload, ok := payloadFromCtx.(telemetry.PayloadContext)
	if !ok {
		return telemetry.PayloadContext{}, fmt.Errorf("%s is of invalid type", payloadCtx)
	}

	return payload, nil
}
