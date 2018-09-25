package core

import (
	"crypto/ecdsa"
	"github.com/airbloc/aero/contracts/binds"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

// Aero is the context of the Aero framework.
type Aero struct {
	// connection to both chain
	Child  *ethclient.Client
	Parent *ethclient.Client

	// binding of Bridge Contracts
	ChildBridge  *contracts.ChildChain
	ParentBridge *contracts.ParentChain

	// private key of the account
	PrivateKey *ecdsa.PrivateKey
}

// NewAero connects to given blockchains
// and creates a context of the Aero framework.
func NewAero(
	childUri string,
	parentUri string,
	childBridgeAddress common.Address,
	parentBridgeAddress common.Address,
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

	// setup contract bind
	childBridge, err := contracts.NewChildChain(childBridgeAddress, child)
	if err != nil {
		return nil, errors.Wrap(err, "failed to locate ChildChain bridge contract")
	}
	parentBridge, err := contracts.NewParentChain(parentBridgeAddress, parent)
	if err != nil {
		return nil, errors.Wrap(err, "failed to locate ParentChain bridge contract")
	}

	return &Aero{
		Child:        child,
		Parent:       parent,
		ChildBridge:  childBridge,
		ParentBridge: parentBridge,
		PrivateKey:   privateKey,
	}, nil
}
