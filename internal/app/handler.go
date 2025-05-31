package app

import (
	"context"
	"github.com/axidex/api-example/internal/api"
)

func (a *App) initHandler(_ context.Context) error {
	a.handler = api.NewGinHandler(a.name, a.cfg.Api, a.logger, a.telemetry)

	return nil
}
