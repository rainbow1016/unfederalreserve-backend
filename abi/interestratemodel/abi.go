// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abiinterestratemodel

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

// AbiinterestratemodelABI is the input ABI used to generate the binding from.
const AbiinterestratemodelABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseRatePerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplierPerYear\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseRatePerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multiplierPerBlock\",\"type\":\"uint256\"}],\"name\":\"NewInterestParams\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blocksPerYear\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"name\":\"getBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"getSupplyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInterestRateModel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"multiplierPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"name\":\"utilizationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// AbiinterestratemodelBin is the compiled bytecode used for deploying new contracts.
var AbiinterestratemodelBin = "0x608060405234801561001057600080fd5b506040516106da3803806106da8339818101604052604081101561003357600080fd5b508051602091820151909161005690839062201480906102ee6100bc821b17901c565b60015561007181622014806100bc602090811b6102ee17901c565b600081905560015460408051918252602082019290925281517ff35fa19c15e9ba782633a5df62a98b20217151addc68e3ff2cd623a48d37ec27929181900390910190a150506101ad565b600061010483836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061010b60201b60201c565b9392505050565b600081836101975760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561015c578181015183820152602001610144565b50505050905090810190601f1680156101895780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385816101a357fe5b0495945050505050565b61051e806101bc6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638726bb891161005b5780638726bb8914610102578063a385fb961461010a578063b816881614610112578063f14039de146101415761007d565b806315f24053146100825780632191f92a146100bd5780636e71e2d8146100d9575b600080fd5b6100ab6004803603606081101561009857600080fd5b5080359060208101359060400135610149565b60408051918252519081900360200190f35b6100c56101a3565b604080519115158252519081900360200190f35b6100ab600480360360608110156100ef57600080fd5b50803590602081013590604001356101a8565b6100ab6101fa565b6100ab610200565b6100ab6004803603608081101561012857600080fd5b5080359060208101359060408101359060600135610207565b6100ab610286565b6000806101578585856101a8565b905061019860015461018c670de0b6b3a76400006101806000548661028c90919063ffffffff16565b9063ffffffff6102ee16565b9063ffffffff61033016565b9150505b9392505050565b600181565b6000826101b75750600061019c565b6101f26101da836101ce878763ffffffff61033016565b9063ffffffff61038a16565b61018085670de0b6b3a764000063ffffffff61028c16565b949350505050565b60005481565b6220148081565b600080610222670de0b6b3a76400008463ffffffff61038a16565b90506000610231878787610149565b90506000610251670de0b6b3a7640000610180848663ffffffff61028c16565b905061027a670de0b6b3a76400006101808361026e8c8c8c6101a8565b9063ffffffff61028c16565b98975050505050505050565b60015481565b60008261029b575060006102e8565b828202828482816102a857fe5b04146102e55760405162461bcd60e51b81526004018080602001828103825260218152602001806104c96021913960400191505060405180910390fd5b90505b92915050565b60006102e583836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152506103cc565b6000828201838110156102e5576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b60006102e583836040518060400160405280601f81526020017f536166654d6174683a207375627472616374696f6e20756e646572666c6f770081525061046e565b600081836104585760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561041d578181015183820152602001610405565b50505050905090810190601f16801561044a5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858161046457fe5b0495945050505050565b600081848411156104c05760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561041d578181015183820152602001610405565b50505090039056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a265627a7a7231582092b1c738ce267719ac08ca45dc4880e4714d22473fc8a0b193dae3a2f1fcdd6a64736f6c63430005100032"

// DeployAbiinterestratemodel deploys a new Ethereum contract, binding an instance of Abiinterestratemodel to it.
func DeployAbiinterestratemodel(auth *bind.TransactOpts, backend bind.ContractBackend, baseRatePerYear *big.Int, multiplierPerYear *big.Int) (common.Address, *types.Transaction, *Abiinterestratemodel, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiinterestratemodelABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbiinterestratemodelBin), backend, baseRatePerYear, multiplierPerYear)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Abiinterestratemodel{AbiinterestratemodelCaller: AbiinterestratemodelCaller{contract: contract}, AbiinterestratemodelTransactor: AbiinterestratemodelTransactor{contract: contract}, AbiinterestratemodelFilterer: AbiinterestratemodelFilterer{contract: contract}}, nil
}

