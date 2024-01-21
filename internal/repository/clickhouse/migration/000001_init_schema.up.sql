CREATE DATABASE IF NOT EXISTS telemetry;

CREATE TABLE IF NOT EXISTS telemetry.actions
(
    ts TIMESTAMP,
    user_id UUID,
    screen String,
    action String,
    app_version String,
    country Nullable(FixedString(2)),
    os Nullable(String)
)
ENGINE = MergeTree()
PRIMARY KEY (ts);