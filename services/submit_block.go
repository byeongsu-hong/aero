package services

import (
	"github.com/airbloc/aero/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"time"
)

func SubmitBlock(aero *core.Aero) error {
	pendingTxnCount, err := aero.ChildBridge.GetPendingTransactionCount(nil)
	if err != nil {
		log.Error("Failed to get pending transaction count from ChildChain", "err", err)
		return err
	}
	if pendingTxnCount.Cmp(big.NewInt(0)) <= 0 {
		log.Trace("No pending Plasma Txn. Skipping submit.")
		return nil
	}

	// gather plasma transaction data
	transactions := make([]*core.PlasmaTx, pendingTxnCount.Int64())
	for i := int64(0); i < pendingTxnCount.Int64(); i++ {
		// TODO: big.Int
		txnHash, err := aero.ChildBridge.PendingTransactions(nil, big.NewInt(i))
		if err != nil {
			log.Error("Failed to get Plasma Txn hash", "err", err)
			return err
		}

		plasmaTxOnchain, err := aero.ChildBridge.Transactions(nil, txnHash)
		if err != nil {
			log.Error("Failed to get Plasma Txn data", "txnHash", txnHash, "err", err)
			return err
		}

		txn := core.NewPlasmaTxFromOnChain(txnHash, plasmaTxOnchain)
		transactions[i] = txn
	}

	// create new plasma block
	block, err := core.NewPlasmaBlock(transactions)
	if err != nil {
		log.Error("Failed to create Plasma Block", "err", err)
		return err
	}

	// submit block to root chain
	_, err = submitToRoot(aero, block)
	if err != nil {
		log.Error("Failed to submit block to the parent chain", "root", block.TxRoot.Hex(), "err", err)
		return err
	}
	return nil
}

func submitToRoot(aero *core.Aero, block *core.PlasmaBlock) (receipt *types.Receipt, err error) {
	tx, err := aero.ParentBridge.SubmitBlock(nil, block.TxRoot)
	if err != nil {
		return
	}
	receipt, err = aero.WaitParentTxToBeMined(tx, 1*time.Minute)
	if err != nil {
		return
	}
	return
}