// Abiinterestratemodel is an auto generated Go binding around an Ethereum contract.
type Abiinterestratemodel struct {
	AbiinterestratemodelCaller     // Read-only binding to the contract
	AbiinterestratemodelTransactor // Write-only binding to the contract
	AbiinterestratemodelFilterer   // Log filterer for contract events
}

// AbiinterestratemodelCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiinterestratemodelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiinterestratemodelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiinterestratemodelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiinterestratemodelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiinterestratemodelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiinterestratemodelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiinterestratemodelSession struct {
	Contract     *Abiinterestratemodel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AbiinterestratemodelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiinterestratemodelCallerSession struct {
	Contract *AbiinterestratemodelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AbiinterestratemodelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiinterestratemodelTransactorSession struct {
	Contract     *AbiinterestratemodelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AbiinterestratemodelRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiinterestratemodelRaw struct {
	Contract *Abiinterestratemodel // Generic contract binding to access the raw methods on
}

// AbiinterestratemodelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiinterestratemodelCallerRaw struct {
	Contract *AbiinterestratemodelCaller // Generic read-only contract binding to access the raw methods on
}

// AbiinterestratemodelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiinterestratemodelTransactorRaw struct {
	Contract *AbiinterestratemodelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbiinterestratemodel creates a new instance of Abiinterestratemodel, bound to a specific deployed contract.
func NewAbiinterestratemodel(address common.Address, backend bind.ContractBackend) (*Abiinterestratemodel, error) {
	contract, err := bindAbiinterestratemodel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abiinterestratemodel{AbiinterestratemodelCaller: AbiinterestratemodelCaller{contract: contract}, AbiinterestratemodelTransactor: AbiinterestratemodelTransactor{contract: contract}, AbiinterestratemodelFilterer: AbiinterestratemodelFilterer{contract: contract}}, nil
}

// NewAbiinterestratemodelCaller creates a new read-only instance of Abiinterestratemodel, bound to a specific deployed contract.
func NewAbiinterestratemodelCaller(address common.Address, caller bind.ContractCaller) (*AbiinterestratemodelCaller, error) {
	contract, err := bindAbiinterestratemodel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiinterestratemodelCaller{contract: contract}, nil
}

// NewAbiinterestratemodelTransactor creates a new write-only instance of Abiinterestratemodel, bound to a specific deployed contract.
func NewAbiinterestratemodelTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiinterestratemodelTransactor, error) {
	contract, err := bindAbiinterestratemodel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiinterestratemodelTransactor{contract: contract}, nil
}

// NewAbiinterestratemodelFilterer creates a new log filterer instance of Abiinterestratemodel, bound to a specific deployed contract.
func NewAbiinterestratemodelFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiinterestratemodelFilterer, error) {
	contract, err := bindAbiinterestratemodel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiinterestratemodelFilterer{contract: contract}, nil
}

// bindAbiinterestratemodel binds a generic wrapper to an already deployed contract.
func bindAbiinterestratemodel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiinterestratemodelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abiinterestratemodel *AbiinterestratemodelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abiinterestratemodel.Contract.AbiinterestratemodelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abiinterestratemodel *AbiinterestratemodelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abiinterestratemodel.Contract.AbiinterestratemodelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abiinterestratemodel *AbiinterestratemodelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abiinterestratemodel.Contract.AbiinterestratemodelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abiinterestratemodel *AbiinterestratemodelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abiinterestratemodel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abiinterestratemodel *AbiinterestratemodelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abiinterestratemodel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abiinterestratemodel *AbiinterestratemodelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abiinterestratemodel.Contract.contract.Transact(opts, method, params...)
}

