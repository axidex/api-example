BUILD_DATE := $(shell date -u +"%d.%m.%Y")
BUILD_COMMIT := $(shell git rev-parse --short HEAD)
BUILD_VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "N/A")
VERSION_PACKAGE := "github.com/axidex/api-example/pkg/version"

# Go
build:
	go build \
		-ldflags="-s -w -X $(VERSION_PACKAGE).buildVersion=$(BUILD_VERSION) -X $(VERSION_PACKAGE).buildDate=$(BUILD_DATE) -X $(VERSION_PACKAGE).buildCommit=$(BUILD_COMMIT)" \
		-o ./api_main cmd/main/main.go

run:
	go run \
		-ldflags="-X $(VERSION_PACKAGE).buildVersion=$(BUILD_VERSION) -X $(VERSION_PACKAGE).buildDate=$(BUILD_DATE) -X $(VERSION_PACKAGE).buildCommit=$(BUILD_COMMIT)" \
 		cmd/main/main.go --debug

tidy:
	go mod tidy
	go fmt ./...

swag:
	swag init -g cmd/main/main.go


# Compose
network:
	docker network create api-example

# Telemetry Docker Compose
telemetry:
	cd ./compose/telemetry && docker compose up -d

telemetry-down:
	cd ./compose/telemetry && docker compose down

telemetry-restart: telemetry-down telemetry


# DB Docker Compose
db:
	cd ./compose/db && docker compose up -d

db-down:
	cd ./compose/db && docker compose down

db-restart: db-down db
