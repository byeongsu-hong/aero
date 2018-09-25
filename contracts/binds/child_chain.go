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

// ChildChainABI is the input ABI used to generate the binding from.
const ChildChainABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastBlockOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingTransactions\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"slotId\",\"type\":\"uint64\"},{\"name\":\"prevBlock\",\"type\":\"uint256\"},{\"name\":\"newOwner\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"txnHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"slotId\",\"type\":\"uint64\"}],\"name\":\"ConfirmTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"slotId\",\"type\":\"uint64\"}],\"name\":\"createTransaction\",\"outputs\":[{\"name\":\"txnHash\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txnHash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"saveWitness\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBlockNumber\",\"type\":\"uint256\"}],\"name\":\"submitNewBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"},{\"name\":\"parentToken\",\"type\":\"address\"},{\"name\":\"slotId\",\"type\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"which\",\"type\":\"uint8\"}],\"name\":\"submitDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPendingTransactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChildChain is an auto generated Go binding around an Ethereum contract.
type ChildChain struct {
	ChildChainCaller     // Read-only binding to the contract
	ChildChainTransactor // Write-only binding to the contract
	ChildChainFilterer   // Log filterer for contract events
}

// ChildChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChildChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChildChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChildChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChildChainSession struct {
	Contract     *ChildChain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChildChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChildChainCallerSession struct {
	Contract *ChildChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ChildChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChildChainTransactorSession struct {
	Contract     *ChildChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ChildChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChildChainRaw struct {
	Contract *ChildChain // Generic contract binding to access the raw methods on
}

// ChildChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChildChainCallerRaw struct {
	Contract *ChildChainCaller // Generic read-only contract binding to access the raw methods on
}

// ChildChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChildChainTransactorRaw struct {
	Contract *ChildChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChildChain creates a new instance of ChildChain, bound to a specific deployed contract.
func NewChildChain(address common.Address, backend bind.ContractBackend) (*ChildChain, error) {
	contract, err := bindChildChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChildChain{ChildChainCaller: ChildChainCaller{contract: contract}, ChildChainTransactor: ChildChainTransactor{contract: contract}, ChildChainFilterer: ChildChainFilterer{contract: contract}}, nil
}

// NewChildChainCaller creates a new read-only instance of ChildChain, bound to a specific deployed contract.
func NewChildChainCaller(address common.Address, caller bind.ContractCaller) (*ChildChainCaller, error) {
	contract, err := bindChildChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChildChainCaller{contract: contract}, nil
}

// NewChildChainTransactor creates a new write-only instance of ChildChain, bound to a specific deployed contract.
func NewChildChainTransactor(address common.Address, transactor bind.ContractTransactor) (*ChildChainTransactor, error) {
	contract, err := bindChildChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChildChainTransactor{contract: contract}, nil
}

// NewChildChainFilterer creates a new log filterer instance of ChildChain, bound to a specific deployed contract.
func NewChildChainFilterer(address common.Address, filterer bind.ContractFilterer) (*ChildChainFilterer, error) {
	contract, err := bindChildChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChildChainFilterer{contract: contract}, nil
}

// bindChildChain binds a generic wrapper to an already deployed contract.
func bindChildChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChildChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildChain *ChildChainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChildChain.Contract.ChildChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildChain *ChildChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildChain.Contract.ChildChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildChain *ChildChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildChain.Contract.ChildChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildChain *ChildChainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChildChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildChain *ChildChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildChain *ChildChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildChain.Contract.contract.Transact(opts, method, params...)
}

// GetPendingTransactionCount is a free data retrieval call binding the contract method 0xddbd5d3f.
//
// Solidity: function getPendingTransactionCount() constant returns(uint256)
func (_ChildChain *ChildChainCaller) GetPendingTransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChildChain.contract.Call(opts, out, "getPendingTransactionCount")
	return *ret0, err
}

// GetPendingTransactionCount is a free data retrieval call binding the contract method 0xddbd5d3f.
//
// Solidity: function getPendingTransactionCount() constant returns(uint256)
func (_ChildChain *ChildChainSession) GetPendingTransactionCount() (*big.Int, error) {
	return _ChildChain.Contract.GetPendingTransactionCount(&_ChildChain.CallOpts)
}

