// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abictoken

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

// AbictokenABI is the input ABI used to generate the binding from.
const AbictokenABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenName\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"_decimalUnits\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_tokenSymbol\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"allocateTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AbictokenBin is the compiled bytecode used for deploying new contracts.
var AbictokenBin = "0x608060405234801561001057600080fd5b50604051610adc380380610adc8339818101604052608081101561003357600080fd5b81516020830180516040519294929383019291908464010000000082111561005a57600080fd5b90830190602082018581111561006f57600080fd5b825164010000000081118282018810171561008957600080fd5b82525081516020918201929091019080838360005b838110156100b657818101518382015260200161009e565b50505050905090810190601f1680156100e35780820380516001836020036101000a031916815260200191505b5060408181526020830151920180519294919391928464010000000082111561010b57600080fd5b90830190602082018581111561012057600080fd5b825164010000000081118282018810171561013a57600080fd5b82525081516020918201929091019080838360005b8381101561016757818101518382015260200161014f565b50505050905090810190601f1680156101945780820380516001836020036101000a031916815260200191505b506040908152600388905533600090815260056020908152918120899055875189955088945087935086926101cd929190860190610203565b5080516101e1906001906020840190610203565b50506002805460ff191660ff929092169190911790555061029e945050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061024457805160ff1916838001178555610271565b82800160010185558215610271579182015b82811115610271578251825591602001919060010190610256565b5061027d929150610281565b5090565b61029b91905b8082111561027d5760008155600101610287565b90565b61082f806102ad6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063313ce56711610066578063313ce567146101de57806370a08231146101fc57806395d89b4114610222578063a9059cbb1461022a578063dd62ed3e146102565761009e565b806306fdde03146100a357806308bca56614610120578063095ea7b31461014e57806318160ddd1461018e57806323b872dd146101a8575b600080fd5b6100ab610284565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100e55781810151838201526020016100cd565b50505050905090810190601f1680156101125780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61014c6004803603604081101561013657600080fd5b506001600160a01b038135169060200135610312565b005b61017a6004803603604081101561016457600080fd5b506001600160a01b038135169060200135610372565b604080519115158252519081900360200190f35b6101966103d8565b60408051918252519081900360200190f35b61017a600480360360608110156101be57600080fd5b506001600160a01b038135811691602081013590911690604001356103de565b6101e661056a565b6040805160ff9092168252519081900360200190f35b6101966004803603602081101561021257600080fd5b50356001600160a01b0316610573565b6100ab610585565b61017a6004803603604081101561024057600080fd5b506001600160a01b0381351690602001356105df565b6101966004803603604081101561026c57600080fd5b506001600160a01b03813581169160200135166106e8565b6000805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561030a5780601f106102df5761010080835404028352916020019161030a565b820191906000526020600020905b8154815290600101906020018083116102ed57829003601f168201915b505050505081565b6001600160a01b03821660008181526005602090815260409182902080548501905560038054850190558151848152915130927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92908290030190a35050565b3360008181526004602090815260408083206001600160a01b038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b60035481565b6040805180820182526016815275496e73756666696369656e7420616c6c6f77616e636560501b6020808301919091526001600160a01b0386166000908152600482528381203382529091529182205461043f91849063ffffffff61070516565b6001600160a01b0385166000818152600460209081526040808320338452825280832094909455835180850185526014815273496e73756666696369656e742062616c616e636560601b818301529282526005905291909120546104aa91849063ffffffff61070516565b6001600160a01b0380861660009081526005602081815260408084209590955584518086018652601081526f42616c616e6365206f766572666c6f7760801b81830152938816835252919091205461050991849063ffffffff61079c16565b6001600160a01b0380851660008181526005602090815260409182902094909455805186815290519193928816927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a35060019392505050565b60025460ff1681565b60056020526000908152604090205481565b60018054604080516020600284861615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561030a5780601f106102df5761010080835404028352916020019161030a565b6040805180820182526014815273496e73756666696369656e742062616c616e636560601b60208083019190915233600090815260059091529182205461062d91849063ffffffff61070516565b3360009081526005602081815260408084209490945583518085018552601081526f42616c616e6365206f766572666c6f7760801b818301526001600160a01b03881684529190529190205461068a91849063ffffffff61079c16565b6001600160a01b0384166000818152600560209081526040918290209390935580518581529051919233927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a350600192915050565b600460209081526000928352604080842090915290825290205481565b600081848411156107945760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610759578181015183820152602001610741565b50505050905090810190601f1680156107865780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600083830182858210156107f15760405162461bcd60e51b8152602060048201818152835160248401528351909283926044909101919085019080838360008315610759578181015183820152602001610741565b5094935050505056fea265627a7a723158208e516c97191d699abe505379f613fe3aa5fb1b52cdfa5a14defa2533ed29279a64736f6c63430005100032"

