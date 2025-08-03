package fx

import (
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/transactions/internal/api"
	"github.com/axidex/api-example/transactions/internal/config"
	"github.com/axidex/api-example/transactions/internal/controller"
	"github.com/axidex/api-example/transactions/internal/handler"
	"github.com/axidex/api-example/transactions/internal/storage"
	internalTon "github.com/axidex/api-example/transactions/pkg/ton"
	"github.com/axidex/api-example/transactions/pkg/eg"
	"github.com/xssnick/tonutils-go/liteclient"
	"go.uber.org/fx"
)

var TransactionsHandlerModule = fx.Module("transactions-handler",
	fx.Provide(
		NewTonController,
		NewTransactionHandler,
	),
)

var ApiHandlerModule = fx.Module("api-handler",
	fx.Provide(NewGinHandler),
)

func NewTonController(
	tonService *internalTon.TransactionService,
	tonConnection *liteclient.ConnectionPool,
	appStorage *storage.AppStorage,
	egStorage *eg.StorageGorm,
	priceService *internalTon.PriceServiceCoinGecko,
	logger logger.Logger,
	cfg *config.TransactionsConfig,
) *controller.TonController {
	return controller.NewTonController(
		tonService,
		tonConnection,
		appStorage,
		egStorage,
		priceService,
		logger,
		cfg.EG,
	)
}

func NewTransactionHandler(ctrl *controller.TonController, logger logger.Logger) handler.Handler {
	return handler.NewTransactionHandler(ctrl, logger)
}

func NewGinHandler(cfg *config.ApiConfig, logger logger.Logger, tel telemetry.Telemetry) *api.GinHandler {
	return api.NewGinHandler("api", cfg.API, logger, tel)
}