// GetPendingTransactionCount is a free data retrieval call binding the contract method 0xddbd5d3f.
//
// Solidity: function getPendingTransactionCount() constant returns(uint256)
func (_ChildChain *ChildChainCallerSession) GetPendingTransactionCount() (*big.Int, error) {
	return _ChildChain.Contract.GetPendingTransactionCount(&_ChildChain.CallOpts)
}

// LastBlockOf is a free data retrieval call binding the contract method 0x3efd0ac2.
//
// Solidity: function lastBlockOf( uint256) constant returns(uint256)
func (_ChildChain *ChildChainCaller) LastBlockOf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChildChain.contract.Call(opts, out, "lastBlockOf", arg0)
	return *ret0, err
}

// LastBlockOf is a free data retrieval call binding the contract method 0x3efd0ac2.
//
// Solidity: function lastBlockOf( uint256) constant returns(uint256)
func (_ChildChain *ChildChainSession) LastBlockOf(arg0 *big.Int) (*big.Int, error) {
	return _ChildChain.Contract.LastBlockOf(&_ChildChain.CallOpts, arg0)
}

// LastBlockOf is a free data retrieval call binding the contract method 0x3efd0ac2.
//
// Solidity: function lastBlockOf( uint256) constant returns(uint256)
func (_ChildChain *ChildChainCallerSession) LastBlockOf(arg0 *big.Int) (*big.Int, error) {
	return _ChildChain.Contract.LastBlockOf(&_ChildChain.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChildChain *ChildChainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChildChain.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChildChain *ChildChainSession) Owner() (common.Address, error) {
	return _ChildChain.Contract.Owner(&_ChildChain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChildChain *ChildChainCallerSession) Owner() (common.Address, error) {
	return _ChildChain.Contract.Owner(&_ChildChain.CallOpts)
}

// PendingTransactions is a free data retrieval call binding the contract method 0x63a8374d.
//
// Solidity: function pendingTransactions( uint256) constant returns(bytes32)
func (_ChildChain *ChildChainCaller) PendingTransactions(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ChildChain.contract.Call(opts, out, "pendingTransactions", arg0)
	return *ret0, err
}

// PendingTransactions is a free data retrieval call binding the contract method 0x63a8374d.
//
// Solidity: function pendingTransactions( uint256) constant returns(bytes32)
func (_ChildChain *ChildChainSession) PendingTransactions(arg0 *big.Int) ([32]byte, error) {
	return _ChildChain.Contract.PendingTransactions(&_ChildChain.CallOpts, arg0)
}

// PendingTransactions is a free data retrieval call binding the contract method 0x63a8374d.
//
// Solidity: function pendingTransactions( uint256) constant returns(bytes32)
func (_ChildChain *ChildChainCallerSession) PendingTransactions(arg0 *big.Int) ([32]byte, error) {
	return _ChildChain.Contract.PendingTransactions(&_ChildChain.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(status uint8, slotId uint64, prevBlock uint256, newOwner address, owner address, signature bytes)
func (_ChildChain *ChildChainCaller) Transactions(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status    uint8
	SlotId    uint64
	PrevBlock *big.Int
	NewOwner  common.Address
	Owner     common.Address
	Signature []byte
}, error) {
	ret := new(struct {
		Status    uint8
		SlotId    uint64
		PrevBlock *big.Int
		NewOwner  common.Address
		Owner     common.Address
		Signature []byte
	})
	out := ret
	err := _ChildChain.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(status uint8, slotId uint64, prevBlock uint256, newOwner address, owner address, signature bytes)
func (_ChildChain *ChildChainSession) Transactions(arg0 [32]byte) (struct {
	Status    uint8
	SlotId    uint64
	PrevBlock *big.Int
	NewOwner  common.Address
	Owner     common.Address
	Signature []byte
}, error) {
	return _ChildChain.Contract.Transactions(&_ChildChain.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(status uint8, slotId uint64, prevBlock uint256, newOwner address, owner address, signature bytes)
func (_ChildChain *ChildChainCallerSession) Transactions(arg0 [32]byte) (struct {
	Status    uint8
	SlotId    uint64
	PrevBlock *big.Int
	NewOwner  common.Address
	Owner     common.Address
	Signature []byte
}, error) {
	return _ChildChain.Contract.Transactions(&_ChildChain.CallOpts, arg0)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xb3d3a508.
//
// Solidity: function createTransaction(from address, to address, slotId uint64) returns(txnHash bytes32)
func (_ChildChain *ChildChainTransactor) CreateTransaction(opts *bind.TransactOpts, from common.Address, to common.Address, slotId uint64) (*types.Transaction, error) {
	return _ChildChain.contract.Transact(opts, "createTransaction", from, to, slotId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xb3d3a508.
//
// Solidity: function createTransaction(from address, to address, slotId uint64) returns(txnHash bytes32)
func (_ChildChain *ChildChainSession) CreateTransaction(from common.Address, to common.Address, slotId uint64) (*types.Transaction, error) {
	return _ChildChain.Contract.CreateTransaction(&_ChildChain.TransactOpts, from, to, slotId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xb3d3a508.
//
// Solidity: function createTransaction(from address, to address, slotId uint64) returns(txnHash bytes32)
func (_ChildChain *ChildChainTransactorSession) CreateTransaction(from common.Address, to common.Address, slotId uint64) (*types.Transaction, error) {
	return _ChildChain.Contract.CreateTransaction(&_ChildChain.TransactOpts, from, to, slotId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChildChain *ChildChainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildChain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChildChain *ChildChainSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChildChain.Contract.RenounceOwnership(&_ChildChain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChildChain *ChildChainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChildChain.Contract.RenounceOwnership(&_ChildChain.TransactOpts)
}

// SaveWitness is a paid mutator transaction binding the contract method 0xb3455f89.
//
// Solidity: function saveWitness(txnHash bytes32, signature bytes) returns()
func (_ChildChain *ChildChainTransactor) SaveWitness(opts *bind.TransactOpts, txnHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ChildChain.contract.Transact(opts, "saveWitness", txnHash, signature)
}

// SaveWitness is a paid mutator transaction binding the contract method 0xb3455f89.
//
// Solidity: function saveWitness(txnHash bytes32, signature bytes) returns()
func (_ChildChain *ChildChainSession) SaveWitness(txnHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ChildChain.Contract.SaveWitness(&_ChildChain.TransactOpts, txnHash, signature)
}

// SaveWitness is a paid mutator transaction binding the contract method 0xb3455f89.
//
// Solidity: function saveWitness(txnHash bytes32, signature bytes) returns()
func (_ChildChain *ChildChainTransactorSession) SaveWitness(txnHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ChildChain.Contract.SaveWitness(&_ChildChain.TransactOpts, txnHash, signature)
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xa560cd7e.
//
// Solidity: function submitDeposit(depositor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildChain *ChildChainTransactor) SubmitDeposit(opts *bind.TransactOpts, depositor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildChain.contract.Transact(opts, "submitDeposit", depositor, parentToken, slotId, amount, which)
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xa560cd7e.
//
// Solidity: function submitDeposit(depositor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildChain *ChildChainSession) SubmitDeposit(depositor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildChain.Contract.SubmitDeposit(&_ChildChain.TransactOpts, depositor, parentToken, slotId, amount, which)
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xa560cd7e.
//
// Solidity: function submitDeposit(depositor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildChain *ChildChainTransactorSession) SubmitDeposit(depositor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildChain.Contract.SubmitDeposit(&_ChildChain.TransactOpts, depositor, parentToken, slotId, amount, which)
}

// SubmitNewBlock is a paid mutator transaction binding the contract method 0xcafb00a0.
//
// Solidity: function submitNewBlock(newBlockNumber uint256) returns()
func (_ChildChain *ChildChainTransactor) SubmitNewBlock(opts *bind.TransactOpts, newBlockNumber *big.Int) (*types.Transaction, error) {
	return _ChildChain.contract.Transact(opts, "submitNewBlock", newBlockNumber)
}

// SubmitNewBlock is a paid mutator transaction binding the contract method 0xcafb00a0.
//
// Solidity: function submitNewBlock(newBlockNumber uint256) returns()
func (_ChildChain *ChildChainSession) SubmitNewBlock(newBlockNumber *big.Int) (*types.Transaction, error) {
	return _ChildChain.Contract.SubmitNewBlock(&_ChildChain.TransactOpts, newBlockNumber)
}

// SubmitNewBlock is a paid mutator transaction binding the contract method 0xcafb00a0.
//
// Solidity: function submitNewBlock(newBlockNumber uint256) returns()
func (_ChildChain *ChildChainTransactorSession) SubmitNewBlock(newBlockNumber *big.Int) (*types.Transaction, error) {
	return _ChildChain.Contract.SubmitNewBlock(&_ChildChain.TransactOpts, newBlockNumber)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ChildChain *ChildChainTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ChildChain.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ChildChain *ChildChainSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ChildChain.Contract.TransferOwnership(&_ChildChain.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ChildChain *ChildChainTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ChildChain.Contract.TransferOwnership(&_ChildChain.TransactOpts, _newOwner)
}

// ChildChainConfirmTransactionIterator is returned from FilterConfirmTransaction and is used to iterate over the raw logs and unpacked data for ConfirmTransaction events raised by the ChildChain contract.
type ChildChainConfirmTransactionIterator struct {
	Event *ChildChainConfirmTransaction // Event containing the contract specifics and raw log

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
func (it *ChildChainConfirmTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildChainConfirmTransaction)
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
		it.Event = new(ChildChainConfirmTransaction)
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
func (it *ChildChainConfirmTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildChainConfirmTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildChainConfirmTransaction represents a ConfirmTransaction event raised by the ChildChain contract.
type ChildChainConfirmTransaction struct {
	TxnHash [32]byte
	Owner   common.Address
	SlotId  uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConfirmTransaction is a free log retrieval operation binding the contract event 0x80e064808a697320936bc53930baab9e01c97b58f18f925c837ff77092d351a1.
//
// Solidity: e ConfirmTransaction(txnHash indexed bytes32, owner indexed address, slotId indexed uint64)
func (_ChildChain *ChildChainFilterer) FilterConfirmTransaction(opts *bind.FilterOpts, txnHash [][32]byte, owner []common.Address, slotId []uint64) (*ChildChainConfirmTransactionIterator, error) {

	var txnHashRule []interface{}
	for _, txnHashItem := range txnHash {
		txnHashRule = append(txnHashRule, txnHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}

	logs, sub, err := _ChildChain.contract.FilterLogs(opts, "ConfirmTransaction", txnHashRule, ownerRule, slotIdRule)
	if err != nil {
		return nil, err
	}
	return &ChildChainConfirmTransactionIterator{contract: _ChildChain.contract, event: "ConfirmTransaction", logs: logs, sub: sub}, nil
}

// WatchConfirmTransaction is a free log subscription operation binding the contract event 0x80e064808a697320936bc53930baab9e01c97b58f18f925c837ff77092d351a1.
//
// Solidity: e ConfirmTransaction(txnHash indexed bytes32, owner indexed address, slotId indexed uint64)
func (_ChildChain *ChildChainFilterer) WatchConfirmTransaction(opts *bind.WatchOpts, sink chan<- *ChildChainConfirmTransaction, txnHash [][32]byte, owner []common.Address, slotId []uint64) (event.Subscription, error) {

	var txnHashRule []interface{}
	for _, txnHashItem := range txnHash {
		txnHashRule = append(txnHashRule, txnHashItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var slotIdRule []interface{}
	for _, slotIdItem := range slotId {
		slotIdRule = append(slotIdRule, slotIdItem)
	}

	logs, sub, err := _ChildChain.contract.WatchLogs(opts, "ConfirmTransaction", txnHashRule, ownerRule, slotIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildChainConfirmTransaction)
				if err := _ChildChain.contract.UnpackLog(event, "ConfirmTransaction", log); err != nil {
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

// ChildChainOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the ChildChain contract.
type ChildChainOwnershipRenouncedIterator struct {
	Event *ChildChainOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ChildChainOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildChainOwnershipRenounced)
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
		it.Event = new(ChildChainOwnershipRenounced)
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
func (it *ChildChainOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildChainOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildChainOwnershipRenounced represents a OwnershipRenounced event raised by the ChildChain contract.
type ChildChainOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ChildChain *ChildChainFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ChildChainOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ChildChain.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChildChainOwnershipRenouncedIterator{contract: _ChildChain.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_ChildChain *ChildChainFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ChildChainOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ChildChain.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildChainOwnershipRenounced)
				if err := _ChildChain.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ChildChainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChildChain contract.
type ChildChainOwnershipTransferredIterator struct {
	Event *ChildChainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChildChainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildChainOwnershipTransferred)
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
		it.Event = new(ChildChainOwnershipTransferred)
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
func (it *ChildChainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildChainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildChainOwnershipTransferred represents a OwnershipTransferred event raised by the ChildChain contract.
type ChildChainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ChildChain *ChildChainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChildChainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChildChain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChildChainOwnershipTransferredIterator{contract: _ChildChain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ChildChain *ChildChainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChildChainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChildChain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildChainOwnershipTransferred)
				if err := _ChildChain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
