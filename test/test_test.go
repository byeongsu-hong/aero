package test

import (
	"context"
	"os"
	"testing"

	"time"

	"github.com/airbloc/aero/bridge/binds"
	"github.com/airbloc/aero/core"
	"github.com/airbloc/aero/operator"
	"github.com/airbloc/aero/validator"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/frostornge/ethornge/account"
	"github.com/frostornge/ethornge/gorange"
	"github.com/frostornge/ethornge/provider"
	"github.com/frostornge/ethornge/utils"
	"github.com/stretchr/testify/assert"
)

func launchTestChain(
	t *testing.T,
	keys account.Keys,
) (
	*gorange.Node,
	*gorange.Node,
) {
	childCfg := gorange.Config{
		NetworkId: 2470,
		HTTPHost:  "127.0.0.1",
		HTTPPort:  2471,
		WSHost:    "127.0.0.1",
		WSPort:    2472,
		Accounts:  keys.GetAddresses(),
		Balances:  int64(100),
		Period:    2,
	}
	parentCfg := gorange.Config{
		NetworkId: 3470,
		HTTPHost:  "127.0.0.1",
		HTTPPort:  3471,
		WSHost:    "127.0.0.1",
		WSPort:    3472,
		Accounts:  keys.GetAddresses(),
		Balances:  int64(100),
		Period:    6,
	}

	child, err := gorange.Launch(childCfg)
	assert.NoError(t, err)

	parent, err := gorange.Launch(parentCfg)
	assert.NoError(t, err)

	return child, parent
}

func getProviders(
	t *testing.T,
	keys account.Keys,
	child *gorange.Node,
	parent *gorange.Node,
) (
	*provider.Provider,
	*provider.Provider,
) {
	ctx := context.Background()
	childPv, err := child.WsProviderWithAccounts(ctx, keys)
	assert.NoError(t, err)

	parentPv, err := parent.WsProviderWithAccounts(ctx, keys)
	assert.NoError(t, err)

	return childPv, parentPv
}

func deployContracts(
	t *testing.T,
	child *provider.Provider,
	parent *provider.Provider,
) (
	*contracts.ChildBridge,
	common.Address,
	*contracts.ParentBridge,
	common.Address,
) {
	cAddr, cTx, childBridge, err := contracts.DeployChildBridge(child.Accounts[0], child, "Airbloc Plasma", "ABLP")
	assert.NoError(t, err)
	pAddr, pTx, parentBridge, err := contracts.DeployParentBridge(parent.Accounts[0], parent)
	assert.NoError(t, err)

	child.WaitDeployedWithTimeout(cTx, 1*time.Minute)
	parent.WaitDeployedWithTimeout(pTx, 1*time.Minute)

	return childBridge, cAddr, parentBridge, pAddr
}

func TestAero(t *testing.T) {
	// setup logger
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stderr, log.TerminalFormat(true)))
	glogger.Verbosity(log.Lvl(3))
	log.Root().SetHandler(glogger)

	var keys account.Keys
	for i := 0; i < 10; i++ {
		key, _ := account.NewKey()
		keys = append(keys, key)
	}

	child, parent := launchTestChain(t, keys)
	defer child.Stop()
	defer parent.Stop()
	log.Info("Test Launch",
		"child", child.WSEndpoint(),
		"parent", parent.WSEndpoint(),
	)

	childPv, parentPv := getProviders(t, keys, child, parent)
	defer childPv.Close()
	defer parentPv.Close()
	log.Info("Test Providers",
		"child owner", childPv.Accounts[0].From.Hex(),
		"parent owner", parentPv.Accounts[0].From.Hex(),
	)

	childBridge, childAddr, parentBridge, parentAddr := deployContracts(t, childPv, parentPv)
	log.Info("Test Contracts",
		"child bridge", childAddr.Hex(),
		"parent bridge", parentAddr.Hex(),
	)

	aero, err := core.NewAero(
		"ws://"+child.WSEndpoint(),
		"ws://"+parent.WSEndpoint(),
		childBridge, childAddr,
		parentBridge, parentAddr,
		keys[0].PrivateKey,
	)
	assert.NoError(t, err)
	log.Info("Test Aero",
		"privKey", "0x"+common.Bytes2Hex(crypto.FromECDSA(keys[0].PrivateKey)))

	op := operator.New(aero)
	op.Start()
	defer op.Stop()
	log.Info("Test Start Operator")

	// does nothing at this time
	va := validator.New(aero)
	va.Start()
	defer va.Stop()
	log.Info("Test Start Validator")

	pOwner := parentPv.Accounts[0]
	//pParty := parentPv.Accounts[1:]
	//cOwner := childPv.Accounts[0]
	//cParty := childPv.Accounts[1:]

	tAddr, tx, tToken, err := contracts.DeployABLToken(pOwner, parentPv)
	assert.NoError(t, err)
	parentPv.WaitDeployedWithTimeout(tx, 1*time.Minute)

	tx, err = tToken.Approve(pOwner, parentAddr, utils.Ether(100))
	assert.NoError(t, err)
	parentPv.WaitMinedWithTimeout(tx, 1*time.Minute)

	alt, _ := tToken.Allowance(nil, pOwner.From, parentAddr)
	log.Info("Test Token Approved",
		"from", pOwner.From.Hex(),
		"to", parentAddr.Hex(),
		"amount", utils.WeiToEth(alt),
	)

	tx, err = parentBridge.DepositFungible(pOwner, tAddr, pOwner.From, utils.Ether(100))
	assert.NoError(t, err)
	parentPv.WaitMinedWithTimeout(tx, 1*time.Minute)
	log.Info("Test Deposit Token",
		"from", pOwner.From.Hex(),
		"to", tAddr.Hex(),
		"amount", 100,
	)

	time.Sleep(1 * time.Minute)
}
