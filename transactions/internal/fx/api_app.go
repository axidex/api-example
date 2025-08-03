package fx

import (
	"context"
	"github.com/axidex/api-example/server/pkg/config_provider"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/transactions/internal/api"
	"github.com/axidex/api-example/transactions/internal/config"
	"github.com/oklog/run"
	"go.uber.org/fx"
	"syscall"
)

func NewApiApp() *fx.App {
	return fx.New(
		ApiConfigModule,
		ApiTelemetryModule,
		ApiLoggerModule,
		ApiHandlerModule,
		fx.Invoke(RunApiApp),
	)
}

func RunApiApp(lc fx.Lifecycle, handler *api.GinHandler, cfg *config.ApiConfig, logger logger.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			config_provider.PrintInfo(cfg, logger.Info)
			go func() {
				g := &run.Group{}
				g.Add(handler.HandleServer(ctx))
				g.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))
				_ = g.Run()
			}()
			return nil
		},
	})
}