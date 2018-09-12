package main

import (
	"github.com/airbloc/ethornge/account"
	"github.com/airbloc/ethornge/gorange"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	log.Println("Hello byeongsu")
	key, _ := account.NewKey()
	println(string(crypto.FromECDSA(key.PrivateKey)))
	node, _ := gorange.Launch(
		gorange.DefaultLocalConfig([]common.Address{key.Address}, 100),
	)
	defer node.Stop()
}