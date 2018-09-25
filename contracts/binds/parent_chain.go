// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ParentChainABI is the input ABI used to generate the binding from.
const ParentChainABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"coinCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"challenges\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"challenger\",\"type\":\"address\"},{\"name\":\"txHash\",\"type\":\"bytes32\"},{\"name\":\"challengingBlockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"childBlocks\",\"outputs\":[{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"timestamp\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHILD_BLOCK_INTERVAL\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentDepositBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"exits\",\"outputs\":[{\"name\":\"prevOwner\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"createdAt\",\"type\":\"uint256\"},{\"name\":\"bond\",\"type\":\"uint256\"},{\"name\":\"prevBlock\",\"type\":\"uint256\"},{\"name\":\"exitBlock\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"depositBlockNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"root\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"BlockSubmit\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"submitBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNextDepositBlockIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"getChildBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ParentChain is an auto generated Go binding around an Ethereum contract.
type ParentChain struct {
	ParentChainCaller     // Read-only binding to the contract
	ParentChainTransactor // Write-only binding to the contract
	ParentChainFilterer   // Log filterer for contract events
}

// ParentChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParentChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParentChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParentChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParentChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParentChainSession struct {
	Contract     *ParentChain      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParentChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParentChainCallerSession struct {
	Contract *ParentChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ParentChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParentChainTransactorSession struct {
	Contract     *ParentChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ParentChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParentChainRaw struct {
	Contract *ParentChain // Generic contract binding to access the raw methods on
}

// ParentChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParentChainCallerRaw struct {
	Contract *ParentChainCaller // Generic read-only contract binding to access the raw methods on
}

// ParentChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParentChainTransactorRaw struct {
	Contract *ParentChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParentChain creates a new instance of ParentChain, bound to a specific deployed contract.
func NewParentChain(address common.Address, backend bind.ContractBackend) (*ParentChain, error) {
	contract, err := bindParentChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ParentChain{ParentChainCaller: ParentChainCaller{contract: contract}, ParentChainTransactor: ParentChainTransactor{contract: contract}, ParentChainFilterer: ParentChainFilterer{contract: contract}}, nil
}

// NewParentChainCaller creates a new read-only instance of ParentChain, bound to a specific deployed contract.
func NewParentChainCaller(address common.Address, caller bind.ContractCaller) (*ParentChainCaller, error) {
	contract, err := bindParentChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParentChainCaller{contract: contract}, nil
}

// NewParentChainTransactor creates a new write-only instance of ParentChain, bound to a specific deployed contract.
func NewParentChainTransactor(address common.Address, transactor bind.ContractTransactor) (*ParentChainTransactor, error) {
	contract, err := bindParentChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParentChainTransactor{contract: contract}, nil
}

// NewParentChainFilterer creates a new log filterer instance of ParentChain, bound to a specific deployed contract.
func NewParentChainFilterer(address common.Address, filterer bind.ContractFilterer) (*ParentChainFilterer, error) {
	contract, err := bindParentChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParentChainFilterer{contract: contract}, nil
}

// bindParentChain binds a generic wrapper to an already deployed contract.
func bindParentChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ParentChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParentChain *ParentChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ParentChain.Contract.ParentChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParentChain *ParentChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParentChain.Contract.ParentChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParentChain *ParentChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParentChain.Contract.ParentChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ParentChain *ParentChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ParentChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ParentChain *ParentChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ParentChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ParentChain *ParentChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ParentChain.Contract.contract.Transact(opts, method, params...)
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_ParentChain *ParentChainCaller) CHILDBLOCKINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParentChain.contract.Call(opts, out, "CHILD_BLOCK_INTERVAL")
	return *ret0, err
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_ParentChain *ParentChainSession) CHILDBLOCKINTERVAL() (*big.Int, error) {
	return _ParentChain.Contract.CHILDBLOCKINTERVAL(&_ParentChain.CallOpts)
}

