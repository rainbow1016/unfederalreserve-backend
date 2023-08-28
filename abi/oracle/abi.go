// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abioracle

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

// AbioracleABI is the input ABI used to generate the binding from.
const AbioracleABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feed\",\"type\":\"address\"}],\"name\":\"ChainlinkPriceFeedPosted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestedPriceMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPriceMantissa\",\"type\":\"uint256\"}],\"name\":\"PricePosted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"assetPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"getUnderlyingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPriceOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"feed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplierMantissa\",\"type\":\"uint256\"}],\"name\":\"setDirectChainlinkFeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setDirectPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"feed\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"multiplierMantissa\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingChainlinkFeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"underlyingPriceMantissa\",\"type\":\"uint256\"}],\"name\":\"setUnderlyingPrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AbioracleBin is the compiled bytecode used for deploying new contracts.
var AbioracleBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03191633179055610db2806100326000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80635e9a523c116100715780635e9a523c1461019857806366331bba146101d0578063b71d1a0c146101ec578063e9c714f214610212578063f851a4401461021a578063fc57d4df14610222576100a9565b806309a8acb0146100ae578063127ffda0146100dc57806326782247146101085780632ad0e2cf1461012c578063359e0d8114610162575b600080fd5b6100da600480360360408110156100c457600080fd5b506001600160a01b038135169060200135610248565b005b6100da600480360360408110156100f257600080fd5b506001600160a01b03813516906020013561031c565b61011061045a565b604080516001600160a01b039092168252519081900360200190f35b6100da6004803603606081101561014257600080fd5b506001600160a01b03813581169160208101359091169060400135610469565b6100da6004803603606081101561017857600080fd5b506001600160a01b03813581169160208101359091169060400135610554565b6101be600480360360208110156101ae57600080fd5b50356001600160a01b0316610622565b60408051918252519081900360200190f35b6101d861077a565b604080519115158252519081900360200190f35b6100da6004803603602081101561020257600080fd5b50356001600160a01b031661077f565b6100da6107ef565b61011061086f565b6101be6004803603602081101561023857600080fd5b50356001600160a01b031661087e565b6000546001600160a01b031633146102a2576040805162461bcd60e51b81526020600482015260186024820152776f6e6c792061646d696e2063616e2073657420707269636560401b604482015290519081900360640190fd5b6001600160a01b0382166000818152600260209081526040918290205482519384529083015281810183905260608201839052517fdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae79181900360800190a16001600160a01b03909116600090815260026020526040902055565b6000546001600160a01b03163314610376576040805162461bcd60e51b81526020600482015260186024820152776f6e6c792061646d696e2063616e2073657420707269636560401b604482015290519081900360640190fd5b6000826001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156103b157600080fd5b505afa1580156103c5573d6000803e3d6000fd5b505050506040513d60208110156103db57600080fd5b50516001600160a01b0381166000818152600260209081526040918290205482519384529083015281810185905260608201859052519192507fdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7919081900360800190a16001600160a01b031660009081526002602052604090205550565b6001546001600160a01b031681565b6000546001600160a01b031633146104c3576040805162461bcd60e51b81526020600482015260186024820152776f6e6c792061646d696e2063616e2073657420707269636560401b604482015290519081900360640190fd5b604080516001600160a01b0380861682528416602082015281517f796834b03c5e453a88d26a47d77d352e37f72ddada3d31c8117add4f445cd8a5929181900390910190a16040805180820182526001600160a01b0393841681526020808201938452948416600090815260039095529320925183546001600160a01b031916921691909117825551600190910155565b6000546001600160a01b031633146105ae576040805162461bcd60e51b81526020600482015260186024820152776f6e6c792061646d696e2063616e2073657420707269636560401b604482015290519081900360640190fd5b61061d836001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b1580156105ea57600080fd5b505afa1580156105fe573d6000803e3d6000fd5b505050506040513d602081101561061457600080fd5b50518383610469565b505050565b6001600160a01b03808216600090815260036020526040812080549192909116156107595760008060008060008560000160009054906101000a90046001600160a01b03166001600160a01b031663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b15801561069f57600080fd5b505afa1580156106b3573d6000803e3d6000fd5b505050506040513d60a08110156106c957600080fd5b50805160208201516040830151606084015160809094015192985090965094509092509050600084121561073e576040805162461bcd60e51b815260206004820152601760248201527670726963652063616e2774206265206e6567617469766560481b604482015290519081900360640190fd5b61074c848760010154610b91565b9650505050505050610775565b50506001600160a01b0381166000908152600260205260409020545b919050565b600181565b6000546001600160a01b031633146107cd576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5e995960a21b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b03163314801561080857503315155b610848576040805162461bcd60e51b815260206004820152600c60248201526b1d5b985d5d1a1bdc9a5e995960a21b604482015290519081900360640190fd5b60018054600080546001600160a01b03199081166001600160a01b03841617909155169055565b6000546001600160a01b031681565b6000806109c7836001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156108bd57600080fd5b505afa1580156108d1573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156108fa57600080fd5b810190808051604051939291908464010000000082111561091a57600080fd5b90830190602082018581111561092f57600080fd5b825164010000000081118282018810171561094957600080fd5b82525081516020918201929091019080838360005b8381101561097657818101518382015260200161095e565b50505050905090810190601f1680156109a35780820380516001836020036101000a031916815260200191505b506040818101905260058152640eadc8aa8960db1b60208201529250610bda915050565b156109d457506000610a3c565b826001600160a01b0316636f307dc36040518163ffffffff1660e01b815260040160206040518083038186803b158015610a0d57600080fd5b505afa158015610a21573d6000803e3d6000fd5b505050506040513d6020811015610a3757600080fd5b505190505b6001600160a01b038082166000908152600360205260409020805490911615610b725760008060008060008560000160009054906101000a90046001600160a01b03166001600160a01b031663feaf968c6040518163ffffffff1660e01b815260040160a06040518083038186803b158015610ab757600080fd5b505afa158015610acb573d6000803e3d6000fd5b505050506040513d60a0811015610ae157600080fd5b508051602082015160408301516060840151608090940151929850909650945090925090506000841215610b56576040805162461bcd60e51b815260206004820152601760248201527670726963652063616e2774206265206e6567617469766560481b604482015290519081900360640190fd5b610b64848760010154610b91565b975050505050505050610775565b506001600160a01b031660009081526002602052604090205492915050565b6000610bd383836040518060400160405280601781526020017f6d756c7469706c69636174696f6e206f766572666c6f77000000000000000000815250610cc1565b9392505050565b6000816040516020018082805190602001908083835b60208310610c0f5780518252601f199092019160209182019101610bf0565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405160208183030381529060405280519060200120836040516020018082805190602001908083835b60208310610c7d5780518252601f199092019160209182019101610c5e565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040516020818303038152906040528051906020012014905092915050565b6000831580610cce575082155b15610cdb57506000610bd3565b83830283858281610ce857fe5b04148390610d745760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610d39578181015183820152602001610d21565b50505050905090810190601f168015610d665780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5094935050505056fea265627a7a72315820f848e1ed687399e7224db735b8e52238b542b234bfdf4f8dc144062696dba04664736f6c63430005100032"

