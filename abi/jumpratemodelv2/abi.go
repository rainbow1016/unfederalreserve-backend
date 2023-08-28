// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abijumpratemodelv2

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

// Abijumpratemodelv2ABI is the input ABI used to generate the binding from.
const Abijumpratemodelv2ABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseRatePerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplierPerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jumpMultiplierPerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"kink_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseRatePerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"multiplierPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"jumpMultiplierPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"kink\",\"type\":\"uint256\"}],\"name\":\"NewInterestParams\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blocksPerYear\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"name\":\"getBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"getSupplyRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInterestRateModel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"jumpMultiplierPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kink\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"multiplierPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baseRatePerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"multiplierPerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jumpMultiplierPerYear\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"kink_\",\"type\":\"uint256\"}],\"name\":\"updateJumpRateModel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"name\":\"utilizationRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// Abijumpratemodelv2Bin is the compiled bytecode used for deploying new contracts.
var Abijumpratemodelv2Bin = "0x608060405234801561001057600080fd5b50604051610a62380380610a62833981810160405260a081101561003357600080fd5b508051602082015160408301516060840151608090940151600080546001600160a01b0319166001600160a01b03831617905592939192909190848484848461007e8585858561008d565b505050505050505050506102bc565b6100a7622014808561017060201b6105bc1790919060201c565b6002556100f66100c562201480836101c1602090811b61056317901c565b6100e4670de0b6b3a7640000866101c160201b6105631790919060201c565b61017060201b6105bc1790919060201c565b6001556101118262201480610170602090811b6105bc17901c565b60038190556004829055600254600154604080519283526020830191909152818101929092526060810183905290517f6960ab234c7ef4b0c9197100f5393cfcde7c453ac910a27bd2000aa1dd4c068d9181900360800190a150505050565b60006101b883836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f00000000000081525061021a60201b60201c565b90505b92915050565b6000826101d0575060006101bb565b828202828482816101dd57fe5b04146101b85760405162461bcd60e51b8152600401808060200182810382526021815260200180610a416021913960400191505060405180910390fd5b600081836102a65760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561026b578181015183820152602001610253565b50505050905090810190601f1680156102985780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385816102b257fe5b0495945050505050565b610776806102cb6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638da5cb5b116100715780638da5cb5b14610167578063a385fb961461018b578063b816881614610193578063b9f9850a146101c2578063f14039de146101ca578063fd2da339146101d2576100a9565b806315f24053146100ae5780632037f3e7146100e95780632191f92a1461011a5780636e71e2d8146101365780638726bb891461015f575b600080fd5b6100d7600480360360608110156100c457600080fd5b50803590602081013590604001356101da565b60408051918252519081900360200190f35b610118600480360360808110156100ff57600080fd5b50803590602081013590604081013590606001356101f1565b005b61012261024c565b604080519115158252519081900360200190f35b6100d76004803603606081101561014c57600080fd5b5080359060208101359060400135610251565b6100d76102a7565b61016f6102ad565b604080516001600160a01b039092168252519081900360200190f35b6100d76102bc565b6100d7600480360360808110156101a957600080fd5b50803590602081013590604081013590606001356102c3565b6100d7610342565b6100d7610348565b6100d761034e565b60006101e7848484610354565b90505b9392505050565b6000546001600160a01b0316331461023a5760405162461bcd60e51b815260040180806020018281038252602681526020018061071c6026913960400191505060405180910390fd5b6102468484848461041d565b50505050565b600181565b600082610260575060006101ea565b6101e761028383610277878763ffffffff6104be16565b9063ffffffff61052116565b61029b85670de0b6b3a764000063ffffffff61056316565b9063ffffffff6105bc16565b60015481565b6000546001600160a01b031681565b6220148081565b6000806102de670de0b6b3a76400008463ffffffff61052116565b905060006102ed878787610354565b9050600061030d670de0b6b3a764000061029b848663ffffffff61056316565b9050610336670de0b6b3a764000061029b8361032a8c8c8c610251565b9063ffffffff61056316565b98975050505050505050565b60035481565b60025481565b60045481565b600080610362858585610251565b905060045481116103a8576103a0600254610394670de0b6b3a764000061029b6001548661056390919063ffffffff16565b9063ffffffff6104be16565b9150506101ea565b60006103d3600254610394670de0b6b3a764000061029b60015460045461056390919063ffffffff16565b905060006103ec6004548461052190919063ffffffff16565b905061041382610394670de0b6b3a764000061029b6003548661056390919063ffffffff16565b93505050506101ea565b610430846220148063ffffffff6105bc16565b600255610449610283622014808363ffffffff61056316565b60015561045f826220148063ffffffff6105bc16565b60038190556004829055600254600154604080519283526020830191909152818101929092526060810183905290517f6960ab234c7ef4b0c9197100f5393cfcde7c453ac910a27bd2000aa1dd4c068d9181900360800190a150505050565b600082820183811015610518576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b90505b92915050565b600061051883836040518060400160405280601f81526020017f536166654d6174683a207375627472616374696f6e20756e646572666c6f77008152506105fe565b6000826105725750600061051b565b8282028284828161057f57fe5b04146105185760405162461bcd60e51b81526004018080602001828103825260218152602001806106fb6021913960400191505060405180910390fd5b600061051883836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250610695565b6000818484111561068d5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561065257818101518382015260200161063a565b50505050905090810190601f16801561067f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600081836106e45760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561065257818101518382015260200161063a565b5060008385816106f057fe5b049594505050505056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f776f6e6c7920746865206f776e6572206d61792063616c6c20746869732066756e6374696f6e2ea265627a7a72315820bf3a0985459c7ed60c8d343d4a1b214dbf4c26a030239149fdda36fd7f8585c264736f6c63430005100032536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77"

