PROGRAM_NAME = main
DB_URL=clickhouse://127.0.0.1:19000?username=default&password=qwerty&x-multi-statement=true
MIGRATION_URL=file://internal/repository/clickhouse/migration

.PHONY: build start test test.coverage lint swag mock migrateup migratedown
.DEFAULT_GOAL := start

build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/${PROGRAM_NAME} ./cmd/app/main.go

start: build
	APP_ENV="local" .bin/main

test:
	APP_ENV="test" GIN_MODE=release go test --short -coverprofile=cover.out -v -count=1 ./...
	make test.coverage

test.coverage:
	go tool cover -func=cover.out | grep "total"

lint:
	golangci-lint run ./... --config=./.golangci.yml

swag:
	swag init -g internal/app/app.go

mock:
	mockgen -source=internal/repository/clickhouse/repository.go -destination=internal/repository/clickhouse/mocks/mock_service.go
	mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock_service.go

migrateup:
	migrate -source ${MIGRATION_URL} -database "${DB_URL}" -verbose up

migratedown:
	migrate -source ${MIGRATION_URL} -database "${DB_URL}" -verbose down