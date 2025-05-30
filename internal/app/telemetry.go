package app

import (
	"context"
	"github.com/axidex/ss-manager/pkg/telemetry"
)

func (a *App) initTelemetry(ctx context.Context) error {
	tel, err := telemetry.NewTelemetry(ctx, a.cfg.Telemetry, a.cfg.App.Name, a.cfg.App.Version)
	if err != nil {
		return err
	}

	a.telemetry = tel
	return nil
}
