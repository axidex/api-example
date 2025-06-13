package app

import (
	"context"
	"github.com/axidex/api-example/pkg/migrations"
	"gorm.io/gorm"
)

func (a *App) initStorage(ctx context.Context) error {
	if err := runMigrations(
		ctx,
		a.dependencies.DB,
		a.cfg.Database.Connection.Schema,
		a.cfg.Database.Connection.Owner,
	); err != nil {
		return err
	}

	return nil
}

func runMigrations(ctx context.Context, db *gorm.DB, schemaName, ownerName string) error {

	return migrations.CreateMigrator(ctx, db).Migrate(schemaName, ownerName)
}
