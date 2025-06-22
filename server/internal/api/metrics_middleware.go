package api

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"time"
)

// MeterRequestDuration is a gin middleware that captures the duration of the request.
func (h *GinHandler) MeterRequestDuration() gin.HandlerFunc {
	// init metric, here we are using histogram for capturing request duration
	histogram, err := h.telemetry.MeterInt64Histogram(telemetry.MetricRequestDurationMillis)
	if err != nil {
		h.logger.Error(context.Background(), fmt.Sprintf("Failed to create histogram: %s", err))
	}

	return func(c *gin.Context) {
		// capture the start time of the request
		startTime := time.Now()

		// execute next http handler
		c.Next()

		// record the request duration
		duration := time.Since(startTime)
		histogram.Record(
			c.Request.Context(),
			duration.Milliseconds(),
			metric.WithAttributes(
				attribute.String("http.method", c.Request.Method),
				attribute.String("http.path", c.Request.URL.Path),
				attribute.String("http.host", c.Request.Host),
			),
		)
	}
}

// MeterRequestsInFlight is a gin middleware that captures the number of requests in flight.
func (h *GinHandler) MeterRequestsInFlight() gin.HandlerFunc {
	// init metric, here we are using counter for capturing request in flight
	counter, err := h.telemetry.MeterInt64UpDownCounter(telemetry.MetricRequestsInFlight)
	if err != nil {
		h.logger.Info(context.Background(), fmt.Sprintf("Failed to create counter: %s", err))
	}

	return func(c *gin.Context) {
		// define metric attributes
		attrs := metric.WithAttributes(
			attribute.String("http.method", c.Request.Method),
			attribute.String("http.path", c.Request.URL.Path),
			attribute.String("http.host", c.Request.Host),
		)
		// increase the number of requests in flight
		counter.Add(c.Request.Context(), 1, attrs)

		// execute next http handler
		c.Next()

		// decrease the number of requests in flight
		counter.Add(c.Request.Context(), -1, attrs)
	}
}
