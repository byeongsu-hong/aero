package operator

import (
	"github.com/airbloc/aero/core"
	"github.com/airbloc/aero/services"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	// chainHeadChanSize is the size of channel listening to ChainHeadEvent.
	chainHeadChanSize = 10
)

type Operator struct {
	aero *core.Aero

	// subscription to the block generation of the child
	newBlockCh  chan *types.Header
	newBlockSub ethereum.Subscription

	// to synchronize exit
	exitCh chan bool
}

func New(aero *core.Aero) *Operator {
	operator := &Operator{
		aero:       aero,
		newBlockCh: make(chan *types.Header, chainHeadChanSize),
		exitCh:     make(chan bool, 1),
	}
	return operator
}

func (operator *Operator) Start() {
	defer operator.newBlockSub.Unsubscribe()
	for {
		select {
		case <-operator.newBlockCh:
			services.SubmitBlock(operator.aero)

		case <-operator.exitCh:
			return
		}
	}
}

func (operator *Operator) Stop() {
	operator.exitCh <- true
}
