// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abichainlink

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

// AbichainlinkABI is the input ABI used to generate the binding from.
const AbichainlinkABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Abichainlink is an auto generated Go binding around an Ethereum contract.
type Abichainlink struct {
	AbichainlinkCaller     // Read-only binding to the contract
	AbichainlinkTransactor // Write-only binding to the contract
	AbichainlinkFilterer   // Log filterer for contract events
}

// AbichainlinkCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbichainlinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbichainlinkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbichainlinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbichainlinkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbichainlinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbichainlinkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbichainlinkSession struct {
	Contract     *Abichainlink     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbichainlinkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbichainlinkCallerSession struct {
	Contract *AbichainlinkCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AbichainlinkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbichainlinkTransactorSession struct {
	Contract     *AbichainlinkTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AbichainlinkRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbichainlinkRaw struct {
	Contract *Abichainlink // Generic contract binding to access the raw methods on
}

// AbichainlinkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbichainlinkCallerRaw struct {
	Contract *AbichainlinkCaller // Generic read-only contract binding to access the raw methods on
}

// AbichainlinkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbichainlinkTransactorRaw struct {
	Contract *AbichainlinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbichainlink creates a new instance of Abichainlink, bound to a specific deployed contract.
func NewAbichainlink(address common.Address, backend bind.ContractBackend) (*Abichainlink, error) {
	contract, err := bindAbichainlink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abichainlink{AbichainlinkCaller: AbichainlinkCaller{contract: contract}, AbichainlinkTransactor: AbichainlinkTransactor{contract: contract}, AbichainlinkFilterer: AbichainlinkFilterer{contract: contract}}, nil
}

// NewAbichainlinkCaller creates a new read-only instance of Abichainlink, bound to a specific deployed contract.
func NewAbichainlinkCaller(address common.Address, caller bind.ContractCaller) (*AbichainlinkCaller, error) {
	contract, err := bindAbichainlink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbichainlinkCaller{contract: contract}, nil
}

// NewAbichainlinkTransactor creates a new write-only instance of Abichainlink, bound to a specific deployed contract.
func NewAbichainlinkTransactor(address common.Address, transactor bind.ContractTransactor) (*AbichainlinkTransactor, error) {
	contract, err := bindAbichainlink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbichainlinkTransactor{contract: contract}, nil
}

// NewAbichainlinkFilterer creates a new log filterer instance of Abichainlink, bound to a specific deployed contract.
func NewAbichainlinkFilterer(address common.Address, filterer bind.ContractFilterer) (*AbichainlinkFilterer, error) {
	contract, err := bindAbichainlink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbichainlinkFilterer{contract: contract}, nil
}

// bindAbichainlink binds a generic wrapper to an already deployed contract.
func bindAbichainlink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbichainlinkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abichainlink *AbichainlinkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abichainlink.Contract.AbichainlinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abichainlink *AbichainlinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abichainlink.Contract.AbichainlinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abichainlink *AbichainlinkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abichainlink.Contract.AbichainlinkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abichainlink *AbichainlinkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abichainlink.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abichainlink *AbichainlinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abichainlink.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abichainlink *AbichainlinkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abichainlink.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abichainlink *AbichainlinkCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abichainlink.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abichainlink *AbichainlinkSession) Decimals() (uint8, error) {
	return _Abichainlink.Contract.Decimals(&_Abichainlink.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abichainlink *AbichainlinkCallerSession) Decimals() (uint8, error) {
	return _Abichainlink.Contract.Decimals(&_Abichainlink.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Abichainlink *AbichainlinkCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abichainlink.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Abichainlink *AbichainlinkSession) Description() (string, error) {
	return _Abichainlink.Contract.Description(&_Abichainlink.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Abichainlink *AbichainlinkCallerSession) Description() (string, error) {
	return _Abichainlink.Contract.Description(&_Abichainlink.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Abichainlink *AbichainlinkCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _Abichainlink.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Abichainlink *AbichainlinkSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Abichainlink.Contract.GetRoundData(&_Abichainlink.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Abichainlink *AbichainlinkCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Abichainlink.Contract.GetRoundData(&_Abichainlink.CallOpts, _roundId)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Abichainlink *AbichainlinkCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _Abichainlink.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Abichainlink *AbichainlinkSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Abichainlink.Contract.LatestRoundData(&_Abichainlink.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Abichainlink *AbichainlinkCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Abichainlink.Contract.LatestRoundData(&_Abichainlink.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Abichainlink *AbichainlinkCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abichainlink.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Abichainlink *AbichainlinkSession) Version() (*big.Int, error) {
	return _Abichainlink.Contract.Version(&_Abichainlink.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Abichainlink *AbichainlinkCallerSession) Version() (*big.Int, error) {
	return _Abichainlink.Contract.Version(&_Abichainlink.CallOpts)
}
