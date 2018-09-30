package core

import (
	"context"
	"crypto/ecdsa"
	"time"

	"github.com/airbloc/aero/bridge/binds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// Aero is the context of the Aero framework.
type Aero struct {
	// connection to both chain
	Child  *ethclient.Client
	Parent *ethclient.Client

	// binding of Bridge Contracts
	ChildBridge      *contracts.ChildBridge
	ChildBridgeAddr  common.Address
	ParentBridge     *contracts.ParentBridge
	ParentBridgeAddr common.Address

	// private key of the account
	PrivateKey *ecdsa.PrivateKey
}

// NewAero connects to given blockchains
// and creates a context of the Aero framework.
func NewAero(
	childUri string,
	parentUri string,
	childBridge *contracts.ChildBridge,
	childBridgeAddr common.Address,
	parentBridge *contracts.ParentBridge,
	parentBridgeAddr common.Address,
	privateKey *ecdsa.PrivateKey,
) (*Aero, error) {
	// connect to blockchains
	child, err := ethclient.Dial(childUri)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to child chain")
	}
	parent, err := ethclient.Dial(parentUri)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to parent chain")
	}

	return &Aero{
		Child:            child,
		Parent:           parent,
		ChildBridge:      childBridge,
		ChildBridgeAddr:  childBridgeAddr,
		ParentBridge:     parentBridge,
		ParentBridgeAddr: parentBridgeAddr,
		PrivateKey:       privateKey,
	}, nil
}

func (aero *Aero) WaitParentTxToBeMined(tx *types.Transaction, timeout time.Duration) (*types.Receipt, error) {
	timeoutContext, cancelTimeout := context.WithTimeout(context.Background(), timeout)
	defer cancelTimeout()
	return bind.WaitMined(timeoutContext, aero.Parent, tx)
}
