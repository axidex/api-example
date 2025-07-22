package provider

import (
	"context"
	"github.com/axidex/api-example/transactions/pkg/ton"
	"github.com/xssnick/tonutils-go/liteclient"
	"gorm.io/gorm"
)

type initFunc = func(context.Context) error

type TransactionsDependencies struct {
	DB              *gorm.DB
	TonService      *ton.TransactionService
	TonConnection   *liteclient.ConnectionPool
	TonPriceService *ton.PriceServiceCoinGecko
}

func (d *TransactionsDependencies) Stop() {
	d.TonConnection.Stop()
}

type ApiDependencies struct{}

func (d *ApiDependencies) Stop() {}
