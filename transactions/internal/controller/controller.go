package controller

import (
	"context"
	"errors"
	"fmt"
	"github.com/axidex/api-example/server/pkg/db"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/transactions/internal/storage"
	"github.com/axidex/api-example/transactions/internal/tables"
	"github.com/axidex/api-example/transactions/pkg/eg"
	"github.com/axidex/api-example/transactions/pkg/ton"
	"github.com/xssnick/tonutils-go/liteclient"
)

type Controller interface {
	Start(ctx context.Context) error
}

type TonController struct {
	tonService   ton.ITransactionService
	conn         *liteclient.ConnectionPool
	storage      *storage.AppStorage
	egStorage    eg.Storage
	priceService ton.PriceService
	logger       logger.Logger
	egAvailable  bool
}

func NewTonController(tonService ton.ITransactionService, conn *liteclient.ConnectionPool, storage *storage.AppStorage, egStorage eg.Storage, priceService ton.PriceService, logger logger.Logger, egAvailable bool) *TonController {
	return &TonController{
		conn:         conn,
		tonService:   tonService,
		storage:      storage,
		egStorage:    egStorage,
		priceService: priceService,
		logger:       logger,
		egAvailable:  egAvailable,
	}
}

func (c *TonController) Start(ctx context.Context) error {
	stickyCtx := c.conn.StickyContext(ctx)
	c.logger.Info(ctx, "Starting transaction handler")

	ltValue, err := c.GetLT(stickyCtx)
	if err != nil {
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

	if transaction.UserId != nil {
		if err := c.egStorage.DepositToUser(ctx, *transaction.UserId, ton.NanoTonsToStars(transaction.Amount, c.priceService.GetPrice())); err != nil {
			c.logger.Error(ctx, fmt.Sprintf("Failed to deposit to user: %s", err.Error()))
		}
	}

	return nil
}

func (c *TonController) GetLT(ctx context.Context) (uint64, error) {
	ltValue, err := c.storage.LogicTimeRepository.Get(ctx)
	if errors.Is(err, db.ErrRecordNotFound) {
		if ltValue, err = c.createCurrentLT(ctx); err != nil {
			return 0, err
		}
	} else if err != nil {
		return 0, err
	}

	// Verify that this LT is still in this block in the TON lite client
	if err := c.tonService.PollTransactions(ctx, ltValue); err != nil {
		if errors.Is(err, ton.TransactionNotFoundInLiteServer) {
			ltValue, err = c.updateCurrentLT(ctx)
			if err != nil {
				return 0, err
			}
		} else {
			return 0, err
		}
	}

	return ltValue, nil
}

func (c *TonController) createCurrentLT(ctx context.Context) (uint64, error) {
	ltValue, err := c.tonService.GetTxLT(ctx)
	if err != nil {
		return 0, err
	}

	if err := c.storage.LogicTimeRepository.Create(ctx, tables.NewLogicTime(ltValue)); err != nil {
		return 0, err
	}

	return ltValue, nil
}

func (c *TonController) updateCurrentLT(ctx context.Context) (uint64, error) {
	ltValue, err := c.tonService.GetTxLT(ctx)
	if err != nil {
		return 0, err
	}

	if err := c.storage.LogicTimeRepository.Update(ctx, ltValue); err != nil {
		return 0, err
	}

	return ltValue, nil
}
