package main

import (
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	ethutils "github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/fdlimit"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"github.com/frostornge/ethornge/account"
	"gopkg.in/urfave/cli.v1"
	"log"
	"math/big"
	"os"
)

func defaultNodeConfig() *node.Config {
	cfg := node.DefaultConfig
	cfg.Name = clientName
	cfg.Version = clientVersion
	cfg.HTTPModules = append(cfg.HTTPModules, "eth", "shh")
	cfg.WSModules = append(cfg.WSModules, "eth", "shh")
	cfg.IPCPath = "aero.ipc"
	return &cfg
}

func developerGenesisBlock(period uint64, faucet common.Address) *core.Genesis {
	config := *params.AllCliqueProtocolChanges
	config.Clique.Period = period

	genesis := &core.Genesis{
		Config:     &config,
		ExtraData:  append(append(make([]byte, 32), faucet[:]...), make([]byte, 65)...),
		GasLimit:   6283185,
		Difficulty: big.NewInt(1),
		Alloc: map[common.Address]core.GenesisAccount{
			common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
			common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
			common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
			common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
			common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
			common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
			common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
			common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
			faucet: {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
		},
	}
	return genesis
}


func loadGenesis(genesisPath string) *core.Genesis {
	if len(genesisPath) == 0 {
		log.Println("No genesis block data is given. Using default...")
		key, _ := account.NewKey()
		genesis := developerGenesisBlock(defaultBlockPeriod, key.Address)

		log.Println("Auto-generated genesis faucet address : ", key.Address.Hex())
		log.Println(" ↳ Private Key : ", hexutil.Encode(crypto.FromECDSA(key.PrivateKey)))
		return genesis
	}

	file, err := os.Open(genesisPath)
	if err != nil {
		utils.Fatalf("Failed to read genesis file: %v", err)
	}
	defer file.Close()

	genesis := new(core.Genesis)
	if err := json.NewDecoder(file).Decode(genesis); err != nil {
		utils.Fatalf("invalid genesis file: %v", err)
	}
	return genesis
}

func makeDatabaseHandles() int {
	limit, err := fdlimit.Current()
	if err != nil {
		log.Fatalf("Failed to retrieve file descriptor allowance: %v", err)
	}
	if limit < 2048 {

		if err := fdlimit.Raise(2048); err != nil {
			log.Fatalf("Failed to raise file descriptor allowance: %v", err)
		}
	}
	if limit > 2048 { // cap database file descriptors even if more is available
		limit = 2048
	}
	return limit / 2 // Leave half for networking and other stuff
}

func setupNode(ctx *cli.Context) (*node.Node, error) {
	genesisPath := ctx.Args().First()
	genesis := loadGenesis(genesisPath)

	nodeConfig := defaultNodeConfig()
	nodeConfig.P2P.ListenAddr = ":17600"
	nodeConfig.DataDir = ""

	// node access config
	nodeConfig.HTTPHost = ctx.String("host")
	nodeConfig.HTTPPort = ctx.Int("port")
	nodeConfig.WSHost = ctx.String("wshost")
	nodeConfig.WSPort = ctx.Int("wsport")

	stack, err := node.New(nodeConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to create the protocol stack: %v", err)
	}

	// unlock account manager
	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	faucet, err := ks.NewAccount("")
	if err != nil {
		return nil, fmt.Errorf("Failed to create account: %v", err)
	}
	if err := ks.Unlock(faucet, ""); err != nil {
		return nil, fmt.Errorf("Failed to unlock account: %v", err)
	}
	log.Println("Etherbase :", faucet.Address.Hex())

	// eth config
	ethConfig := &eth.DefaultConfig
	ethConfig.Genesis = genesis
	ethConfig.NetworkId = genesis.Config.ChainID.Uint64()
	ethConfig.EnablePreimageRecording = true
	ethConfig.MinerGasPrice = big.NewInt(1)
	ethConfig.DatabaseHandles = makeDatabaseHandles()

	ethutils.RegisterEthService(stack, ethConfig)
	//ethutils.RegisterShhService(stack, &whisper.DefaultConfig)

	return stack, nil
}