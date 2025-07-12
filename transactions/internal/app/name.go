package app

import (
	"context"
)

func (a *App) initName(_ context.Context) error {
	a.name = "transactions"
	return nil
}
