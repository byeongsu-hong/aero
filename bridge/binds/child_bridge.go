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

// ChildBridgeABI is the input ABI used to generate the binding from.
const ChildBridgeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastBlockOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingTransactions\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"slotId\",\"type\":\"uint64\"},{\"name\":\"prevBlock\",\"type\":\"uint256\"},{\"name\":\"newOwner\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"txnHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"slotId\",\"type\":\"uint64\"}],\"name\":\"ConfirmTransaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"slotId\",\"type\":\"uint64\"}],\"name\":\"createTransaction\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txnHash\",\"type\":\"bytes32\"},{\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"saveWitness\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBlockNumber\",\"type\":\"uint256\"}],\"name\":\"submitNewBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\"},{\"name\":\"parentToken\",\"type\":\"address\"},{\"name\":\"slotId\",\"type\":\"uint64\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"which\",\"type\":\"uint8\"}],\"name\":\"submitDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getPendingTransactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ChildBridgeBin is the compiled bytecode used for deploying new contracts.
const ChildBridgeBin = `0x608060405234801561001057600080fd5b50604051602080610e18833981016040525160008054600160a060020a0319908116331790915560018054600160a060020a0390931692909116919091179055610db98061005f6000396000f3006080604052600436106100ae5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633efd0ac281146100b357806363a8374d146100dd578063642f2eaf146100f5578063715018a6146101ee5780638da5cb5b14610205578063a560cd7e14610236578063b3455f8914610273578063b3d3a508146102d1578063cafb00a014610305578063ddbd5d3f1461031d578063f2fde38b14610332575b600080fd5b3480156100bf57600080fd5b506100cb600435610353565b60408051918252519081900360200190f35b3480156100e957600080fd5b506100cb600435610365565b34801561010157600080fd5b5061010d600435610384565b6040518087600181111561011d57fe5b60ff1681526020018667ffffffffffffffff1667ffffffffffffffff16815260200185815260200184600160a060020a0316600160a060020a0316815260200183600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101ae578181015183820152602001610196565b50505050905090810190601f1680156101db5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b3480156101fa57600080fd5b5061020361045b565b005b34801561021157600080fd5b5061021a6104c7565b60408051600160a060020a039092168252519081900360200190f35b34801561024257600080fd5b50610203600160a060020a036004358116906024351667ffffffffffffffff6044351660643560ff608435166104d6565b34801561027f57600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526102039583359536956044949193909101919081908401838280828437509497506106c69650505050505050565b3480156102dd57600080fd5b506100cb600160a060020a036004358116906024351667ffffffffffffffff604435166107e0565b34801561031157600080fd5b50610203600435610a80565b34801561032957600080fd5b506100cb610b5f565b34801561033e57600080fd5b50610203600160a060020a0360043516610b66565b60036020526000908152604090205481565b600480548290811061037357fe5b600091825260209091200154905081565b60026020818152600092835260409283902080546001808301548386015460038501546004860180548a516101009682161587026000190190911699909904601f810189900489028a018901909a5289895260ff8616999490950467ffffffffffffffff16979296600160a060020a039283169692909116949193908301828280156104515780601f1061042657610100808354040283529160200191610451565b820191906000526020600020905b81548152906001019060200180831161043457829003601f168201915b5050505050905086565b600054600160a060020a0316331461047257600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b600080548190600160a060020a031633146104f057600080fd5b600154600160a060020a03878116911614610555576040805160e560020a62461bcd02815260206004820152601360248201527f556e7265676973746572656420746f6b656e2e00000000000000000000000000604482015290519081900360640190fd5b600083600181111561056357fe5b141561061a5785915081600160a060020a031663ea92283788866040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156105e857600080fd5b505af11580156105fc573d6000803e3d6000fd5b505050506040513d602081101561061257600080fd5b506106bd9050565b50604080517f079cec9f000000000000000000000000000000000000000000000000000000008152600160a060020a03888116600483015267ffffffffffffffff871660248301529151879283169163079cec9f9160448083019260209291908290030181600087803b15801561069057600080fd5b505af11580156106a4573d6000803e3d6000fd5b505050506040513d60208110156106ba57600080fd5b50505b50505050505050565b600082815260026020526040812054610100900467ffffffffffffffff16151561073a576040805160e560020a62461bcd02815260206004820152601860248201527f4e6f205472616e73616374696f6e20494420666f756e642e0000000000000000604482015290519081900360640190fd5b5060008281526002602052604090206003810154600160a060020a0316610767848463ffffffff610b8916565b600160a060020a0316146107c5576040805160e560020a62461bcd02815260206004820152601360248201527f5369676e6174757265206d69736d617463682e00000000000000000000000000604482015290519081900360640190fd5b81516107da9060048301906020850190610cdb565b50505050565b60015460009060609082908190600160a060020a0316331461084c576040805160e560020a62461bcd02815260206004820152601b60248201527f4469726563742063616c6c206973206e6f7420616c6c6f7765642e0000000000604482015290519081900360640190fd5b67ffffffffffffffff85166000818152600360209081526040918290205482517ff857000000000000000000000000000000000000000000000000000000000000818401527f88000000000000000000000000000000000000000000000000000000000000006022820152780100000000000000000000000000000000000000000000000090940260238501527fa000000000000000000000000000000000000000000000000000000000000000602b850152602c8401527f9400000000000000000000000000000000000000000000000000000000000000604c8401526c01000000000000000000000000600160a060020a038a1602604d840152815160418185030181526061909301918290528251929550859282918401908083835b6020831061098a5780518252601f19909201916020918201910161096b565b518151600019602094850361010090810a9190910191821691199290921617909152604080519590930185900390942060008181526002808452848220805467ffffffffffffffff909f1697880268ffffffffffffffff0019909f169e909e178e559581526003928390529283205460018d810191909155948c018054600160a060020a039e8f1673ffffffffffffffffffffffffffffffffffffffff199182161790915591909b0180549d909c169c169b909b17909955600480549182018155909952505050507f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b9094018290555092915050565b6000805481908190600160a060020a03163314610a9c57600080fd5b600092505b600454831015610b53576004805484908110610ab957fe5b600091825260208083209190910154808352600282526040808420805460ff19166001178082556101009081900467ffffffffffffffff90811687526003958690528387208b9055825495830154935194985091965090930490921692600160a060020a039092169185917f80e064808a697320936bc53930baab9e01c97b58f18f925c837ff77092d351a19190a4600190920191610aa1565b6107da60046000610d59565b6004545b90565b600054600160a060020a03163314610b7d57600080fd5b610b8681610c5e565b50565b60008060008084516041141515610ba35760009350610c55565b50505060208201516040830151606084015160001a601b60ff82161015610bc857601b015b8060ff16601b14158015610be057508060ff16601c14155b15610bee5760009350610c55565b60408051600080825260208083018085528a905260ff8516838501526060830187905260808301869052925160019360a0808501949193601f19840193928390039091019190865af1158015610c48573d6000803e3d6000fd5b5050506020604051035193505b50505092915050565b600160a060020a0381161515610c7357600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610d1c57805160ff1916838001178555610d49565b82800160010185558215610d49579182015b82811115610d49578251825591602001919060010190610d2e565b50610d55929150610d73565b5090565b5080546000825590600052602060002090810190610b8691905b610b6391905b80821115610d555760008155600101610d795600a165627a7a723058201a3c6a805374357ab7fb401b823f326737c41d273a4170033bb4a7cfcc76af2f0029`

