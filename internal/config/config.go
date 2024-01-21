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

	ClickHouseTest struct {
		Host     string `env:"TEST_CLICKHOUSE_HOST"`
		Port     uint16 `env:"TEST_CLICKHOUSE_PORT"`
		Database string `env:"TEST_CLICKHOUSE_DATABASE"`
		Username string `env:"TEST_CLICKHOUSE_USERNAME"`
		Password string `env:"TEST_CLICKHOUSE_PASSWORD"`
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
)

func InitConfig(configPath string) *Config {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("main")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}

	if os.Getenv("APP_ENV") == EnvLocal {
		if err := godotenv.Load(); err != nil {
			log.Fatal(err)
		}
	}

	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	return &cfg
}

func InitTestConfig(envPath string) (*ClickHouseTest, error) {
	var cfg ClickHouseTest

	if os.Getenv("APP_ENV") == EnvTest {
		if err := godotenv.Load(envPath); err != nil {
			return nil, err
		}
	}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