// CHILDBLOCKINTERVAL is a free data retrieval call binding the contract method 0xa831fa07.
//
// Solidity: function CHILD_BLOCK_INTERVAL() constant returns(uint256)
func (_ParentChain *ParentChainCallerSession) CHILDBLOCKINTERVAL() (*big.Int, error) {
	return _ParentChain.Contract.CHILDBLOCKINTERVAL(&_ParentChain.CallOpts)
}

// Challenges is a free data retrieval call binding the contract method 0x68e6a92e.
//
// Solidity: function challenges( uint64,  uint256) constant returns(owner address, challenger address, txHash bytes32, challengingBlockNumber uint256)
func (_ParentChain *ParentChainCaller) Challenges(opts *bind.CallOpts, arg0 uint64, arg1 *big.Int) (struct {
	Owner                  common.Address
	Challenger             common.Address
	TxHash                 [32]byte
	ChallengingBlockNumber *big.Int
}, error) {
	ret := new(struct {
		Owner                  common.Address
		Challenger             common.Address
		TxHash                 [32]byte
		ChallengingBlockNumber *big.Int
	})
	out := ret
	err := _ParentChain.contract.Call(opts, out, "challenges", arg0, arg1)
	return *ret, err
}

// Challenges is a free data retrieval call binding the contract method 0x68e6a92e.
//
// Solidity: function challenges( uint64,  uint256) constant returns(owner address, challenger address, txHash bytes32, challengingBlockNumber uint256)
func (_ParentChain *ParentChainSession) Challenges(arg0 uint64, arg1 *big.Int) (struct {
	Owner                  common.Address
	Challenger             common.Address
	TxHash                 [32]byte
	ChallengingBlockNumber *big.Int
}, error) {
	return _ParentChain.Contract.Challenges(&_ParentChain.CallOpts, arg0, arg1)
}

// Challenges is a free data retrieval call binding the contract method 0x68e6a92e.
//
// Solidity: function challenges( uint64,  uint256) constant returns(owner address, challenger address, txHash bytes32, challengingBlockNumber uint256)
func (_ParentChain *ParentChainCallerSession) Challenges(arg0 uint64, arg1 *big.Int) (struct {
	Owner                  common.Address
	Challenger             common.Address
	TxHash                 [32]byte
	ChallengingBlockNumber *big.Int
}, error) {
	return _ParentChain.Contract.Challenges(&_ParentChain.CallOpts, arg0, arg1)
}

// ChildBlocks is a free data retrieval call binding the contract method 0x97ef999b.
//
// Solidity: function childBlocks( uint256) constant returns(root bytes32, timestamp uint256)
func (_ParentChain *ParentChainCaller) ChildBlocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	ret := new(struct {
		Root      [32]byte
		Timestamp *big.Int
	})
	out := ret
	err := _ParentChain.contract.Call(opts, out, "childBlocks", arg0)
	return *ret, err
}

// ChildBlocks is a free data retrieval call binding the contract method 0x97ef999b.
//
// Solidity: function childBlocks( uint256) constant returns(root bytes32, timestamp uint256)
func (_ParentChain *ParentChainSession) ChildBlocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	return _ParentChain.Contract.ChildBlocks(&_ParentChain.CallOpts, arg0)
}

// ChildBlocks is a free data retrieval call binding the contract method 0x97ef999b.
//
// Solidity: function childBlocks( uint256) constant returns(root bytes32, timestamp uint256)
func (_ParentChain *ParentChainCallerSession) ChildBlocks(arg0 *big.Int) (struct {
	Root      [32]byte
	Timestamp *big.Int
}, error) {
	return _ParentChain.Contract.ChildBlocks(&_ParentChain.CallOpts, arg0)
}

// CoinCount is a free data retrieval call binding the contract method 0x678fd289.
//
// Solidity: function coinCount() constant returns(uint64)
func (_ParentChain *ParentChainCaller) CoinCount(opts *bind.CallOpts) (uint64, error) {
	var (
		ret0 = new(uint64)
	)
	out := ret0
	err := _ParentChain.contract.Call(opts, out, "coinCount")
	return *ret0, err
}

