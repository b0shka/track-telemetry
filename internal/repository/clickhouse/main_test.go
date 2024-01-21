package repository

import (
	"fmt"
	"os"
	"testing"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/vanya/backend/internal/config"
	"github.com/vanya/backend/pkg/logging"
)

var testRepos *Repositories

func TestMain(m *testing.M) {
	logger := logging.NewLogger(config.EnvTest)

	cfg := config.InitTestEnv("../../../.env.test")

	option := &clickhouse.Options{
		Addr: []string{
			fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		},
		Auth: clickhouse.Auth{
			Database: cfg.Database,
			Username: cfg.Username,
			Password: cfg.Password,
		},
	}

	conn, err := clickhouse.Open(option)
	if err != nil {
		logger.Errorf("cannot to connect to database: %s", err)

		return
	}

	testRepos = NewRepositories(conn)

	os.Exit(m.Run())
}
