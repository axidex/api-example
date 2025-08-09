package fx

import (
	"context"
	"github.com/axidex/api-example/server/pkg/migrations"
	"github.com/axidex/api-example/transactions/internal/config"
	"github.com/axidex/api-example/transactions/internal/storage"
	"github.com/axidex/api-example/transactions/internal/tables"
	"github.com/axidex/api-example/transactions/pkg/eg"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var StorageModule = fx.Module("storage",
	fx.Provide(
		NewAppStorage,
		NewEGStorage,
	),
	fx.Invoke(runMigrations),
)

func NewAppStorage(db *gorm.DB) *storage.AppStorage {
	return storage.NewApiStorage(db)
}

func NewEGStorage(db *gorm.DB) *eg.StorageGorm {
	return eg.NewEGStorage(db)
}

func runMigrations(lc fx.Lifecycle, db *gorm.DB, cfg *config.TransactionsConfig) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return migrations.CreateMigrator(ctx, db).
				Migrate(
					[]interface{}{
						&tables.Transaction{},
						&tables.LogicTime{},
					},
					cfg.Database.Connection.Schema,
					cfg.Database.Connection.Owner,
				)
		},
	})
}
