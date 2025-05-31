package logger

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/log"
	"go.opentelemetry.io/otel/trace"
)

type OtelLogger struct {
	inner log.Logger
}

func NewOtelLogger(otelLogger log.Logger) Logger {
	return &OtelLogger{
		inner: otelLogger,
	}
}

func (l *OtelLogger) log(ctx context.Context, lvl log.Severity, msg string) {
	var attrs []log.KeyValue

	if spanCtx := trace.SpanContextFromContext(ctx); spanCtx.IsValid() {
		attrs = append(attrs,
			log.String("trace_id", spanCtx.TraceID().String()),
			log.String("span_id", spanCtx.SpanID().String()),
		)
	}

	rec := log.Record{}
	rec.SetBody(log.StringValue(msg))
	rec.SetSeverity(lvl)
	rec.AddAttributes(attrs...)

	l.inner.Emit(ctx, rec)
}

func (l *OtelLogger) Trace(ctx context.Context, msg string, args ...interface{}) {
	l.log(ctx, log.SeverityTrace, fmt.Sprintf(msg, args...))
}

func (l *OtelLogger) Debug(ctx context.Context, msg string, args ...interface{}) {
	l.log(ctx, log.SeverityDebug, fmt.Sprintf(msg, args...))
}

func (l *OtelLogger) Info(ctx context.Context, msg string, args ...interface{}) {
	l.log(ctx, log.SeverityInfo, fmt.Sprintf(msg, args...))
}

func (l *OtelLogger) Warn(ctx context.Context, msg string, args ...interface{}) {
	l.log(ctx, log.SeverityWarn, fmt.Sprintf(msg, args...))
}

func (l *OtelLogger) Error(ctx context.Context, msg string, args ...interface{}) {
	l.log(ctx, log.SeverityError, fmt.Sprintf(msg, args...))
}

func (l *OtelLogger) Fatal(ctx context.Context, msg string, args ...interface{}) {
	l.log(ctx, log.SeverityFatal, fmt.Sprintf(msg, args...))
}
