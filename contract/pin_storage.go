// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/v3/abi"
	"github.com/FISCO-BCOS/go-sdk/v3/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/v3/types"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
)

// PinStorageABI is the input ABI used to generate the binding from.
const PinStorageABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_initialPin\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldPin\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newPin\",\"type\":\"string\"}],\"name\":\"PinChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pin\",\"type\":\"string\"}],\"name\":\"PinSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getPin\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newPin\",\"type\":\"string\"}],\"name\":\"setPin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PinStorageBin is the compiled bytecode used for deploying new contracts.
var PinStorageBin = "0x60806040523480156200001157600080fd5b5060405162000b2e38038062000b2e83398181016040528101906200003791906200031e565b80600090805190602001906200004f929190620000d1565b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f2618cc7c17807bca48ed015a9f3e204f5e7d65d25589705ab2b03c580236f28181604051620000c29190620003cc565b60405180910390a15062000455565b828054620000df906200041f565b90600052602060002090601f0160209004810192826200010357600085556200014f565b82601f106200011e57805160ff19168380011785556200014f565b828001600101855582156200014f579182015b828111156200014e57825182559160200191906001019062000131565b5b5090506200015e919062000162565b5090565b5b808211156200017d57600081600090555060010162000163565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620001ea826200019f565b810181811067ffffffffffffffff821117156200020c576200020b620001b0565b5b80604052505050565b60006200022162000181565b90506200022f8282620001df565b919050565b600067ffffffffffffffff821115620002525762000251620001b0565b5b6200025d826200019f565b9050602081019050919050565b60005b838110156200028a5780820151818401526020810190506200026d565b838111156200029a576000848401525b50505050565b6000620002b7620002b18462000234565b62000215565b905082815260208101848484011115620002d657620002d56200019a565b5b620002e38482856200026a565b509392505050565b600082601f83011262000303576200030262000195565b5b815162000315848260208601620002a0565b91505092915050565b6000602082840312156200033757620003366200018b565b5b600082015167ffffffffffffffff81111562000358576200035762000190565b5b6200036684828501620002eb565b91505092915050565b600081519050919050565b600082825260208201905092915050565b600062000398826200036f565b620003a481856200037a565b9350620003b68185602086016200026a565b620003c1816200019f565b840191505092915050565b60006020820190508181036000830152620003e881846200038b565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200043857607f821691505b602082108114156200044f576200044e620003f0565b5b50919050565b6106c980620004656000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063bec93fb11461003b578063c8c4b43d14610057575b600080fd5b61005560048036038101906100509190610476565b610075565b005b61005f6101e7565b60405161006c9190610547565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610105576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100fc906105db565b60405180910390fd5b60008080546101139061062a565b80601f016020809104026020016040519081016040528092919081815260200182805461013f9061062a565b801561018c5780601f106101615761010080835404028352916020019161018c565b820191906000526020600020905b81548152906001019060200180831161016f57829003601f168201915b5050505050905081600090805190602001906101a9929190610279565b507f667bd9d7575800a134d097034305e4d9bf371817ab26cba54239ad73ff1923f581836040516101db92919061065c565b60405180910390a15050565b6060600080546101f69061062a565b80601f01602080910402602001604051908101604052809291908181526020018280546102229061062a565b801561026f5780601f106102445761010080835404028352916020019161026f565b820191906000526020600020905b81548152906001019060200180831161025257829003601f168201915b5050505050905090565b8280546102859061062a565b90600052602060002090601f0160209004810192826102a757600085556102ee565b82601f106102c057805160ff19168380011785556102ee565b828001600101855582156102ee579182015b828111156102ed5782518255916020019190600101906102d2565b5b5090506102fb91906102ff565b5090565b5b80821115610318576000816000905550600101610300565b5090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6103838261033a565b810181811067ffffffffffffffff821117156103a2576103a161034b565b5b80604052505050565b60006103b561031c565b90506103c1828261037a565b919050565b600067ffffffffffffffff8211156103e1576103e061034b565b5b6103ea8261033a565b9050602081019050919050565b82818337600083830152505050565b6000610419610414846103c6565b6103ab565b90508281526020810184848401111561043557610434610335565b5b6104408482856103f7565b509392505050565b600082601f83011261045d5761045c610330565b5b813561046d848260208601610406565b91505092915050565b60006020828403121561048c5761048b610326565b5b600082013567ffffffffffffffff8111156104aa576104a961032b565b5b6104b684828501610448565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156104f95780820151818401526020810190506104de565b83811115610508576000848401525b50505050565b6000610519826104bf565b61052381856104ca565b93506105338185602086016104db565b61053c8161033a565b840191505092915050565b60006020820190508181036000830152610561818461050e565b905092915050565b7f4f6e6c7920746865206f776e65722063616e206368616e67652074686520504960008201527f4e2e000000000000000000000000000000000000000000000000000000000000602082015250565b60006105c56022836104ca565b91506105d082610569565b604082019050919050565b600060208201905081810360008301526105f4816105b8565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061064257607f821691505b60208210811415610656576106556105fb565b5b50919050565b60006040820190508181036000830152610676818561050e565b9050818103602083015261068a818461050e565b9050939250505056fea264697066735822122034ddf3a28d997da05f44a685e1b827be7680e90bd4d8f399815d12b71d3e4bf164736f6c634300080b0033"
var PinStorageSMBin = "0x"