// DeployAbioracle deploys a new Ethereum contract, binding an instance of Abioracle to it.
func DeployAbioracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Abioracle, error) {
	parsed, err := abi.JSON(strings.NewReader(AbioracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbioracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Abioracle{AbioracleCaller: AbioracleCaller{contract: contract}, AbioracleTransactor: AbioracleTransactor{contract: contract}, AbioracleFilterer: AbioracleFilterer{contract: contract}}, nil
}

// Abioracle is an auto generated Go binding around an Ethereum contract.
type Abioracle struct {
	AbioracleCaller     // Read-only binding to the contract
	AbioracleTransactor // Write-only binding to the contract
	AbioracleFilterer   // Log filterer for contract events
}

// AbioracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbioracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbioracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbioracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbioracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbioracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbioracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbioracleSession struct {
	Contract     *Abioracle        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbioracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbioracleCallerSession struct {
	Contract *AbioracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AbioracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbioracleTransactorSession struct {
	Contract     *AbioracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AbioracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbioracleRaw struct {
	Contract *Abioracle // Generic contract binding to access the raw methods on
}

// AbioracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbioracleCallerRaw struct {
	Contract *AbioracleCaller // Generic read-only contract binding to access the raw methods on
}

// AbioracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbioracleTransactorRaw struct {
	Contract *AbioracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbioracle creates a new instance of Abioracle, bound to a specific deployed contract.
func NewAbioracle(address common.Address, backend bind.ContractBackend) (*Abioracle, error) {
	contract, err := bindAbioracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abioracle{AbioracleCaller: AbioracleCaller{contract: contract}, AbioracleTransactor: AbioracleTransactor{contract: contract}, AbioracleFilterer: AbioracleFilterer{contract: contract}}, nil
}

// NewAbioracleCaller creates a new read-only instance of Abioracle, bound to a specific deployed contract.
func NewAbioracleCaller(address common.Address, caller bind.ContractCaller) (*AbioracleCaller, error) {
	contract, err := bindAbioracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbioracleCaller{contract: contract}, nil
}

// NewAbioracleTransactor creates a new write-only instance of Abioracle, bound to a specific deployed contract.
func NewAbioracleTransactor(address common.Address, transactor bind.ContractTransactor) (*AbioracleTransactor, error) {
	contract, err := bindAbioracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbioracleTransactor{contract: contract}, nil
}

// NewAbioracleFilterer creates a new log filterer instance of Abioracle, bound to a specific deployed contract.
func NewAbioracleFilterer(address common.Address, filterer bind.ContractFilterer) (*AbioracleFilterer, error) {
	contract, err := bindAbioracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbioracleFilterer{contract: contract}, nil
}

// bindAbioracle binds a generic wrapper to an already deployed contract.
func bindAbioracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbioracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abioracle *AbioracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abioracle.Contract.AbioracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abioracle *AbioracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abioracle.Contract.AbioracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abioracle *AbioracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abioracle.Contract.AbioracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abioracle *AbioracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abioracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abioracle *AbioracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abioracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abioracle *AbioracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abioracle.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abioracle *AbioracleCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abioracle.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abioracle *AbioracleSession) Admin() (common.Address, error) {
	return _Abioracle.Contract.Admin(&_Abioracle.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abioracle *AbioracleCallerSession) Admin() (common.Address, error) {
	return _Abioracle.Contract.Admin(&_Abioracle.CallOpts)
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_Abioracle *AbioracleCaller) AssetPrices(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abioracle.contract.Call(opts, &out, "assetPrices", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_Abioracle *AbioracleSession) AssetPrices(asset common.Address) (*big.Int, error) {
	return _Abioracle.Contract.AssetPrices(&_Abioracle.CallOpts, asset)
}

// AssetPrices is a free data retrieval call binding the contract method 0x5e9a523c.
//
// Solidity: function assetPrices(address asset) view returns(uint256)
func (_Abioracle *AbioracleCallerSession) AssetPrices(asset common.Address) (*big.Int, error) {
	return _Abioracle.Contract.AssetPrices(&_Abioracle.CallOpts, asset)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_Abioracle *AbioracleCaller) GetUnderlyingPrice(opts *bind.CallOpts, cToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abioracle.contract.Call(opts, &out, "getUnderlyingPrice", cToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_Abioracle *AbioracleSession) GetUnderlyingPrice(cToken common.Address) (*big.Int, error) {
	return _Abioracle.Contract.GetUnderlyingPrice(&_Abioracle.CallOpts, cToken)
}

// GetUnderlyingPrice is a free data retrieval call binding the contract method 0xfc57d4df.
//
// Solidity: function getUnderlyingPrice(address cToken) view returns(uint256)
func (_Abioracle *AbioracleCallerSession) GetUnderlyingPrice(cToken common.Address) (*big.Int, error) {
	return _Abioracle.Contract.GetUnderlyingPrice(&_Abioracle.CallOpts, cToken)
}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_Abioracle *AbioracleCaller) IsPriceOracle(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abioracle.contract.Call(opts, &out, "isPriceOracle")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_Abioracle *AbioracleSession) IsPriceOracle() (bool, error) {
	return _Abioracle.Contract.IsPriceOracle(&_Abioracle.CallOpts)
}

// IsPriceOracle is a free data retrieval call binding the contract method 0x66331bba.
//
// Solidity: function isPriceOracle() view returns(bool)
func (_Abioracle *AbioracleCallerSession) IsPriceOracle() (bool, error) {
	return _Abioracle.Contract.IsPriceOracle(&_Abioracle.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abioracle *AbioracleCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abioracle.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abioracle *AbioracleSession) PendingAdmin() (common.Address, error) {
	return _Abioracle.Contract.PendingAdmin(&_Abioracle.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abioracle *AbioracleCallerSession) PendingAdmin() (common.Address, error) {
	return _Abioracle.Contract.PendingAdmin(&_Abioracle.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns()
func (_Abioracle *AbioracleTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abioracle.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns()
func (_Abioracle *AbioracleSession) AcceptAdmin() (*types.Transaction, error) {
	return _Abioracle.Contract.AcceptAdmin(&_Abioracle.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns()
func (_Abioracle *AbioracleTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Abioracle.Contract.AcceptAdmin(&_Abioracle.TransactOpts)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns()
func (_Abioracle *AbioracleTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abioracle.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns()
func (_Abioracle *AbioracleSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abioracle.Contract.SetPendingAdmin(&_Abioracle.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns()
func (_Abioracle *AbioracleTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abioracle.Contract.SetPendingAdmin(&_Abioracle.TransactOpts, newPendingAdmin)
}

// SetDirectChainlinkFeed is a paid mutator transaction binding the contract method 0x2ad0e2cf.
//
// Solidity: function setDirectChainlinkFeed(address asset, address feed, uint256 multiplierMantissa) returns()
func (_Abioracle *AbioracleTransactor) SetDirectChainlinkFeed(opts *bind.TransactOpts, asset common.Address, feed common.Address, multiplierMantissa *big.Int) (*types.Transaction, error) {
	return _Abioracle.contract.Transact(opts, "setDirectChainlinkFeed", asset, feed, multiplierMantissa)
}

// SetDirectChainlinkFeed is a paid mutator transaction binding the contract method 0x2ad0e2cf.
//
// Solidity: function setDirectChainlinkFeed(address asset, address feed, uint256 multiplierMantissa) returns()
func (_Abioracle *AbioracleSession) SetDirectChainlinkFeed(asset common.Address, feed common.Address, multiplierMantissa *big.Int) (*types.Transaction, error) {
	return _Abioracle.Contract.SetDirectChainlinkFeed(&_Abioracle.TransactOpts, asset, feed, multiplierMantissa)
}

// SetDirectChainlinkFeed is a paid mutator transaction binding the contract method 0x2ad0e2cf.
//
// Solidity: function setDirectChainlinkFeed(address asset, address feed, uint256 multiplierMantissa) returns()
func (_Abioracle *AbioracleTransactorSession) SetDirectChainlinkFeed(asset common.Address, feed common.Address, multiplierMantissa *big.Int) (*types.Transaction, error) {
	return _Abioracle.Contract.SetDirectChainlinkFeed(&_Abioracle.TransactOpts, asset, feed, multiplierMantissa)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_Abioracle *AbioracleTransactor) SetDirectPrice(opts *bind.TransactOpts, asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _Abioracle.contract.Transact(opts, "setDirectPrice", asset, price)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_Abioracle *AbioracleSession) SetDirectPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _Abioracle.Contract.SetDirectPrice(&_Abioracle.TransactOpts, asset, price)
}

// SetDirectPrice is a paid mutator transaction binding the contract method 0x09a8acb0.
//
// Solidity: function setDirectPrice(address asset, uint256 price) returns()
func (_Abioracle *AbioracleTransactorSession) SetDirectPrice(asset common.Address, price *big.Int) (*types.Transaction, error) {
	return _Abioracle.Contract.SetDirectPrice(&_Abioracle.TransactOpts, asset, price)
}

// SetUnderlyingChainlinkFeed is a paid mutator transaction binding the contract method 0x359e0d81.
//
// Solidity: function setUnderlyingChainlinkFeed(address cToken, address feed, uint256 multiplierMantissa) returns()
func (_Abioracle *AbioracleTransactor) SetUnderlyingChainlinkFeed(opts *bind.TransactOpts, cToken common.Address, feed common.Address, multiplierMantissa *big.Int) (*types.Transaction, error) {
	return _Abioracle.contract.Transact(opts, "setUnderlyingChainlinkFeed", cToken, feed, multiplierMantissa)
}

// SetUnderlyingChainlinkFeed is a paid mutator transaction binding the contract method 0x359e0d81.
//
// Solidity: function setUnderlyingChainlinkFeed(address cToken, address feed, uint256 multiplierMantissa) returns()
func (_Abioracle *AbioracleSession) SetUnderlyingChainlinkFeed(cToken common.Address, feed common.Address, multiplierMantissa *big.Int) (*types.Transaction, error) {
	return _Abioracle.Contract.SetUnderlyingChainlinkFeed(&_Abioracle.TransactOpts, cToken, feed, multiplierMantissa)
}

// SetUnderlyingChainlinkFeed is a paid mutator transaction binding the contract method 0x359e0d81.
//
// Solidity: function setUnderlyingChainlinkFeed(address cToken, address feed, uint256 multiplierMantissa) returns()
func (_Abioracle *AbioracleTransactorSession) SetUnderlyingChainlinkFeed(cToken common.Address, feed common.Address, multiplierMantissa *big.Int) (*types.Transaction, error) {
	return _Abioracle.Contract.SetUnderlyingChainlinkFeed(&_Abioracle.TransactOpts, cToken, feed, multiplierMantissa)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address cToken, uint256 underlyingPriceMantissa) returns()
func (_Abioracle *AbioracleTransactor) SetUnderlyingPrice(opts *bind.TransactOpts, cToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _Abioracle.contract.Transact(opts, "setUnderlyingPrice", cToken, underlyingPriceMantissa)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address cToken, uint256 underlyingPriceMantissa) returns()
func (_Abioracle *AbioracleSession) SetUnderlyingPrice(cToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _Abioracle.Contract.SetUnderlyingPrice(&_Abioracle.TransactOpts, cToken, underlyingPriceMantissa)
}

// SetUnderlyingPrice is a paid mutator transaction binding the contract method 0x127ffda0.
//
// Solidity: function setUnderlyingPrice(address cToken, uint256 underlyingPriceMantissa) returns()
func (_Abioracle *AbioracleTransactorSession) SetUnderlyingPrice(cToken common.Address, underlyingPriceMantissa *big.Int) (*types.Transaction, error) {
	return _Abioracle.Contract.SetUnderlyingPrice(&_Abioracle.TransactOpts, cToken, underlyingPriceMantissa)
}

// AbioracleChainlinkPriceFeedPostedIterator is returned from FilterChainlinkPriceFeedPosted and is used to iterate over the raw logs and unpacked data for ChainlinkPriceFeedPosted events raised by the Abioracle contract.
type AbioracleChainlinkPriceFeedPostedIterator struct {
	Event *AbioracleChainlinkPriceFeedPosted // Event containing the contract specifics and raw log

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
func (it *AbioracleChainlinkPriceFeedPostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbioracleChainlinkPriceFeedPosted)
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
		it.Event = new(AbioracleChainlinkPriceFeedPosted)
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
func (it *AbioracleChainlinkPriceFeedPostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbioracleChainlinkPriceFeedPostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbioracleChainlinkPriceFeedPosted represents a ChainlinkPriceFeedPosted event raised by the Abioracle contract.
type AbioracleChainlinkPriceFeedPosted struct {
	Asset common.Address
	Feed  common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterChainlinkPriceFeedPosted is a free log retrieval operation binding the contract event 0x796834b03c5e453a88d26a47d77d352e37f72ddada3d31c8117add4f445cd8a5.
//
// Solidity: event ChainlinkPriceFeedPosted(address asset, address feed)
func (_Abioracle *AbioracleFilterer) FilterChainlinkPriceFeedPosted(opts *bind.FilterOpts) (*AbioracleChainlinkPriceFeedPostedIterator, error) {

	logs, sub, err := _Abioracle.contract.FilterLogs(opts, "ChainlinkPriceFeedPosted")
	if err != nil {
		return nil, err
	}
	return &AbioracleChainlinkPriceFeedPostedIterator{contract: _Abioracle.contract, event: "ChainlinkPriceFeedPosted", logs: logs, sub: sub}, nil
}

// WatchChainlinkPriceFeedPosted is a free log subscription operation binding the contract event 0x796834b03c5e453a88d26a47d77d352e37f72ddada3d31c8117add4f445cd8a5.
//
// Solidity: event ChainlinkPriceFeedPosted(address asset, address feed)
func (_Abioracle *AbioracleFilterer) WatchChainlinkPriceFeedPosted(opts *bind.WatchOpts, sink chan<- *AbioracleChainlinkPriceFeedPosted) (event.Subscription, error) {

	logs, sub, err := _Abioracle.contract.WatchLogs(opts, "ChainlinkPriceFeedPosted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbioracleChainlinkPriceFeedPosted)
				if err := _Abioracle.contract.UnpackLog(event, "ChainlinkPriceFeedPosted", log); err != nil {
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

// ParseChainlinkPriceFeedPosted is a log parse operation binding the contract event 0x796834b03c5e453a88d26a47d77d352e37f72ddada3d31c8117add4f445cd8a5.
//
// Solidity: event ChainlinkPriceFeedPosted(address asset, address feed)
func (_Abioracle *AbioracleFilterer) ParseChainlinkPriceFeedPosted(log types.Log) (*AbioracleChainlinkPriceFeedPosted, error) {
	event := new(AbioracleChainlinkPriceFeedPosted)
	if err := _Abioracle.contract.UnpackLog(event, "ChainlinkPriceFeedPosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbioraclePricePostedIterator is returned from FilterPricePosted and is used to iterate over the raw logs and unpacked data for PricePosted events raised by the Abioracle contract.
type AbioraclePricePostedIterator struct {
	Event *AbioraclePricePosted // Event containing the contract specifics and raw log

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
func (it *AbioraclePricePostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbioraclePricePosted)
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
		it.Event = new(AbioraclePricePosted)
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
func (it *AbioraclePricePostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbioraclePricePostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbioraclePricePosted represents a PricePosted event raised by the Abioracle contract.
type AbioraclePricePosted struct {
	Asset                  common.Address
	PreviousPriceMantissa  *big.Int
	RequestedPriceMantissa *big.Int
	NewPriceMantissa       *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterPricePosted is a free log retrieval operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_Abioracle *AbioracleFilterer) FilterPricePosted(opts *bind.FilterOpts) (*AbioraclePricePostedIterator, error) {

	logs, sub, err := _Abioracle.contract.FilterLogs(opts, "PricePosted")
	if err != nil {
		return nil, err
	}
	return &AbioraclePricePostedIterator{contract: _Abioracle.contract, event: "PricePosted", logs: logs, sub: sub}, nil
}

// WatchPricePosted is a free log subscription operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_Abioracle *AbioracleFilterer) WatchPricePosted(opts *bind.WatchOpts, sink chan<- *AbioraclePricePosted) (event.Subscription, error) {

	logs, sub, err := _Abioracle.contract.WatchLogs(opts, "PricePosted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbioraclePricePosted)
				if err := _Abioracle.contract.UnpackLog(event, "PricePosted", log); err != nil {
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

// ParsePricePosted is a log parse operation binding the contract event 0xdd71a1d19fcba687442a1d5c58578f1e409af71a79d10fd95a4d66efd8fa9ae7.
//
// Solidity: event PricePosted(address asset, uint256 previousPriceMantissa, uint256 requestedPriceMantissa, uint256 newPriceMantissa)
func (_Abioracle *AbioracleFilterer) ParsePricePosted(log types.Log) (*AbioraclePricePosted, error) {
	event := new(AbioraclePricePosted)
	if err := _Abioracle.contract.UnpackLog(event, "PricePosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
