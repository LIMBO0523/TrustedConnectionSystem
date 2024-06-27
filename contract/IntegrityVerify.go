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

// IntegrityVerifyABI is the input ABI used to generate the binding from.
const IntegrityVerifyABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"sigArray\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"senderPublicKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"hashArray\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"HashArrayStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"measureFinished\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToBoolMap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"addressToPublicKeyMap\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"fileToHashMap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAddressToPublicKey\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isAddressTrue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"value\",\"type\":\"bool\"}],\"name\":\"setAddressToBool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"setAddressToPublicKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"values\",\"type\":\"bytes32[]\"}],\"name\":\"setMultipleHashes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"keys\",\"type\":\"string[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"values\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"hashArrayToStore\",\"type\":\"bytes32[]\"}],\"name\":\"verifyAndStoreHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"verifyHash\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IntegrityVerifyBin is the compiled bytecode used for deploying new contracts.
var IntegrityVerifyBin = "0x608060405234801561001057600080fd5b5061174c806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c806397dbbca01161007157806397dbbca01461019d5780639858a4c6146101b9578063a320839c146101e9578063b89207e014610219578063d223f17f14610249578063f042b69714610279576100b4565b806309ea9251146100b95780635b6beeb9146100d55780636c521237146101055780637850d93a1461012157806385ec0c57146101515780638eaf72441461016d575b600080fd5b6100d360048036038101906100ce9190610b21565b6102a9565b005b6100ef60048036038101906100ea9190610ca7565b610330565b6040516100fc9190610d09565b60405180910390f35b61011f600480360381019061011a9190610dc5565b610357565b005b61013b60048036038101906101369190610e21565b6103af565b6040516101489190610ed6565b60405180910390f35b61016b60048036038101906101669190610f24565b61044f565b005b61018760048036038101906101829190610e21565b610476565b6040516101949190610f8f565b60405180910390f35b6101b760048036038101906101b29190611153565b610496565b005b6101d360048036038101906101ce9190610ca7565b610555565b6040516101e09190610d09565b60405180910390f35b61020360048036038101906101fe9190610e21565b610583565b6040516102109190610ed6565b60405180910390f35b610233600480360381019061022e9190610f24565b610654565b6040516102409190610f8f565b60405180910390f35b610263600480360381019061025e9190610e21565b610684565b6040516102709190610f8f565b60405180910390f35b610293600480360381019061028e91906111cb565b6106da565b6040516102a09190610f8f565b60405180910390f35b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fb9731c194d904fe4f70b3d61d77bd55fda2e9e934ac8b7cc16d68cb7af85604e60405160405180910390a15050565b6000808260405161034191906112b9565b9081526020016040518091039020549050919050565b80600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090805190602001906103aa9291906109d4565b505050565b600160205280600052604060002060009150905080546103ce906112ff565b80601f01602080910402602001604051908101604052809291908181526020018280546103fa906112ff565b80156104475780601f1061041c57610100808354040283529160200191610447565b820191906000526020600020905b81548152906001019060200180831161042a57829003601f168201915b505050505081565b8060008360405161046091906112b9565b9081526020016040518091039020819055505050565b60026020528060005260406000206000915054906101000a900460ff1681565b80518251146104da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d1906113b4565b60405180910390fd5b60005b8251811015610550578181815181106104f9576104f86113d4565b5b60200260200101516000848381518110610516576105156113d4565b5b602002602001015160405161052b91906112b9565b90815260200160405180910390208190555080806105489061143c565b9150506104dd565b505050565b6000818051602081018201805184825260208301602085012081835280955050505050506000915090505481565b6060600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080546105cf906112ff565b80601f01602080910402602001604051908101604052809291908181526020018280546105fb906112ff565b80156106485780601f1061061d57610100808354040283529160200191610648565b820191906000526020600020905b81548152906001019060200180831161062b57829003601f168201915b50505050509050919050565b60008060008460405161066791906112b9565b908152602001604051809103902054905082811491505092915050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b60008251845114610720576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610717906114f7565b60405180910390fd5b600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690556001600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060005b845181101561087a578381815181106107e6576107e56113d4565b5b60200260200101516000868381518110610803576108026113d4565b5b602002602001015160405161081891906112b9565b90815260200160405180910390205414610867576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085e90611563565b60405180910390fd5b80806108729061143c565b9150506107ca565b506000600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080546108c7906112ff565b80601f01602080910402602001604051908101604052809291908181526020018280546108f3906112ff565b80156109405780601f1061091557610100808354040283529160200191610940565b820191906000526020600020905b81548152906001019060200180831161092357829003601f168201915b50505050509050600081511161098b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610982906115cf565b60405180910390fd5b7fdde8f1c4b3994910d6f2775c2ae898bce9327a67a2c04a711b584e951d76828c838286336040516109c094939291906116bc565b60405180910390a160019150509392505050565b8280546109e0906112ff565b90600052602060002090601f016020900481019282610a025760008555610a49565b82601f10610a1b57805160ff1916838001178555610a49565b82800160010185558215610a49579182015b82811115610a48578251825591602001919060010190610a2d565b5b509050610a569190610a5a565b5090565b5b80821115610a73576000816000905550600101610a5b565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ab682610a8b565b9050919050565b610ac681610aab565b8114610ad157600080fd5b50565b600081359050610ae381610abd565b92915050565b60008115159050919050565b610afe81610ae9565b8114610b0957600080fd5b50565b600081359050610b1b81610af5565b92915050565b60008060408385031215610b3857610b37610a81565b5b6000610b4685828601610ad4565b9250506020610b5785828601610b0c565b9150509250929050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610bb482610b6b565b810181811067ffffffffffffffff82111715610bd357610bd2610b7c565b5b80604052505050565b6000610be6610a77565b9050610bf28282610bab565b919050565b600067ffffffffffffffff821115610c1257610c11610b7c565b5b610c1b82610b6b565b9050602081019050919050565b82818337600083830152505050565b6000610c4a610c4584610bf7565b610bdc565b905082815260208101848484011115610c6657610c65610b66565b5b610c71848285610c28565b509392505050565b600082601f830112610c8e57610c8d610b61565b5b8135610c9e848260208601610c37565b91505092915050565b600060208284031215610cbd57610cbc610a81565b5b600082013567ffffffffffffffff811115610cdb57610cda610a86565b5b610ce784828501610c79565b91505092915050565b6000819050919050565b610d0381610cf0565b82525050565b6000602082019050610d1e6000830184610cfa565b92915050565b600067ffffffffffffffff821115610d3f57610d3e610b7c565b5b610d4882610b6b565b9050602081019050919050565b6000610d68610d6384610d24565b610bdc565b905082815260208101848484011115610d8457610d83610b66565b5b610d8f848285610c28565b509392505050565b600082601f830112610dac57610dab610b61565b5b8135610dbc848260208601610d55565b91505092915050565b60008060408385031215610ddc57610ddb610a81565b5b6000610dea85828601610ad4565b925050602083013567ffffffffffffffff811115610e0b57610e0a610a86565b5b610e1785828601610d97565b9150509250929050565b600060208284031215610e3757610e36610a81565b5b6000610e4584828501610ad4565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610e88578082015181840152602081019050610e6d565b83811115610e97576000848401525b50505050565b6000610ea882610e4e565b610eb28185610e59565b9350610ec2818560208601610e6a565b610ecb81610b6b565b840191505092915050565b60006020820190508181036000830152610ef08184610e9d565b905092915050565b610f0181610cf0565b8114610f0c57600080fd5b50565b600081359050610f1e81610ef8565b92915050565b60008060408385031215610f3b57610f3a610a81565b5b600083013567ffffffffffffffff811115610f5957610f58610a86565b5b610f6585828601610c79565b9250506020610f7685828601610f0f565b9150509250929050565b610f8981610ae9565b82525050565b6000602082019050610fa46000830184610f80565b92915050565b600067ffffffffffffffff821115610fc557610fc4610b7c565b5b602082029050602081019050919050565b600080fd5b6000610fee610fe984610faa565b610bdc565b9050808382526020820190506020840283018581111561101157611010610fd6565b5b835b8181101561105857803567ffffffffffffffff81111561103657611035610b61565b5b8086016110438982610c79565b85526020850194505050602081019050611013565b5050509392505050565b600082601f83011261107757611076610b61565b5b8135611087848260208601610fdb565b91505092915050565b600067ffffffffffffffff8211156110ab576110aa610b7c565b5b602082029050602081019050919050565b60006110cf6110ca84611090565b610bdc565b905080838252602082019050602084028301858111156110f2576110f1610fd6565b5b835b8181101561111b57806111078882610f0f565b8452602084019350506020810190506110f4565b5050509392505050565b600082601f83011261113a57611139610b61565b5b813561114a8482602086016110bc565b91505092915050565b6000806040838503121561116a57611169610a81565b5b600083013567ffffffffffffffff81111561118857611187610a86565b5b61119485828601611062565b925050602083013567ffffffffffffffff8111156111b5576111b4610a86565b5b6111c185828601611125565b9150509250929050565b6000806000606084860312156111e4576111e3610a81565b5b600084013567ffffffffffffffff81111561120257611201610a86565b5b61120e86828701611062565b935050602084013567ffffffffffffffff81111561122f5761122e610a86565b5b61123b86828701611125565b925050604084013567ffffffffffffffff81111561125c5761125b610a86565b5b61126886828701611125565b9150509250925092565b600081519050919050565b600081905092915050565b600061129382611272565b61129d818561127d565b93506112ad818560208601610e6a565b80840191505092915050565b60006112c58284611288565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061131757607f821691505b6020821081141561132b5761132a6112d0565b5b50919050565b600082825260208201905092915050565b7f4e756d626572206f66206b657973206d757374206d61746368206e756d62657260008201527f206f662076616c75657300000000000000000000000000000000000000000000602082015250565b600061139e602a83611331565b91506113a982611342565b604082019050919050565b600060208201905081810360008301526113cd81611391565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b600061144782611432565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561147a57611479611403565b5b600182019050919050565b7f4b65797320616e642076616c756573206172726179206c656e677468206d757360008201527f74206d6174636800000000000000000000000000000000000000000000000000602082015250565b60006114e1602783611331565b91506114ec82611485565b604082019050919050565b60006020820190508181036000830152611510816114d4565b9050919050565b7f4b657920616e642076616c7565206d69736d6174636800000000000000000000600082015250565b600061154d601683611331565b915061155882611517565b602082019050919050565b6000602082019050818103600083015261157c81611540565b9050919050565b7f53656e646572207075626c6963206b6579206e6f742073657400000000000000600082015250565b60006115b9601983611331565b91506115c482611583565b602082019050919050565b600060208201905081810360008301526115e8816115ac565b9050919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61162481610cf0565b82525050565b6000611636838361161b565b60208301905092915050565b6000602082019050919050565b600061165a826115ef565b61166481856115fa565b935061166f8361160b565b8060005b838110156116a0578151611687888261162a565b975061169283611642565b925050600181019050611673565b5085935050505092915050565b6116b681610aab565b82525050565b600060808201905081810360008301526116d6818761164f565b905081810360208301526116ea8186610e9d565b905081810360408301526116fe818561164f565b905061170d60608301846116ad565b9594505050505056fea2646970667358221220a6c0d18131e87d11867236874e371e538010ee2e81e8653b97a4e04435e1108964736f6c634300080b0033"
var IntegrityVerifySMBin = "0x"

