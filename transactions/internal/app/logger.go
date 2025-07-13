package app

import (
	"context"
	"github.com/axidex/api-example/server/pkg/logger"
)

func (a *TransactionsApp) initLogger(_ context.Context) error {
	log, err := logger.NewZapLogger(a.cfg.Logger, a.name, a.telemetry.GetLoggerProvider())
	if err != nil {
		return err
	}

	a.logger = log
	return nil
}

func (a *ApiApp) initLogger(_ context.Context) error {
	log, err := logger.NewZapLogger(a.cfg.Logger, a.name, a.telemetry.GetLoggerProvider())
	if err != nil {
		return err
	}

	a.logger = log
	return nil
}
