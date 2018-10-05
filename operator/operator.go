package operator

import (
	"context"

	"fmt"

	"github.com/airbloc/aero/bridge/binds"
	"github.com/airbloc/aero/core"
	"github.com/airbloc/aero/services"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

const (
	blockInterval = 3
)

type Operator struct {
	aero   *core.Aero
	ctx    context.Context
	cancel context.CancelFunc
}

func New(aero *core.Aero) *Operator {
	ctx, cancel := context.WithCancel(context.Background())
	operator := &Operator{
		aero:   aero,
		ctx:    ctx,
		cancel: cancel,
	}
	return operator
}

func (operator *Operator) Start() {
	go func(aero *core.Aero) {
		exit := make(chan *contracts.ParentBridgeExitFinalized)
		deposit := make(chan *contracts.ParentBridgeDeposit)
		newChildBlock := make(chan *types.Header)
		defer close(exit)
		defer close(deposit)
		defer close(newChildBlock)

		depositSub, err := aero.ParentBridge.WatchDeposit(nil, deposit, nil, nil)
		if err != nil {
			log.Error("Operator: failed to watch deposit event", "err", err.Error())
			return
		}
		defer depositSub.Unsubscribe()

		exitSub, err := aero.ParentBridge.WatchExitFinalized(nil, exit, nil, nil)
		if err != nil {
			log.Error("Operator: failed to watch exit-finalized event", "err", err.Error())
			return
		}
		defer exitSub.Unsubscribe()

		newChildBlockSub, err := aero.Child.SubscribeNewHead(context.Background(), newChildBlock)
		if err != nil {
			log.Error("failed to subscribe to child block event", "err", err.Error())
			return
		}
		defer newChildBlockSub.Unsubscribe()

		events := make(chan []interface{})
		queue := services.NewQueue(events)
		queue.Run(aero.Parent.Client)
		defer queue.Close()

		for {
			// parent chain
			select {
			case evt := <-deposit:
				log.Info("Operator: deposit",
					"owner", evt.Owner.Hex(),
					"slotId", evt.SlotId,
				)
				queue.Push(evt)
			case evt := <-exit:
				log.Info("Operator: exit",
					"owner", evt.Owner.Hex(),
					"slotId", evt.SlotId,
					"token", fmt.Sprintln(evt.Token.Hex(), "(Type = ", evt.Typ, ")"))
				queue.Push(evt)
			case task := <-events:
				submitParentEvent(aero, task)
			default:
			}

			// child chain
			select {
			case <-newChildBlock:
				log.Info("Operator: new child block")
				submitChildEvent(aero)
			default:
			}

			// done
			select {
			case <-operator.ctx.Done():
				log.Info("Operator: closing")
				return
			default:
			}
		}
	}(operator.aero)
}

func (operator *Operator) Stop() {
	operator.cancel()
}
