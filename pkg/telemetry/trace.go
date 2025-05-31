package telemetry

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	oteltracer "go.opentelemetry.io/otel/trace"
)

type TraceInfo struct {
	TraceId string
	SpanId  string
}

func GetTraceId(ctx context.Context) (TraceInfo, bool) {
	span := trace.SpanFromContext(ctx)
	if !span.IsRecording() {
		return TraceInfo{}, false
	}

	return TraceInfo{
		TraceId: span.SpanContext().TraceID().String(),
		SpanId:  span.SpanContext().SpanID().String(),
	}, true
}

// TraceStart starts a new span with the given name. The span must be ended by calling End.
func (t *OpenTelemetry) TraceStart(ctx context.Context, name string) (context.Context, oteltracer.Span) {
	return t.tracer.Start(ctx, name)
}
