package provider

import (
	"context"
	"github.com/axidex/api-example/server/pkg/db"
)

func (p *TransactionsProvider) initDatabase(_ context.Context) error {
	psqlEngine, err := db.NewGormConnection(p.cfg.Database, p.logger)
	if err != nil {
		return err
	}

	engineWithTelemetry, err := db.WithTelemetry(psqlEngine, p.telemetry.GetTracerProvider())
	if err != nil {
		return err
	}

	p.dependencies.DB = engineWithTelemetry

	return nil
}
