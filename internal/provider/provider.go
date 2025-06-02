package provider

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/internal/config"
	"github.com/axidex/api-example/pkg/logger"
)

type IServiceProvider interface {
	InitDependencies(ctx context.Context) error
	GetDependencies() *Dependencies
}

type Provider struct {
	dependencies Dependencies
	cfg          *config.Config
	logger       logger.Logger
	debug        bool
}

func NewServiceProvider(cfg *config.Config, logger logger.Logger) IServiceProvider {
	return &Provider{cfg: cfg, logger: logger}
}

func (s *Provider) InitDependencies(ctx context.Context) error {
	inits := map[string]initFunc{}

	for name, init := range inits {
		if err := init(ctx); err != nil {
			return fmt.Errorf("error got in %s: %w", name, err)
		}
	}

	return nil
}

func (s *Provider) GetDependencies() *Dependencies {
	return &s.dependencies
}
