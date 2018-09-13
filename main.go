package main

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/frostornge/ethornge/account"
	"github.com/frostornge/ethornge/gorange"
)

func main() {
	key, _ := account.NewKey()
	node, err := gorange.Launch(
		gorange.DefaultLocalConfig([]common.Address{key.Address}, 100))
	if err != nil {
		panic(err)
	}
	defer node.Stop()
}