package app

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/transactions/internal/provider"
)

func (a *TransactionsApp) initDependencies(ctx context.Context) error {
	serviceProvider := provider.NewTransactionsProvider(a.cfg, a.logger, a.telemetry)

	if err := serviceProvider.InitDependencies(ctx); err != nil {
		return fmt.Errorf("can't init dependencies: %w", err)
	}

	dependencies := serviceProvider.GetDependencies()
	a.dependencies = dependencies

	return nil
}

func (a *ApiApp) initDependencies(ctx context.Context) error {
	serviceProvider := provider.NewApiProvider(a.cfg, a.logger, a.telemetry)

	if err := serviceProvider.InitDependencies(ctx); err != nil {
		return fmt.Errorf("can't init dependencies: %w", err)
	}

	dependencies := serviceProvider.GetDependencies()
	a.dependencies = dependencies

	return nil
}
