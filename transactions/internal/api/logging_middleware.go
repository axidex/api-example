package api

import (
	"context"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type LoggingFunc func(ctx context.Context, msg string, args ...logger.Attribute)

type RequestInfo struct {
	path       string
	statusCode int
	latency    time.Duration
	method     string
	clientIP   string
	errors     []string
}

func NewRequestInfo(c *gin.Context) RequestInfo {
	// Start timer
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery

	// Process request
	c.Next()

	// Stop timer
	end := time.Now()
	latency := end.Sub(start)

	clientIP := c.ClientIP()
	method := c.Request.Method
	statusCode := c.Writer.Status()
	//errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()
	if raw != "" {
		path = path + "?" + raw
	}

	return RequestInfo{
		path:       path,
		statusCode: statusCode,
		latency:    latency,
		method:     method,
		clientIP:   clientIP,
		errors:     c.Errors.Errors(),
	}
}

type LoggerMiddleware struct {
	logger logger.Logger
}

func NewLoggerMiddleware(logger logger.Logger) *LoggerMiddleware {
	return &LoggerMiddleware{logger: logger}
}

func (m *LoggerMiddleware) Default() gin.HandlerFunc {
	return func(c *gin.Context) {
		info := NewRequestInfo(c)
		m.LogInfo(c.Request.Context(), info, m.logger.Info)
		m.LogErrors(c.Request.Context(), info, m.logger.Error)
	}
}

func (m *LoggerMiddleware) Debug() gin.HandlerFunc {
	return func(c *gin.Context) {
		info := NewRequestInfo(c)
		m.LogInfo(c.Request.Context(), info, m.logger.Debug)
		m.LogErrors(c.Request.Context(), info, m.logger.Error)
	}
}

func (m *LoggerMiddleware) LogErrors(ctx context.Context, request RequestInfo, logging LoggingFunc) {
	for _, errorMsg := range request.errors {
		logging(
			ctx,
			"REQUEST FAILED",
			logger.NewAttribute("path", request.path),
			logger.NewAttribute("status_code", request.statusCode),
			logger.NewAttribute("method", request.method),
			logger.NewAttribute("client_ip", request.clientIP),
			logger.NewAttribute("latency", request.latency),
			logger.NewAttribute("error_msg", errorMsg),
		)
	}
}

func (m *LoggerMiddleware) LogInfo(ctx context.Context, request RequestInfo, logging LoggingFunc) {
	logging(
		ctx,
		"REQUEST INFO",
		logger.NewAttribute("path", request.path),
		logger.NewAttribute("status_code", request.statusCode),
		logger.NewAttribute("method", request.method),
		logger.NewAttribute("client_ip", request.clientIP),
		logger.NewAttribute("latency", request.latency),
	)
}
