package fx

import (
	"context"
	"github.com/axidex/api-example/server/pkg/config_provider"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/transactions/internal/config"
	"github.com/axidex/api-example/transactions/internal/handler"
	"github.com/oklog/run"
	"go.uber.org/fx"
	"syscall"
)

func NewTransactionsApp() *fx.App {
	return fx.New(
		TransactionsConfigModule,
		TransactionsTelemetryModule,
		TransactionsLoggerModule,
		DatabaseModule,
		TonModule,
		StorageModule,
		TransactionsHandlerModule,
		fx.Invoke(RunTransactionsApp),
	)
}

func RunTransactionsApp(lc fx.Lifecycle, handler handler.Handler, cfg *config.TransactionsConfig, logger logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			config_provider.PrintInfo(cfg, logger.Info)
			go func() {
				g := &run.Group{}
				g.Add(handler.Handle(ctx))
				g.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))
				_ = g.Run()
			}()
			return nil
		},
	})
}