// BaseRatePerBlock is a free data retrieval call binding the contract method 0xf14039de.
//
// Solidity: function baseRatePerBlock() view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCaller) BaseRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abiinterestratemodel.contract.Call(opts, &out, "baseRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRatePerBlock is a free data retrieval call binding the contract method 0xf14039de.
//
// Solidity: function baseRatePerBlock() view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelSession) BaseRatePerBlock() (*big.Int, error) {
	return _Abiinterestratemodel.Contract.BaseRatePerBlock(&_Abiinterestratemodel.CallOpts)
}

// BaseRatePerBlock is a free data retrieval call binding the contract method 0xf14039de.
//
// Solidity: function baseRatePerBlock() view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCallerSession) BaseRatePerBlock() (*big.Int, error) {
	return _Abiinterestratemodel.Contract.BaseRatePerBlock(&_Abiinterestratemodel.CallOpts)
}

// BlocksPerYear is a free data retrieval call binding the contract method 0xa385fb96.
//
// Solidity: function blocksPerYear() view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCaller) BlocksPerYear(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abiinterestratemodel.contract.Call(opts, &out, "blocksPerYear")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksPerYear is a free data retrieval call binding the contract method 0xa385fb96.
//
// Solidity: function blocksPerYear() view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelSession) BlocksPerYear() (*big.Int, error) {
	return _Abiinterestratemodel.Contract.BlocksPerYear(&_Abiinterestratemodel.CallOpts)
}

// BlocksPerYear is a free data retrieval call binding the contract method 0xa385fb96.
//
// Solidity: function blocksPerYear() view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCallerSession) BlocksPerYear() (*big.Int, error) {
	return _Abiinterestratemodel.Contract.BlocksPerYear(&_Abiinterestratemodel.CallOpts)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCaller) GetBorrowRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abiinterestratemodel.contract.Call(opts, &out, "getBorrowRate", cash, borrows, reserves)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelSession) GetBorrowRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _Abiinterestratemodel.Contract.GetBorrowRate(&_Abiinterestratemodel.CallOpts, cash, borrows, reserves)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCallerSession) GetBorrowRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _Abiinterestratemodel.Contract.GetBorrowRate(&_Abiinterestratemodel.CallOpts, cash, borrows, reserves)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCaller) GetSupplyRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abiinterestratemodel.contract.Call(opts, &out, "getSupplyRate", cash, borrows, reserves, reserveFactorMantissa)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelSession) GetSupplyRate(cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	return _Abiinterestratemodel.Contract.GetSupplyRate(&_Abiinterestratemodel.CallOpts, cash, borrows, reserves, reserveFactorMantissa)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCallerSession) GetSupplyRate(cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	return _Abiinterestratemodel.Contract.GetSupplyRate(&_Abiinterestratemodel.CallOpts, cash, borrows, reserves, reserveFactorMantissa)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_Abiinterestratemodel *AbiinterestratemodelCaller) IsInterestRateModel(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abiinterestratemodel.contract.Call(opts, &out, "isInterestRateModel")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_Abiinterestratemodel *AbiinterestratemodelSession) IsInterestRateModel() (bool, error) {
	return _Abiinterestratemodel.Contract.IsInterestRateModel(&_Abiinterestratemodel.CallOpts)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_Abiinterestratemodel *AbiinterestratemodelCallerSession) IsInterestRateModel() (bool, error) {
	return _Abiinterestratemodel.Contract.IsInterestRateModel(&_Abiinterestratemodel.CallOpts)
}

// MultiplierPerBlock is a free data retrieval call binding the contract method 0x8726bb89.
//
// Solidity: function multiplierPerBlock() view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCaller) MultiplierPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abiinterestratemodel.contract.Call(opts, &out, "multiplierPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MultiplierPerBlock is a free data retrieval call binding the contract method 0x8726bb89.
//
// Solidity: function multiplierPerBlock() view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelSession) MultiplierPerBlock() (*big.Int, error) {
	return _Abiinterestratemodel.Contract.MultiplierPerBlock(&_Abiinterestratemodel.CallOpts)
}

