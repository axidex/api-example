package storage

import (
	"context"
	"github.com/axidex/api-example/server/pkg/db"
	"github.com/axidex/api-example/transactions/internal/tables"
	"gorm.io/gorm"
)

type LogicTimeRepository interface {
	Create(ctx context.Context, lt *tables.LogicTime) error
	Get(ctx context.Context) (uint64, error)
	Update(ctx context.Context, value uint64) error
}

type logicTimeRepository struct {
	db *gorm.DB
}

func NewLogicTimeRepository(db *gorm.DB) LogicTimeRepository {
	return &logicTimeRepository{db: db}
}

func (r *logicTimeRepository) Get(ctx context.Context) (uint64, error) {
	var ltTables []tables.LogicTime

	if err := r.db.WithContext(ctx).Find(&ltTables).Error; err != nil {
		return 0, err
	}

	if len(ltTables) == 0 {
		return 0, db.ErrRecordNotFound
	}

	return ltTables[0].Value, nil
}

func (r *logicTimeRepository) Create(ctx context.Context, lt *tables.LogicTime) error {
	if err := r.db.WithContext(ctx).Create(lt).Error; err != nil {
		return err
	}

	return nil
}

func (r *logicTimeRepository) Update(ctx context.Context, value uint64) error {
	var lt tables.LogicTime

	if err := r.db.WithContext(ctx).First(&lt).Error; err != nil {
		return err
	}

	if err := r.db.WithContext(ctx).
		Model(&lt).
		Update("value", value).Error; err != nil {
		return err
	}

	return nil
}
