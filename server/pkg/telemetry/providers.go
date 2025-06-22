package telemetry

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.30.0"
	"os"
)

func newLoggerProvider(ctx context.Context, res *resource.Resource, collectorUrl string) (*log.LoggerProvider, error) {
	if collectorUrl == "" {
		return log.NewLoggerProvider(
			log.WithResource(res),
		), nil
	}

	exporter, err := otlploggrpc.New(ctx, otlploggrpc.WithEndpoint(collectorUrl), otlploggrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	processor := log.NewBatchProcessor(exporter)
	lp := log.NewLoggerProvider(
		log.WithProcessor(processor),
		log.WithResource(res),
	)

	return lp, nil
}

func newMeterProvider(ctx context.Context, res *resource.Resource, collectorUrl string) (*metric.MeterProvider, error) {
	if collectorUrl == "" {
		return metric.NewMeterProvider(
			metric.WithResource(res),
		), nil
	}

	exporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithEndpoint(collectorUrl), otlpmetricgrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	mp := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(exporter)),
		metric.WithResource(res),
	)

	otel.SetMeterProvider(mp)
	return mp, nil
}

func newTracerProvider(ctx context.Context, res *resource.Resource, collectorUrl string) (*trace.TracerProvider, error) {
	if collectorUrl == "" {
		return trace.NewTracerProvider(
			trace.WithResource(res),
		), nil
	}

	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint(collectorUrl), otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(res),
	)
	otel.SetTracerProvider(tp)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return tp, nil
}

func newResource(serviceName, serviceVersion string) *resource.Resource {
	hostName, _ := os.Hostname()

	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceName(serviceName),
		semconv.ServiceVersion(serviceVersion),
		semconv.HostName(hostName),
		attribute.String("library.language", "go"),
	)
}
