package validator

import (
	"context"

	"github.com/airbloc/aero/core"
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

}

func (v *Validator) Stop() {
	v.cancel()
}
