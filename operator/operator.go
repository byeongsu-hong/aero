package operator

import (
	"fmt"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
)

const (
	// chainHeadChanSize is the size of channel listening to ChainHeadEvent.
	chainHeadChanSize = 10
)

type Operator struct {
	Chain        *eth.Ethereum
	ParentChain  *ethclient.Client
	chainHeadCh  chan core.ChainHeadEvent
	chainHeadSub event.Subscription
	exitCh       chan bool
}

func New(chain *eth.Ethereum, parentChain *ethclient.Client) *Operator {
	operator := &Operator{
		Chain:       chain,
		ParentChain: parentChain,
		chainHeadCh: make(chan core.ChainHeadEvent, chainHeadChanSize),
		exitCh:      make(chan bool, 1),
	}
	operator.chainHeadSub = chain.BlockChain().SubscribeChainHeadEvent(operator.chainHeadCh)
	return operator
}

func (operator *Operator) Start() {
	defer operator.chainHeadSub.Unsubscribe()

	for {
		select {
		case head := <-operator.chainHeadCh:
			// a new block is created on child, so operator submits it onto the parent.

			root := head.Block.Header().Root
			fmt.Println("New created block root is : ", root.Hex())
			head.Block.Transactions().GetRlp()
		}
	}
}

func (operator *Operator) Stop() {
	operator.exitCh <- true
}
