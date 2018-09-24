package services

import (
	"github.com/airbloc/aero/core"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
	"math/big"
)

func SubmitBlock(aero *core.Aero) error {
	currentBlock := &core.PlasmaBlock{}

	pendingTxnCount, err := aero.ChildBridge.GetPendingTransactionCount(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get pending transaction count from ChildChain")
	}
	if pendingTxnCount.Cmp(big.NewInt(0)) <= 0 {
		log.Trace("No pending Plasma Txn. Skipping submit.")
		return nil
	}

	for i := int64(0); i < pendingTxnCount.Int64(); i++ {
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
		currentBlock.AddTransaction(txn)
	}
	// TODO: merklize, add root
	return nil
}
