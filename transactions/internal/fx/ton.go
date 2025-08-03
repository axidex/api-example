package fx

import (
	"context"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/transactions/internal/config"
	internalTon "github.com/axidex/api-example/transactions/pkg/ton"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/ton"
	"go.uber.org/fx"
)

var TonModule = fx.Module("ton",
	fx.Provide(
		NewTonConnection,
		NewTonClient,
		NewTonTransactionService,
		NewTonPriceService,
	),
)

func NewTonConnection(lc fx.Lifecycle, cfg *config.TransactionsConfig) (*liteclient.ConnectionPool, error) {
	conn := liteclient.NewConnectionPool()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return conn.AddConnectionsFromConfigUrl(ctx, cfg.Ton.ConfigUrl)
		},
		OnStop: func(ctx context.Context) error {
			conn.Stop()
			return nil
		},
	})

	return conn, nil
}

func NewTonClient(conn *liteclient.ConnectionPool) *ton.APIClient {
	return ton.NewAPIClient(conn)
}

func NewTonTransactionService(cfg *config.TransactionsConfig, client *ton.APIClient, logger logger.Logger) (*internalTon.TransactionService, error) {
	return internalTon.NewTonTransactionService(cfg.Ton.WalletAddress, client, logger)
}

func NewTonPriceService(lc fx.Lifecycle, logger logger.Logger) (*internalTon.PriceServiceCoinGecko, error) {
	service, err := internalTon.NewPriceService(context.Background(), logger)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			service.Start()
			return nil
		},
	})

	return service, nil
}