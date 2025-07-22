package config

import (
	"github.com/axidex/api-example/server/pkg/db"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/transactions/internal/api"
	"github.com/axidex/api-example/transactions/pkg/ton"
)

type TransactionsConfig struct {
	Telemetry telemetry.Config `mapstructure:",squash"`
	Logger    logger.Config    `mapstructure:",squash"`
	Database  db.Config        `mapstructure:",squash"`
	Ton       ton.Config       `mapstructure:",squash"`
	EG        bool             `mapstructure:"eg"`
}

type ApiConfig struct {
	Telemetry telemetry.Config `mapstructure:",squash"`
	Logger    logger.Config    `mapstructure:",squash"`
	API       api.Config       `mapstructure:",squash"`
}
