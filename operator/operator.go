package operator

import (
	"github.com/airbloc/aero/core"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
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

	// a block that'll be committed to the parent chain soon
	currentBlock := &core.PlasmaBlock{}

	for {
		select {
		case <-operator.newBlockCh:
			// a new block is created on child, so operator submits it onto the parent.

			pendingTxnCount, err := operator.aero.ChildBridge.GetPendingTransactionCount(nil)
			if err != nil {
				log.Error("Failed to get pending transaction count from ChildChain")
				continue
			}
			if pendingTxnCount.Cmp(big.NewInt(0)) <= 0 {
				log.Trace("No pending Plasma Txn. Skipping submit.")
				continue
			}

			for i := int64(0); i < pendingTxnCount.Int64(); i++ {
				txnHash, err := operator.aero.ChildBridge.PendingTransactions(nil, big.NewInt(i))
				if err != nil {
					log.Error("Failed to get pending transaction count from ChildChain")
					continue
				}

				plasmaTxOnchain, err := operator.aero.ChildBridge.Transactions(nil, txnHash)
				if err != nil {
					log.Error("Failed to get", i, "th transaction from ChildChain")
					continue
				}

				txn := core.NewPlasmaTxFromOnChain(txnHash, plasmaTxOnchain)
				currentBlock.AddTransaction(txn)
			}

			// commit to the parent chain

		case <-operator.exitCh:
			return
		}
	}
}

func (operator *Operator) Stop() {
	operator.exitCh <- true
}
