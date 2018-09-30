package test

import (
	"context"
	"os"
	"testing"

	"github.com/airbloc/aero/bridge/binds"
	"github.com/airbloc/aero/core"
	"github.com/airbloc/aero/operator"
	"github.com/airbloc/aero/validator"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/frostornge/ethornge/account"
	"github.com/frostornge/ethornge/gorange"
	"github.com/frostornge/ethornge/provider"
	"github.com/frostornge/ethornge/utils"
	"github.com/stretchr/testify/assert"
)

type contractInfo struct {
	child      *contracts.ChildBridge
	childAddr  common.Address
	parent     *contracts.ParentBridge
	parentAddr common.Address
}

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
	cAddr, _, childBridge, err := contracts.DeployChildBridge(child.Accounts[0], child, "Airbloc Plasma", "ABLP")
	assert.NoError(t, err)
	pAddr, _, parentBridge, err := contracts.DeployParentBridge(parent.Accounts[0], parent)
	assert.NoError(t, err)

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
	log.Info("launch")

	childPv, parentPv := getProviders(t, keys, child, parent)
	defer childPv.Close()
	defer parentPv.Close()
	log.Info("providers")

	childBridge, childAddr, parentBridge, parentAddr := deployContracts(t, childPv, parentPv)
	log.Info("contracts")

	aero, err := core.NewAero(
		"ws://"+child.WSEndpoint(),
		"ws://"+parent.WSEndpoint(),
		childBridge, childAddr,
		parentBridge, parentAddr,
		keys[0].PrivateKey,
	)
	assert.NoError(t, err)
	log.Info("aero")

	op := operator.New(aero)
	op.Start()
	defer op.Stop()
	log.Info("operator")

	// does nothing at this time
	va := validator.New(aero)
	va.Start()
	defer va.Stop()
	log.Info("validator")

	tAddr, _, tToken, err := contracts.DeployABLToken(childPv.Accounts[0], childPv)
	assert.NoError(t, err)

	tToken.Approve()

	tx, err := parentBridge.DepositFungible(
		childPv.Accounts[0],
		tAddr,
		childPv.Accounts[0].From,
		utils.Ether(100),
	)
	assert.NoError(t, err)
}
