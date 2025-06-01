package telemetry

import (
	"context"
	otelmetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltracer "go.opentelemetry.io/otel/trace"
)

type MockTelemetry struct {
	tp *trace.TracerProvider
	mp *metric.MeterProvider
	lp *log.LoggerProvider
}

func NewMockTelemetry() Telemetry {
	ctx := context.Background()
	res, _ := resource.New(ctx)

	lp := log.NewLoggerProvider(
		log.WithResource(res),
	)

	mp := metric.NewMeterProvider(
		metric.WithResource(res),
	)

	tp := trace.NewTracerProvider(
		trace.WithResource(res),
	)

	return &MockTelemetry{
		lp: lp,
		mp: mp,
		tp: tp,
	}
}

func (t *MockTelemetry) GetLoggerProvider() *log.LoggerProvider {
	return t.lp
}

func (t *MockTelemetry) GetMeterProvider() *metric.MeterProvider {
	return t.mp
}

func (t *MockTelemetry) GetTracerProvider() *trace.TracerProvider {
	return t.tp
}

func (t *MockTelemetry) TraceStart(ctx context.Context, name string) (context.Context, oteltracer.Span) {
	return nil, nil
}

func (t *MockTelemetry) MeterInt64UpDownCounter(metric Metric) (otelmetric.Int64UpDownCounter, error) {
	return nil, nil
}

func (t *MockTelemetry) MeterInt64Histogram(metric Metric) (otelmetric.Int64Histogram, error) {
	return nil, nil
}

func (t *MockTelemetry) Stop(context.Context) {}
