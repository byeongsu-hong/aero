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
	"github.com/ethereum/go-ethereum/log"
	"github.com/frostornge/ethornge/gorange"
	"github.com/frostornge/ethornge/provider"
	"github.com/frostornge/ethornge/utils"
	"github.com/stretchr/testify/assert"
)

func launchTestChain(
	t *testing.T,
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
		Period:    2,
	}
	parentCfg := gorange.Config{
		NetworkId: 3470,
		HTTPHost:  "127.0.0.1",
		HTTPPort:  3471,
		WSHost:    "127.0.0.1",
		WSPort:    3472,
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
	child *gorange.Node,
	parent *gorange.Node,
) (
	*provider.Provider,
	*provider.Provider,
) {
	ctx := context.Background()
	childPv, err := child.WsProvider(ctx)
	assert.NoError(t, err)

	parentPv, err := parent.WsProvider(ctx)
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

	child, parent := launchTestChain(t)
	defer child.Stop()
	defer parent.Stop()
	log.Info("Test Launch",
		"child", child.WSEndpoint(),
		"parent", parent.WSEndpoint(),
	)

	childPv, parentPv := getProviders(t, child, parent)
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

	aero := &core.Aero{
		Child:            childPv.Client,
		Parent:           parentPv.Client,
		ChildBridge:      childBridge,
		ChildBridgeAddr:  childAddr,
		ParentBridge:     parentBridge,
		ParentBridgeAddr: parentAddr,
		ChildOpt:         childPv.Accounts[0],
		ParentOpt:        parentPv.Accounts[0],
	}

	op := operator.New(aero)
	op.Start()
	defer op.Stop()
	log.Info("Test Start Operator")

	// does nothing at this time
	va := validator.New(aero)
	va.Start()
	defer va.Stop()
	log.Info("Test Start Validator")

	tAddr, tx, tToken, err := contracts.DeployABLToken(aero.ParentOpt, parentPv)
	assert.NoError(t, err)
	parentPv.WaitDeployedWithTimeout(tx, 1*time.Minute)

	tx, err = tToken.Approve(aero.ParentOpt, parentAddr, utils.Ether(100))
	assert.NoError(t, err)
	parentPv.WaitMinedWithTimeout(tx, 1*time.Minute)

	alt, _ := tToken.Allowance(nil, aero.ParentOpt.From, parentAddr)
	log.Info("Test Token Approved",
		"from", aero.ParentOpt.From.Hex(),
		"to", parentAddr.Hex(),
		"amount", utils.WeiToEth(alt),
	)

	for i := 1; i < 10; i++ {
		tx, err = parentBridge.DepositFungible(aero.ParentOpt, tAddr, aero.ParentOpt.From, utils.Ether(int64(i)))
		assert.NoError(t, err)
		parentPv.WaitMinedWithTimeout(tx, 1*time.Minute)
		log.Info("Test Deposit Token",
			"from", aero.ParentOpt.From.Hex(),
			"to", tAddr.Hex(),
			"amount", i,
		)
	}

	childTokenAddr, _ := childBridge.Token(nil)
	childToken, _ := contracts.NewPeggedERC721(childTokenAddr, aero.Child)
	childBalance, _ := childToken.BalanceOf(nil, aero.ParentOpt.From)
	log.Info("Test: balance",
		"owner", aero.ParentOpt.From.Hex(),
		"amount", childBalance,
	)

	time.Sleep(5 * time.Second)
}
