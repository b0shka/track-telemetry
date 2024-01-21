package config

import (
	"os"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestInitConfig(t *testing.T) {
	type env struct {
		clickhouseHost         string
		clickhousePort         uint16
		clickhouseDatabase     string
		clickhouseUsername     string
		clickhousePassword     string
		clickhouseMigrationURL string
		environment            string
		httpHost               string
	}

	type args struct {
		path string
		env  env
	}

	setEnv := func(env env) {
		os.Setenv("ENV", env.environment)
		os.Setenv("HTTP_HOST", env.httpHost)
		os.Setenv("CLICKHOUSE_HOST", env.clickhouseHost)
		os.Setenv("CLICKHOUSE_PORT", strconv.Itoa(int(env.clickhousePort)))
		os.Setenv("CLICKHOUSE_DATABASE", env.clickhouseDatabase)
		os.Setenv("CLICKHOUSE_USERNAME", env.clickhouseUsername)
		os.Setenv("CLICKHOUSE_PASSWORD", env.clickhousePassword)
		os.Setenv("CLICKHOUSE_MIGRATION_URL", env.clickhouseMigrationURL)
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "test config",
			args: args{
				path: "fixtures",
				env: env{
					environment:            "local",
					httpHost:               "localhost",
					clickhouseHost:         "127.0.0.1",
					clickhousePort:         19000,
					clickhouseDatabase:     "database",
					clickhouseUsername:     "username",
					clickhousePassword:     "password",
					clickhouseMigrationURL: "migrationURL",
				},
			},
			want: &Config{
				Environment: "local",
				ClickHouse: ClickHouseConfig{
					Host:            "127.0.0.1",
					Port:            19000,
					Database:        "database",
					Username:        "username",
					Password:        "password",
					MigrationURL:    "migrationURL",
					DialTimeout:     time.Second,
					MaxOpenConns:    5,
					MaxIdleConns:    5,
					ConnMaxLifetime: time.Hour,
				},
				HTTP: HTTPConfig{
					Host:               "localhost",
					Port:               8080,
					MaxHeaderMegabytes: 1,
					ReadTimeout:        time.Second * 10,
					WriteTimeout:       time.Second * 10,
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			setEnv(testCase.args.env)

			got := InitConfig(testCase.args.path)

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("InitConfig() got = %v, want = %v", got, testCase.want)
			}
		})
	}
}
