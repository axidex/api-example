package app

import (
	"context"
	"github.com/axidex/api-example/server/internal/api"
	"github.com/axidex/api-example/server/internal/config"
	"github.com/axidex/api-example/server/internal/provider"
	"github.com/axidex/api-example/server/internal/storage"
	"github.com/axidex/api-example/server/pkg/config_provider"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/oklog/run"

	"syscall"
)

type initFunc func(context.Context) error

type IApp interface {
	Run(ctx context.Context) error
}

type App struct {
	handler      *api.GinHandler
	storage      storage.ApiStorage
	telemetry    telemetry.Telemetry
	dependencies *provider.Dependencies
	cfg          *config.Config
	logger       logger.Logger
	name         string
}

func NewApp() IApp {
	return &App{}
}

func (a *App) Run(ctx context.Context) error {
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

func (a *App) init(ctx context.Context) error {

	inits := []initFunc{
		a.initConfig,
		a.initName,
		a.initTelemetry,
		a.initLogger,
		a.initDependencies,
		a.initStorage,
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

func (a *App) stop() {
	a.dependencies.Stop()
	a.telemetry.Stop(context.Background())
}
