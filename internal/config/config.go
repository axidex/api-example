package config

import (
	"github.com/axidex/api-example/internal/api"
	"github.com/axidex/api-example/pkg/logger"
	"github.com/axidex/api-example/pkg/telemetry"
)

type Config struct {
	Telemetry telemetry.Config `mapstructure:",squash"`
	Logger    logger.Config    `mapstructure:",squash"`
	Api       api.Config       `mapstructure:",squash"`
}
