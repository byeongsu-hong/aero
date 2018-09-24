package core

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type PlasmaTx struct {
	TokenId     *big.Int        `json:"tokenId"`
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
	TokenId   *big.Int
	PrevBlock *big.Int
	NewOwner  common.Address
	Owner     common.Address
	Signature []byte
}

func NewPlasmaTxFromOnChain(hash common.Hash, onchainTx PlasmaTxOnchain) *PlasmaTx {
	tx := &PlasmaTx{
		TokenId:     onchainTx.TokenId,
		PrevBlockId: onchainTx.PrevBlock,
		NewOwner:    &onchainTx.NewOwner,
		Hash:        &hash,
	}
	return tx
}
