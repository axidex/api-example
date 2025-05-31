# Example API repository

The purpose of this repo is for be an example of how to create modern API in Golang


## .env example

```
TELEMETRY_TRACE_COLLECTOR=localhost:4317
TELEMETRY_METRIC_COLLECTOR=localhost:4317
TELEMETRY_LOG_COLLECTOR=localhost:4317

LOGGER_LEVEL=info
LOGGER_PATH=logs/api-example.log
LOGGER_ROTATION_MAX_SIZE=20
LOGGER_ROTATION_MAX_AGE=30
LOGGER_ROTATION_MAX_BACKUPS=0
LOGGER_ROTATION_COMPRESS=false
```