// DeployAbijumpratemodelv2 deploys a new Ethereum contract, binding an instance of Abijumpratemodelv2 to it.
func DeployAbijumpratemodelv2(auth *bind.TransactOpts, backend bind.ContractBackend, baseRatePerYear *big.Int, multiplierPerYear *big.Int, jumpMultiplierPerYear *big.Int, kink_ *big.Int, owner_ common.Address) (common.Address, *types.Transaction, *Abijumpratemodelv2, error) {
	parsed, err := abi.JSON(strings.NewReader(Abijumpratemodelv2ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Abijumpratemodelv2Bin), backend, baseRatePerYear, multiplierPerYear, jumpMultiplierPerYear, kink_, owner_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Abijumpratemodelv2{Abijumpratemodelv2Caller: Abijumpratemodelv2Caller{contract: contract}, Abijumpratemodelv2Transactor: Abijumpratemodelv2Transactor{contract: contract}, Abijumpratemodelv2Filterer: Abijumpratemodelv2Filterer{contract: contract}}, nil
}

// Abijumpratemodelv2 is an auto generated Go binding around an Ethereum contract.
type Abijumpratemodelv2 struct {
	Abijumpratemodelv2Caller     // Read-only binding to the contract
	Abijumpratemodelv2Transactor // Write-only binding to the contract
	Abijumpratemodelv2Filterer   // Log filterer for contract events
}

// Abijumpratemodelv2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Abijumpratemodelv2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abijumpratemodelv2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Abijumpratemodelv2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abijumpratemodelv2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Abijumpratemodelv2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abijumpratemodelv2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Abijumpratemodelv2Session struct {
	Contract     *Abijumpratemodelv2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Abijumpratemodelv2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Abijumpratemodelv2CallerSession struct {
	Contract *Abijumpratemodelv2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// Abijumpratemodelv2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Abijumpratemodelv2TransactorSession struct {
	Contract     *Abijumpratemodelv2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Abijumpratemodelv2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Abijumpratemodelv2Raw struct {
	Contract *Abijumpratemodelv2 // Generic contract binding to access the raw methods on
}

// Abijumpratemodelv2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Abijumpratemodelv2CallerRaw struct {
	Contract *Abijumpratemodelv2Caller // Generic read-only contract binding to access the raw methods on
}

// Abijumpratemodelv2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Abijumpratemodelv2TransactorRaw struct {
	Contract *Abijumpratemodelv2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAbijumpratemodelv2 creates a new instance of Abijumpratemodelv2, bound to a specific deployed contract.
func NewAbijumpratemodelv2(address common.Address, backend bind.ContractBackend) (*Abijumpratemodelv2, error) {
	contract, err := bindAbijumpratemodelv2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abijumpratemodelv2{Abijumpratemodelv2Caller: Abijumpratemodelv2Caller{contract: contract}, Abijumpratemodelv2Transactor: Abijumpratemodelv2Transactor{contract: contract}, Abijumpratemodelv2Filterer: Abijumpratemodelv2Filterer{contract: contract}}, nil
}

// NewAbijumpratemodelv2Caller creates a new read-only instance of Abijumpratemodelv2, bound to a specific deployed contract.
func NewAbijumpratemodelv2Caller(address common.Address, caller bind.ContractCaller) (*Abijumpratemodelv2Caller, error) {
	contract, err := bindAbijumpratemodelv2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Abijumpratemodelv2Caller{contract: contract}, nil
}

// NewAbijumpratemodelv2Transactor creates a new write-only instance of Abijumpratemodelv2, bound to a specific deployed contract.
func NewAbijumpratemodelv2Transactor(address common.Address, transactor bind.ContractTransactor) (*Abijumpratemodelv2Transactor, error) {
	contract, err := bindAbijumpratemodelv2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Abijumpratemodelv2Transactor{contract: contract}, nil
}

// NewAbijumpratemodelv2Filterer creates a new log filterer instance of Abijumpratemodelv2, bound to a specific deployed contract.
func NewAbijumpratemodelv2Filterer(address common.Address, filterer bind.ContractFilterer) (*Abijumpratemodelv2Filterer, error) {
	contract, err := bindAbijumpratemodelv2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Abijumpratemodelv2Filterer{contract: contract}, nil
}

// bindAbijumpratemodelv2 binds a generic wrapper to an already deployed contract.
func bindAbijumpratemodelv2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Abijumpratemodelv2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abijumpratemodelv2 *Abijumpratemodelv2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abijumpratemodelv2.Contract.Abijumpratemodelv2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abijumpratemodelv2 *Abijumpratemodelv2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abijumpratemodelv2.Contract.Abijumpratemodelv2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abijumpratemodelv2 *Abijumpratemodelv2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abijumpratemodelv2.Contract.Abijumpratemodelv2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abijumpratemodelv2 *Abijumpratemodelv2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abijumpratemodelv2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abijumpratemodelv2 *Abijumpratemodelv2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abijumpratemodelv2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abijumpratemodelv2 *Abijumpratemodelv2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abijumpratemodelv2.Contract.contract.Transact(opts, method, params...)
}