// DeployAbictoken deploys a new Ethereum contract, binding an instance of Abictoken to it.
func DeployAbictoken(auth *bind.TransactOpts, backend bind.ContractBackend, _initialAmount *big.Int, _tokenName string, _decimalUnits uint8, _tokenSymbol string) (common.Address, *types.Transaction, *Abictoken, error) {
	parsed, err := abi.JSON(strings.NewReader(AbictokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbictokenBin), backend, _initialAmount, _tokenName, _decimalUnits, _tokenSymbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Abictoken{AbictokenCaller: AbictokenCaller{contract: contract}, AbictokenTransactor: AbictokenTransactor{contract: contract}, AbictokenFilterer: AbictokenFilterer{contract: contract}}, nil
}

// Abictoken is an auto generated Go binding around an Ethereum contract.
type Abictoken struct {
	AbictokenCaller     // Read-only binding to the contract
	AbictokenTransactor // Write-only binding to the contract
	AbictokenFilterer   // Log filterer for contract events
}

// AbictokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbictokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbictokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbictokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbictokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbictokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbictokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbictokenSession struct {
	Contract     *Abictoken        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbictokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbictokenCallerSession struct {
	Contract *AbictokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AbictokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbictokenTransactorSession struct {
	Contract     *AbictokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AbictokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbictokenRaw struct {
	Contract *Abictoken // Generic contract binding to access the raw methods on
}

// AbictokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbictokenCallerRaw struct {
	Contract *AbictokenCaller // Generic read-only contract binding to access the raw methods on
}

// AbictokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbictokenTransactorRaw struct {
	Contract *AbictokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbictoken creates a new instance of Abictoken, bound to a specific deployed contract.
func NewAbictoken(address common.Address, backend bind.ContractBackend) (*Abictoken, error) {
	contract, err := bindAbictoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abictoken{AbictokenCaller: AbictokenCaller{contract: contract}, AbictokenTransactor: AbictokenTransactor{contract: contract}, AbictokenFilterer: AbictokenFilterer{contract: contract}}, nil
}

// NewAbictokenCaller creates a new read-only instance of Abictoken, bound to a specific deployed contract.
func NewAbictokenCaller(address common.Address, caller bind.ContractCaller) (*AbictokenCaller, error) {
	contract, err := bindAbictoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbictokenCaller{contract: contract}, nil
}

// NewAbictokenTransactor creates a new write-only instance of Abictoken, bound to a specific deployed contract.
func NewAbictokenTransactor(address common.Address, transactor bind.ContractTransactor) (*AbictokenTransactor, error) {
	contract, err := bindAbictoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbictokenTransactor{contract: contract}, nil
}

// NewAbictokenFilterer creates a new log filterer instance of Abictoken, bound to a specific deployed contract.
func NewAbictokenFilterer(address common.Address, filterer bind.ContractFilterer) (*AbictokenFilterer, error) {
	contract, err := bindAbictoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbictokenFilterer{contract: contract}, nil
}

// bindAbictoken binds a generic wrapper to an already deployed contract.
func bindAbictoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbictokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abictoken *AbictokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abictoken.Contract.AbictokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abictoken *AbictokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abictoken.Contract.AbictokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abictoken *AbictokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abictoken.Contract.AbictokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abictoken *AbictokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abictoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abictoken *AbictokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abictoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abictoken *AbictokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abictoken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Abictoken *AbictokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abictoken.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Abictoken *AbictokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Abictoken.Contract.Allowance(&_Abictoken.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Abictoken *AbictokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Abictoken.Contract.Allowance(&_Abictoken.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Abictoken *AbictokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abictoken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Abictoken *AbictokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Abictoken.Contract.BalanceOf(&_Abictoken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Abictoken *AbictokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Abictoken.Contract.BalanceOf(&_Abictoken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abictoken *AbictokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abictoken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abictoken *AbictokenSession) Decimals() (uint8, error) {
	return _Abictoken.Contract.Decimals(&_Abictoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abictoken *AbictokenCallerSession) Decimals() (uint8, error) {
	return _Abictoken.Contract.Decimals(&_Abictoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abictoken *AbictokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abictoken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abictoken *AbictokenSession) Name() (string, error) {
	return _Abictoken.Contract.Name(&_Abictoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abictoken *AbictokenCallerSession) Name() (string, error) {
	return _Abictoken.Contract.Name(&_Abictoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abictoken *AbictokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abictoken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abictoken *AbictokenSession) Symbol() (string, error) {
	return _Abictoken.Contract.Symbol(&_Abictoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abictoken *AbictokenCallerSession) Symbol() (string, error) {
	return _Abictoken.Contract.Symbol(&_Abictoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abictoken *AbictokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abictoken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abictoken *AbictokenSession) TotalSupply() (*big.Int, error) {
	return _Abictoken.Contract.TotalSupply(&_Abictoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abictoken *AbictokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Abictoken.Contract.TotalSupply(&_Abictoken.CallOpts)
}

// AllocateTo is a paid mutator transaction binding the contract method 0x08bca566.
//
// Solidity: function allocateTo(address _owner, uint256 value) returns()
func (_Abictoken *AbictokenTransactor) AllocateTo(opts *bind.TransactOpts, _owner common.Address, value *big.Int) (*types.Transaction, error) {
	return _Abictoken.contract.Transact(opts, "allocateTo", _owner, value)
}

// AllocateTo is a paid mutator transaction binding the contract method 0x08bca566.
//
// Solidity: function allocateTo(address _owner, uint256 value) returns()
func (_Abictoken *AbictokenSession) AllocateTo(_owner common.Address, value *big.Int) (*types.Transaction, error) {
	return _Abictoken.Contract.AllocateTo(&_Abictoken.TransactOpts, _owner, value)
}

// AllocateTo is a paid mutator transaction binding the contract method 0x08bca566.
//
// Solidity: function allocateTo(address _owner, uint256 value) returns()
func (_Abictoken *AbictokenTransactorSession) AllocateTo(_owner common.Address, value *big.Int) (*types.Transaction, error) {
	return _Abictoken.Contract.AllocateTo(&_Abictoken.TransactOpts, _owner, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 amount) returns(bool)
func (_Abictoken *AbictokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abictoken.contract.Transact(opts, "approve", _spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 amount) returns(bool)
func (_Abictoken *AbictokenSession) Approve(_spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abictoken.Contract.Approve(&_Abictoken.TransactOpts, _spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 amount) returns(bool)
func (_Abictoken *AbictokenTransactorSession) Approve(_spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abictoken.Contract.Approve(&_Abictoken.TransactOpts, _spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Abictoken *AbictokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abictoken.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Abictoken *AbictokenSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abictoken.Contract.Transfer(&_Abictoken.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Abictoken *AbictokenTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abictoken.Contract.Transfer(&_Abictoken.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Abictoken *AbictokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abictoken.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Abictoken *AbictokenSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abictoken.Contract.TransferFrom(&_Abictoken.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Abictoken *AbictokenTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abictoken.Contract.TransferFrom(&_Abictoken.TransactOpts, src, dst, amount)
}

// AbictokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Abictoken contract.
type AbictokenApprovalIterator struct {
	Event *AbictokenApproval // Event containing the contract specifics and raw log

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
func (it *AbictokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbictokenApproval)
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
		it.Event = new(AbictokenApproval)
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
func (it *AbictokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbictokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbictokenApproval represents a Approval event raised by the Abictoken contract.
type AbictokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Abictoken *AbictokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AbictokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Abictoken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AbictokenApprovalIterator{contract: _Abictoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Abictoken *AbictokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AbictokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Abictoken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbictokenApproval)
				if err := _Abictoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Abictoken *AbictokenFilterer) ParseApproval(log types.Log) (*AbictokenApproval, error) {
	event := new(AbictokenApproval)
	if err := _Abictoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbictokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Abictoken contract.
type AbictokenTransferIterator struct {
	Event *AbictokenTransfer // Event containing the contract specifics and raw log

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
func (it *AbictokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbictokenTransfer)
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
		it.Event = new(AbictokenTransfer)
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
func (it *AbictokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbictokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbictokenTransfer represents a Transfer event raised by the Abictoken contract.
type AbictokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Abictoken *AbictokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AbictokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Abictoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AbictokenTransferIterator{contract: _Abictoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Abictoken *AbictokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AbictokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Abictoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbictokenTransfer)
				if err := _Abictoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Abictoken *AbictokenFilterer) ParseTransfer(log types.Log) (*AbictokenTransfer, error) {
	event := new(AbictokenTransfer)
	if err := _Abictoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
