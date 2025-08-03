package fx

import (
	"context"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/server/pkg/version"
	"github.com/axidex/api-example/transactions/internal/config"
	"go.uber.org/fx"
)

var TransactionsTelemetryModule = fx.Module("transactions-telemetry",
	fx.Provide(NewTransactionsTelemetry),
)

var ApiTelemetryModule = fx.Module("api-telemetry",
	fx.Provide(NewApiTelemetry),
)

func NewTransactionsTelemetry(lc fx.Lifecycle, cfg *config.TransactionsConfig) (telemetry.Telemetry, error) {
	tel, err := telemetry.NewTelemetry(context.Background(), cfg.Telemetry, "transactions", version.NewVersion().Version())
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			tel.Stop(ctx)
			return nil
		},
	})

	return tel, nil
}

func NewApiTelemetry(lc fx.Lifecycle, cfg *config.ApiConfig) (telemetry.Telemetry, error) {
	tel, err := telemetry.NewTelemetry(context.Background(), cfg.Telemetry, "api", version.NewVersion().Version())
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			tel.Stop(ctx)
			return nil
		},
	})

	return tel, nil
}