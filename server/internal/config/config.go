package config

import (
	"github.com/axidex/api-example/server/internal/api"
	"github.com/axidex/api-example/server/pkg/db"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/server/pkg/telemetry"
)

type Config struct {
	Telemetry telemetry.Config `mapstructure:",squash"`
	Logger    logger.Config    `mapstructure:",squash"`
	Api       api.Config       `mapstructure:",squash"`
	Database  db.Config        `mapstructure:",squash"`
}
