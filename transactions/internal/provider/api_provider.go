package provider

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/transactions/internal/config"
)

// nolint
type ApiProvider struct {
	dependencies ApiDependencies
	cfg          *config.ApiConfig
	logger       logger.Logger
	telemetry    telemetry.Telemetry
	debug        bool
}

func NewApiProvider(cfg *config.ApiConfig, logger logger.Logger, telemetry telemetry.Telemetry) *ApiProvider {
	return &ApiProvider{cfg: cfg, logger: logger, telemetry: telemetry}
}

func (p *ApiProvider) InitDependencies(ctx context.Context) error {
	inits := map[string]initFunc{}
	for name, init := range inits {
		if err := init(ctx); err != nil {
			return fmt.Errorf("error got in %s: %w", name, err)
		}
	}

	return nil
}

func (p *ApiProvider) GetDependencies() *ApiDependencies {
	return &p.dependencies
}
