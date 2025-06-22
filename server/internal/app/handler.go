package app

import (
	"context"
	"github.com/axidex/api-example/server/internal/api"
	"github.com/axidex/api-example/server/internal/controller"
)

func (a *App) initHandler(_ context.Context) error {
	a.handler = api.NewGinHandler(a.name, a.cfg.Api, a.logger, a.telemetry, controller.NewApiController(a.storage))

	return nil
}
