package core

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type PlasmaTx struct {
	SlotId      uint64          `json:"slotId"`
	PrevBlockId *big.Int        `json:"prevBlock"`
	NewOwner    *common.Address `json:"owner"`

	// Plasma Debit
	Value     *big.Int `json:"v"`
	Available *big.Int `json:"a"`

	// This is only used when marshaling to JSON.
	Hash *common.Hash `json:"hash" rlp:"-"`
}

type PlasmaTxOnchain struct {
	Status    uint8
	SlotId    uint64
	PrevBlock *big.Int
	NewOwner  common.Address
	Owner     common.Address
	Signature []byte
}

func NewPlasmaTxFromOnChain(hash common.Hash, onchainTx PlasmaTxOnchain) *PlasmaTx {
	tx := &PlasmaTx{
		SlotId:      onchainTx.SlotId,
		PrevBlockId: onchainTx.PrevBlock,
		NewOwner:    &onchainTx.NewOwner,
		Hash:        &hash,
	}
	return tx
}
