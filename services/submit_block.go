package services

import (
	"github.com/airbloc/aero/core"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
	"math/big"
	"time"
)

func SubmitBlock(aero *core.Aero) (*core.PlasmaBlock, error) {
	pendingTxnCount, err := aero.ChildBridge.GetPendingTransactionCount(nil)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get pending transaction count from ChildChain")
	}
	if pendingTxnCount.Cmp(big.NewInt(0)) <= 0 {
		log.Trace("No pending Plasma Txn. Skipping submit.")
		return nil, nil
	}

	// gather plasma transaction data
	transactions := make([]*core.PlasmaTx, pendingTxnCount.Int64())
	for i := int64(0); i < pendingTxnCount.Int64(); i++ {
		// TODO: big.Int
		txnHash, err := aero.ChildBridge.PendingTransactions(nil, big.NewInt(i))
		if err != nil {
			log.Error("Failed to get pending transaction count from ChildChain")
			continue
		}

		plasmaTxOnchain, err := aero.ChildBridge.Transactions(nil, txnHash)
		if err != nil {
			log.Error("Failed to get", i, "th transaction from ChildChain")
			continue
		}

		txn := core.NewPlasmaTxFromOnChain(txnHash, plasmaTxOnchain)
		transactions[i] = txn
	}

	// create new plasma block
	block, err := core.NewPlasmaBlock(transactions)
	if err != nil {
		return nil, err
	}

	// submit block to root chain
	tx, err := aero.ParentBridge.SubmitBlock(nil, block.TxRoot)
	if err != nil {
		return nil, errors.Wrap(err, "failed to submit block to root")
	}
	_, err = aero.WaitParentTxToBeMined(tx, 1*time.Minute)
	if err != nil {
		return nil, err
	}
	return block, nil
}
