package validator

import (
	"context"

	"github.com/airbloc/aero/bridge/binds"
	"github.com/airbloc/aero/core"
	"github.com/ethereum/go-ethereum/log"
)

type Validator struct {
	aero   *core.Aero
	ctx    context.Context
	cancel context.CancelFunc
}

func New(aero *core.Aero) *Validator {
	ctx, cancel := context.WithCancel(context.Background())
	return &Validator{
		aero,
		ctx,
		cancel,
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
				log.Info("Exit started", "By", evt.Owner.Hex(), "Slot", evt.SlotId)
			case evt := <-exitRejected:
				log.Info("Exit rejected", "By", evt.Claimer.Hex(), "Slot", evt.SlotId)
			case evt := <-exitFinalized:
				log.Info("Exit Finalized", "By", evt.Owner.Hex(), "Slot", evt.SlotId)
			case <-v.ctx.Done():
				return
			}
		}
	}(v.aero)
}

func (v *Validator) Stop() {
	v.cancel()
}
