package storage

import (
	"context"
	"github.com/axidex/api-example/transactions/internal/tables"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(ctx context.Context, transaction *tables.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Save(ctx context.Context, transaction *tables.Transaction) error {
	if err := r.db.WithContext(ctx).Create(transaction).Error; err != nil {
		return err
	}

	return nil
}
