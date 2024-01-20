CREATE DATABASE IF NOT EXISTS telemetry;

CREATE TABLE IF NOT EXISTS telemetry.actions
(
    ts TIMESTAMP,
    user_id UUID,
    screen_id UInt32,
    action String,
    country Nullable(FixedString(2)),
    os Nullable(String)
)
ENGINE = MergeTree()
PRIMARY KEY (ts);