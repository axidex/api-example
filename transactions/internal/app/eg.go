package app

import (
	"context"
	"github.com/axidex/api-example/transactions/pkg/eg"
)

func (a *TransactionsApp) initEg(_ context.Context) error {
	a.egStorage = eg.NewEGStorage(a.dependencies.DB)

	return nil
}
