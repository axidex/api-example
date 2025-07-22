package provider

import (
	"context"
	internalTon "github.com/axidex/api-example/transactions/pkg/ton"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
)

func (p *TransactionsProvider) initTonTransactions(ctx context.Context) error {

	conn := liteclient.NewConnectionPool()

	if err := conn.AddConnectionsFromConfigUrl(ctx, p.cfg.Ton.ConfigUrl); err != nil {
		return err
	}

	client := ton.NewAPIClient(conn)

	transactionService, err := internalTon.NewTonTransactionService(p.cfg.Ton.WalletAddress, client, p.logger)
	if err != nil {
		return err
	}

	p.dependencies.TonService = transactionService
	p.dependencies.TonConnection = conn

	return nil
}

func (p *TransactionsProvider) initTonPriceService(ctx context.Context) error {
	service, err := internalTon.NewPriceService(ctx, p.logger)
	if err != nil {
		return err
	}
	service.Start()

	p.dependencies.TonPriceService = service

	return nil
}
