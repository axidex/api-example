package db

import (
	"fmt"
	"github.com/axidex/api-example/server/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Credentials Credentials `mapstructure:",squash"`
	Connection  Connection  `mapstructure:",squash"`
}

type Credentials struct {
	Username string `mapstructure:"POSTGRES_USERNAME"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
}

type Connection struct {
	Hosts    string `mapstructure:"POSTGRES_HOSTS"`
	Database string `mapstructure:"POSTGRES_DATABASE"`
	Params   string `mapstructure:"POSTGRES_PARAMS"`
	Schema   string `mapstructure:"POSTGRES_SCHEMA"`
	Owner    string `mapstructure:"POSTGRES_OWNER"`
}

func (c Connection) Info() string {
	return fmt.Sprintf(
		"Postgres: Hosts - %s, Database - %s, Params - %s, Schema - %s, Owner - %s",
		c.Hosts, c.Database, c.Params, c.Schema, c.Owner,
	)
}

func NewGormConnection(config Config, logger logger.Logger) (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s/%s?%s",
		config.Credentials.Username,
		config.Credentials.Password,
		config.Connection.Hosts,
		config.Connection.Database,
		config.Connection.Params,
	)

	logConfig := gorm.Config{Logger: NewGormLogger(logger)}

	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dataSourceName, PreferSimpleProtocol: true}), &logConfig)
	if err != nil {
		return nil, fmt.Errorf("can't create postgres connection: %w", err)
	}

	return db, nil
}
