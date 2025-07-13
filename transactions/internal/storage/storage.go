package storage

import (
	"context"
	"github.com/axidex/api-example/transactions/internal/tables"
	"github.com/axidex/api-example/transactions/pkg/ton"
	"gorm.io/gorm"
)

type AppStorage struct {
	TransactionRepository TransactionRepository
	LogicTimeRepository   LogicTimeRepository
	db                    *gorm.DB
}

func NewApiStorage(db *gorm.DB) *AppStorage {
	return &AppStorage{
		TransactionRepository: NewTransactionRepository(db),
		LogicTimeRepository:   NewLogicTimeRepository(db),
		db:                    db,
	}
}

func (a AppStorage) SaveTransaction(ctx context.Context, transaction ton.Transaction) error {
	if err := a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		transactionRepoTx := NewTransactionRepository(tx)
		logicTimeRepoTx := NewLogicTimeRepository(tx)

		if err := transactionRepoTx.Save(ctx, tables.NewTransaction(transaction.Source, transaction.Amount)); err != nil {
			return err
		}

		if err := logicTimeRepoTx.Update(ctx, transaction.LT); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
