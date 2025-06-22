package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/axidex/api-example/server/pkg/logger"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

const (
	infoStr      = "%s [info] "
	warnStr      = "%s [warn] "
	errStr       = "%s [error] "
	traceStr     = "%s [%.3fms] [rows:%v] %s"
	traceWarnStr = "%s %s [%.3fms] [rows:%v] %s"
	traceErrStr  = "%s %s [%.3fms] [rows:%v] %s"
)

type GormLogger struct {
	logger        logger.Logger
	slowThreshold time.Duration
}

func NewGormLogger(logger logger.Logger) gormLogger.Interface {
	return &GormLogger{
		logger:        logger,
		slowThreshold: 200 * time.Millisecond,
	}
}

func (l *GormLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	return l
}

// Info print info
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Debug(ctx, fmt.Sprintf(infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
}

// Warn print warn messages
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Debug(ctx, fmt.Sprintf(warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
}

// Error print error messages
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	l.logger.Debug(ctx, fmt.Sprintf(errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...))
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	switch {
	case err != nil && (!errors.Is(err, gorm.ErrRecordNotFound)):
		sql, rows := fc()
		if rows == -1 {
			l.logger.Debug(ctx, fmt.Sprintf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql))
		} else {
			l.logger.Debug(ctx, fmt.Sprintf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}
	case elapsed > l.slowThreshold && l.slowThreshold != 0:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.slowThreshold)
		if rows == -1 {
			l.logger.Debug(ctx, fmt.Sprintf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql))
		} else {
			l.logger.Debug(ctx, fmt.Sprintf(traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}
	default:
		sql, rows := fc()
		if rows == -1 {
			l.logger.Debug(ctx, fmt.Sprintf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql))
		} else {
			l.logger.Debug(ctx, fmt.Sprintf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql))
		}
	}
}
