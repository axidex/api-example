package logger

import (
	"context"
	"github.com/axidex/api-example/server/pkg/version"
	"go.opentelemetry.io/contrib/bridges/otelzap"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/trace"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	cfg       Config
	rawLogger *zap.Logger
	//sugarLogger *zap.SugaredLogger
}

func NewZapLogger(cfg Config, serviceName string, lp *log.LoggerProvider) (Logger, error) {
	logger := &ZapLogger{cfg: cfg}
	logger.InitLogger(lp, serviceName)

	return logger, nil
}

// For mapping config.go logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *ZapLogger) GetRawLogger() *zap.Logger {
	return l.rawLogger
}

func (l *ZapLogger) getLoggerLevel(cfg Config) zapcore.Level {
	level, exist := loggerLevelMap[strings.ToLower(cfg.Level)]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// CustomTimeEncoder Custom time encoder
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format(RFC3339))
}

func RFC832TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format(RFC832))
}

// CustomLevelEncoder Custom level encoder with uppercase
func CustomLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(l.CapitalString())
}

func GetEncoderCfg(encoder zapcore.LevelEncoder, timeEncoder zapcore.TimeEncoder) zapcore.EncoderConfig {
	encoderCfg := zap.NewProductionEncoderConfig()

	encoderCfg.LevelKey = "level"
	encoderCfg.CallerKey = "caller"
	encoderCfg.TimeKey = "time"
	encoderCfg.NameKey = "name"
	encoderCfg.MessageKey = "message"
	encoderCfg.EncodeTime = timeEncoder
	encoderCfg.EncodeLevel = encoder
	return encoderCfg
}

// InitLogger Init logger
func (l *ZapLogger) InitLogger(lp *log.LoggerProvider, serviceName string) {
	logLevel := l.getLoggerLevel(l.cfg)

	fileEncoderCfg := GetEncoderCfg(CustomLevelEncoder, CustomTimeEncoder)
	consoleEncoderCfg := GetEncoderCfg(zapcore.CapitalColorLevelEncoder, RFC832TimeEncoder)

	encoder := zapcore.NewConsoleEncoder(consoleEncoderCfg)
	fileEncoder := zapcore.NewJSONEncoder(fileEncoderCfg)

	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, zapcore.AddSync(&lumberjack.Logger{
			Filename:   l.cfg.Path,
			MaxSize:    l.cfg.Rotation.MaxSize, // megabytes
			MaxBackups: l.cfg.Rotation.MaxBackups,
			MaxAge:     l.cfg.Rotation.MaxAge,   //days
			Compress:   l.cfg.Rotation.Compress, // disabled by default
		}), zap.NewAtomicLevelAt(logLevel)),
		zapcore.NewCore(encoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), zap.NewAtomicLevelAt(logLevel)),
		otelzap.NewCore(serviceName, otelzap.WithLoggerProvider(lp), otelzap.WithVersion(version.NewVersion().Version())),
	}

	core := zapcore.NewTee(cores...)
	rawLogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.rawLogger = rawLogger

	//l.sugarLogger = l.rawLogger.Sugar()
	//err := l.sugarLogger.Sync()
	//if err != nil && !errors.Is(err, syscall.ENOTTY) && !errors.Is(err, syscall.EINVAL) {
	//	l.sugarLogger.Error(err)
	//}
}

func (l *ZapLogger) withCtx(ctx context.Context) *zap.Logger {
	return withTrace(l.rawLogger, ctx)
}

func withTrace(logger *zap.Logger, ctx context.Context) *zap.Logger {
	spanCtx := trace.SpanContextFromContext(ctx)
	if !spanCtx.IsValid() {
		return logger
	}

	return logger.With(
		zap.String("trace_id", spanCtx.TraceID().String()),
		zap.String("span_id", spanCtx.SpanID().String()),
	)
}

func (l *ZapLogger) Trace(ctx context.Context, msg string, attrs ...Attribute) {
	l.withCtx(ctx).Debug(msg, transformAttributes(attrs)...)
}

func (l *ZapLogger) Debug(ctx context.Context, msg string, attrs ...Attribute) {
	l.withCtx(ctx).Debug(msg, transformAttributes(attrs)...)
}

func (l *ZapLogger) Info(ctx context.Context, msg string, attrs ...Attribute) {
	l.withCtx(ctx).Info(msg, transformAttributes(attrs)...)
}

func (l *ZapLogger) Warn(ctx context.Context, msg string, attrs ...Attribute) {
	l.withCtx(ctx).Warn(msg, transformAttributes(attrs)...)
}

func (l *ZapLogger) Error(ctx context.Context, msg string, attrs ...Attribute) {
	l.withCtx(ctx).Error(msg, transformAttributes(attrs)...)
}

func (l *ZapLogger) Fatal(ctx context.Context, msg string, attrs ...Attribute) {
	l.withCtx(ctx).Fatal(msg, transformAttributes(attrs)...)
}
