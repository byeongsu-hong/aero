package operator

import (
	"time"

	"github.com/airbloc/aero/bridge/binds"
	"github.com/airbloc/aero/core"
)

func submitParentEvent(aero *core.Aero, evts []interface{}) error {
	for _, evt := range evts {
		switch v := evt.(type) {
		case *contracts.ParentBridgeDeposit:
			if err := deposit(aero, v); err != nil {
				return err
			}
		case *contracts.ParentBridgeExitFinalized:
			if err := withdraw(aero, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func deposit(aero *core.Aero, evt *contracts.ParentBridgeDeposit) error {
	tx, err := aero.ChildBridge.SubmitDeposit(aero.Child.Accounts[0], evt.Owner, evt.Token, evt.SlotId, evt.Amount, evt.Typ)
	if err != nil {
		return err
	}
	_, err = aero.Child.WaitMinedWithTimeout(tx, 5*time.Minute)
	return err
}

func withdraw(aero *core.Aero, evt *contracts.ParentBridgeExitFinalized) error {
	tx, err := aero.ChildBridge.SubmitWithdraw(aero.Child.Accounts[0], evt.Owner, evt.Token, evt.SlotId, evt.Amount, evt.Typ)
	if err != nil {
		return err
	}
	_, err = aero.Child.WaitMinedWithTimeout(tx, 5*time.Minute)
	return err
}
