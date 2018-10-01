package validator

import (
	"context"

	"math/big"

	"time"

	"github.com/airbloc/aero/bridge/binds"
	"github.com/airbloc/aero/core"
	"github.com/ethereum/go-ethereum/log"
)

// TODO: manage confirmation by event queue
const (
	MinConfirm = 6
)

type Validator struct {
	aero *core.Aero

	// TODO: manage exits (finalize)
	exits []uint64

	ctx    context.Context
	cancel context.CancelFunc
}

func New(aero *core.Aero) *Validator {
	ctx, cancel := context.WithCancel(context.Background())
	return &Validator{
		aero:   aero,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (v *Validator) Start() {
	go func(aero *core.Aero) {
		exitStarted := make(chan *contracts.ParentBridgeExitStarted)
		exitRejected := make(chan *contracts.ParentBridgeExitRejected)
		exitFinalized := make(chan *contracts.ParentBridgeExitFinalized)
		defer close(exitStarted)
		defer close(exitRejected)
		defer close(exitFinalized)

		exitStartedSub, err := aero.ParentBridge.WatchExitStarted(nil, exitStarted, nil, nil)
		if err != nil {
			log.Error("failed to subscribe to exit start event", err.Error())
			return
		}
		defer exitStartedSub.Unsubscribe()

		exitRejectedSub, err := aero.ParentBridge.WatchExitRejected(nil, exitRejected, nil, nil)
		if err != nil {
			log.Error("failed to subscribe to exit reject event", err.Error())
			return
		}
		defer exitRejectedSub.Unsubscribe()

		exitFinalizedSub, err := aero.ParentBridge.WatchExitFinalized(nil, exitFinalized, nil, nil)
		if err != nil {
			log.Error("failed to subscribe to exit finalize event", err.Error())
			return
		}
		defer exitFinalizedSub.Unsubscribe()

		for {
			select {
			case evt := <-exitStarted:
				v.exits = append(v.exits, evt.SlotId)
				log.Info("Validator: Exit started",
					"By", evt.Owner.Hex(),
					"Slot", evt.SlotId,
				)
			case evt := <-exitRejected:
				for index, exit := range v.exits {
					if evt.SlotId == exit {
						v.exits = append(v.exits[:index], v.exits[index+1:]...)
						break
					}
				}
				log.Info("Validator: Exit rejected",
					"By", evt.Claimer.Hex(),
					"Slot", evt.SlotId,
				)
			case evt := <-exitFinalized:
				for index, exit := range v.exits {
					if evt.SlotId == exit {
						v.exits = append(v.exits[:index], v.exits[index+1:]...)
						break
					}
				}

				tx, err := aero.ChildBridge.SubmitWithdraw(aero.ChildOpt, evt.Owner, evt.Token, evt.SlotId, big.NewInt(0), evt.Typ)
				if err != nil {
					log.Error("Validator:", "create tx error", err)
				}
				if _, err = aero.WaitChildTxMined(tx, 1*time.Minute); err != nil {
					log.Error("Validator:", "deploy tx error", err)
				}

				log.Info("Validator: Exit Finalized",
					"By", evt.Owner.Hex(),
					"Slot", evt.SlotId,
				)
			case <-v.ctx.Done():
				log.Info("Validator: Stopping...")
				return
			}
		}
	}(v.aero)
}

func (v *Validator) Stop() {
	v.cancel()
}