// BaseRatePerBlock is a free data retrieval call binding the contract method 0xf14039de.
//
// Solidity: function baseRatePerBlock() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Caller) BaseRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abijumpratemodelv2.contract.Call(opts, &out, "baseRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseRatePerBlock is a free data retrieval call binding the contract method 0xf14039de.
//
// Solidity: function baseRatePerBlock() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Session) BaseRatePerBlock() (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.BaseRatePerBlock(&_Abijumpratemodelv2.CallOpts)
}

// BaseRatePerBlock is a free data retrieval call binding the contract method 0xf14039de.
//
// Solidity: function baseRatePerBlock() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2CallerSession) BaseRatePerBlock() (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.BaseRatePerBlock(&_Abijumpratemodelv2.CallOpts)
}

// BlocksPerYear is a free data retrieval call binding the contract method 0xa385fb96.
//
// Solidity: function blocksPerYear() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Caller) BlocksPerYear(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abijumpratemodelv2.contract.Call(opts, &out, "blocksPerYear")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlocksPerYear is a free data retrieval call binding the contract method 0xa385fb96.
//
// Solidity: function blocksPerYear() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Session) BlocksPerYear() (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.BlocksPerYear(&_Abijumpratemodelv2.CallOpts)
}

// BlocksPerYear is a free data retrieval call binding the contract method 0xa385fb96.
//
// Solidity: function blocksPerYear() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2CallerSession) BlocksPerYear() (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.BlocksPerYear(&_Abijumpratemodelv2.CallOpts)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Caller) GetBorrowRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abijumpratemodelv2.contract.Call(opts, &out, "getBorrowRate", cash, borrows, reserves)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Session) GetBorrowRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.GetBorrowRate(&_Abijumpratemodelv2.CallOpts, cash, borrows, reserves)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2CallerSession) GetBorrowRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.GetBorrowRate(&_Abijumpratemodelv2.CallOpts, cash, borrows, reserves)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Caller) GetSupplyRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abijumpratemodelv2.contract.Call(opts, &out, "getSupplyRate", cash, borrows, reserves, reserveFactorMantissa)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Session) GetSupplyRate(cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.GetSupplyRate(&_Abijumpratemodelv2.CallOpts, cash, borrows, reserves, reserveFactorMantissa)
}