// DeployPinStorage deploys a new contract, binding an instance of PinStorage to it.
func DeployPinStorage(auth *bind.TransactOpts, backend bind.ContractBackend, _initialPin string) (common.Address, *types.Receipt, *PinStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(PinStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(PinStorageSMBin)
	} else {
		bytecode = common.FromHex(PinStorageBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, PinStorageABI, backend, _initialPin)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &PinStorage{PinStorageCaller: PinStorageCaller{contract: contract}, PinStorageTransactor: PinStorageTransactor{contract: contract}, PinStorageFilterer: PinStorageFilterer{contract: contract}}, nil
}

func AsyncDeployPinStorage(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _initialPin string) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(PinStorageABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(PinStorageSMBin)
	} else {
		bytecode = common.FromHex(PinStorageBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, PinStorageABI, backend, _initialPin)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// PinStorage is an auto generated Go binding around a Solidity contract.
type PinStorage struct {
	PinStorageCaller     // Read-only binding to the contract
	PinStorageTransactor // Write-only binding to the contract
	PinStorageFilterer   // Log filterer for contract events
}

// PinStorageCaller is an auto generated read-only Go binding around a Solidity contract.
type PinStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinStorageTransactor is an auto generated write-only Go binding around a Solidity contract.
type PinStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinStorageFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type PinStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinStorageSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type PinStorageSession struct {
	Contract     *PinStorage       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PinStorageCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type PinStorageCallerSession struct {
	Contract *PinStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PinStorageTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type PinStorageTransactorSession struct {
	Contract     *PinStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PinStorageRaw is an auto generated low-level Go binding around a Solidity contract.
type PinStorageRaw struct {
	Contract *PinStorage // Generic contract binding to access the raw methods on
}

// PinStorageCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type PinStorageCallerRaw struct {
	Contract *PinStorageCaller // Generic read-only contract binding to access the raw methods on
}

// PinStorageTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type PinStorageTransactorRaw struct {
	Contract *PinStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPinStorage creates a new instance of PinStorage, bound to a specific deployed contract.
func NewPinStorage(address common.Address, backend bind.ContractBackend) (*PinStorage, error) {
	contract, err := bindPinStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PinStorage{PinStorageCaller: PinStorageCaller{contract: contract}, PinStorageTransactor: PinStorageTransactor{contract: contract}, PinStorageFilterer: PinStorageFilterer{contract: contract}}, nil
}

// NewPinStorageCaller creates a new read-only instance of PinStorage, bound to a specific deployed contract.
func NewPinStorageCaller(address common.Address, caller bind.ContractCaller) (*PinStorageCaller, error) {
	contract, err := bindPinStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PinStorageCaller{contract: contract}, nil
}

// NewPinStorageTransactor creates a new write-only instance of PinStorage, bound to a specific deployed contract.
func NewPinStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*PinStorageTransactor, error) {
	contract, err := bindPinStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PinStorageTransactor{contract: contract}, nil
}

// NewPinStorageFilterer creates a new log filterer instance of PinStorage, bound to a specific deployed contract.
func NewPinStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*PinStorageFilterer, error) {
	contract, err := bindPinStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PinStorageFilterer{contract: contract}, nil
}

// bindPinStorage binds a generic wrapper to an already deployed contract.
func bindPinStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PinStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PinStorage *PinStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PinStorage.Contract.PinStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PinStorage *PinStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _PinStorage.Contract.PinStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PinStorage *PinStorageRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _PinStorage.Contract.PinStorageTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PinStorage *PinStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PinStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PinStorage *PinStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _PinStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PinStorage *PinStorageTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _PinStorage.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// GetPin is a free data retrieval call binding the contract method 0xc8c4b43d.
//
// Solidity: function getPin() constant returns(string)
func (_PinStorage *PinStorageCaller) GetPin(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PinStorage.contract.Call(opts, out, "getPin")
	return *ret0, err
}

// GetPin is a free data retrieval call binding the contract method 0xc8c4b43d.
//
// Solidity: function getPin() constant returns(string)
func (_PinStorage *PinStorageSession) GetPin() (string, error) {
	return _PinStorage.Contract.GetPin(&_PinStorage.CallOpts)
}

// GetPin is a free data retrieval call binding the contract method 0xc8c4b43d.
//
// Solidity: function getPin() constant returns(string)
func (_PinStorage *PinStorageCallerSession) GetPin() (string, error) {
	return _PinStorage.Contract.GetPin(&_PinStorage.CallOpts)
}

// SetPin is a paid mutator transaction binding the contract method 0xbec93fb1.
//
// Solidity: function setPin(string _newPin) returns()
func (_PinStorage *PinStorageTransactor) SetPin(opts *bind.TransactOpts, _newPin string) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _PinStorage.contract.TransactWithResult(opts, out, "setPin", _newPin)
	return transaction, receipt, err
}

func (_PinStorage *PinStorageTransactor) AsyncSetPin(handler func(*types.Receipt, error), opts *bind.TransactOpts, _newPin string) (*types.Transaction, error) {
	return _PinStorage.contract.AsyncTransact(opts, handler, "setPin", _newPin)
}

// SetPin is a paid mutator transaction binding the contract method 0xbec93fb1.
//
// Solidity: function setPin(string _newPin) returns()
func (_PinStorage *PinStorageSession) SetPin(_newPin string) (*types.Transaction, *types.Receipt, error) {
	return _PinStorage.Contract.SetPin(&_PinStorage.TransactOpts, _newPin)
}

func (_PinStorage *PinStorageSession) AsyncSetPin(handler func(*types.Receipt, error), _newPin string) (*types.Transaction, error) {
	return _PinStorage.Contract.AsyncSetPin(handler, &_PinStorage.TransactOpts, _newPin)
}

// SetPin is a paid mutator transaction binding the contract method 0xbec93fb1.
//
// Solidity: function setPin(string _newPin) returns()
func (_PinStorage *PinStorageTransactorSession) SetPin(_newPin string) (*types.Transaction, *types.Receipt, error) {
	return _PinStorage.Contract.SetPin(&_PinStorage.TransactOpts, _newPin)
}

func (_PinStorage *PinStorageTransactorSession) AsyncSetPin(handler func(*types.Receipt, error), _newPin string) (*types.Transaction, error) {
	return _PinStorage.Contract.AsyncSetPin(handler, &_PinStorage.TransactOpts, _newPin)
}

// PinStoragePinChanged represents a PinChanged event raised by the PinStorage contract.
type PinStoragePinChanged struct {
	OldPin string
	NewPin string
	Raw    types.Log // Blockchain specific contextual infos
}

// WatchPinChanged is a free log subscription operation binding the contract event 0x667bd9d7575800a134d097034305e4d9bf371817ab26cba54239ad73ff1923f5.
//
// Solidity: event PinChanged(string oldPin, string newPin)
func (_PinStorage *PinStorageFilterer) WatchPinChanged(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _PinStorage.contract.WatchLogs(fromBlock, handler, "PinChanged")
}

func (_PinStorage *PinStorageFilterer) WatchAllPinChanged(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _PinStorage.contract.WatchLogs(fromBlock, handler, "PinChanged")
}

// ParsePinChanged is a log parse operation binding the contract event 0x667bd9d7575800a134d097034305e4d9bf371817ab26cba54239ad73ff1923f5.
//
// Solidity: event PinChanged(string oldPin, string newPin)
func (_PinStorage *PinStorageFilterer) ParsePinChanged(log types.Log) (*PinStoragePinChanged, error) {
	event := new(PinStoragePinChanged)
	if err := _PinStorage.contract.UnpackLog(event, "PinChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchPinChanged is a free log subscription operation binding the contract event 0x667bd9d7575800a134d097034305e4d9bf371817ab26cba54239ad73ff1923f5.
//
// Solidity: event PinChanged(string oldPin, string newPin)
func (_PinStorage *PinStorageSession) WatchPinChanged(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _PinStorage.Contract.WatchPinChanged(fromBlock, handler)
}

func (_PinStorage *PinStorageSession) WatchAllPinChanged(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _PinStorage.Contract.WatchAllPinChanged(fromBlock, handler)
}

// ParsePinChanged is a log parse operation binding the contract event 0x667bd9d7575800a134d097034305e4d9bf371817ab26cba54239ad73ff1923f5.
//
// Solidity: event PinChanged(string oldPin, string newPin)
func (_PinStorage *PinStorageSession) ParsePinChanged(log types.Log) (*PinStoragePinChanged, error) {
	return _PinStorage.Contract.ParsePinChanged(log)
}

// PinStoragePinSet represents a PinSet event raised by the PinStorage contract.
type PinStoragePinSet struct {
	Pin string
	Raw types.Log // Blockchain specific contextual infos
}

// WatchPinSet is a free log subscription operation binding the contract event 0x2618cc7c17807bca48ed015a9f3e204f5e7d65d25589705ab2b03c580236f281.
//
// Solidity: event PinSet(string pin)
func (_PinStorage *PinStorageFilterer) WatchPinSet(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _PinStorage.contract.WatchLogs(fromBlock, handler, "PinSet")
}

func (_PinStorage *PinStorageFilterer) WatchAllPinSet(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _PinStorage.contract.WatchLogs(fromBlock, handler, "PinSet")
}

// ParsePinSet is a log parse operation binding the contract event 0x2618cc7c17807bca48ed015a9f3e204f5e7d65d25589705ab2b03c580236f281.
//
// Solidity: event PinSet(string pin)
func (_PinStorage *PinStorageFilterer) ParsePinSet(log types.Log) (*PinStoragePinSet, error) {
	event := new(PinStoragePinSet)
	if err := _PinStorage.contract.UnpackLog(event, "PinSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchPinSet is a free log subscription operation binding the contract event 0x2618cc7c17807bca48ed015a9f3e204f5e7d65d25589705ab2b03c580236f281.
//
// Solidity: event PinSet(string pin)
func (_PinStorage *PinStorageSession) WatchPinSet(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _PinStorage.Contract.WatchPinSet(fromBlock, handler)
}

func (_PinStorage *PinStorageSession) WatchAllPinSet(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _PinStorage.Contract.WatchAllPinSet(fromBlock, handler)
}

// ParsePinSet is a log parse operation binding the contract event 0x2618cc7c17807bca48ed015a9f3e204f5e7d65d25589705ab2b03c580236f281.
//
// Solidity: event PinSet(string pin)
func (_PinStorage *PinStorageSession) ParsePinSet(log types.Log) (*PinStoragePinSet, error) {
	return _PinStorage.Contract.ParsePinSet(log)
}
