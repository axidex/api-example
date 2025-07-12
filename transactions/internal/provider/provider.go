package provider

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/transactions/internal/config"
)

type IDependenciesProvider interface {
	InitDependencies(ctx context.Context) error
	GetDependencies() *Dependencies
}

// nolint
type Provider struct {
	dependencies Dependencies
	cfg          *config.Config
	logger       logger.Logger
	telemetry    telemetry.Telemetry
	debug        bool
}

func NewServiceProvider(cfg *config.Config, logger logger.Logger, telemetry telemetry.Telemetry) IDependenciesProvider {
	return &Provider{cfg: cfg, logger: logger, telemetry: telemetry}
}

func (p *Provider) InitDependencies(ctx context.Context) error {
	inits := map[string]initFunc{
		"database": p.initDatabase,
		"ton":      p.initTon,
	}
	for name, init := range inits {
		if err := init(ctx); err != nil {
			return fmt.Errorf("error got in %s: %w", name, err)
		}
	}

	return nil
}

func (p *Provider) GetDependencies() *Dependencies {
	return &p.dependencies
}
