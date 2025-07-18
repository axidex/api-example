package app

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/server/internal/provider"
)

func (a *App) initDependencies(ctx context.Context) error {
	serviceProvider := provider.NewServiceProvider(a.cfg, a.logger, a.telemetry)

	if err := serviceProvider.InitDependencies(ctx); err != nil {
		return fmt.Errorf("can't init dependencies: %w", err)
	}

	dependencies := serviceProvider.GetDependencies()
	a.dependencies = dependencies

	return nil
}
