package app

import (
	"context"
)

func (a *TransactionsApp) initName(_ context.Context) error {
	a.name = "transactions"
	return nil
}

func (a *ApiApp) initName(_ context.Context) error {
	a.name = "api"
	return nil
}
