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
	aero *core.Aero

	// to synchronize exit
	exitCh chan bool
}

func New(aero *core.Aero) *Operator {
	operator := &Operator{
		aero:   aero,
		exitCh: make(chan bool, 1),
	}
	return operator
}

func (operator *Operator) Start() {
	newParentBlock := make(chan *types.Header, chainHeadChanSize)
	newChildBlock := make(chan *types.Header, chainHeadChanSize)

	newParentBlockSub, err := operator.aero.Parent.SubscribeNewHead(context.Background(), newParentBlock)
	if err != nil {
		log.Error("failed to subscribe to parent block event", err.Error())
		return
	}
	newChildBlockSub, err := operator.aero.Child.SubscribeNewHead(context.Background(), newChildBlock)
	if err != nil {
		log.Error("failed to subscribe to parent block event", err.Error())
		return
	}
	defer newParentBlockSub.Unsubscribe()
	defer newChildBlockSub.Unsubscribe()

	for {
		select {
		case head := <-newParentBlock:
			services.FetchDeposits(
				operator.aero,
				head.Number.Uint64()-(minConfirmations*2+1),
				head.Number.Uint64()-(minConfirmations-1),
			)

		case <-newChildBlock:
			services.SubmitBlock(operator.aero)

		case <-operator.exitCh:
			return
		}
	}
}

func (operator *Operator) Stop() {
	operator.exitCh <- true
}
