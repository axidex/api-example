package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/axidex/api-example/server/pkg/db"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/transactions/internal/storage"
	"github.com/axidex/api-example/transactions/internal/tables"
	"github.com/axidex/api-example/transactions/pkg/ton"
	"github.com/xssnick/tonutils-go/liteclient"
)

type Controller interface {
	Start(ctx context.Context) error
}

type TonController struct {
	tonService ton.ITransactionService
	conn       *liteclient.ConnectionPool
	storage    *storage.AppStorage
	logger     logger.Logger
}

func NewTonController(tonService ton.ITransactionService, conn *liteclient.ConnectionPool, storage *storage.AppStorage, logger logger.Logger) *TonController {
	return &TonController{
		conn:       conn,
		tonService: tonService,
		storage:    storage,
		logger:     logger,
	}
}

func (c *TonController) Start(ctx context.Context) error {
	stickyCtx := c.conn.StickyContext(ctx)
	c.logger.Info(ctx, "Starting transaction handler")

	ltValue, err := c.storage.LogicTimeRepository.Get(ctx)
	if errors.Is(err, db.ErrRecordNotFound) {
		ltValue, err = c.tonService.GetTxLT(stickyCtx)
		if err != nil {
			return err
		}
		if err := c.storage.LogicTimeRepository.Create(ctx, tables.NewLogicTime(ltValue)); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	c.logger.Info(ctx, fmt.Sprintf("Transaction handler started with LT: %d", ltValue))
	internalTxChan := make(chan ton.Transaction)
	errChan := make(chan error)
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
		case errChan <- c.tonService.StartListenTransactions(stickyCtx, ltValue, internalTxChan):
		}
	}(stickyCtx)

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errChan:
			return err
		case tx, ok := <-internalTxChan:
			if !ok {
				return errors.New("internal transaction channel closed")
			}

			if err := c.saveTransaction(ctx, tx); err != nil {
				return err
			}
		}
	}
}

func (c *TonController) saveTransaction(ctx context.Context, transaction ton.Transaction) error {
	if err := c.storage.SaveTransaction(ctx, transaction); err != nil {
		return err
	}

	return nil
}
