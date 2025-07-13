package app

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/transactions/internal/config"

	"github.com/axidex/api-example/server/pkg/config_provider"
)

func (a *TransactionsApp) initConfig(_ context.Context) error {
	configPath := ".env" // os.Getenv("config")

	cfg, err := config_provider.ParseConfig[config.TransactionsConfig](configPath, config_provider.EnvConfig)
	if err != nil {
		return fmt.Errorf("can't parse config: %w", err)
	}

	a.cfg = cfg
	return nil
}

func (a *ApiApp) initConfig(_ context.Context) error {
	configPath := ".env" // os.Getenv("config")

	cfg, err := config_provider.ParseConfig[config.ApiConfig](configPath, config_provider.EnvConfig)
	if err != nil {
		return fmt.Errorf("can't parse config: %w", err)
	}

	a.cfg = cfg
	return nil
}
