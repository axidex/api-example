package app

import (
	"context"
	"github.com/axidex/api-example/server/pkg/config_provider"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/transactions/internal/config"
	"github.com/axidex/api-example/transactions/internal/handler"
	"github.com/axidex/api-example/transactions/internal/provider"
	"github.com/axidex/api-example/transactions/internal/storage"
	"github.com/axidex/api-example/transactions/pkg/eg"
	"github.com/oklog/run"

	"syscall"
)

type TransactionsApp struct {
	storage      *storage.AppStorage
	egStorage    *eg.StorageGorm
	telemetry    telemetry.Telemetry
	dependencies *provider.TransactionsDependencies
	cfg          *config.TransactionsConfig
	logger       logger.Logger
	handler      handler.Handler
	name         string
}

func NewTransactionsApp() IApp {
	return &TransactionsApp{}
}

func (a *TransactionsApp) Run(ctx context.Context) error {
	if err := a.init(ctx); err != nil {
		return err
	}
	defer a.stop()

	g := &run.Group{}
	g.Add(a.handler.Handle(ctx))
	g.Add(run.SignalHandler(ctx, syscall.SIGINT, syscall.SIGTERM))

	if err := g.Run(); err != nil {
		//a.logger.Errorf("Error occured | %s", err.Error())
		return err
	}

	return nil
}

func (a *TransactionsApp) init(ctx context.Context) error {

	inits := []initFunc{
		a.initConfig,
		a.initName,
		a.initTelemetry,
		a.initLogger,
		a.initDependencies,
		a.initStorage,
		a.initEg,
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

func (a *TransactionsApp) stop() {
	a.dependencies.Stop()
	a.telemetry.Stop(context.Background())
}