// DeployIntegrityVerify deploys a new contract, binding an instance of IntegrityVerify to it.
func DeployIntegrityVerify(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Receipt, *IntegrityVerify, error) {
	parsed, err := abi.JSON(strings.NewReader(IntegrityVerifyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(IntegrityVerifySMBin)
	} else {
		bytecode = common.FromHex(IntegrityVerifyBin)
	}
	if len(bytecode) == 0 {
		return common.Address{}, nil, nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	address, receipt, contract, err := bind.DeployContract(auth, parsed, bytecode, IntegrityVerifyABI, backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, receipt, &IntegrityVerify{IntegrityVerifyCaller: IntegrityVerifyCaller{contract: contract}, IntegrityVerifyTransactor: IntegrityVerifyTransactor{contract: contract}, IntegrityVerifyFilterer: IntegrityVerifyFilterer{contract: contract}}, nil
}

func AsyncDeployIntegrityVerify(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(IntegrityVerifyABI))
	if err != nil {
		return nil, err
	}

	var bytecode []byte
	if backend.SMCrypto() {
		bytecode = common.FromHex(IntegrityVerifySMBin)
	} else {
		bytecode = common.FromHex(IntegrityVerifyBin)
	}
	if len(bytecode) == 0 {
		return nil, fmt.Errorf("cannot deploy empty bytecode")
	}
	tx, err := bind.AsyncDeployContract(auth, handler, parsed, bytecode, IntegrityVerifyABI, backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// IntegrityVerify is an auto generated Go binding around a Solidity contract.
type IntegrityVerify struct {
	IntegrityVerifyCaller     // Read-only binding to the contract
	IntegrityVerifyTransactor // Write-only binding to the contract
	IntegrityVerifyFilterer   // Log filterer for contract events
}

// IntegrityVerifyCaller is an auto generated read-only Go binding around a Solidity contract.
type IntegrityVerifyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntegrityVerifyTransactor is an auto generated write-only Go binding around a Solidity contract.
type IntegrityVerifyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntegrityVerifyFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IntegrityVerifyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IntegrityVerifySession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IntegrityVerifySession struct {
	Contract     *IntegrityVerify  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IntegrityVerifyCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IntegrityVerifyCallerSession struct {
	Contract *IntegrityVerifyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IntegrityVerifyTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IntegrityVerifyTransactorSession struct {
	Contract     *IntegrityVerifyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IntegrityVerifyRaw is an auto generated low-level Go binding around a Solidity contract.
type IntegrityVerifyRaw struct {
	Contract *IntegrityVerify // Generic contract binding to access the raw methods on
}

// IntegrityVerifyCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IntegrityVerifyCallerRaw struct {
	Contract *IntegrityVerifyCaller // Generic read-only contract binding to access the raw methods on
}

// IntegrityVerifyTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IntegrityVerifyTransactorRaw struct {
	Contract *IntegrityVerifyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIntegrityVerify creates a new instance of IntegrityVerify, bound to a specific deployed contract.
func NewIntegrityVerify(address common.Address, backend bind.ContractBackend) (*IntegrityVerify, error) {
	contract, err := bindIntegrityVerify(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IntegrityVerify{IntegrityVerifyCaller: IntegrityVerifyCaller{contract: contract}, IntegrityVerifyTransactor: IntegrityVerifyTransactor{contract: contract}, IntegrityVerifyFilterer: IntegrityVerifyFilterer{contract: contract}}, nil
}

// NewIntegrityVerifyCaller creates a new read-only instance of IntegrityVerify, bound to a specific deployed contract.
func NewIntegrityVerifyCaller(address common.Address, caller bind.ContractCaller) (*IntegrityVerifyCaller, error) {
	contract, err := bindIntegrityVerify(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IntegrityVerifyCaller{contract: contract}, nil
}

// NewIntegrityVerifyTransactor creates a new write-only instance of IntegrityVerify, bound to a specific deployed contract.
func NewIntegrityVerifyTransactor(address common.Address, transactor bind.ContractTransactor) (*IntegrityVerifyTransactor, error) {
	contract, err := bindIntegrityVerify(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IntegrityVerifyTransactor{contract: contract}, nil
}

// NewIntegrityVerifyFilterer creates a new log filterer instance of IntegrityVerify, bound to a specific deployed contract.
func NewIntegrityVerifyFilterer(address common.Address, filterer bind.ContractFilterer) (*IntegrityVerifyFilterer, error) {
	contract, err := bindIntegrityVerify(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IntegrityVerifyFilterer{contract: contract}, nil
}

// bindIntegrityVerify binds a generic wrapper to an already deployed contract.
func bindIntegrityVerify(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IntegrityVerifyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IntegrityVerify *IntegrityVerifyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IntegrityVerify.Contract.IntegrityVerifyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IntegrityVerify *IntegrityVerifyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.IntegrityVerifyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IntegrityVerify *IntegrityVerifyRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.IntegrityVerifyTransactor.contract.TransactWithResult(opts, result, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IntegrityVerify *IntegrityVerifyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IntegrityVerify.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IntegrityVerify *IntegrityVerifyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IntegrityVerify *IntegrityVerifyTransactorRaw) TransactWithResult(opts *bind.TransactOpts, result interface{}, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.contract.TransactWithResult(opts, result, method, params...)
}

// AddressToBoolMap is a free data retrieval call binding the contract method 0x8eaf7244.
//
// Solidity: function addressToBoolMap(address ) constant returns(bool)
func (_IntegrityVerify *IntegrityVerifyCaller) AddressToBoolMap(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IntegrityVerify.contract.Call(opts, out, "addressToBoolMap", arg0)
	return *ret0, err
}

// AddressToBoolMap is a free data retrieval call binding the contract method 0x8eaf7244.
//
// Solidity: function addressToBoolMap(address ) constant returns(bool)
func (_IntegrityVerify *IntegrityVerifySession) AddressToBoolMap(arg0 common.Address) (bool, error) {
	return _IntegrityVerify.Contract.AddressToBoolMap(&_IntegrityVerify.CallOpts, arg0)
}

// AddressToBoolMap is a free data retrieval call binding the contract method 0x8eaf7244.
//
// Solidity: function addressToBoolMap(address ) constant returns(bool)
func (_IntegrityVerify *IntegrityVerifyCallerSession) AddressToBoolMap(arg0 common.Address) (bool, error) {
	return _IntegrityVerify.Contract.AddressToBoolMap(&_IntegrityVerify.CallOpts, arg0)
}

// AddressToPublicKeyMap is a free data retrieval call binding the contract method 0x7850d93a.
//
// Solidity: function addressToPublicKeyMap(address ) constant returns(bytes)
func (_IntegrityVerify *IntegrityVerifyCaller) AddressToPublicKeyMap(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IntegrityVerify.contract.Call(opts, out, "addressToPublicKeyMap", arg0)
	return *ret0, err
}

// AddressToPublicKeyMap is a free data retrieval call binding the contract method 0x7850d93a.
//
// Solidity: function addressToPublicKeyMap(address ) constant returns(bytes)
func (_IntegrityVerify *IntegrityVerifySession) AddressToPublicKeyMap(arg0 common.Address) ([]byte, error) {
	return _IntegrityVerify.Contract.AddressToPublicKeyMap(&_IntegrityVerify.CallOpts, arg0)
}

// AddressToPublicKeyMap is a free data retrieval call binding the contract method 0x7850d93a.
//
// Solidity: function addressToPublicKeyMap(address ) constant returns(bytes)
func (_IntegrityVerify *IntegrityVerifyCallerSession) AddressToPublicKeyMap(arg0 common.Address) ([]byte, error) {
	return _IntegrityVerify.Contract.AddressToPublicKeyMap(&_IntegrityVerify.CallOpts, arg0)
}

// FileToHashMap is a free data retrieval call binding the contract method 0x9858a4c6.
//
// Solidity: function fileToHashMap(string ) constant returns(bytes32)
func (_IntegrityVerify *IntegrityVerifyCaller) FileToHashMap(opts *bind.CallOpts, arg0 string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IntegrityVerify.contract.Call(opts, out, "fileToHashMap", arg0)
	return *ret0, err
}

// FileToHashMap is a free data retrieval call binding the contract method 0x9858a4c6.
//
// Solidity: function fileToHashMap(string ) constant returns(bytes32)
func (_IntegrityVerify *IntegrityVerifySession) FileToHashMap(arg0 string) ([32]byte, error) {
	return _IntegrityVerify.Contract.FileToHashMap(&_IntegrityVerify.CallOpts, arg0)
}

// FileToHashMap is a free data retrieval call binding the contract method 0x9858a4c6.
//
// Solidity: function fileToHashMap(string ) constant returns(bytes32)
func (_IntegrityVerify *IntegrityVerifyCallerSession) FileToHashMap(arg0 string) ([32]byte, error) {
	return _IntegrityVerify.Contract.FileToHashMap(&_IntegrityVerify.CallOpts, arg0)
}

// GetAddressToPublicKey is a free data retrieval call binding the contract method 0xa320839c.
//
// Solidity: function getAddressToPublicKey(address addr) constant returns(bytes)
func (_IntegrityVerify *IntegrityVerifyCaller) GetAddressToPublicKey(opts *bind.CallOpts, addr common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _IntegrityVerify.contract.Call(opts, out, "getAddressToPublicKey", addr)
	return *ret0, err
}

// GetAddressToPublicKey is a free data retrieval call binding the contract method 0xa320839c.
//
// Solidity: function getAddressToPublicKey(address addr) constant returns(bytes)
func (_IntegrityVerify *IntegrityVerifySession) GetAddressToPublicKey(addr common.Address) ([]byte, error) {
	return _IntegrityVerify.Contract.GetAddressToPublicKey(&_IntegrityVerify.CallOpts, addr)
}

// GetAddressToPublicKey is a free data retrieval call binding the contract method 0xa320839c.
//
// Solidity: function getAddressToPublicKey(address addr) constant returns(bytes)
func (_IntegrityVerify *IntegrityVerifyCallerSession) GetAddressToPublicKey(addr common.Address) ([]byte, error) {
	return _IntegrityVerify.Contract.GetAddressToPublicKey(&_IntegrityVerify.CallOpts, addr)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string key) constant returns(bytes32)
func (_IntegrityVerify *IntegrityVerifyCaller) GetHash(opts *bind.CallOpts, key string) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _IntegrityVerify.contract.Call(opts, out, "getHash", key)
	return *ret0, err
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string key) constant returns(bytes32)
func (_IntegrityVerify *IntegrityVerifySession) GetHash(key string) ([32]byte, error) {
	return _IntegrityVerify.Contract.GetHash(&_IntegrityVerify.CallOpts, key)
}

// GetHash is a free data retrieval call binding the contract method 0x5b6beeb9.
//
// Solidity: function getHash(string key) constant returns(bytes32)
func (_IntegrityVerify *IntegrityVerifyCallerSession) GetHash(key string) ([32]byte, error) {
	return _IntegrityVerify.Contract.GetHash(&_IntegrityVerify.CallOpts, key)
}

// IsAddressTrue is a free data retrieval call binding the contract method 0xd223f17f.
//
// Solidity: function isAddressTrue(address addr) constant returns(bool)
func (_IntegrityVerify *IntegrityVerifyCaller) IsAddressTrue(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IntegrityVerify.contract.Call(opts, out, "isAddressTrue", addr)
	return *ret0, err
}

// IsAddressTrue is a free data retrieval call binding the contract method 0xd223f17f.
//
// Solidity: function isAddressTrue(address addr) constant returns(bool)
func (_IntegrityVerify *IntegrityVerifySession) IsAddressTrue(addr common.Address) (bool, error) {
	return _IntegrityVerify.Contract.IsAddressTrue(&_IntegrityVerify.CallOpts, addr)
}

// IsAddressTrue is a free data retrieval call binding the contract method 0xd223f17f.
//
// Solidity: function isAddressTrue(address addr) constant returns(bool)
func (_IntegrityVerify *IntegrityVerifyCallerSession) IsAddressTrue(addr common.Address) (bool, error) {
	return _IntegrityVerify.Contract.IsAddressTrue(&_IntegrityVerify.CallOpts, addr)
}

// VerifyHash is a free data retrieval call binding the contract method 0xb89207e0.
//
// Solidity: function verifyHash(string key, bytes32 value) constant returns(bool)
func (_IntegrityVerify *IntegrityVerifyCaller) VerifyHash(opts *bind.CallOpts, key string, value [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IntegrityVerify.contract.Call(opts, out, "verifyHash", key, value)
	return *ret0, err
}

// VerifyHash is a free data retrieval call binding the contract method 0xb89207e0.
//
// Solidity: function verifyHash(string key, bytes32 value) constant returns(bool)
func (_IntegrityVerify *IntegrityVerifySession) VerifyHash(key string, value [32]byte) (bool, error) {
	return _IntegrityVerify.Contract.VerifyHash(&_IntegrityVerify.CallOpts, key, value)
}

// VerifyHash is a free data retrieval call binding the contract method 0xb89207e0.
//
// Solidity: function verifyHash(string key, bytes32 value) constant returns(bool)
func (_IntegrityVerify *IntegrityVerifyCallerSession) VerifyHash(key string, value [32]byte) (bool, error) {
	return _IntegrityVerify.Contract.VerifyHash(&_IntegrityVerify.CallOpts, key, value)
}

// SetAddressToBool is a paid mutator transaction binding the contract method 0x09ea9251.
//
// Solidity: function setAddressToBool(address addr, bool value) returns()
func (_IntegrityVerify *IntegrityVerifyTransactor) SetAddressToBool(opts *bind.TransactOpts, addr common.Address, value bool) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _IntegrityVerify.contract.TransactWithResult(opts, out, "setAddressToBool", addr, value)
	return transaction, receipt, err
}

func (_IntegrityVerify *IntegrityVerifyTransactor) AsyncSetAddressToBool(handler func(*types.Receipt, error), opts *bind.TransactOpts, addr common.Address, value bool) (*types.Transaction, error) {
	return _IntegrityVerify.contract.AsyncTransact(opts, handler, "setAddressToBool", addr, value)
}

// SetAddressToBool is a paid mutator transaction binding the contract method 0x09ea9251.
//
// Solidity: function setAddressToBool(address addr, bool value) returns()
func (_IntegrityVerify *IntegrityVerifySession) SetAddressToBool(addr common.Address, value bool) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.SetAddressToBool(&_IntegrityVerify.TransactOpts, addr, value)
}

func (_IntegrityVerify *IntegrityVerifySession) AsyncSetAddressToBool(handler func(*types.Receipt, error), addr common.Address, value bool) (*types.Transaction, error) {
	return _IntegrityVerify.Contract.AsyncSetAddressToBool(handler, &_IntegrityVerify.TransactOpts, addr, value)
}

// SetAddressToBool is a paid mutator transaction binding the contract method 0x09ea9251.
//
// Solidity: function setAddressToBool(address addr, bool value) returns()
func (_IntegrityVerify *IntegrityVerifyTransactorSession) SetAddressToBool(addr common.Address, value bool) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.SetAddressToBool(&_IntegrityVerify.TransactOpts, addr, value)
}

func (_IntegrityVerify *IntegrityVerifyTransactorSession) AsyncSetAddressToBool(handler func(*types.Receipt, error), addr common.Address, value bool) (*types.Transaction, error) {
	return _IntegrityVerify.Contract.AsyncSetAddressToBool(handler, &_IntegrityVerify.TransactOpts, addr, value)
}

// SetAddressToPublicKey is a paid mutator transaction binding the contract method 0x6c521237.
//
// Solidity: function setAddressToPublicKey(address addr, bytes publicKey) returns()
func (_IntegrityVerify *IntegrityVerifyTransactor) SetAddressToPublicKey(opts *bind.TransactOpts, addr common.Address, publicKey []byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _IntegrityVerify.contract.TransactWithResult(opts, out, "setAddressToPublicKey", addr, publicKey)
	return transaction, receipt, err
}

func (_IntegrityVerify *IntegrityVerifyTransactor) AsyncSetAddressToPublicKey(handler func(*types.Receipt, error), opts *bind.TransactOpts, addr common.Address, publicKey []byte) (*types.Transaction, error) {
	return _IntegrityVerify.contract.AsyncTransact(opts, handler, "setAddressToPublicKey", addr, publicKey)
}

// SetAddressToPublicKey is a paid mutator transaction binding the contract method 0x6c521237.
//
// Solidity: function setAddressToPublicKey(address addr, bytes publicKey) returns()
func (_IntegrityVerify *IntegrityVerifySession) SetAddressToPublicKey(addr common.Address, publicKey []byte) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.SetAddressToPublicKey(&_IntegrityVerify.TransactOpts, addr, publicKey)
}

func (_IntegrityVerify *IntegrityVerifySession) AsyncSetAddressToPublicKey(handler func(*types.Receipt, error), addr common.Address, publicKey []byte) (*types.Transaction, error) {
	return _IntegrityVerify.Contract.AsyncSetAddressToPublicKey(handler, &_IntegrityVerify.TransactOpts, addr, publicKey)
}

// SetAddressToPublicKey is a paid mutator transaction binding the contract method 0x6c521237.
//
// Solidity: function setAddressToPublicKey(address addr, bytes publicKey) returns()
func (_IntegrityVerify *IntegrityVerifyTransactorSession) SetAddressToPublicKey(addr common.Address, publicKey []byte) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.SetAddressToPublicKey(&_IntegrityVerify.TransactOpts, addr, publicKey)
}

func (_IntegrityVerify *IntegrityVerifyTransactorSession) AsyncSetAddressToPublicKey(handler func(*types.Receipt, error), addr common.Address, publicKey []byte) (*types.Transaction, error) {
	return _IntegrityVerify.Contract.AsyncSetAddressToPublicKey(handler, &_IntegrityVerify.TransactOpts, addr, publicKey)
}

// SetHash is a paid mutator transaction binding the contract method 0x85ec0c57.
//
// Solidity: function setHash(string key, bytes32 value) returns()
func (_IntegrityVerify *IntegrityVerifyTransactor) SetHash(opts *bind.TransactOpts, key string, value [32]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _IntegrityVerify.contract.TransactWithResult(opts, out, "setHash", key, value)
	return transaction, receipt, err
}

func (_IntegrityVerify *IntegrityVerifyTransactor) AsyncSetHash(handler func(*types.Receipt, error), opts *bind.TransactOpts, key string, value [32]byte) (*types.Transaction, error) {
	return _IntegrityVerify.contract.AsyncTransact(opts, handler, "setHash", key, value)
}

// SetHash is a paid mutator transaction binding the contract method 0x85ec0c57.
//
// Solidity: function setHash(string key, bytes32 value) returns()
func (_IntegrityVerify *IntegrityVerifySession) SetHash(key string, value [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.SetHash(&_IntegrityVerify.TransactOpts, key, value)
}

func (_IntegrityVerify *IntegrityVerifySession) AsyncSetHash(handler func(*types.Receipt, error), key string, value [32]byte) (*types.Transaction, error) {
	return _IntegrityVerify.Contract.AsyncSetHash(handler, &_IntegrityVerify.TransactOpts, key, value)
}

// SetHash is a paid mutator transaction binding the contract method 0x85ec0c57.
//
// Solidity: function setHash(string key, bytes32 value) returns()
func (_IntegrityVerify *IntegrityVerifyTransactorSession) SetHash(key string, value [32]byte) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.SetHash(&_IntegrityVerify.TransactOpts, key, value)
}

func (_IntegrityVerify *IntegrityVerifyTransactorSession) AsyncSetHash(handler func(*types.Receipt, error), key string, value [32]byte) (*types.Transaction, error) {
	return _IntegrityVerify.Contract.AsyncSetHash(handler, &_IntegrityVerify.TransactOpts, key, value)
}

// SetMultipleHashes is a paid mutator transaction binding the contract method 0x97dbbca0.
//
// Solidity: function setMultipleHashes(string[] keys, bytes32[] values) returns()
func (_IntegrityVerify *IntegrityVerifyTransactor) SetMultipleHashes(opts *bind.TransactOpts, keys []string, values [][32]byte) (*types.Transaction, *types.Receipt, error) {
	var ()
	out := &[]interface{}{}
	transaction, receipt, err := _IntegrityVerify.contract.TransactWithResult(opts, out, "setMultipleHashes", keys, values)
	return transaction, receipt, err
}

func (_IntegrityVerify *IntegrityVerifyTransactor) AsyncSetMultipleHashes(handler func(*types.Receipt, error), opts *bind.TransactOpts, keys []string, values [][32]byte) (*types.Transaction, error) {
	return _IntegrityVerify.contract.AsyncTransact(opts, handler, "setMultipleHashes", keys, values)
}

// SetMultipleHashes is a paid mutator transaction binding the contract method 0x97dbbca0.
//
// Solidity: function setMultipleHashes(string[] keys, bytes32[] values) returns()
func (_IntegrityVerify *IntegrityVerifySession) SetMultipleHashes(keys []string, values [][32]byte) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.SetMultipleHashes(&_IntegrityVerify.TransactOpts, keys, values)
}

func (_IntegrityVerify *IntegrityVerifySession) AsyncSetMultipleHashes(handler func(*types.Receipt, error), keys []string, values [][32]byte) (*types.Transaction, error) {
	return _IntegrityVerify.Contract.AsyncSetMultipleHashes(handler, &_IntegrityVerify.TransactOpts, keys, values)
}

// SetMultipleHashes is a paid mutator transaction binding the contract method 0x97dbbca0.
//
// Solidity: function setMultipleHashes(string[] keys, bytes32[] values) returns()
func (_IntegrityVerify *IntegrityVerifyTransactorSession) SetMultipleHashes(keys []string, values [][32]byte) (*types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.SetMultipleHashes(&_IntegrityVerify.TransactOpts, keys, values)
}

func (_IntegrityVerify *IntegrityVerifyTransactorSession) AsyncSetMultipleHashes(handler func(*types.Receipt, error), keys []string, values [][32]byte) (*types.Transaction, error) {
	return _IntegrityVerify.Contract.AsyncSetMultipleHashes(handler, &_IntegrityVerify.TransactOpts, keys, values)
}

// VerifyAndStoreHashes is a paid mutator transaction binding the contract method 0xf042b697.
//
// Solidity: function verifyAndStoreHashes(string[] keys, bytes32[] values, bytes32[] hashArrayToStore) returns(bool)
func (_IntegrityVerify *IntegrityVerifyTransactor) VerifyAndStoreHashes(opts *bind.TransactOpts, keys []string, values [][32]byte, hashArrayToStore [][32]byte) (bool, *types.Transaction, *types.Receipt, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	transaction, receipt, err := _IntegrityVerify.contract.TransactWithResult(opts, out, "verifyAndStoreHashes", keys, values, hashArrayToStore)
	return *ret0, transaction, receipt, err
}

func (_IntegrityVerify *IntegrityVerifyTransactor) AsyncVerifyAndStoreHashes(handler func(*types.Receipt, error), opts *bind.TransactOpts, keys []string, values [][32]byte, hashArrayToStore [][32]byte) (*types.Transaction, error) {
	return _IntegrityVerify.contract.AsyncTransact(opts, handler, "verifyAndStoreHashes", keys, values, hashArrayToStore)
}

// VerifyAndStoreHashes is a paid mutator transaction binding the contract method 0xf042b697.
//
// Solidity: function verifyAndStoreHashes(string[] keys, bytes32[] values, bytes32[] hashArrayToStore) returns(bool)
func (_IntegrityVerify *IntegrityVerifySession) VerifyAndStoreHashes(keys []string, values [][32]byte, hashArrayToStore [][32]byte) (bool, *types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.VerifyAndStoreHashes(&_IntegrityVerify.TransactOpts, keys, values, hashArrayToStore)
}

func (_IntegrityVerify *IntegrityVerifySession) AsyncVerifyAndStoreHashes(handler func(*types.Receipt, error), keys []string, values [][32]byte, hashArrayToStore [][32]byte) (*types.Transaction, error) {
	return _IntegrityVerify.Contract.AsyncVerifyAndStoreHashes(handler, &_IntegrityVerify.TransactOpts, keys, values, hashArrayToStore)
}

// VerifyAndStoreHashes is a paid mutator transaction binding the contract method 0xf042b697.
//
// Solidity: function verifyAndStoreHashes(string[] keys, bytes32[] values, bytes32[] hashArrayToStore) returns(bool)
func (_IntegrityVerify *IntegrityVerifyTransactorSession) VerifyAndStoreHashes(keys []string, values [][32]byte, hashArrayToStore [][32]byte) (bool, *types.Transaction, *types.Receipt, error) {
	return _IntegrityVerify.Contract.VerifyAndStoreHashes(&_IntegrityVerify.TransactOpts, keys, values, hashArrayToStore)
}

func (_IntegrityVerify *IntegrityVerifyTransactorSession) AsyncVerifyAndStoreHashes(handler func(*types.Receipt, error), keys []string, values [][32]byte, hashArrayToStore [][32]byte) (*types.Transaction, error) {
	return _IntegrityVerify.Contract.AsyncVerifyAndStoreHashes(handler, &_IntegrityVerify.TransactOpts, keys, values, hashArrayToStore)
}

// IntegrityVerifyHashArrayStored represents a HashArrayStored event raised by the IntegrityVerify contract.
type IntegrityVerifyHashArrayStored struct {
	SigArray        [][32]byte
	SenderPublicKey []byte
	HashArray       [][32]byte
	Sender          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// WatchHashArrayStored is a free log subscription operation binding the contract event 0xdde8f1c4b3994910d6f2775c2ae898bce9327a67a2c04a711b584e951d76828c.
//
// Solidity: event HashArrayStored(bytes32[] sigArray, bytes senderPublicKey, bytes32[] hashArray, address sender)
func (_IntegrityVerify *IntegrityVerifyFilterer) WatchHashArrayStored(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _IntegrityVerify.contract.WatchLogs(fromBlock, handler, "HashArrayStored")
}

func (_IntegrityVerify *IntegrityVerifyFilterer) WatchAllHashArrayStored(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _IntegrityVerify.contract.WatchLogs(fromBlock, handler, "HashArrayStored")
}

// ParseHashArrayStored is a log parse operation binding the contract event 0xdde8f1c4b3994910d6f2775c2ae898bce9327a67a2c04a711b584e951d76828c.
//
// Solidity: event HashArrayStored(bytes32[] sigArray, bytes senderPublicKey, bytes32[] hashArray, address sender)
func (_IntegrityVerify *IntegrityVerifyFilterer) ParseHashArrayStored(log types.Log) (*IntegrityVerifyHashArrayStored, error) {
	event := new(IntegrityVerifyHashArrayStored)
	if err := _IntegrityVerify.contract.UnpackLog(event, "HashArrayStored", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchHashArrayStored is a free log subscription operation binding the contract event 0xdde8f1c4b3994910d6f2775c2ae898bce9327a67a2c04a711b584e951d76828c.
//
// Solidity: event HashArrayStored(bytes32[] sigArray, bytes senderPublicKey, bytes32[] hashArray, address sender)
func (_IntegrityVerify *IntegrityVerifySession) WatchHashArrayStored(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _IntegrityVerify.Contract.WatchHashArrayStored(fromBlock, handler)
}

func (_IntegrityVerify *IntegrityVerifySession) WatchAllHashArrayStored(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _IntegrityVerify.Contract.WatchAllHashArrayStored(fromBlock, handler)
}

// ParseHashArrayStored is a log parse operation binding the contract event 0xdde8f1c4b3994910d6f2775c2ae898bce9327a67a2c04a711b584e951d76828c.
//
// Solidity: event HashArrayStored(bytes32[] sigArray, bytes senderPublicKey, bytes32[] hashArray, address sender)
func (_IntegrityVerify *IntegrityVerifySession) ParseHashArrayStored(log types.Log) (*IntegrityVerifyHashArrayStored, error) {
	return _IntegrityVerify.Contract.ParseHashArrayStored(log)
}

// IntegrityVerifyMeasureFinished represents a MeasureFinished event raised by the IntegrityVerify contract.
type IntegrityVerifyMeasureFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// WatchMeasureFinished is a free log subscription operation binding the contract event 0xb9731c194d904fe4f70b3d61d77bd55fda2e9e934ac8b7cc16d68cb7af85604e.
//
// Solidity: event measureFinished()
func (_IntegrityVerify *IntegrityVerifyFilterer) WatchMeasureFinished(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _IntegrityVerify.contract.WatchLogs(fromBlock, handler, "measureFinished")
}

func (_IntegrityVerify *IntegrityVerifyFilterer) WatchAllMeasureFinished(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _IntegrityVerify.contract.WatchLogs(fromBlock, handler, "measureFinished")
}

// ParseMeasureFinished is a log parse operation binding the contract event 0xb9731c194d904fe4f70b3d61d77bd55fda2e9e934ac8b7cc16d68cb7af85604e.
//
// Solidity: event measureFinished()
func (_IntegrityVerify *IntegrityVerifyFilterer) ParseMeasureFinished(log types.Log) (*IntegrityVerifyMeasureFinished, error) {
	event := new(IntegrityVerifyMeasureFinished)
	if err := _IntegrityVerify.contract.UnpackLog(event, "measureFinished", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchMeasureFinished is a free log subscription operation binding the contract event 0xb9731c194d904fe4f70b3d61d77bd55fda2e9e934ac8b7cc16d68cb7af85604e.
//
// Solidity: event measureFinished()
func (_IntegrityVerify *IntegrityVerifySession) WatchMeasureFinished(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _IntegrityVerify.Contract.WatchMeasureFinished(fromBlock, handler)
}

func (_IntegrityVerify *IntegrityVerifySession) WatchAllMeasureFinished(fromBlock *int64, handler func(int, []types.Log)) (string, error) {
	return _IntegrityVerify.Contract.WatchAllMeasureFinished(fromBlock, handler)
}

// ParseMeasureFinished is a log parse operation binding the contract event 0xb9731c194d904fe4f70b3d61d77bd55fda2e9e934ac8b7cc16d68cb7af85604e.
//
// Solidity: event measureFinished()
func (_IntegrityVerify *IntegrityVerifySession) ParseMeasureFinished(log types.Log) (*IntegrityVerifyMeasureFinished, error) {
	return _IntegrityVerify.Contract.ParseMeasureFinished(log)
}
