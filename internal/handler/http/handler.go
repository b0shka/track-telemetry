package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/vanya/backend/docs"
	"github.com/vanya/backend/internal/config"
	"github.com/vanya/backend/internal/service"
)

type Handler struct {
	Services *service.Services
	Geoip    *geoip2.Reader
	Logger   *logrus.Logger
}

func NewHandler(
	services *service.Services,
	geoip *geoip2.Reader,
	logger *logrus.Logger,
) *Handler {
	return &Handler{
		Services: services,
		Geoip:    geoip,
		Logger:   logger,
	}
}

func (h *Handler) InitRoutes(cfg *config.Config) *gin.Engine {
	router := gin.Default()

	router.Use(
		gin.Recovery(),
		gin.Logger(),
		corsMiddleware,
	)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", cfg.HTTP.Host, cfg.HTTP.Port)

	if cfg.Environment != config.EnvLocal {
		docs.SwaggerInfo.Host = cfg.HTTP.Host
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	api := router.Group("/api/v1")
	{
		h.initTelemetryRoutes(api)
	}

	return router
}
