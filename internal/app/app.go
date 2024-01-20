package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/clickhouse" // for connect to clickhouse
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/gommon/log"
	"github.com/oschwald/geoip2-golang"
	"github.com/vanya/backend/internal/config"
	handler "github.com/vanya/backend/internal/handler/http"
	repository "github.com/vanya/backend/internal/repository/clickhouse"
	"github.com/vanya/backend/internal/server"
	"github.com/vanya/backend/internal/service"
	"github.com/vanya/backend/pkg/database/clickhouse"
	"github.com/vanya/backend/pkg/logger"
)

//	@title			Track API
//	@version		1.0
//	@description	REST API for Track App

//	@host		localhost:8080
//	@BasePath	/api/v1/

func Run(configPath string) {
	cfg, err := config.InitConfig(configPath)
	if err != nil {
		logger.Error(err)

		return
	}

	reader, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		logger.Error(err)

		return
	}

	err = runDBMigration(cfg.ClickHouse)
	if err != nil {
		logger.Error(err)

		return
	}

	logger.Info("ClickHouse migrated successfully")

	clickhouseConn, err := clickhouse.Connect(cfg.ClickHouse)
	if err != nil {
		logger.Errorf("Cannot connect to ClickHouse: %s", err)

		return
	}

	logger.Info("Success connect to ClickHouse")

	repos := repository.NewRepositories(clickhouseConn)
	services := service.NewServices(service.Deps{
		Repos: repos,
	})
	handlers := handler.NewHandler(services, reader)
	routes := handlers.InitRoutes(cfg)
	srv := server.NewServer(cfg, routes)

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	gracefulShutdown(srv, clickhouseConn, reader)
}

func gracefulShutdown(
	srv *server.Server,
	clickHouse driver.Conn,
	reader *geoip2.Reader,
) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second
	ctx, shutdown := context.WithTimeout(context.Background(), timeout)

	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		log.Errorf("Failed to stop server: %v", err)
	}

	logger.Info("Server stopped")

	clickHouse.Close()
	reader.Close()
}

func runDBMigration(cfg config.ClickHouseConfig) error {
	migration, err := migrate.New(
		cfg.MigrationURL,
		fmt.Sprintf(
			"clickhouse://%s:%d?username=%s&password=%s&x-multi-statement=true",
			cfg.Host, cfg.Port, cfg.Username, cfg.Password,
		),
	)
	if err != nil {
		return err
	}

	if err := migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
