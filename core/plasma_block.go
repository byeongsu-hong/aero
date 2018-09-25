package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/mamamerkle"
	"math/big"
)

type PlasmaBlock struct {
	Transactions []*PlasmaTx
	TxRoot       common.Hash
	Tree         mamamerkle.SparseMerkleTree

	Id big.Int
}

func NewPlasmaBlock(transactions []*PlasmaTx) *PlasmaBlock {
	return &PlasmaBlock{
		Transactions: transactions,
	}
}
