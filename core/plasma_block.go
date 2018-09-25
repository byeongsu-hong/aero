package core

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/loomnetwork/mamamerkle"
	"github.com/pkg/errors"
	"math/big"
)

type PlasmaBlock struct {
	Transactions []*PlasmaTx
	TxRoot       common.Hash
	TxTree       *mamamerkle.SparseMerkleTree

	Id big.Int
}

func NewPlasmaBlock(transactions []*PlasmaTx) (*PlasmaBlock, error) {
	// build sparse merkle tree and calculate root
	leaves := make(map[uint64][]byte)
	for _, tx := range transactions {
		leaves[tx.SlotId] = tx.Hash.Bytes()
	}

	tree, err := mamamerkle.NewSparseMerkleTree(64, leaves)
	if err != nil {
		return nil, errors.Wrap(err, "failed to construct SMT")
	}
	return &PlasmaBlock{
		Transactions: transactions,
		TxRoot:       common.BytesToHash(tree.Root()),
		TxTree:       tree,
	}, nil
}
