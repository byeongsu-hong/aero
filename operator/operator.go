package operator

import (
	"context"
	"crypto/ecdsa"
	"github.com/airbloc/aero/core"
	"github.com/airbloc/aero/ethcontract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
	"math/big"
)

const (
	// chainHeadChanSize is the size of channel listening to ChainHeadEvent.
	chainHeadChanSize = 10
)

type Operator struct {
	// Child chain client
	Child *ethclient.Client
	// Parent chain (Ethereum) client
	Parent *ethclient.Client
	// Private key of the operator account on the parent Chain
	PrivateKey *ecdsa.PrivateKey

	// Binding of Bridge Contracts
	ChildChain  *ethcontract.ChildChain
	ParentChain *ethcontract.ParentChain

	// subscription to the block generation of the child
	newBlockCh  chan *types.Header
	newBlockSub ethereum.Subscription

	// to synchronize exit
	exitCh chan bool
}

func New(
	childChain *ethclient.Client,
	parentChain *ethclient.Client,
	childBridgeAddr common.Address,
	parentBridgeAddr common.Address,
) (*Operator, error) {
	var err error
	operator := &Operator{
		Child:      childChain,
		Parent:     parentChain,
		newBlockCh: make(chan *types.Header, chainHeadChanSize),
		exitCh:     make(chan bool, 1),
	}

	operator.newBlockSub, err = childChain.SubscribeNewHead(context.Background(), operator.newBlockCh)
	if err != nil {
		return nil, errors.Wrap(err, "failed to subscribe child chain event")
	}

	// setup contract bind
	operator.ParentChain, err = ethcontract.NewParentChain(parentBridgeAddr, parentChain)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup ParentChain contract")
	}
	operator.ChildChain, err = ethcontract.NewChildChain(childBridgeAddr, childChain)
	if err != nil {
		return nil, errors.Wrap(err, "failed to setup ChildChain contract")
	}

	return operator, nil
}

func (operator *Operator) Start() {
	defer operator.newBlockSub.Unsubscribe()

	// a block that'll be committed to the parent chain soon
	currentBlock := &core.PlasmaBlock{}

	for {
		select {
		case <-operator.newBlockCh:
			// a new block is created on child, so operator submits it onto the parent.

			pendingTxnCount, err := operator.ChildChain.GetPendingTransactionCount(nil)
			if err != nil {
				log.Error("Failed to get pending transaction count from ChildChain")
				continue
			}
			if pendingTxnCount.Cmp(big.NewInt(0)) <= 0 {
				log.Trace("No pending Plasma Txn. Skipping submit.")
				continue
			}

			for i := int64(0); i < pendingTxnCount.Int64(); i++ {
				txnHash, err := operator.ChildChain.PendingTransactions(nil, big.NewInt(i))
				if err != nil {
					log.Error("Failed to get pending transaction count from ChildChain")
					continue
				}

				plasmaTxOnchain, err := operator.ChildChain.Transactions(nil, txnHash)
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
