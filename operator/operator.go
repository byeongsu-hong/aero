package operator

import (
	"context"

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
		newParentBlock := make(chan *types.Header, chainHeadChanSize)
		newChildBlock := make(chan *types.Header, chainHeadChanSize)

		newParentBlockSub, err := aero.Parent.SubscribeNewHead(context.Background(), newParentBlock)
		if err != nil {
			log.Error("failed to subscribe to parent block event", err.Error())
			return
		}

		newChildBlockSub, err := aero.Child.SubscribeNewHead(context.Background(), newChildBlock)
		if err != nil {
			log.Error("failed to subscribe to child block event", err.Error())
			return
		}
		defer newParentBlockSub.Unsubscribe()
		defer newChildBlockSub.Unsubscribe()

		for {
			select {
			case head := <-newParentBlock:
				services.FetchDeposits(
					aero,
					head.Number.Uint64()-(minConfirmations*2+1),
					head.Number.Uint64()-(minConfirmations-1),
				)

			case <-newChildBlock:
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
