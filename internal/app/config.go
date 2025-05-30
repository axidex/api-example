package app

import (
	"context"
	"fmt"
	"github.com/axidex/ss-manager/pkg/config_provider"
	"os"
)

func (a *App) initConfig(_ context.Context) error {
	configPath := os.Getenv("config")

	cfg, err := config_provider.ParseConfig[config.Config](configPath)
	if err != nil {
		return fmt.Errorf("can't parse config: %w", err)
	}
	config_provider.WaitSecrets(cfg.Secrets)

	a.cfg = cfg
	return nil
}