// GetSupplyRate is a free data retrieval call binding the contract method 0xb8168816.
//
// Solidity: function getSupplyRate(uint256 cash, uint256 borrows, uint256 reserves, uint256 reserveFactorMantissa) view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2CallerSession) GetSupplyRate(cash *big.Int, borrows *big.Int, reserves *big.Int, reserveFactorMantissa *big.Int) (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.GetSupplyRate(&_Abijumpratemodelv2.CallOpts, cash, borrows, reserves, reserveFactorMantissa)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Caller) IsInterestRateModel(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abijumpratemodelv2.contract.Call(opts, &out, "isInterestRateModel")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Session) IsInterestRateModel() (bool, error) {
	return _Abijumpratemodelv2.Contract.IsInterestRateModel(&_Abijumpratemodelv2.CallOpts)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_Abijumpratemodelv2 *Abijumpratemodelv2CallerSession) IsInterestRateModel() (bool, error) {
	return _Abijumpratemodelv2.Contract.IsInterestRateModel(&_Abijumpratemodelv2.CallOpts)
}

// JumpMultiplierPerBlock is a free data retrieval call binding the contract method 0xb9f9850a.
//
// Solidity: function jumpMultiplierPerBlock() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Caller) JumpMultiplierPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abijumpratemodelv2.contract.Call(opts, &out, "jumpMultiplierPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// JumpMultiplierPerBlock is a free data retrieval call binding the contract method 0xb9f9850a.
//
// Solidity: function jumpMultiplierPerBlock() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Session) JumpMultiplierPerBlock() (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.JumpMultiplierPerBlock(&_Abijumpratemodelv2.CallOpts)
}

// JumpMultiplierPerBlock is a free data retrieval call binding the contract method 0xb9f9850a.
//
// Solidity: function jumpMultiplierPerBlock() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2CallerSession) JumpMultiplierPerBlock() (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.JumpMultiplierPerBlock(&_Abijumpratemodelv2.CallOpts)
}

// Kink is a free data retrieval call binding the contract method 0xfd2da339.
//
// Solidity: function kink() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Caller) Kink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abijumpratemodelv2.contract.Call(opts, &out, "kink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Kink is a free data retrieval call binding the contract method 0xfd2da339.
//
// Solidity: function kink() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Session) Kink() (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.Kink(&_Abijumpratemodelv2.CallOpts)
}

