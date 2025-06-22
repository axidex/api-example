package telemetry

import (
	"context"
	otelmetric "go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	oteltracer "go.opentelemetry.io/otel/trace"
)

type Telemetry interface {
	GetLoggerProvider() *log.LoggerProvider
	GetMeterProvider() *metric.MeterProvider
	GetTracerProvider() *trace.TracerProvider

	TraceStart(ctx context.Context, name string) (context.Context, oteltracer.Span)

	MeterInt64Histogram(metric Metric) (otelmetric.Int64Histogram, error)
	MeterInt64UpDownCounter(metric Metric) (otelmetric.Int64UpDownCounter, error)

	Stop(ctx context.Context)
}

type OpenTelemetry struct {
	lp     *log.LoggerProvider
	mp     *metric.MeterProvider
	tp     *trace.TracerProvider
	meter  otelmetric.Meter
	tracer oteltracer.Tracer
}

func NewTelemetry(ctx context.Context, config Config, serviceName, serviceVersion string) (Telemetry, error) {
	rp := newResource(serviceName, serviceVersion)

	lp, err := newLoggerProvider(ctx, rp, config.LogCollector)
	if err != nil {
		return nil, err
	}

	mp, err := newMeterProvider(ctx, rp, config.MetricCollector)
	if err != nil {
		return nil, err
	}

	tp, err := newTracerProvider(ctx, rp, config.TraceCollector)
	if err != nil {
		return nil, err
	}

	return &OpenTelemetry{
		lp:     lp,
		mp:     mp,
		tp:     tp,
		meter:  mp.Meter(serviceName),
		tracer: tp.Tracer(serviceName),
	}, nil
}

func (t *OpenTelemetry) GetLoggerProvider() *log.LoggerProvider {
	return t.lp
}

func (t *OpenTelemetry) GetMeterProvider() *metric.MeterProvider {
	return t.mp
}

func (t *OpenTelemetry) GetTracerProvider() *trace.TracerProvider {
	return t.tp
}

// nolint
func (t *OpenTelemetry) Stop(ctx context.Context) {
	t.lp.Shutdown(ctx) // #nosec G104
	t.mp.Shutdown(ctx) // #nosec G104
	t.tp.Shutdown(ctx) // #nosec G104
}
