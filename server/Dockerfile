FROM golang:1.24.2-alpine AS builder

USER 0

ARG BUILD_VERSION
ARG BUILD_DATE
ARG BUILD_COMMIT
ARG VERSION_PACKAGE="github.com/axidex/api-example/server/pkg/config_provider"

WORKDIR /build

COPY . .

RUN go mod edit -go 1.24

ENV PATH="/opt/app-root/src/go/bin:${PATH}"
RUN go env -w GOBIN="/opt/app-root/src/go/bin"

RUN go get ./...
RUN go install github.com/swaggo/swag/cmd/swag@latest && swag init -g ./cmd/main/main.go

RUN go test ./...
RUN go build \
    -ldflags="-s -w -X ${VERSION_PACKAGE}.buildVersion=${BUILD_VERSION} -X ${VERSION_PACKAGE}.buildDate=${BUILD_DATE} -X ${VERSION_PACKAGE}.buildCommit=${BUILD_COMMIT}" \
    -o app ./cmd/main/main.go

FROM gcr.io/distroless/static

WORKDIR /opt/ss-manager

COPY --from=builder /build/app app

ENTRYPOINT ["./app"]