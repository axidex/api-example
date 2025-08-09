package fx

import (
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/transactions/internal/config"
	"go.uber.org/fx"
)

var TransactionsLoggerModule = fx.Module("transactions-logger",
	fx.Provide(NewTransactionsLogger),
)

var ApiLoggerModule = fx.Module("api-logger",
	fx.Provide(NewApiLogger),
)

func NewTransactionsLogger(cfg *config.TransactionsConfig, tel telemetry.Telemetry) (logger.Logger, error) {
	return logger.NewZapLogger(cfg.Logger, "transactions", tel.GetLoggerProvider())
}

func NewApiLogger(cfg *config.ApiConfig, tel telemetry.Telemetry) (logger.Logger, error) {
	return logger.NewZapLogger(cfg.Logger, "api", tel.GetLoggerProvider())
}
