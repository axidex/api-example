package api

import (
	"context"
	"github.com/axidex/api-example/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type LoggingFunc func(ctx context.Context, msg string, args ...interface{})

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
			"path: %s | statusCode: %d | method: %s | clientIP: %s | latency: %s | errorMsg: %s",
			request.path,
			request.statusCode,
			request.method,
			request.clientIP,
			request.latency.String(),
			errorMsg,
		)
	}
}

func (m *LoggerMiddleware) LogInfo(ctx context.Context, request RequestInfo, logging LoggingFunc) {
	logging(
		ctx,
		"path: %s | statusCode: %d | method: %s | clientIP: %s | latency: %s",
		request.path,
		request.statusCode,
		request.method,
		request.clientIP,
		request.latency.String(),
	)
}
