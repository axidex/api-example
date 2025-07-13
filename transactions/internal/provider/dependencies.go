package provider

import (
	"context"
	"github.com/axidex/api-example/transactions/pkg/ton"
	"github.com/xssnick/tonutils-go/liteclient"
	"gorm.io/gorm"
)

type initFunc = func(context.Context) error

type Dependencies struct {
	DB            *gorm.DB
	TonService    *ton.TransactionService
	TonConnection *liteclient.ConnectionPool
}

func (d *Dependencies) Stop() {
	d.TonConnection.Stop()
}
