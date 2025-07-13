package app

import (
	"context"
	"github.com/axidex/api-example/transactions/internal/api"
	"github.com/axidex/api-example/transactions/internal/controller"
	"github.com/axidex/api-example/transactions/internal/handler"
)

func (a *TransactionsApp) initHandler(_ context.Context) error {
	ctrl := controller.NewTonController(
		a.dependencies.TonService,
		a.dependencies.TonConnection,
		a.storage,
		a.logger,
	)

	a.handler = handler.NewTransactionHandler(ctrl, a.logger)

	return nil
}

func (a *ApiApp) initHandler(_ context.Context) error {
	a.handler = api.NewGinHandler(a.name, a.cfg.API, a.logger, a.telemetry)

	return nil
}
