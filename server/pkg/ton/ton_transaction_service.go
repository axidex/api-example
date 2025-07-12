package ton

import (
	"context"
	"github.com/axidex/api-example/server/pkg/logger"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
)

type ITransactionService interface {
	StartListenTransactions(ctx context.Context) error
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
func (service *TransactionService) StartListenTransactions(ctx context.Context, lastProcessedLT *uint64, internalTxChan chan<- Transaction) error {
	txChan := make(chan *tlb.Transaction)

	txLt, err := service.getTxLT(ctx, lastProcessedLT)
	if err != nil {
		return err
	}

	go service.client.SubscribeOnTransactions(ctx, service.address, txLt, txChan)

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case tx := <-txChan:
				if tx.IO.In != nil && tx.IO.In.MsgType == tlb.MsgTypeInternal {
					ti := tx.IO.In.AsInternal()

					if dsc, ok := tx.Description.(tlb.TransactionDescriptionOrdinary); ok && dsc.BouncePhase != nil {
						if _, ok = dsc.BouncePhase.Phase.(tlb.BouncePhaseOk); ok {
							// transaction was bounced, and coins were returned to sender
							// this can happen mostly on custom contracts
							continue
						}
					}

					if ti.Amount.Nano().Sign() > 0 {
						service.logger.Info(
							ctx, "received transaction",
							logger.NewAttribute("amount", ti.Amount.String()),
							logger.NewAttribute("from", ti.SrcAddr.StringRaw()),
						)

						txLt = tx.LT
						internalTxChan <- NewTransaction(ti.SrcAddr.StringRaw(), ti.Amount.Nano().Uint64(), txLt)
					}
				}
			}
		}
	}(ctx)

	// nolint
	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (service *TransactionService) getTxLT(ctx context.Context, lastProcessedLT *uint64) (uint64, error) {
	if lastProcessedLT != nil {
		return *lastProcessedLT, nil
	}

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