// MultiplierPerBlock is a free data retrieval call binding the contract method 0x8726bb89.
//
// Solidity: function multiplierPerBlock() view returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCallerSession) MultiplierPerBlock() (*big.Int, error) {
	return _Abiinterestratemodel.Contract.MultiplierPerBlock(&_Abiinterestratemodel.CallOpts)
}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCaller) UtilizationRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abiinterestratemodel.contract.Call(opts, &out, "utilizationRate", cash, borrows, reserves)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelSession) UtilizationRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _Abiinterestratemodel.Contract.UtilizationRate(&_Abiinterestratemodel.CallOpts, cash, borrows, reserves)
}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_Abiinterestratemodel *AbiinterestratemodelCallerSession) UtilizationRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _Abiinterestratemodel.Contract.UtilizationRate(&_Abiinterestratemodel.CallOpts, cash, borrows, reserves)
}

// AbiinterestratemodelNewInterestParamsIterator is returned from FilterNewInterestParams and is used to iterate over the raw logs and unpacked data for NewInterestParams events raised by the Abiinterestratemodel contract.
type AbiinterestratemodelNewInterestParamsIterator struct {
	Event *AbiinterestratemodelNewInterestParams // Event containing the contract specifics and raw log

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
func (it *AbiinterestratemodelNewInterestParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbiinterestratemodelNewInterestParams)
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
		it.Event = new(AbiinterestratemodelNewInterestParams)
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
func (it *AbiinterestratemodelNewInterestParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbiinterestratemodelNewInterestParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbiinterestratemodelNewInterestParams represents a NewInterestParams event raised by the Abiinterestratemodel contract.
type AbiinterestratemodelNewInterestParams struct {
	BaseRatePerBlock   *big.Int
	MultiplierPerBlock *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewInterestParams is a free log retrieval operation binding the contract event 0xf35fa19c15e9ba782633a5df62a98b20217151addc68e3ff2cd623a48d37ec27.
//
// Solidity: event NewInterestParams(uint256 baseRatePerBlock, uint256 multiplierPerBlock)
func (_Abiinterestratemodel *AbiinterestratemodelFilterer) FilterNewInterestParams(opts *bind.FilterOpts) (*AbiinterestratemodelNewInterestParamsIterator, error) {

	logs, sub, err := _Abiinterestratemodel.contract.FilterLogs(opts, "NewInterestParams")
	if err != nil {
		return nil, err
	}
	return &AbiinterestratemodelNewInterestParamsIterator{contract: _Abiinterestratemodel.contract, event: "NewInterestParams", logs: logs, sub: sub}, nil
}

// WatchNewInterestParams is a free log subscription operation binding the contract event 0xf35fa19c15e9ba782633a5df62a98b20217151addc68e3ff2cd623a48d37ec27.
//
// Solidity: event NewInterestParams(uint256 baseRatePerBlock, uint256 multiplierPerBlock)
func (_Abiinterestratemodel *AbiinterestratemodelFilterer) WatchNewInterestParams(opts *bind.WatchOpts, sink chan<- *AbiinterestratemodelNewInterestParams) (event.Subscription, error) {

	logs, sub, err := _Abiinterestratemodel.contract.WatchLogs(opts, "NewInterestParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbiinterestratemodelNewInterestParams)
				if err := _Abiinterestratemodel.contract.UnpackLog(event, "NewInterestParams", log); err != nil {
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

// ParseNewInterestParams is a log parse operation binding the contract event 0xf35fa19c15e9ba782633a5df62a98b20217151addc68e3ff2cd623a48d37ec27.
//
// Solidity: event NewInterestParams(uint256 baseRatePerBlock, uint256 multiplierPerBlock)
func (_Abiinterestratemodel *AbiinterestratemodelFilterer) ParseNewInterestParams(log types.Log) (*AbiinterestratemodelNewInterestParams, error) {
	event := new(AbiinterestratemodelNewInterestParams)
	if err := _Abiinterestratemodel.contract.UnpackLog(event, "NewInterestParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
