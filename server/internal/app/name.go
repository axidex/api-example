package app

import (
	"context"
)

func (a *App) initName(_ context.Context) error {
	a.name = "api-example"
	return nil
}