// CoinCount is a free data retrieval call binding the contract method 0x678fd289.
//
// Solidity: function coinCount() constant returns(uint64)
func (_ParentChain *ParentChainSession) CoinCount() (uint64, error) {
	return _ParentChain.Contract.CoinCount(&_ParentChain.CallOpts)
}

// CoinCount is a free data retrieval call binding the contract method 0x678fd289.
//
// Solidity: function coinCount() constant returns(uint64)
func (_ParentChain *ParentChainCallerSession) CoinCount() (uint64, error) {
	return _ParentChain.Contract.CoinCount(&_ParentChain.CallOpts)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_ParentChain *ParentChainCaller) CurrentChildBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParentChain.contract.Call(opts, out, "currentChildBlock")
	return *ret0, err
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_ParentChain *ParentChainSession) CurrentChildBlock() (*big.Int, error) {
	return _ParentChain.Contract.CurrentChildBlock(&_ParentChain.CallOpts)
}

// CurrentChildBlock is a free data retrieval call binding the contract method 0x7a95f1e8.
//
// Solidity: function currentChildBlock() constant returns(uint256)
func (_ParentChain *ParentChainCallerSession) CurrentChildBlock() (*big.Int, error) {
	return _ParentChain.Contract.CurrentChildBlock(&_ParentChain.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_ParentChain *ParentChainCaller) CurrentDepositBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParentChain.contract.Call(opts, out, "currentDepositBlock")
	return *ret0, err
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_ParentChain *ParentChainSession) CurrentDepositBlock() (*big.Int, error) {
	return _ParentChain.Contract.CurrentDepositBlock(&_ParentChain.CallOpts)
}

// CurrentDepositBlock is a free data retrieval call binding the contract method 0xa98c7f2c.
//
// Solidity: function currentDepositBlock() constant returns(uint256)
func (_ParentChain *ParentChainCallerSession) CurrentDepositBlock() (*big.Int, error) {
	return _ParentChain.Contract.CurrentDepositBlock(&_ParentChain.CallOpts)
}

// Exits is a free data retrieval call binding the contract method 0xd6463d40.
//
// Solidity: function exits( uint64) constant returns(prevOwner address, owner address, amount uint256, createdAt uint256, bond uint256, prevBlock uint256, exitBlock uint256)
func (_ParentChain *ParentChainCaller) Exits(opts *bind.CallOpts, arg0 uint64) (struct {
	PrevOwner common.Address
	Owner     common.Address
	Amount    *big.Int
	CreatedAt *big.Int
	Bond      *big.Int
	PrevBlock *big.Int
	ExitBlock *big.Int
}, error) {
	ret := new(struct {
		PrevOwner common.Address
		Owner     common.Address
		Amount    *big.Int
		CreatedAt *big.Int
		Bond      *big.Int
		PrevBlock *big.Int
		ExitBlock *big.Int
	})
	out := ret
	err := _ParentChain.contract.Call(opts, out, "exits", arg0)
	return *ret, err
}

// Exits is a free data retrieval call binding the contract method 0xd6463d40.
//
// Solidity: function exits( uint64) constant returns(prevOwner address, owner address, amount uint256, createdAt uint256, bond uint256, prevBlock uint256, exitBlock uint256)
func (_ParentChain *ParentChainSession) Exits(arg0 uint64) (struct {
	PrevOwner common.Address
	Owner     common.Address
	Amount    *big.Int
	CreatedAt *big.Int
	Bond      *big.Int
	PrevBlock *big.Int
	ExitBlock *big.Int
}, error) {
	return _ParentChain.Contract.Exits(&_ParentChain.CallOpts, arg0)
}

// Exits is a free data retrieval call binding the contract method 0xd6463d40.
//
// Solidity: function exits( uint64) constant returns(prevOwner address, owner address, amount uint256, createdAt uint256, bond uint256, prevBlock uint256, exitBlock uint256)
func (_ParentChain *ParentChainCallerSession) Exits(arg0 uint64) (struct {
	PrevOwner common.Address
	Owner     common.Address
	Amount    *big.Int
	CreatedAt *big.Int
	Bond      *big.Int
	PrevBlock *big.Int
	ExitBlock *big.Int
}, error) {
	return _ParentChain.Contract.Exits(&_ParentChain.CallOpts, arg0)
}

