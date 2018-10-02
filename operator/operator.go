package operator

import (
	"context"

	"math/big"

	"time"

	"github.com/airbloc/aero/bridge/binds"
	"github.com/airbloc/aero/core"
	"github.com/airbloc/aero/services"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

const (
	// chainHeadChanSize is the size of channel listening to ChainHeadEvent.
	chainHeadChanSize = 10
	minConfirmations  = 3
)

type Operator struct {
	aero   *core.Aero
	ctx    context.Context
	cancel context.CancelFunc
}

func New(aero *core.Aero) *Operator {
	ctx, cancel := context.WithCancel(context.Background())
	operator := &Operator{
		aero:   aero,
		ctx:    ctx,
		cancel: cancel,
	}
	return operator
}

func (operator *Operator) Start() {
	go func(aero *core.Aero) {

		//newParentBlock := make(chan *types.Header, chainHeadChanSize)
		//newParentBlockSub, err := aero.Parent.SubscribeNewHead(context.Background(), newParentBlock)
		//if err != nil {
		//	log.Error("failed to subscribe to parent block event", err.Error())
		//	return
		//}
		//defer newParentBlockSub.Unsubscribe()

		deposit := make(chan *contracts.ParentBridgeDeposit)
		newChildBlock := make(chan *types.Header, chainHeadChanSize)

		depositSub, err := aero.ParentBridge.WatchDeposit(nil, deposit, nil, nil)
		if err != nil {
			log.Error("Operator: failed to watch deposit event", "err", err.Error())
		}
		defer depositSub.Unsubscribe()

		newChildBlockSub, err := aero.Child.SubscribeNewHead(context.Background(), newChildBlock)
		if err != nil {
			log.Error("failed to subscribe to child block event", "err", err.Error())
			return
		}
		defer newChildBlockSub.Unsubscribe()

		for {
			select {
			//case head := <-newParentBlock:
			//	log.Info("Operator: new parent block",
			//		"number", head.Number,
			//		"txRoot", head.Root.Hex(),
			//	)
			//	services.FetchDeposits(
			//		aero,
			//		head.Number.Uint64()-(minConfirmations*2+1),
			//		head.Number.Uint64()-(minConfirmations-1),
			//	)
			case evt := <-deposit:
				log.Info("Operator: deposit",
					"owner", evt.Owner.Hex(),
					"slotId", evt.SlotId,
				)

				tx, err := aero.ChildBridge.SubmitDeposit(aero.ChildOpt, evt.Owner, evt.Token, evt.SlotId, big.NewInt(0), 1)
				if err != nil {
					log.Error("Operator:", "create tx error", err)
				}
				if _, err = aero.WaitChildTxMined(tx, 1*time.Minute); err != nil {
					log.Error("Operator:", "deploy tx error", err)
				}
			case <-newChildBlock:
				log.Info("Operator: new child block")
				services.SubmitBlock(aero)

			case <-operator.ctx.Done():
				return
			}
		}
	}(operator.aero)
}

func (operator *Operator) Stop() {
	operator.cancel()
}
