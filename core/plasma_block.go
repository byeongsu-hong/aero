package core

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type PlasmaBlock struct {
	Transactions []*PlasmaTx

	Id big.Int

	// merkle root of the block
	Root common.Hash
}

func (block *PlasmaBlock) AddTransaction(tx *PlasmaTx) {
	block.Transactions = append(block.Transactions, tx)
}
