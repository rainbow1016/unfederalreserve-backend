// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abiunitroller

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AbiunitrollerABI is the input ABI used to generate the binding from.
const AbiunitrollerABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingImplementation\",\"type\":\"address\"}],\"name\":\"NewPendingImplementation\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptImplementation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPendingImplementation\",\"type\":\"address\"}],\"name\":\"_setPendingImplementation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AbiunitrollerBin is the compiled bytecode used for deploying new contracts.
var AbiunitrollerBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556105e4806100326000396000f3fe60806040526004361061007b5760003560e01c8063dcfbc0c71161004e578063dcfbc0c71461019e578063e992a041146101b3578063e9c714f2146101e6578063f851a440146101fb5761007b565b806326782247146100fe578063b71d1a0c1461012f578063bb82aa5e14610174578063c1e8033414610189575b6002546040516000916001600160a01b031690829036908083838082843760405192019450600093509091505080830381855af49150503d80600081146100de576040519150601f19603f3d011682016040523d82523d6000602084013e6100e3565b606091505b505090506040513d6000823e8180156100fa573d82f35b3d82fd5b34801561010a57600080fd5b50610113610210565b604080516001600160a01b039092168252519081900360200190f35b34801561013b57600080fd5b506101626004803603602081101561015257600080fd5b50356001600160a01b031661021f565b60408051918252519081900360200190f35b34801561018057600080fd5b506101136102b0565b34801561019557600080fd5b506101626102bf565b3480156101aa57600080fd5b506101136103ba565b3480156101bf57600080fd5b50610162600480360360208110156101d657600080fd5b50356001600160a01b03166103c9565b3480156101f257600080fd5b5061016261044d565b34801561020757600080fd5b50610113610533565b6001546001600160a01b031681565b600080546001600160a01b031633146102455761023e6001600e610542565b90506102ab565b600180546001600160a01b038481166001600160a01b0319831681179093556040805191909216808252602082019390935281517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9929181900390910190a160005b9150505b919050565b6002546001600160a01b031681565b6003546000906001600160a01b0316331415806102e557506003546001600160a01b0316155b156102fc576102f5600180610542565b90506103b7565b60028054600380546001600160a01b038082166001600160a01b031980861682179687905590921690925560408051938316808552949092166020840152815190927fd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a92908290030190a1600354604080516001600160a01b038085168252909216602083015280517fe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d8159281900390910190a160005b925050505b90565b6003546001600160a01b031681565b600080546001600160a01b031633146103e85761023e6001600f610542565b600380546001600160a01b038481166001600160a01b0319831617928390556040805192821680845293909116602083015280517fe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d8159281900390910190a160006102a7565b6001546000906001600160a01b031633141580610468575033155b15610479576102f560016000610542565b60008054600180546001600160a01b038082166001600160a01b031980861682179687905590921690925560408051938316808552949092166020840152815190927ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc92908290030190a1600154604080516001600160a01b038085168252909216602083015280517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a99281900390910190a160006103b2565b6000546001600160a01b031681565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa083601181111561057157fe5b83601481111561057d57fe5b604080519283526020830191909152600082820152519081900360600190a18260118111156105a857fe5b939250505056fea265627a7a72315820cc0a1a66430087241905ee085a969670967a4a5c47c61296d40587791a546d2364736f6c63430005100032"

// DeployAbiunitroller deploys a new Ethereum contract, binding an instance of Abiunitroller to it.
func DeployAbiunitroller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Abiunitroller, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiunitrollerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbiunitrollerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Abiunitroller{AbiunitrollerCaller: AbiunitrollerCaller{contract: contract}, AbiunitrollerTransactor: AbiunitrollerTransactor{contract: contract}, AbiunitrollerFilterer: AbiunitrollerFilterer{contract: contract}}, nil
}

// Abiunitroller is an auto generated Go binding around an Ethereum contract.
type Abiunitroller struct {
	AbiunitrollerCaller     // Read-only binding to the contract
	AbiunitrollerTransactor // Write-only binding to the contract
	AbiunitrollerFilterer   // Log filterer for contract events
}

// AbiunitrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiunitrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiunitrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiunitrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiunitrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiunitrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiunitrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiunitrollerSession struct {
	Contract     *Abiunitroller    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiunitrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiunitrollerCallerSession struct {
	Contract *AbiunitrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AbiunitrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiunitrollerTransactorSession struct {
	Contract     *AbiunitrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AbiunitrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiunitrollerRaw struct {
	Contract *Abiunitroller // Generic contract binding to access the raw methods on
}

// AbiunitrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiunitrollerCallerRaw struct {
	Contract *AbiunitrollerCaller // Generic read-only contract binding to access the raw methods on
}

// AbiunitrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiunitrollerTransactorRaw struct {
	Contract *AbiunitrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbiunitroller creates a new instance of Abiunitroller, bound to a specific deployed contract.
func NewAbiunitroller(address common.Address, backend bind.ContractBackend) (*Abiunitroller, error) {
	contract, err := bindAbiunitroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abiunitroller{AbiunitrollerCaller: AbiunitrollerCaller{contract: contract}, AbiunitrollerTransactor: AbiunitrollerTransactor{contract: contract}, AbiunitrollerFilterer: AbiunitrollerFilterer{contract: contract}}, nil
}

// NewAbiunitrollerCaller creates a new read-only instance of Abiunitroller, bound to a specific deployed contract.
func NewAbiunitrollerCaller(address common.Address, caller bind.ContractCaller) (*AbiunitrollerCaller, error) {
	contract, err := bindAbiunitroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiunitrollerCaller{contract: contract}, nil
}

// NewAbiunitrollerTransactor creates a new write-only instance of Abiunitroller, bound to a specific deployed contract.
func NewAbiunitrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiunitrollerTransactor, error) {
	contract, err := bindAbiunitroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiunitrollerTransactor{contract: contract}, nil
}

// NewAbiunitrollerFilterer creates a new log filterer instance of Abiunitroller, bound to a specific deployed contract.
func NewAbiunitrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiunitrollerFilterer, error) {
	contract, err := bindAbiunitroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiunitrollerFilterer{contract: contract}, nil
}