// DeployChildBridge deploys a new Ethereum contract, binding an instance of ChildBridge to it.
func DeployChildBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address) (common.Address, *types.Transaction, *ChildBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(ChildBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChildBridgeBin), backend, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChildBridge{ChildBridgeCaller: ChildBridgeCaller{contract: contract}, ChildBridgeTransactor: ChildBridgeTransactor{contract: contract}, ChildBridgeFilterer: ChildBridgeFilterer{contract: contract}}, nil
}

// ChildBridge is an auto generated Go binding around an Ethereum contract.
type ChildBridge struct {
	ChildBridgeCaller     // Read-only binding to the contract
	ChildBridgeTransactor // Write-only binding to the contract
	ChildBridgeFilterer   // Log filterer for contract events
}

// ChildBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChildBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChildBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChildBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChildBridgeSession struct {
	Contract     *ChildBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChildBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChildBridgeCallerSession struct {
	Contract *ChildBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ChildBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChildBridgeTransactorSession struct {
	Contract     *ChildBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ChildBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChildBridgeRaw struct {
	Contract *ChildBridge // Generic contract binding to access the raw methods on
}

// ChildBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChildBridgeCallerRaw struct {
	Contract *ChildBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// ChildBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChildBridgeTransactorRaw struct {
	Contract *ChildBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChildBridge creates a new instance of ChildBridge, bound to a specific deployed contract.
func NewChildBridge(address common.Address, backend bind.ContractBackend) (*ChildBridge, error) {
	contract, err := bindChildBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChildBridge{ChildBridgeCaller: ChildBridgeCaller{contract: contract}, ChildBridgeTransactor: ChildBridgeTransactor{contract: contract}, ChildBridgeFilterer: ChildBridgeFilterer{contract: contract}}, nil
}

// NewChildBridgeCaller creates a new read-only instance of ChildBridge, bound to a specific deployed contract.
func NewChildBridgeCaller(address common.Address, caller bind.ContractCaller) (*ChildBridgeCaller, error) {
	contract, err := bindChildBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeCaller{contract: contract}, nil
}

// NewChildBridgeTransactor creates a new write-only instance of ChildBridge, bound to a specific deployed contract.
func NewChildBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChildBridgeTransactor, error) {
	contract, err := bindChildBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeTransactor{contract: contract}, nil
}

// NewChildBridgeFilterer creates a new log filterer instance of ChildBridge, bound to a specific deployed contract.
func NewChildBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChildBridgeFilterer, error) {
	contract, err := bindChildBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeFilterer{contract: contract}, nil
}

