package logger

import "fmt"

type Config struct {
	Level    string   `mapstructure:"LOGGER_LEVEL"`
	Path     string   `mapstructure:"LOGGER_PATH"`
	Rotation Rotation `mapstructure:",squash"`
}

func (c Config) Info() string {
	return fmt.Sprintf(
		"Logger: Level - %s, Path - %s, RotationMaxSize - %d, RotationMaxAge - %d, RotationMaxBackups - %d, Compress - %t",
		c.Level, c.Path, c.Rotation.MaxSize, c.Rotation.MaxAge, c.Rotation.MaxBackups, c.Rotation.Compress,
	)
}

type Rotation struct {
	MaxSize    int  `mapstructure:"LOGGER_ROTATION_MAX_SIZE"`
	MaxAge     int  `mapstructure:"LOGGER_ROTATION_MAX_AGE"`
	MaxBackups int  `mapstructure:"LOGGER_ROTATION_MAX_BACKUPS"`
	Compress   bool `mapstructure:"LOGGER_ROTATION_COMPRESS"`
}