// Kink is a free data retrieval call binding the contract method 0xfd2da339.
//
// Solidity: function kink() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2CallerSession) Kink() (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.Kink(&_Abijumpratemodelv2.CallOpts)
}

// MultiplierPerBlock is a free data retrieval call binding the contract method 0x8726bb89.
//
// Solidity: function multiplierPerBlock() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Caller) MultiplierPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abijumpratemodelv2.contract.Call(opts, &out, "multiplierPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MultiplierPerBlock is a free data retrieval call binding the contract method 0x8726bb89.
//
// Solidity: function multiplierPerBlock() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Session) MultiplierPerBlock() (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.MultiplierPerBlock(&_Abijumpratemodelv2.CallOpts)
}

// MultiplierPerBlock is a free data retrieval call binding the contract method 0x8726bb89.
//
// Solidity: function multiplierPerBlock() view returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2CallerSession) MultiplierPerBlock() (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.MultiplierPerBlock(&_Abijumpratemodelv2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abijumpratemodelv2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Session) Owner() (common.Address, error) {
	return _Abijumpratemodelv2.Contract.Owner(&_Abijumpratemodelv2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abijumpratemodelv2 *Abijumpratemodelv2CallerSession) Owner() (common.Address, error) {
	return _Abijumpratemodelv2.Contract.Owner(&_Abijumpratemodelv2.CallOpts)
}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Caller) UtilizationRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Abijumpratemodelv2.contract.Call(opts, &out, "utilizationRate", cash, borrows, reserves)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Session) UtilizationRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.UtilizationRate(&_Abijumpratemodelv2.CallOpts, cash, borrows, reserves)
}

// UtilizationRate is a free data retrieval call binding the contract method 0x6e71e2d8.
//
// Solidity: function utilizationRate(uint256 cash, uint256 borrows, uint256 reserves) pure returns(uint256)
func (_Abijumpratemodelv2 *Abijumpratemodelv2CallerSession) UtilizationRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, error) {
	return _Abijumpratemodelv2.Contract.UtilizationRate(&_Abijumpratemodelv2.CallOpts, cash, borrows, reserves)
}

// UpdateJumpRateModel is a paid mutator transaction binding the contract method 0x2037f3e7.
//
// Solidity: function updateJumpRateModel(uint256 baseRatePerYear, uint256 multiplierPerYear, uint256 jumpMultiplierPerYear, uint256 kink_) returns()
func (_Abijumpratemodelv2 *Abijumpratemodelv2Transactor) UpdateJumpRateModel(opts *bind.TransactOpts, baseRatePerYear *big.Int, multiplierPerYear *big.Int, jumpMultiplierPerYear *big.Int, kink_ *big.Int) (*types.Transaction, error) {
	return _Abijumpratemodelv2.contract.Transact(opts, "updateJumpRateModel", baseRatePerYear, multiplierPerYear, jumpMultiplierPerYear, kink_)
}

// UpdateJumpRateModel is a paid mutator transaction binding the contract method 0x2037f3e7.
//
// Solidity: function updateJumpRateModel(uint256 baseRatePerYear, uint256 multiplierPerYear, uint256 jumpMultiplierPerYear, uint256 kink_) returns()
func (_Abijumpratemodelv2 *Abijumpratemodelv2Session) UpdateJumpRateModel(baseRatePerYear *big.Int, multiplierPerYear *big.Int, jumpMultiplierPerYear *big.Int, kink_ *big.Int) (*types.Transaction, error) {
	return _Abijumpratemodelv2.Contract.UpdateJumpRateModel(&_Abijumpratemodelv2.TransactOpts, baseRatePerYear, multiplierPerYear, jumpMultiplierPerYear, kink_)
}

