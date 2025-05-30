package telemetry

import (
	"context"
	"go.opentelemetry.io/otel/trace"
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
