package config

import (
	"log"
	"os"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	EnvLocal = "local"
	EnvTest  = "test"
	EnvProd  = "prod"
)

type (
	Config struct {
		Environment string           `env:"ENV"`
		ClickHouse  ClickHouseConfig `mapstructure:"clickhouse"`
		HTTP        HTTPConfig       `mapstructure:"http"`
		Geoip2File  string           `mapstructure:"geoip2_file"`
	}

	ClickHouseConfig struct {
		Host            string        `env:"CLICKHOUSE_HOST"`
		Port            uint16        `env:"CLICKHOUSE_PORT"`
		Database        string        `env:"CLICKHOUSE_DATABASE"`
		Username        string        `env:"CLICKHOUSE_USERNAME"`
		Password        string        `env:"CLICKHOUSE_PASSWORD"`
		MigrationURL    string        `env:"CLICKHOUSE_MIGRATION_URL"`
		DialTimeout     time.Duration `mapstructure:"dial_timeout"`
		MaxOpenConns    int           `mapstructure:"max_open_conns"`
		MaxIdleConns    int           `mapstructure:"max_idle_conns"`
		ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	}

	HTTPConfig struct {
		Host               string        `env:"HTTP_HOST"`
		Port               uint16        `mapstructure:"port"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
	}

	ClickHouseTest struct {
		Host     string `env:"TEST_CLICKHOUSE_HOST"`
		Port     uint16 `env:"TEST_CLICKHOUSE_PORT"`
		Database string `env:"TEST_CLICKHOUSE_DATABASE"`
		Username string `env:"TEST_CLICKHOUSE_USERNAME"`
		Password string `env:"TEST_CLICKHOUSE_PASSWORD"`
	}

	HandlerTest struct {
		Geoip2File string `mapstructure:"geoip2_file_test"`
	}
)

func InitConfig(configPath string) *Config {
	var cfg Config

	parseFileConfig(configPath, "main", &cfg)
	parseFileEnv(".env", EnvLocal, &cfg)

	return &cfg
}

func InitTestConfig(configPath string) *HandlerTest {
	var cfg HandlerTest

	parseFileConfig(configPath, "test", &cfg)

	return &cfg
}

func InitTestEnv(envPath string) *ClickHouseTest {
	var cfg ClickHouseTest

	parseFileEnv(envPath, EnvTest, &cfg)

	return &cfg
}

func parseFileConfig[T Config | HandlerTest](configPath, nameFile string, cfg *T) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(nameFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}
}

func parseFileEnv[T Config | ClickHouseTest](envPath, envMode string, cfg *T) {
	if os.Getenv("APP_ENV") == envMode {
		if err := godotenv.Load(envPath); err != nil {
			log.Fatal(err)
		}
	}

	if err := env.Parse(cfg); err != nil {
		log.Fatal(err)
	}
}
