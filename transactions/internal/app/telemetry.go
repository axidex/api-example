package app

import (
	"context"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/server/pkg/version"
)

func (a *App) initTelemetry(ctx context.Context) error {
	tel, err := telemetry.NewTelemetry(ctx, a.cfg.Telemetry, a.name, version.NewVersion().Version())
	if err != nil {
		return err
	}

	a.telemetry = tel

	return nil
}
