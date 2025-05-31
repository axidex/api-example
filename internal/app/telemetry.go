package app

import (
	"context"
	"github.com/axidex/api-example/pkg/config_provider"
	"github.com/axidex/api-example/pkg/telemetry"
)

func (a *App) initTelemetry(ctx context.Context) error {
	tel, err := telemetry.NewTelemetry(ctx, a.cfg.Telemetry, a.name, config_provider.NewVersion().Version())
	if err != nil {
		return err
	}

	a.telemetry = tel

	return nil
}
