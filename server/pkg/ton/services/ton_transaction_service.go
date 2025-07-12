package services

import (
	"context"
	"github.com/axidex/api-example/server/pkg/ton/dto"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"log"
)

type TransactionService interface {
	StartListenTransactions(ctx context.Context) error
}

type TonTransactionService struct {
	address *address.Address
	client  ton.APIClientWrapped
}

func NewTonTransactionService(walletAddr string, client ton.APIClientWrapped) (*TonTransactionService, error) {
	addr, err := address.ParseAddr(walletAddr)
	if err != nil {
		return nil, err
	}

	return &TonTransactionService{
		address: addr,
		client:  client,
	}, nil
}

// StartListenTransactions if nil then lasTxLT would be acc.LastTxLT
func (service *TonTransactionService) StartListenTransactions(ctx context.Context, lastProcessedLT *uint64, internalTxChan chan<- dto.Transaction) error {
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
					src := ti.SrcAddr

					if dsc, ok := tx.Description.(tlb.TransactionDescriptionOrdinary); ok && dsc.BouncePhase != nil {
						if _, ok = dsc.BouncePhase.Phase.(tlb.BouncePhaseOk); ok {
							// transaction was bounced, and coins were returned to sender
							// this can happen mostly on custom contracts
							continue
						}
					}

					if !ti.ExtraCurrencies.IsEmpty() {
						kv, err := ti.ExtraCurrencies.LoadAll()
						if err != nil {
							log.Fatalln("load extra currencies err: ", err.Error())
							return
						}

						for _, dictKV := range kv {
							currencyId := dictKV.Key.MustLoadUInt(32)
							amount := dictKV.Value.MustLoadVarUInt(32)

							log.Println("received", amount.String(), "ExtraCurrency with id", currencyId, "from", src.StringRaw())
						}
					}

					if ti.Amount.Nano().Sign() > 0 {
						// show received ton amount

						log.Println("received", ti.Amount.String(), "TON from", ti.SrcAddr.StringRaw())

						txLt = tx.LT
						internalTxChan <- dto.NewTransaction(ti.SrcAddr.StringRaw(), ti.Amount.Nano().Uint64(), txLt)
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

func (service *TonTransactionService) getTxLT(ctx context.Context, lastProcessedLT *uint64) (uint64, error) {
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
