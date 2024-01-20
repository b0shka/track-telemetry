# track-telemetry

![Go](https://img.shields.io/static/v1?label=GO&message=v1.21&color=blue)

---

## Installation

#### Prerequisites

- Go 1.21
- Docker & Docker Compose
- mockgen (used to start mock generation for unit tests)
- golang-migrate (used to run migrations in the database)
- golangci-lint (used to run code checks)
- swag (used to re-generate swagger documentation)

Create `.env` file in root directory and add following values:

```
ENV=local|production
HTTP_HOST=localhost

CLICKHOUSE_HOST=<host>
CLICKHOUSE_PORT=<port>
CLICKHOUSE_DATABASE=<database>
CLICKHOUSE_USERNAME=<username>
CLICKHOUSE_PASSWORD=<password>
CLICKHOUSE_MIGRATION_URL=file://<path to folder with migrate files>
```

---

## Build & Run

To start, run

```
make start
```