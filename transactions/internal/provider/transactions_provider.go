package provider

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/transactions/internal/config"
)

// nolint
type TransactionsProvider struct {
	dependencies TransactionsDependencies
	cfg          *config.TransactionsConfig
	logger       logger.Logger
	telemetry    telemetry.Telemetry
	debug        bool
}

func NewTransactionsProvider(cfg *config.TransactionsConfig, logger logger.Logger, telemetry telemetry.Telemetry) *TransactionsProvider {
	return &TransactionsProvider{cfg: cfg, logger: logger, telemetry: telemetry}
}

func (p *TransactionsProvider) InitDependencies(ctx context.Context) error {
	inits := map[string]initFunc{
		"database": p.initDatabase,
		"ton":      p.initTonTransactions,
	}
	for name, init := range inits {
		if err := init(ctx); err != nil {
			return fmt.Errorf("error got in %s: %w", name, err)
		}
	}

	return nil
}

func (p *TransactionsProvider) GetDependencies() *TransactionsDependencies {
	return &p.dependencies
}
