package services

import (
	"github.com/airbloc/aero/core"
	"github.com/airbloc/aero/ethcontract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/log"
	"github.com/pkg/errors"
)

func FetchDeposits(aero *core.Aero, startBlock uint64, endBlock uint64) error {
	filterOpts := &bind.FilterOpts{
		Start: startBlock,
		End:   &endBlock,
	}

	it, err := aero.ParentBridge.FilterDeposit(filterOpts, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to filter deposit log from ParentChain")
	}

	for {
		ok := it.Next()
		if ok {
			// TODO: queued? batch deposit
			if err := reportDeposit(aero, it.Event); err != nil {
				return errors.Wrapf(err, "failed to report deposit : %v", *it.Event)
			}
		} else {
			if err := it.Error(); err != nil {
				return errors.Wrap(err, "failed to get event data of deposit from ParentChain")
			}
			it.Close()
			break
		}
	}
	return nil
}

func reportDeposit(aero *core.Aero, event *ethcontract.ParentChainDeposit) error {
	auth := bind.NewKeyedTransactor(aero.PrivateKey)
	tx, err := aero.ChildBridge.SubmitDeposit(auth, event.Depositor, event.Token, event.Amount, 0)
	if err != nil {
		return err
	}
	log.Debug("submitDeposit(%s, %d) -> %s", event.Depositor.Hex(), event.Amount.Int64(), tx.Hash().Hex())
	return nil
}
