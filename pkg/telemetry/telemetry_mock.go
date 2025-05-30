package telemetry

import (
	"context"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
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

func (t *MockTelemetry) Stop(context.Context) {}
