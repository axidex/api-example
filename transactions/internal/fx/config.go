package fx

import (
	"fmt"
	"github.com/axidex/api-example/server/pkg/config_provider"
	"github.com/axidex/api-example/transactions/internal/config"
	"go.uber.org/fx"
)

var TransactionsConfigModule = fx.Module("transactions-config",
	fx.Provide(NewTransactionsConfig),
)

var ApiConfigModule = fx.Module("api-config",
	fx.Provide(NewApiConfig),
)

func NewTransactionsConfig() (*config.TransactionsConfig, error) {
	configPath := ".env"
	cfg, err := config_provider.ParseConfig[config.TransactionsConfig](configPath, config_provider.EnvConfig)
	if err != nil {
		return nil, fmt.Errorf("can't parse transactions config: %w", err)
	}
	return cfg, nil
}

func NewApiConfig() (*config.ApiConfig, error) {
	configPath := ".env"
	cfg, err := config_provider.ParseConfig[config.ApiConfig](configPath, config_provider.EnvConfig)
	if err != nil {
		return nil, fmt.Errorf("can't parse api config: %w", err)
	}
	return cfg, nil
}