package ton

import (
	"context"
	"errors"
	"fmt"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
)

var (
	TransactionNotFoundInLiteServer = errors.New("transaction not found in lite server")
	UnknownLiteServerError          = errors.New("unknown lite server error")
)

type ITransactionService interface {
	StartListenTransactions(ctx context.Context, lastProcessedLT uint64, internalTxChan chan<- Transaction) error
	GetTxLT(ctx context.Context) (uint64, error)
	PollTransactions(ctx context.Context, lastProcessedLT uint64) error
}

type TransactionService struct {
	address *address.Address
	client  ton.APIClientWrapped
	logger  logger.Logger
}

func NewTonTransactionService(walletAddr string, client ton.APIClientWrapped, logger logger.Logger) (*TransactionService, error) {
	addr, err := address.ParseAddr(walletAddr)
	if err != nil {
		return nil, err
	}

	return &TransactionService{
		address: addr,
		client:  client,
		logger:  logger,
	}, nil
}

// StartListenTransactions if nil then lasTxLT would be acc.LastTxLT
func (service *TransactionService) StartListenTransactions(ctx context.Context, lastProcessedLT uint64, internalTxChan chan<- Transaction) error {
	txChan := make(chan *tlb.Transaction)
	errChan := make(chan error)

	go service.client.SubscribeOnTransactions(ctx, service.address, lastProcessedLT, txChan)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case tx, ok := <-txChan:
				if !ok {
					service.logger.Error(ctx, "txChan closed, stopping transaction listener")
					errChan <- errors.New("txChan closed, stopping transaction listener")
					return
				}

				if tx.IO.In != nil && tx.IO.In.MsgType == tlb.MsgTypeInternal {
					ti := tx.IO.In.AsInternal()

					if dsc, ok := tx.Description.(tlb.TransactionDescriptionOrdinary); ok && dsc.BouncePhase != nil {
						if _, ok = dsc.BouncePhase.Phase.(tlb.BouncePhaseOk); ok {
							// transaction was bounced, and coins were returned to sender
							// this can happen mostly on custom contracts
							continue
						}
					}

					userId, err := DecodeStringPayload(ti.Payload())
					if err != nil {
						service.logger.Warn(ctx, fmt.Sprintf("Ton payload decode failed: %s", err.Error()))
					}

					if ti.Amount.Nano().Sign() > 0 {
						validatedUserId, isValid := ValidateUserID(userId)
						if !isValid {
							service.logger.Warn(ctx,
								"Invalid userId in transaction payload, skipping transaction",
								logger.NewAttribute("invalid_userId", fmt.Sprintf("%q", userId)),
								logger.NewAttribute("from", ti.SrcAddr.StringRaw()),
								logger.NewAttribute("amount", ti.Amount.String()),
							)
						}

						service.logger.Info(
							ctx, "received transaction",
							logger.NewAttribute("userId", validatedUserId),
							logger.NewAttribute("amount", ti.Amount.String()),
							logger.NewAttribute("from", ti.SrcAddr.StringRaw()),
							logger.NewAttribute("lt", tx.LT),
						)

						lastProcessedLT = tx.LT
						internalTxChan <- NewTransaction(ti.SrcAddr.StringRaw(), validatedUserId, ti.Amount.Nano().Uint64(), lastProcessedLT)
					}
				}
			}
		}
	}(ctx)

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

func (service *TransactionService) GetTxLT(ctx context.Context) (uint64, error) {
	master, err := service.client.CurrentMasterchainInfo(ctx) // we fetch block just to trigger chain proof check
	if err != nil {
		return 0, err
	}

	acc, err := service.client.GetAccount(ctx, master, service.address)
	if err != nil {
		return 0, err
	}

	return acc.LastTxLT, nil
}

func (service *TransactionService) PollTransactions(ctx context.Context, lastProcessedLT uint64) error {
	_, err := service.client.ListTransactions(ctx, service.address, 1, lastProcessedLT, nil)
	if err != nil {
		var lsErr ton.LSError
		if errors.As(err, &lsErr) {
			service.logger.Warn(ctx, fmt.Sprintf("LSError received: Code=%d, Text=%s", lsErr.Code, lsErr.Text), logger.NewAttribute("LT", lastProcessedLT))

			switch lsErr.Code {
			case -400:
				if lsErr.Text == "transaction hash mismatch" {
					return nil
				}
				return TransactionNotFoundInLiteServer
			case 0:
				// no transactions found
				return nil
			default:
				service.logger.Error(ctx, fmt.Sprintf("liteserver error (code %d): %s", lsErr.Code, lsErr.Text), logger.NewAttribute("LT", lastProcessedLT))
				return UnknownLiteServerError
			}
		}
	}

	return nil
}