// bindChildBridge binds a generic wrapper to an already deployed contract.
func bindChildBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChildBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildBridge *ChildBridgeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChildBridge.Contract.ChildBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildBridge *ChildBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildBridge.Contract.ChildBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildBridge *ChildBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildBridge.Contract.ChildBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChildBridge *ChildBridgeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChildBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChildBridge *ChildBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChildBridge *ChildBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChildBridge.Contract.contract.Transact(opts, method, params...)
}

// GetPendingTransactionCount is a free data retrieval call binding the contract method 0xddbd5d3f.
//
// Solidity: function getPendingTransactionCount() constant returns(uint256)
func (_ChildBridge *ChildBridgeCaller) GetPendingTransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChildBridge.contract.Call(opts, out, "getPendingTransactionCount")
	return *ret0, err
}

// GetPendingTransactionCount is a free data retrieval call binding the contract method 0xddbd5d3f.
//
// Solidity: function getPendingTransactionCount() constant returns(uint256)
func (_ChildBridge *ChildBridgeSession) GetPendingTransactionCount() (*big.Int, error) {
	return _ChildBridge.Contract.GetPendingTransactionCount(&_ChildBridge.CallOpts)
}

// GetPendingTransactionCount is a free data retrieval call binding the contract method 0xddbd5d3f.
//
// Solidity: function getPendingTransactionCount() constant returns(uint256)
func (_ChildBridge *ChildBridgeCallerSession) GetPendingTransactionCount() (*big.Int, error) {
	return _ChildBridge.Contract.GetPendingTransactionCount(&_ChildBridge.CallOpts)
}

// LastBlockOf is a free data retrieval call binding the contract method 0x3efd0ac2.
//
// Solidity: function lastBlockOf( uint256) constant returns(uint256)
func (_ChildBridge *ChildBridgeCaller) LastBlockOf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ChildBridge.contract.Call(opts, out, "lastBlockOf", arg0)
	return *ret0, err
}

