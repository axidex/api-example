package handler

import (
	"context"
	"fmt"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/axidex/api-example/transactions/internal/controller"
)

type Handler interface {
	Handle(ctx context.Context) (func() error, func(error))
}

type TransactionHandler struct {
	controller controller.Controller
	logger     logger.Logger
}

func NewTransactionHandler(controller controller.Controller, logger logger.Logger) *TransactionHandler {

	return &TransactionHandler{
		controller: controller,
		logger:     logger,
	}
}

func (h *TransactionHandler) Handle(ctx context.Context) (func() error, func(error)) {
	ctx, cancel := context.WithCancel(ctx)
	errorChan := make(chan error, 1)
	return func() error {
			go func() {
				errorChan <- h.controller.Start(ctx)
			}()
			select {
			case <-ctx.Done():
				return ctx.Err()
			case err := <-errorChan:
				return err
			}
		},
		func(err error) {
			cancel()
			h.logger.Warn(ctx, fmt.Sprintf("Shutting down server: %s", err))
		}
}
