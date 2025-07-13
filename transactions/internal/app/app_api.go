package app

import (
	"context"
	"github.com/axidex/api-example/server/pkg/config_provider"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/transactions/internal/api"
	"github.com/axidex/api-example/transactions/internal/config"
	"github.com/axidex/api-example/transactions/internal/provider"
	"github.com/oklog/run"
	"syscall"
)

type ApiApp struct {
	telemetry    telemetry.Telemetry
	dependencies *provider.ApiDependencies
	cfg          *config.ApiConfig
	logger       logger.Logger
	handler      *api.GinHandler
	name         string
}

func NewApiApp() IApp {
	return &ApiApp{}
}

func (a *ApiApp) Run(ctx context.Context) error {
	if err := a.init(ctx); err != nil {
		return err
	}
	defer a.stop()

	g := &run.Group{}
	g.Add(a.handler.HandleServer(ctx))
	g.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))

	if err := g.Run(); err != nil {
		//a.logger.Errorf("Error occured | %s", err.Error())
		return err
	}

	return nil
}

func (a *ApiApp) init(ctx context.Context) error {

	inits := []initFunc{
		a.initConfig,
		a.initName,
		a.initTelemetry,
		a.initLogger,
		a.initDependencies,
		a.initHandler,
	}

	for _, init := range inits {
		if err := init(ctx); err != nil {
			return err
		}
	}

	config_provider.PrintInfo(a.cfg, a.logger.Info)

	return nil
}

func (a *ApiApp) stop() {
	a.dependencies.Stop()
	a.telemetry.Stop(context.Background())
}
