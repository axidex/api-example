BUILD_DATE := $(shell date -u +"%d.%m.%Y")
BUILD_COMMIT := $(shell git rev-parse --short HEAD)
BUILD_VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "N/A")
VERSION_PACKAGE := "github.com/axidex/api-example/server/pkg/version"

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

lint:
	golangci-lint run ./...

swag:
	swag init -g cmd/main/main.go -o server/
