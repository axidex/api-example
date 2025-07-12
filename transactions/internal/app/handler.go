package app

import (
	"context"
	"github.com/axidex/api-example/transactions/internal/controller"
	"github.com/axidex/api-example/transactions/internal/handler"
)

func (a *App) initHandler(_ context.Context) error {
	ctrl := controller.NewTonController(
		a.dependencies.TonService,
		a.dependencies.TonConnection,
		a.storage,
		a.logger,
	)

	a.handler = handler.NewTransactionHandler(ctrl, a.logger)

	return nil
}