// bindAbiunitroller binds a generic wrapper to an already deployed contract.
func bindAbiunitroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiunitrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abiunitroller *AbiunitrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abiunitroller.Contract.AbiunitrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abiunitroller *AbiunitrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abiunitroller.Contract.AbiunitrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abiunitroller *AbiunitrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abiunitroller.Contract.AbiunitrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abiunitroller *AbiunitrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abiunitroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abiunitroller *AbiunitrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abiunitroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abiunitroller *AbiunitrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abiunitroller.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abiunitroller *AbiunitrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abiunitroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abiunitroller *AbiunitrollerSession) Admin() (common.Address, error) {
	return _Abiunitroller.Contract.Admin(&_Abiunitroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abiunitroller *AbiunitrollerCallerSession) Admin() (common.Address, error) {
	return _Abiunitroller.Contract.Admin(&_Abiunitroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Abiunitroller *AbiunitrollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abiunitroller.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Abiunitroller *AbiunitrollerSession) ComptrollerImplementation() (common.Address, error) {
	return _Abiunitroller.Contract.ComptrollerImplementation(&_Abiunitroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Abiunitroller *AbiunitrollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _Abiunitroller.Contract.ComptrollerImplementation(&_Abiunitroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abiunitroller *AbiunitrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abiunitroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abiunitroller *AbiunitrollerSession) PendingAdmin() (common.Address, error) {
	return _Abiunitroller.Contract.PendingAdmin(&_Abiunitroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abiunitroller *AbiunitrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _Abiunitroller.Contract.PendingAdmin(&_Abiunitroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Abiunitroller *AbiunitrollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abiunitroller.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Abiunitroller *AbiunitrollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Abiunitroller.Contract.PendingComptrollerImplementation(&_Abiunitroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Abiunitroller *AbiunitrollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Abiunitroller.Contract.PendingComptrollerImplementation(&_Abiunitroller.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Abiunitroller *AbiunitrollerTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abiunitroller.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Abiunitroller *AbiunitrollerSession) AcceptAdmin() (*types.Transaction, error) {
	return _Abiunitroller.Contract.AcceptAdmin(&_Abiunitroller.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Abiunitroller *AbiunitrollerTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Abiunitroller.Contract.AcceptAdmin(&_Abiunitroller.TransactOpts)
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_Abiunitroller *AbiunitrollerTransactor) AcceptImplementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abiunitroller.contract.Transact(opts, "_acceptImplementation")
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_Abiunitroller *AbiunitrollerSession) AcceptImplementation() (*types.Transaction, error) {
	return _Abiunitroller.Contract.AcceptImplementation(&_Abiunitroller.TransactOpts)
}

// AcceptImplementation is a paid mutator transaction binding the contract method 0xc1e80334.
//
// Solidity: function _acceptImplementation() returns(uint256)
func (_Abiunitroller *AbiunitrollerTransactorSession) AcceptImplementation() (*types.Transaction, error) {
	return _Abiunitroller.Contract.AcceptImplementation(&_Abiunitroller.TransactOpts)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Abiunitroller *AbiunitrollerTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abiunitroller.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Abiunitroller *AbiunitrollerSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abiunitroller.Contract.SetPendingAdmin(&_Abiunitroller.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Abiunitroller *AbiunitrollerTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abiunitroller.Contract.SetPendingAdmin(&_Abiunitroller.TransactOpts, newPendingAdmin)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_Abiunitroller *AbiunitrollerTransactor) SetPendingImplementation(opts *bind.TransactOpts, newPendingImplementation common.Address) (*types.Transaction, error) {
	return _Abiunitroller.contract.Transact(opts, "_setPendingImplementation", newPendingImplementation)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_Abiunitroller *AbiunitrollerSession) SetPendingImplementation(newPendingImplementation common.Address) (*types.Transaction, error) {
	return _Abiunitroller.Contract.SetPendingImplementation(&_Abiunitroller.TransactOpts, newPendingImplementation)
}

// SetPendingImplementation is a paid mutator transaction binding the contract method 0xe992a041.
//
// Solidity: function _setPendingImplementation(address newPendingImplementation) returns(uint256)
func (_Abiunitroller *AbiunitrollerTransactorSession) SetPendingImplementation(newPendingImplementation common.Address) (*types.Transaction, error) {
	return _Abiunitroller.Contract.SetPendingImplementation(&_Abiunitroller.TransactOpts, newPendingImplementation)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abiunitroller *AbiunitrollerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Abiunitroller.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abiunitroller *AbiunitrollerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Abiunitroller.Contract.Fallback(&_Abiunitroller.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abiunitroller *AbiunitrollerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Abiunitroller.Contract.Fallback(&_Abiunitroller.TransactOpts, calldata)
}

// AbiunitrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Abiunitroller contract.
type AbiunitrollerFailureIterator struct {
	Event *AbiunitrollerFailure // Event containing the contract specifics and raw log

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
func (it *AbiunitrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiunitrollerFailure)
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
		it.Event = new(AbiunitrollerFailure)
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
func (it *AbiunitrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiunitrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiunitrollerFailure represents a Failure event raised by the Abiunitroller contract.
type AbiunitrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Abiunitroller *AbiunitrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*AbiunitrollerFailureIterator, error) {

	logs, sub, err := _Abiunitroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &AbiunitrollerFailureIterator{contract: _Abiunitroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Abiunitroller *AbiunitrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *AbiunitrollerFailure) (event.Subscription, error) {

	logs, sub, err := _Abiunitroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiunitrollerFailure)
				if err := _Abiunitroller.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Abiunitroller *AbiunitrollerFilterer) ParseFailure(log types.Log) (*AbiunitrollerFailure, error) {
	event := new(AbiunitrollerFailure)
	if err := _Abiunitroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiunitrollerNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Abiunitroller contract.
type AbiunitrollerNewAdminIterator struct {
	Event *AbiunitrollerNewAdmin // Event containing the contract specifics and raw log

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
func (it *AbiunitrollerNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiunitrollerNewAdmin)
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
		it.Event = new(AbiunitrollerNewAdmin)
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
func (it *AbiunitrollerNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiunitrollerNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiunitrollerNewAdmin represents a NewAdmin event raised by the Abiunitroller contract.
type AbiunitrollerNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Abiunitroller *AbiunitrollerFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*AbiunitrollerNewAdminIterator, error) {

	logs, sub, err := _Abiunitroller.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &AbiunitrollerNewAdminIterator{contract: _Abiunitroller.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Abiunitroller *AbiunitrollerFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *AbiunitrollerNewAdmin) (event.Subscription, error) {

	logs, sub, err := _Abiunitroller.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiunitrollerNewAdmin)
				if err := _Abiunitroller.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Abiunitroller *AbiunitrollerFilterer) ParseNewAdmin(log types.Log) (*AbiunitrollerNewAdmin, error) {
	event := new(AbiunitrollerNewAdmin)
	if err := _Abiunitroller.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiunitrollerNewImplementationIterator is returned from FilterNewImplementation and is used to iterate over the raw logs and unpacked data for NewImplementation events raised by the Abiunitroller contract.
type AbiunitrollerNewImplementationIterator struct {
	Event *AbiunitrollerNewImplementation // Event containing the contract specifics and raw log

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
func (it *AbiunitrollerNewImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiunitrollerNewImplementation)
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
		it.Event = new(AbiunitrollerNewImplementation)
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
func (it *AbiunitrollerNewImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiunitrollerNewImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiunitrollerNewImplementation represents a NewImplementation event raised by the Abiunitroller contract.
type AbiunitrollerNewImplementation struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewImplementation is a free log retrieval operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Abiunitroller *AbiunitrollerFilterer) FilterNewImplementation(opts *bind.FilterOpts) (*AbiunitrollerNewImplementationIterator, error) {

	logs, sub, err := _Abiunitroller.contract.FilterLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return &AbiunitrollerNewImplementationIterator{contract: _Abiunitroller.contract, event: "NewImplementation", logs: logs, sub: sub}, nil
}

// WatchNewImplementation is a free log subscription operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Abiunitroller *AbiunitrollerFilterer) WatchNewImplementation(opts *bind.WatchOpts, sink chan<- *AbiunitrollerNewImplementation) (event.Subscription, error) {

	logs, sub, err := _Abiunitroller.contract.WatchLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiunitrollerNewImplementation)
				if err := _Abiunitroller.contract.UnpackLog(event, "NewImplementation", log); err != nil {
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

// ParseNewImplementation is a log parse operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Abiunitroller *AbiunitrollerFilterer) ParseNewImplementation(log types.Log) (*AbiunitrollerNewImplementation, error) {
	event := new(AbiunitrollerNewImplementation)
	if err := _Abiunitroller.contract.UnpackLog(event, "NewImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiunitrollerNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Abiunitroller contract.
type AbiunitrollerNewPendingAdminIterator struct {
	Event *AbiunitrollerNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *AbiunitrollerNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiunitrollerNewPendingAdmin)
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
		it.Event = new(AbiunitrollerNewPendingAdmin)
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
func (it *AbiunitrollerNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiunitrollerNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiunitrollerNewPendingAdmin represents a NewPendingAdmin event raised by the Abiunitroller contract.
type AbiunitrollerNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Abiunitroller *AbiunitrollerFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*AbiunitrollerNewPendingAdminIterator, error) {

	logs, sub, err := _Abiunitroller.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &AbiunitrollerNewPendingAdminIterator{contract: _Abiunitroller.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Abiunitroller *AbiunitrollerFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *AbiunitrollerNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Abiunitroller.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiunitrollerNewPendingAdmin)
				if err := _Abiunitroller.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Abiunitroller *AbiunitrollerFilterer) ParseNewPendingAdmin(log types.Log) (*AbiunitrollerNewPendingAdmin, error) {
	event := new(AbiunitrollerNewPendingAdmin)
	if err := _Abiunitroller.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbiunitrollerNewPendingImplementationIterator is returned from FilterNewPendingImplementation and is used to iterate over the raw logs and unpacked data for NewPendingImplementation events raised by the Abiunitroller contract.
type AbiunitrollerNewPendingImplementationIterator struct {
	Event *AbiunitrollerNewPendingImplementation // Event containing the contract specifics and raw log

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
func (it *AbiunitrollerNewPendingImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiunitrollerNewPendingImplementation)
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
		it.Event = new(AbiunitrollerNewPendingImplementation)
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
func (it *AbiunitrollerNewPendingImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiunitrollerNewPendingImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiunitrollerNewPendingImplementation represents a NewPendingImplementation event raised by the Abiunitroller contract.
type AbiunitrollerNewPendingImplementation struct {
	OldPendingImplementation common.Address
	NewPendingImplementation common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewPendingImplementation is a free log retrieval operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address oldPendingImplementation, address newPendingImplementation)
func (_Abiunitroller *AbiunitrollerFilterer) FilterNewPendingImplementation(opts *bind.FilterOpts) (*AbiunitrollerNewPendingImplementationIterator, error) {

	logs, sub, err := _Abiunitroller.contract.FilterLogs(opts, "NewPendingImplementation")
	if err != nil {
		return nil, err
	}
	return &AbiunitrollerNewPendingImplementationIterator{contract: _Abiunitroller.contract, event: "NewPendingImplementation", logs: logs, sub: sub}, nil
}

// WatchNewPendingImplementation is a free log subscription operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address oldPendingImplementation, address newPendingImplementation)
func (_Abiunitroller *AbiunitrollerFilterer) WatchNewPendingImplementation(opts *bind.WatchOpts, sink chan<- *AbiunitrollerNewPendingImplementation) (event.Subscription, error) {

	logs, sub, err := _Abiunitroller.contract.WatchLogs(opts, "NewPendingImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiunitrollerNewPendingImplementation)
				if err := _Abiunitroller.contract.UnpackLog(event, "NewPendingImplementation", log); err != nil {
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

// ParseNewPendingImplementation is a log parse operation binding the contract event 0xe945ccee5d701fc83f9b8aa8ca94ea4219ec1fcbd4f4cab4f0ea57c5c3e1d815.
//
// Solidity: event NewPendingImplementation(address oldPendingImplementation, address newPendingImplementation)
func (_Abiunitroller *AbiunitrollerFilterer) ParseNewPendingImplementation(log types.Log) (*AbiunitrollerNewPendingImplementation, error) {
	event := new(AbiunitrollerNewPendingImplementation)
	if err := _Abiunitroller.contract.UnpackLog(event, "NewPendingImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
