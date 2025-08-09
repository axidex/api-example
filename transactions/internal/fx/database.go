package fx

import (
	"github.com/axidex/api-example/server/pkg/db"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/transactions/internal/config"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var DatabaseModule = fx.Module("database",
	fx.Provide(NewDatabase),
)

func NewDatabase(cfg *config.TransactionsConfig, logger logger.Logger, tel telemetry.Telemetry) (*gorm.DB, error) {
	psqlEngine, err := db.NewGormConnection(cfg.Database, logger)
	if err != nil {
		return nil, err
	}

	engineWithTelemetry, err := db.WithTelemetry(psqlEngine, tel.GetTracerProvider())
	if err != nil {
		return nil, err
	}

	return engineWithTelemetry, nil
}