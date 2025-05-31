package telemetry

import (
	"fmt"
	otelmetric "go.opentelemetry.io/otel/metric"
)

// Metric represents a metric that can be collected by the server.
type Metric struct {
	Name        string
	Unit        string
	Description string
}

// MetricRequestDurationMillis is a metric that measures the latency of HTTP requests processed by the server, in milliseconds.
var MetricRequestDurationMillis = Metric{
	Name:        "request_duration_millis",
	Unit:        "ms",
	Description: "Measures the latency of HTTP requests processed by the server, in milliseconds.",
}

// MetricRequestsInFlight is a metric that measures the number of requests currently being processed by the server.
var MetricRequestsInFlight = Metric{
	Name:        "requests_inflight",
	Unit:        "{count}",
	Description: "Measures the number of requests currently being processed by the server.",
}

// MeterInt64Histogram creates a new int64 histogram metric.
func (t *OpenTelemetry) MeterInt64Histogram(metric Metric) (otelmetric.Int64Histogram, error) {
	histogram, err := t.meter.Int64Histogram(
		metric.Name,
		otelmetric.WithDescription(metric.Description),
		otelmetric.WithUnit(metric.Unit),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create histogram: %w", err)
	}

	return histogram, nil
}

// MeterInt64UpDownCounter creates a new int64 up-down-counter metric.
func (t *OpenTelemetry) MeterInt64UpDownCounter(metric Metric) (otelmetric.Int64UpDownCounter, error) {
	counter, err := t.meter.Int64UpDownCounter(
		metric.Name,
		otelmetric.WithDescription(metric.Description),
		otelmetric.WithUnit(metric.Unit),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create counter: %w", err)
	}

	return counter, nil
}