// GetChildBlock is a free data retrieval call binding the contract method 0xbb2dc863.
//
// Solidity: function getChildBlock(_blockNumber uint256) constant returns(bytes32, uint256)
func (_ParentChain *ParentChainCaller) GetChildBlock(opts *bind.CallOpts, _blockNumber *big.Int) ([32]byte, *big.Int, error) {
	var (
		ret0 = new([32]byte)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ParentChain.contract.Call(opts, out, "getChildBlock", _blockNumber)
	return *ret0, *ret1, err
}

// GetChildBlock is a free data retrieval call binding the contract method 0xbb2dc863.
//
// Solidity: function getChildBlock(_blockNumber uint256) constant returns(bytes32, uint256)
func (_ParentChain *ParentChainSession) GetChildBlock(_blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _ParentChain.Contract.GetChildBlock(&_ParentChain.CallOpts, _blockNumber)
}

// GetChildBlock is a free data retrieval call binding the contract method 0xbb2dc863.
//
// Solidity: function getChildBlock(_blockNumber uint256) constant returns(bytes32, uint256)
func (_ParentChain *ParentChainCallerSession) GetChildBlock(_blockNumber *big.Int) ([32]byte, *big.Int, error) {
	return _ParentChain.Contract.GetChildBlock(&_ParentChain.CallOpts, _blockNumber)
}

// GetNextDepositBlockIndex is a free data retrieval call binding the contract method 0x00a3216c.
//
// Solidity: function getNextDepositBlockIndex() constant returns(uint256)
func (_ParentChain *ParentChainCaller) GetNextDepositBlockIndex(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ParentChain.contract.Call(opts, out, "getNextDepositBlockIndex")
	return *ret0, err
}

// GetNextDepositBlockIndex is a free data retrieval call binding the contract method 0x00a3216c.
//
// Solidity: function getNextDepositBlockIndex() constant returns(uint256)
func (_ParentChain *ParentChainSession) GetNextDepositBlockIndex() (*big.Int, error) {
	return _ParentChain.Contract.GetNextDepositBlockIndex(&_ParentChain.CallOpts)
}

// GetNextDepositBlockIndex is a free data retrieval call binding the contract method 0x00a3216c.
//
// Solidity: function getNextDepositBlockIndex() constant returns(uint256)
func (_ParentChain *ParentChainCallerSession) GetNextDepositBlockIndex() (*big.Int, error) {
	return _ParentChain.Contract.GetNextDepositBlockIndex(&_ParentChain.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(from address, token address, value uint256) returns()
func (_ParentChain *ParentChainTransactor) Deposit(opts *bind.TransactOpts, from common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _ParentChain.contract.Transact(opts, "deposit", from, token, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(from address, token address, value uint256) returns()
func (_ParentChain *ParentChainSession) Deposit(from common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _ParentChain.Contract.Deposit(&_ParentChain.TransactOpts, from, token, value)
}

// Deposit is a paid mutator transaction binding the contract method 0x8340f549.
//
// Solidity: function deposit(from address, token address, value uint256) returns()
func (_ParentChain *ParentChainTransactorSession) Deposit(from common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _ParentChain.Contract.Deposit(&_ParentChain.TransactOpts, from, token, value)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns()
func (_ParentChain *ParentChainTransactor) SubmitBlock(opts *bind.TransactOpts, _root [32]byte) (*types.Transaction, error) {
	return _ParentChain.contract.Transact(opts, "submitBlock", _root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns()
func (_ParentChain *ParentChainSession) SubmitBlock(_root [32]byte) (*types.Transaction, error) {
	return _ParentChain.Contract.SubmitBlock(&_ParentChain.TransactOpts, _root)
}

// SubmitBlock is a paid mutator transaction binding the contract method 0xbaa47694.
//
// Solidity: function submitBlock(_root bytes32) returns()
func (_ParentChain *ParentChainTransactorSession) SubmitBlock(_root [32]byte) (*types.Transaction, error) {
	return _ParentChain.Contract.SubmitBlock(&_ParentChain.TransactOpts, _root)
}

// ParentChainBlockSubmitIterator is returned from FilterBlockSubmit and is used to iterate over the raw logs and unpacked data for BlockSubmit events raised by the ParentChain contract.
type ParentChainBlockSubmitIterator struct {
	Event *ParentChainBlockSubmit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParentChainBlockSubmitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParentChainBlockSubmit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParentChainBlockSubmit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParentChainBlockSubmitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParentChainBlockSubmitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParentChainBlockSubmit represents a BlockSubmit event raised by the ParentChain contract.
type ParentChainBlockSubmit struct {
	Root      [32]byte
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBlockSubmit is a free log retrieval operation binding the contract event 0x79f92be16163b2592e15e92db55fab076ee3c8b88792ef547d89febb58170792.
//
// Solidity: e BlockSubmit(root bytes32, timestamp uint256)
func (_ParentChain *ParentChainFilterer) FilterBlockSubmit(opts *bind.FilterOpts) (*ParentChainBlockSubmitIterator, error) {

	logs, sub, err := _ParentChain.contract.FilterLogs(opts, "BlockSubmit")
	if err != nil {
		return nil, err
	}
	return &ParentChainBlockSubmitIterator{contract: _ParentChain.contract, event: "BlockSubmit", logs: logs, sub: sub}, nil
}

// WatchBlockSubmit is a free log subscription operation binding the contract event 0x79f92be16163b2592e15e92db55fab076ee3c8b88792ef547d89febb58170792.
//
// Solidity: e BlockSubmit(root bytes32, timestamp uint256)
func (_ParentChain *ParentChainFilterer) WatchBlockSubmit(opts *bind.WatchOpts, sink chan<- *ParentChainBlockSubmit) (event.Subscription, error) {

	logs, sub, err := _ParentChain.contract.WatchLogs(opts, "BlockSubmit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParentChainBlockSubmit)
				if err := _ParentChain.contract.UnpackLog(event, "BlockSubmit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParentChainDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ParentChain contract.
type ParentChainDepositIterator struct {
	Event *ParentChainDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ParentChainDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ParentChainDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ParentChainDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ParentChainDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ParentChainDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ParentChainDeposit represents a Deposit event raised by the ParentChain contract.
type ParentChainDeposit struct {
	Depositor          common.Address
	DepositBlockNumber *big.Int
	Token              common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xd2f8022f659fd9c8c558f30c00fd5ee7038f7cb56da45095c3e0e7d48b3e0c4b.
//
// Solidity: e Deposit(depositor indexed address, depositBlockNumber indexed uint256, token address, amount uint256)
func (_ParentChain *ParentChainFilterer) FilterDeposit(opts *bind.FilterOpts, depositor []common.Address, depositBlockNumber []*big.Int) (*ParentChainDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositBlockNumberRule []interface{}
	for _, depositBlockNumberItem := range depositBlockNumber {
		depositBlockNumberRule = append(depositBlockNumberRule, depositBlockNumberItem)
	}

	logs, sub, err := _ParentChain.contract.FilterLogs(opts, "Deposit", depositorRule, depositBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &ParentChainDepositIterator{contract: _ParentChain.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xd2f8022f659fd9c8c558f30c00fd5ee7038f7cb56da45095c3e0e7d48b3e0c4b.
//
// Solidity: e Deposit(depositor indexed address, depositBlockNumber indexed uint256, token address, amount uint256)
func (_ParentChain *ParentChainFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ParentChainDeposit, depositor []common.Address, depositBlockNumber []*big.Int) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var depositBlockNumberRule []interface{}
	for _, depositBlockNumberItem := range depositBlockNumber {
		depositBlockNumberRule = append(depositBlockNumberRule, depositBlockNumberItem)
	}

	logs, sub, err := _ParentChain.contract.WatchLogs(opts, "Deposit", depositorRule, depositBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ParentChainDeposit)
				if err := _ParentChain.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
