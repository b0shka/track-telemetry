# track-telemetry

![Go](https://img.shields.io/static/v1?label=GO&message=v1.21&color=blue)

---

## Описание

Сервис для сбора телеметрии с мобильного приложения. Пока что обрабатывает и хранит самую простую телеметрию, в дальнейшем ее можно без проблем расширить.

Для хранения телеметрии используется **ClickHouse**, поскольку данная БД очень хорошо подходит именно для хранения метрик.

---

## Установка

#### Необходимые компоненты

- Go 1.21
- Docker & Docker Compose
- mockgen (используется для запуска генерации mock для модульных тестов)
- golang-migrate (используется для запуска миграций в базе данных)
- golangci-lint (используется для выполнения проверок кода)
- swag (используется для создания документации swagger)

Создайте файл `.env` в корневом каталоге и добавьте следующие значения:

```
ENV=local|prod
HTTP_HOST=localhost

CLICKHOUSE_HOST=<host>
CLICKHOUSE_PORT=<port>
CLICKHOUSE_DATABASE=<database>
CLICKHOUSE_USERNAME=<username>
CLICKHOUSE_PASSWORD=<password>
CLICKHOUSE_MIGRATION_URL=file://<path to folder with migrate files>
```

---

## Сборка и запуск

Используйте `make start` для сборки и запуска проекта, `make lint` для проверки кода с помощью linter.