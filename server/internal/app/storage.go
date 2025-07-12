package app

import (
	"context"
	"github.com/axidex/api-example/server/internal/storage"
	"github.com/axidex/api-example/server/pkg/migrations"
	"github.com/axidex/api-example/server/pkg/tables"
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

	a.storage = storage.NewApiStorage(a.dependencies.DB)

	return nil
}

func runMigrations(ctx context.Context, db *gorm.DB, schemaName, ownerName string) error {

	return migrations.CreateMigrator(ctx, db).
		Migrate(
			[]interface{}{
				&tables.User{},
			},
			schemaName, ownerName,
		)
}
