package logger

import "context"

type Logger interface {
	Trace(ctx context.Context, msg string, attrs ...Attribute)
	Debug(ctx context.Context, msg string, attrs ...Attribute)
	Info(ctx context.Context, msg string, attrs ...Attribute)
	Warn(ctx context.Context, msg string, attrs ...Attribute)
	Error(ctx context.Context, msg string, attrs ...Attribute)
	Fatal(ctx context.Context, msg string, attrs ...Attribute)
}
