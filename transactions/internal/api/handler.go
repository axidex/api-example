package api

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	_ "github.com/axidex/api-example/transactions/docs"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type GinHandler struct {
	name      string
	config    Config
	logger    logger.Logger
	telemetry telemetry.Telemetry
}

func NewGinHandler(name string, config Config, logger logger.Logger, telemetry telemetry.Telemetry) *GinHandler {
	gin.SetMode(gin.ReleaseMode)

	return &GinHandler{
		name:      name,
		config:    config,
		logger:    logger,
		telemetry: telemetry,
	}
}

func (h *GinHandler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(otelgin.Middleware(h.name, otelgin.WithMeterProvider(h.telemetry.GetMeterProvider()), otelgin.WithTracerProvider(h.telemetry.GetTracerProvider())))
	router.Use(h.MeterRequestsInFlight())

	loggerMiddleware := NewLoggerMiddleware(h.logger)
	router.Use(CustomRecoveryFunc(h.logger))

	router.GET("/swagger/*any", h.Swagger)

	v1 := router.Group("/v1")
	{
		v1.Use(loggerMiddleware.Default())

		v1.POST("/cell", h.CreateCell)
	}

	health := router.Group("/health")
	{
		health.GET("/ping", h.Ping)
	}

	h.listRoutes(router)
	return router
}

func (h *GinHandler) listRoutes(router *gin.Engine) {
	for _, route := range router.Routes() {
		h.logger.Info(context.Background(), fmt.Sprintf("Method: %s | Path: %s | Handler: %s", route.Method, route.Path, route.Handler))
	}
}

func (h *GinHandler) HandleServer(ctx context.Context) (func() error, func(error)) {
	ctx, cancel := context.WithCancel(ctx)
	engine := h.InitRoutes()
	errorChan := make(chan error, 1)
	return func() error {
			go func() {
				errorChan <- engine.Run(fmt.Sprintf(":%d", h.config.Port))
			}()
			select {
			case <-ctx.Done():
				return ctx.Err()
			case err := <-errorChan:
				return err
			}
		},
		func(err error) {
			cancel()
			h.logger.Warn(ctx, fmt.Sprintf("Shutting down server: %s", err))
		}
}
