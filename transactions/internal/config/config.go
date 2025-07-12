package config

import (
	"github.com/axidex/api-example/server/pkg/db"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
	"github.com/axidex/api-example/server/pkg/ton"
)

type Config struct {
	Telemetry telemetry.Config `mapstructure:",squash"`
	Logger    logger.Config    `mapstructure:",squash"`
	Database  db.Config        `mapstructure:",squash"`
	Ton       ton.Config       `mapstructure:",squash"`
}
