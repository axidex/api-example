package app

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/internal/config"
	"github.com/axidex/api-example/pkg/config_provider"
)

func (a *App) initConfig(_ context.Context) error {
	configPath := ".env" // os.Getenv("config")

	cfg, err := config_provider.ParseConfig[config.Config](configPath, config_provider.EnvConfig)
	if err != nil {
		return fmt.Errorf("can't parse config: %w", err)
	}

	a.cfg = cfg
	return nil
}
