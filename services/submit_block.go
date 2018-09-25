package services

import (
	"github.com/airbloc/aero/core"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
	"math/big"
)

func SubmitBlock(aero *core.Aero) error {

	pendingTxnCount, err := aero.ChildBridge.GetPendingTransactionCount(nil)
	if err != nil {
		return errors.Wrap(err, "Failed to get pending transaction count from ChildChain")
	}
	if pendingTxnCount.Cmp(big.NewInt(0)) <= 0 {
		log.Trace("No pending Plasma Txn. Skipping submit.")
		return nil
	}

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

	block := core.NewPlasmaBlock(transactions)

	return nil
}
