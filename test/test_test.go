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

func setup(t *testing.T, child *gorange.Node, parent *gorange.Node) *core.Aero {
	childPv, parentPv := getProviders(t, child, parent)
	log.Info("Test Providers",
		"child owner", childPv.Accounts[0].From.Hex(),
		"parent owner", parentPv.Accounts[0].From.Hex(),
	)

	childBridge, childAddr, parentBridge, parentAddr := deployContracts(t, childPv, parentPv)
	log.Info("Test Contracts",
		"child bridge", childAddr.Hex(),
		"parent bridge", parentAddr.Hex(),
	)

	return &core.Aero{
		Child:            childPv,
		Parent:           parentPv,
		ChildBridge:      childBridge,
		ChildBridgeAddr:  childAddr,
		ParentBridge:     parentBridge,
		ParentBridgeAddr: parentAddr,
	}
}

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

	aero := setup(t, child, parent)
	defer aero.Close()

	op := operator.New(aero)
	op.Start()
	defer op.Stop()
	log.Info("Test Start Operator")

	// does nothing at this time
	va := validator.New(aero)
	va.Start()
	defer va.Stop()
	log.Info("Test Start Validator")

	parentOwner := aero.Parent.Accounts[0]
	//childOwner := aero.Child.Accounts[0]

	tAddr, tx, tToken, err := contracts.DeployABLToken(parentOwner, aero.Parent)
	assert.NoError(t, err)
	aero.WaitParentTxMined(tx, 1*time.Minute)

	tx, err = tToken.Approve(parentOwner, aero.ParentBridgeAddr, utils.Ether(100))
	assert.NoError(t, err)
	aero.WaitParentTxMined(tx, 1*time.Minute)

	alt, _ := tToken.Allowance(nil, parentOwner.From, aero.ParentBridgeAddr)
	log.Info("Test Token Approved",
		"from", parentOwner.From.Hex(),
		"to", aero.ParentBridgeAddr.Hex(),
		"amount", utils.WeiToEth(alt),
	)

	for i := 1; i < 3; i++ {
		tx, err = aero.ParentBridge.DepositFungible(parentOwner, tAddr, aero.Parent.Accounts[0].From, utils.Ether(int64(i)))
		assert.NoError(t, err)
		aero.WaitParentTxMined(tx, 1*time.Minute)
		log.Info("Test Deposit Token",
			"from", parentOwner.From.Hex(),
			"to", tAddr.Hex(),
			"amount", i,
		)
	}

	childTokenAddr, _ := aero.ChildBridge.Token(nil)
	childToken, _ := contracts.NewPeggedERC721(childTokenAddr, aero.Child)
	childBalance, _ := childToken.BalanceOf(nil, parentOwner.From)
	log.Info("Test: balance",
		"owner", parentOwner.From.Hex(),
		"amount", childBalance,
	)

	time.Sleep(5 * time.Second)
}