// LastBlockOf is a free data retrieval call binding the contract method 0x3efd0ac2.
//
// Solidity: function lastBlockOf( uint256) constant returns(uint256)
func (_ChildBridge *ChildBridgeSession) LastBlockOf(arg0 *big.Int) (*big.Int, error) {
	return _ChildBridge.Contract.LastBlockOf(&_ChildBridge.CallOpts, arg0)
}

// LastBlockOf is a free data retrieval call binding the contract method 0x3efd0ac2.
//
// Solidity: function lastBlockOf( uint256) constant returns(uint256)
func (_ChildBridge *ChildBridgeCallerSession) LastBlockOf(arg0 *big.Int) (*big.Int, error) {
	return _ChildBridge.Contract.LastBlockOf(&_ChildBridge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChildBridge *ChildBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ChildBridge.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChildBridge *ChildBridgeSession) Owner() (common.Address, error) {
	return _ChildBridge.Contract.Owner(&_ChildBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ChildBridge *ChildBridgeCallerSession) Owner() (common.Address, error) {
	return _ChildBridge.Contract.Owner(&_ChildBridge.CallOpts)
}

// PendingTransactions is a free data retrieval call binding the contract method 0x63a8374d.
//
// Solidity: function pendingTransactions( uint256) constant returns(bytes32)
func (_ChildBridge *ChildBridgeCaller) PendingTransactions(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ChildBridge.contract.Call(opts, out, "pendingTransactions", arg0)
	return *ret0, err
}

// PendingTransactions is a free data retrieval call binding the contract method 0x63a8374d.
//
// Solidity: function pendingTransactions( uint256) constant returns(bytes32)
func (_ChildBridge *ChildBridgeSession) PendingTransactions(arg0 *big.Int) ([32]byte, error) {
	return _ChildBridge.Contract.PendingTransactions(&_ChildBridge.CallOpts, arg0)
}

// PendingTransactions is a free data retrieval call binding the contract method 0x63a8374d.
//
// Solidity: function pendingTransactions( uint256) constant returns(bytes32)
func (_ChildBridge *ChildBridgeCallerSession) PendingTransactions(arg0 *big.Int) ([32]byte, error) {
	return _ChildBridge.Contract.PendingTransactions(&_ChildBridge.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(status uint8, slotId uint64, prevBlock uint256, newOwner address, owner address, signature bytes)
func (_ChildBridge *ChildBridgeCaller) Transactions(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _ChildBridge.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(status uint8, slotId uint64, prevBlock uint256, newOwner address, owner address, signature bytes)
func (_ChildBridge *ChildBridgeSession) Transactions(arg0 [32]byte) (struct {
	Status    uint8
	SlotId    uint64
	PrevBlock *big.Int
	NewOwner  common.Address
	Owner     common.Address
	Signature []byte
}, error) {
	return _ChildBridge.Contract.Transactions(&_ChildBridge.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x642f2eaf.
//
// Solidity: function transactions( bytes32) constant returns(status uint8, slotId uint64, prevBlock uint256, newOwner address, owner address, signature bytes)
func (_ChildBridge *ChildBridgeCallerSession) Transactions(arg0 [32]byte) (struct {
	Status    uint8
	SlotId    uint64
	PrevBlock *big.Int
	NewOwner  common.Address
	Owner     common.Address
	Signature []byte
}, error) {
	return _ChildBridge.Contract.Transactions(&_ChildBridge.CallOpts, arg0)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xb3d3a508.
//
// Solidity: function createTransaction(from address, to address, slotId uint64) returns(bytes32)
func (_ChildBridge *ChildBridgeTransactor) CreateTransaction(opts *bind.TransactOpts, from common.Address, to common.Address, slotId uint64) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "createTransaction", from, to, slotId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xb3d3a508.
//
// Solidity: function createTransaction(from address, to address, slotId uint64) returns(bytes32)
func (_ChildBridge *ChildBridgeSession) CreateTransaction(from common.Address, to common.Address, slotId uint64) (*types.Transaction, error) {
	return _ChildBridge.Contract.CreateTransaction(&_ChildBridge.TransactOpts, from, to, slotId)
}

// CreateTransaction is a paid mutator transaction binding the contract method 0xb3d3a508.
//
// Solidity: function createTransaction(from address, to address, slotId uint64) returns(bytes32)
func (_ChildBridge *ChildBridgeTransactorSession) CreateTransaction(from common.Address, to common.Address, slotId uint64) (*types.Transaction, error) {
	return _ChildBridge.Contract.CreateTransaction(&_ChildBridge.TransactOpts, from, to, slotId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChildBridge *ChildBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChildBridge *ChildBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChildBridge.Contract.RenounceOwnership(&_ChildBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ChildBridge *ChildBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ChildBridge.Contract.RenounceOwnership(&_ChildBridge.TransactOpts)
}

// SaveWitness is a paid mutator transaction binding the contract method 0xb3455f89.
//
// Solidity: function saveWitness(txnHash bytes32, signature bytes) returns()
func (_ChildBridge *ChildBridgeTransactor) SaveWitness(opts *bind.TransactOpts, txnHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "saveWitness", txnHash, signature)
}

// SaveWitness is a paid mutator transaction binding the contract method 0xb3455f89.
//
// Solidity: function saveWitness(txnHash bytes32, signature bytes) returns()
func (_ChildBridge *ChildBridgeSession) SaveWitness(txnHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ChildBridge.Contract.SaveWitness(&_ChildBridge.TransactOpts, txnHash, signature)
}

// SaveWitness is a paid mutator transaction binding the contract method 0xb3455f89.
//
// Solidity: function saveWitness(txnHash bytes32, signature bytes) returns()
func (_ChildBridge *ChildBridgeTransactorSession) SaveWitness(txnHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _ChildBridge.Contract.SaveWitness(&_ChildBridge.TransactOpts, txnHash, signature)
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xa560cd7e.
//
// Solidity: function submitDeposit(depositor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildBridge *ChildBridgeTransactor) SubmitDeposit(opts *bind.TransactOpts, depositor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "submitDeposit", depositor, parentToken, slotId, amount, which)
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xa560cd7e.
//
// Solidity: function submitDeposit(depositor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildBridge *ChildBridgeSession) SubmitDeposit(depositor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildBridge.Contract.SubmitDeposit(&_ChildBridge.TransactOpts, depositor, parentToken, slotId, amount, which)
}

// SubmitDeposit is a paid mutator transaction binding the contract method 0xa560cd7e.
//
// Solidity: function submitDeposit(depositor address, parentToken address, slotId uint64, amount uint256, which uint8) returns()
func (_ChildBridge *ChildBridgeTransactorSession) SubmitDeposit(depositor common.Address, parentToken common.Address, slotId uint64, amount *big.Int, which uint8) (*types.Transaction, error) {
	return _ChildBridge.Contract.SubmitDeposit(&_ChildBridge.TransactOpts, depositor, parentToken, slotId, amount, which)
}

// SubmitNewBlock is a paid mutator transaction binding the contract method 0xcafb00a0.
//
// Solidity: function submitNewBlock(newBlockNumber uint256) returns()
func (_ChildBridge *ChildBridgeTransactor) SubmitNewBlock(opts *bind.TransactOpts, newBlockNumber *big.Int) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "submitNewBlock", newBlockNumber)
}

// SubmitNewBlock is a paid mutator transaction binding the contract method 0xcafb00a0.
//
// Solidity: function submitNewBlock(newBlockNumber uint256) returns()
func (_ChildBridge *ChildBridgeSession) SubmitNewBlock(newBlockNumber *big.Int) (*types.Transaction, error) {
	return _ChildBridge.Contract.SubmitNewBlock(&_ChildBridge.TransactOpts, newBlockNumber)
}

// SubmitNewBlock is a paid mutator transaction binding the contract method 0xcafb00a0.
//
// Solidity: function submitNewBlock(newBlockNumber uint256) returns()
func (_ChildBridge *ChildBridgeTransactorSession) SubmitNewBlock(newBlockNumber *big.Int) (*types.Transaction, error) {
	return _ChildBridge.Contract.SubmitNewBlock(&_ChildBridge.TransactOpts, newBlockNumber)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ChildBridge *ChildBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ChildBridge.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ChildBridge *ChildBridgeSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ChildBridge.Contract.TransferOwnership(&_ChildBridge.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_ChildBridge *ChildBridgeTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ChildBridge.Contract.TransferOwnership(&_ChildBridge.TransactOpts, _newOwner)
}

// ChildBridgeConfirmTransactionIterator is returned from FilterConfirmTransaction and is used to iterate over the raw logs and unpacked data for ConfirmTransaction events raised by the ChildBridge contract.
type ChildBridgeConfirmTransactionIterator struct {
	Event *ChildBridgeConfirmTransaction // Event containing the contract specifics and raw log

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
func (it *ChildBridgeConfirmTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildBridgeConfirmTransaction)
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
		it.Event = new(ChildBridgeConfirmTransaction)
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
func (it *ChildBridgeConfirmTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildBridgeConfirmTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildBridgeConfirmTransaction represents a ConfirmTransaction event raised by the ChildBridge contract.
type ChildBridgeConfirmTransaction struct {
	TxnHash [32]byte
	Owner   common.Address
	SlotId  uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConfirmTransaction is a free log retrieval operation binding the contract event 0x80e064808a697320936bc53930baab9e01c97b58f18f925c837ff77092d351a1.
//
// Solidity: event ConfirmTransaction(txnHash indexed bytes32, owner indexed address, slotId indexed uint64)
func (_ChildBridge *ChildBridgeFilterer) FilterConfirmTransaction(opts *bind.FilterOpts, txnHash [][32]byte, owner []common.Address, slotId []uint64) (*ChildBridgeConfirmTransactionIterator, error) {

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

	logs, sub, err := _ChildBridge.contract.FilterLogs(opts, "ConfirmTransaction", txnHashRule, ownerRule, slotIdRule)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeConfirmTransactionIterator{contract: _ChildBridge.contract, event: "ConfirmTransaction", logs: logs, sub: sub}, nil
}

// WatchConfirmTransaction is a free log subscription operation binding the contract event 0x80e064808a697320936bc53930baab9e01c97b58f18f925c837ff77092d351a1.
//
// Solidity: event ConfirmTransaction(txnHash indexed bytes32, owner indexed address, slotId indexed uint64)
func (_ChildBridge *ChildBridgeFilterer) WatchConfirmTransaction(opts *bind.WatchOpts, sink chan<- *ChildBridgeConfirmTransaction, txnHash [][32]byte, owner []common.Address, slotId []uint64) (event.Subscription, error) {

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

	logs, sub, err := _ChildBridge.contract.WatchLogs(opts, "ConfirmTransaction", txnHashRule, ownerRule, slotIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildBridgeConfirmTransaction)
				if err := _ChildBridge.contract.UnpackLog(event, "ConfirmTransaction", log); err != nil {
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

// ChildBridgeOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the ChildBridge contract.
type ChildBridgeOwnershipRenouncedIterator struct {
	Event *ChildBridgeOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *ChildBridgeOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildBridgeOwnershipRenounced)
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
		it.Event = new(ChildBridgeOwnershipRenounced)
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
func (it *ChildBridgeOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildBridgeOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildBridgeOwnershipRenounced represents a OwnershipRenounced event raised by the ChildBridge contract.
type ChildBridgeOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_ChildBridge *ChildBridgeFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*ChildBridgeOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ChildBridge.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeOwnershipRenouncedIterator{contract: _ChildBridge.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_ChildBridge *ChildBridgeFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *ChildBridgeOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _ChildBridge.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildBridgeOwnershipRenounced)
				if err := _ChildBridge.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ChildBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ChildBridge contract.
type ChildBridgeOwnershipTransferredIterator struct {
	Event *ChildBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChildBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildBridgeOwnershipTransferred)
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
		it.Event = new(ChildBridgeOwnershipTransferred)
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
func (it *ChildBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the ChildBridge contract.
type ChildBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ChildBridge *ChildBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChildBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChildBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChildBridgeOwnershipTransferredIterator{contract: _ChildBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_ChildBridge *ChildBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChildBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ChildBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildBridgeOwnershipTransferred)
				if err := _ChildBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