// UpdateJumpRateModel is a paid mutator transaction binding the contract method 0x2037f3e7.
//
// Solidity: function updateJumpRateModel(uint256 baseRatePerYear, uint256 multiplierPerYear, uint256 jumpMultiplierPerYear, uint256 kink_) returns()
func (_Abijumpratemodelv2 *Abijumpratemodelv2TransactorSession) UpdateJumpRateModel(baseRatePerYear *big.Int, multiplierPerYear *big.Int, jumpMultiplierPerYear *big.Int, kink_ *big.Int) (*types.Transaction, error) {
	return _Abijumpratemodelv2.Contract.UpdateJumpRateModel(&_Abijumpratemodelv2.TransactOpts, baseRatePerYear, multiplierPerYear, jumpMultiplierPerYear, kink_)
}

// Abijumpratemodelv2NewInterestParamsIterator is returned from FilterNewInterestParams and is used to iterate over the raw logs and unpacked data for NewInterestParams events raised by the Abijumpratemodelv2 contract.
type Abijumpratemodelv2NewInterestParamsIterator struct {
	Event *Abijumpratemodelv2NewInterestParams // Event containing the contract specifics and raw log

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
func (it *Abijumpratemodelv2NewInterestParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abijumpratemodelv2NewInterestParams)
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
		it.Event = new(Abijumpratemodelv2NewInterestParams)
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
func (it *Abijumpratemodelv2NewInterestParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abijumpratemodelv2NewInterestParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abijumpratemodelv2NewInterestParams represents a NewInterestParams event raised by the Abijumpratemodelv2 contract.
type Abijumpratemodelv2NewInterestParams struct {
	BaseRatePerBlock       *big.Int
	MultiplierPerBlock     *big.Int
	JumpMultiplierPerBlock *big.Int
	Kink                   *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewInterestParams is a free log retrieval operation binding the contract event 0x6960ab234c7ef4b0c9197100f5393cfcde7c453ac910a27bd2000aa1dd4c068d.
//
// Solidity: event NewInterestParams(uint256 baseRatePerBlock, uint256 multiplierPerBlock, uint256 jumpMultiplierPerBlock, uint256 kink)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Filterer) FilterNewInterestParams(opts *bind.FilterOpts) (*Abijumpratemodelv2NewInterestParamsIterator, error) {

	logs, sub, err := _Abijumpratemodelv2.contract.FilterLogs(opts, "NewInterestParams")
	if err != nil {
		return nil, err
	}
	return &Abijumpratemodelv2NewInterestParamsIterator{contract: _Abijumpratemodelv2.contract, event: "NewInterestParams", logs: logs, sub: sub}, nil
}

// WatchNewInterestParams is a free log subscription operation binding the contract event 0x6960ab234c7ef4b0c9197100f5393cfcde7c453ac910a27bd2000aa1dd4c068d.
//
// Solidity: event NewInterestParams(uint256 baseRatePerBlock, uint256 multiplierPerBlock, uint256 jumpMultiplierPerBlock, uint256 kink)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Filterer) WatchNewInterestParams(opts *bind.WatchOpts, sink chan<- *Abijumpratemodelv2NewInterestParams) (event.Subscription, error) {

	logs, sub, err := _Abijumpratemodelv2.contract.WatchLogs(opts, "NewInterestParams")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abijumpratemodelv2NewInterestParams)
				if err := _Abijumpratemodelv2.contract.UnpackLog(event, "NewInterestParams", log); err != nil {
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

// ParseNewInterestParams is a log parse operation binding the contract event 0x6960ab234c7ef4b0c9197100f5393cfcde7c453ac910a27bd2000aa1dd4c068d.
//
// Solidity: event NewInterestParams(uint256 baseRatePerBlock, uint256 multiplierPerBlock, uint256 jumpMultiplierPerBlock, uint256 kink)
func (_Abijumpratemodelv2 *Abijumpratemodelv2Filterer) ParseNewInterestParams(log types.Log) (*Abijumpratemodelv2NewInterestParams, error) {
	event := new(Abijumpratemodelv2NewInterestParams)
	if err := _Abijumpratemodelv2.contract.UnpackLog(event, "NewInterestParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
