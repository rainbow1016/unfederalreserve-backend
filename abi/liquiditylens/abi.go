// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abiliquiditylens

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

// LensCTokenBalance is an auto generated low-level Go binding around an user-defined struct.
type LensCTokenBalance struct {
	CToken              common.Address
	BalanceOf           *big.Int
	BorrowBalance       *big.Int
	BalanceOfUnderlying *big.Int
	CollateralValue     *big.Int
	BorrowValue         *big.Int
	CollateralFactor    *big.Int
}

// LensCTokenBalances is an auto generated low-level Go binding around an user-defined struct.
type LensCTokenBalances struct {
	Balances      []LensCTokenBalance
	SumCollateral *big.Int
	SumBorrow     *big.Int
}

// AbiliquiditylensABI is the input ABI used to generate the binding from.
const AbiliquiditylensABI = "[{\"inputs\":[{\"internalType\":\"contractComptroller\",\"name\":\"comp\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"featureflags\",\"type\":\"uint8\"}],\"name\":\"cTokenBalances\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balanceOf\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceOfUnderlying\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactor\",\"type\":\"uint256\"}],\"internalType\":\"structLens.CTokenBalance[]\",\"name\":\"balances\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"sumCollateral\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sumBorrow\",\"type\":\"uint256\"}],\"internalType\":\"structLens.CTokenBalances\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getBorrowBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AbiliquiditylensBin is the compiled bytecode used for deploying new contracts.
var AbiliquiditylensBin = "0x608060405234801561001057600080fd5b50610df3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063118e31b7146100465780639cb376271461006c578063ac41865a1461008c575b600080fd5b610059610054366004610b0d565b61009f565b6040519081526020015b60405180910390f35b61007f61007a366004610b45565b610127565b6040516100639190610c2c565b61005961009a366004610b0d565b610640565b6040516305eff7ef60e21b81526001600160a01b038281166004830152600091908416906317bfdfbc90602401602060405180830381600087803b1580156100e657600080fd5b505af11580156100fa573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061011e9190610b94565b90505b92915050565b61014b60405180606001604052806060815260200160008152602001600081525090565b6000846001600160a01b031663b0772d0b6040518163ffffffff1660e01b815260040160006040518083038186803b15801561018657600080fd5b505afa15801561019a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526101c29190810190610a0b565b90506000856001600160a01b0316637dc0d1d06040518163ffffffff1660e01b815260040160206040518083038186803b1580156101ff57600080fd5b505afa158015610213573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061023791906109ef565b905061025d60405180606001604052806060815260200160008152602001600081525090565b825167ffffffffffffffff81111561028557634e487b7160e01b600052604160045260246000fd5b6040519080825280602002602001820160405280156102be57816020015b6102ab610989565b8152602001906001900390816102a35790505b50815260005b83518110156106335760006001600160a01b03168482815181106102f857634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b0316141561031457610621565b600084828151811061033657634e487b7160e01b600052603260045260246000fd5b60200260200101519050610348610989565b6001600160a01b038281168083526040516305eff7ef60e21b8152918b166004830152906317bfdfbc90602401602060405180830381600087803b15801561038f57600080fd5b505af11580156103a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103c79190610b94565b60408281019190915251633af9e66960e01b81526001600160a01b038a81166004830152831690633af9e66990602401602060405180830381600087803b15801561041157600080fd5b505af1158015610425573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104499190610b94565b60608201526040516370a0823160e01b81526001600160a01b038a811660048301528316906370a082319060240160206040518083038186803b15801561048f57600080fd5b505afa1580156104a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104c79190610b94565b6020820152604051638e8f294b60e01b81526001600160a01b0383811660048301528b1690638e8f294b9060240160606040518083038186803b15801561050d57600080fd5b505afa158015610521573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105459190610ad2565b5060c083019081526040805160208101909152905181529050600061056a8785610640565b90508061057a5750505050610621565b6040805160208101909152818152600061059484836107d3565b90506105a481866060015161081b565b6080860181905260208901516105ba919061083b565b602089015260408501516105cf90839061081b565b60a0860181905260408901516105e5919061083b565b6040890152875180518691908990811061060f57634e487b7160e01b600052603260045260246000fd5b60200260200101819052505050505050505b8061062b81610d5e565b9150506102c4565b50925050505b9392505050565b604080516001600160a01b0383811660208301526000928392839287169163fc57d4df60e01b910160408051601f19818403018152908290526106869291602001610bac565b60408051601f19818403018152908290526106a091610bdd565b6000604051808303816000865af19150503d80600081146106dd576040519150601f19603f3d011682016040523d82523d6000602084013e6106e2565b606091505b50915091508180156106f5575060008151115b15610717578080602001905181019061070e9190610b94565b92505050610121565b60408051600481526024810182526020810180516001600160e01b0316636f307dc360e01b17905290516001600160a01b0386169161075591610bdd565b6000604051808303816000865af19150503d8060008114610792576040519150601f19603f3d011682016040523d82523d6000602084013e610797565b606091505b5090925090508180156107ab575060008151115b156107c85761070e858280602001905181019061009a91906109ef565b506000949350505050565b6040805160208101909152600081526040518060200160405280670de0b6b3a764000061080886600001518660000151610871565b6108129190610cef565b90529392505050565b60008061082884846108b3565b9050610833816108db565b949350505050565b600061011e8383604051806040016040528060118152602001706164646974696f6e206f766572666c6f7760781b8152506108f3565b600061011e83836040518060400160405280601781526020017f6d756c7469706c69636174696f6e206f766572666c6f77000000000000000000815250610936565b6040805160208101909152600081526040518060200160405280610812856000015185610871565b805160009061012190670de0b6b3a764000090610cef565b6000806109008486610cd7565b9050828582101561092d5760405162461bcd60e51b81526004016109249190610bf9565b60405180910390fd5b50949350505050565b6000831580610943575082155b1561095057506000610639565b600061095c8486610d0f565b9050836109698683610cef565b14839061092d5760405162461bcd60e51b81526004016109249190610bf9565b6040518060e0016040528060006001600160a01b031681526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b805180151581146109df57600080fd5b919050565b80516109df81610da5565b600060208284031215610a00578081fd5b815161063981610da5565b60006020808385031215610a1d578182fd5b825167ffffffffffffffff80821115610a34578384fd5b818501915085601f830112610a47578384fd5b815181811115610a5957610a59610d8f565b8060051b604051601f19603f83011681018181108582111715610a7e57610a7e610d8f565b604052828152858101935084860182860187018a1015610a9c578788fd5b8795505b83861015610ac557610ab1816109e4565b855260019590950194938601938601610aa0565b5098975050505050505050565b600080600060608486031215610ae6578182fd5b610aef846109cf565b925060208401519150610b04604085016109cf565b90509250925092565b60008060408385031215610b1f578182fd5b8235610b2a81610da5565b91506020830135610b3a81610da5565b809150509250929050565b600080600060608486031215610b59578283fd5b8335610b6481610da5565b92506020840135610b7481610da5565b9150604084013560ff81168114610b89578182fd5b809150509250925092565b600060208284031215610ba5578081fd5b5051919050565b6001600160e01b0319831681528151600090610bcf816004850160208701610d2e565b919091016004019392505050565b60008251610bef818460208701610d2e565b9190910192915050565b6020815260008251806020840152610c18816040850160208701610d2e565b601f01601f19169190910160400192915050565b60006020808352608080840185516060808588015282825180855260a0945084890191508684019350875b81811015610cb457845180516001600160a01b0316845288810151898501526040808201519085015284810151858501528781015188850152868101518785015260c090810151908401529387019360e090920191600101610c57565b505085890151604089015260408901518289015280965050505050505092915050565b60008219821115610cea57610cea610d79565b500190565b600082610d0a57634e487b7160e01b81526012600452602481fd5b500490565b6000816000190483118215151615610d2957610d29610d79565b500290565b60005b83811015610d49578181015183820152602001610d31565b83811115610d58576000848401525b50505050565b6000600019821415610d7257610d72610d79565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b0381168114610dba57600080fd5b5056fea2646970667358221220eb1c6c9d026334e0fc41c5dea96f26df10f5c409beaff0f97fe36ca4fb8c9fcc64736f6c63430008040033"

// DeployAbiliquiditylens deploys a new Ethereum contract, binding an instance of Abiliquiditylens to it.
func DeployAbiliquiditylens(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Abiliquiditylens, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiliquiditylensABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbiliquiditylensBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Abiliquiditylens{AbiliquiditylensCaller: AbiliquiditylensCaller{contract: contract}, AbiliquiditylensTransactor: AbiliquiditylensTransactor{contract: contract}, AbiliquiditylensFilterer: AbiliquiditylensFilterer{contract: contract}}, nil
}

// Abiliquiditylens is an auto generated Go binding around an Ethereum contract.
type Abiliquiditylens struct {
	AbiliquiditylensCaller     // Read-only binding to the contract
	AbiliquiditylensTransactor // Write-only binding to the contract
	AbiliquiditylensFilterer   // Log filterer for contract events
}

// AbiliquiditylensCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbiliquiditylensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiliquiditylensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbiliquiditylensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiliquiditylensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbiliquiditylensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbiliquiditylensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbiliquiditylensSession struct {
	Contract     *Abiliquiditylens // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbiliquiditylensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbiliquiditylensCallerSession struct {
	Contract *AbiliquiditylensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// AbiliquiditylensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbiliquiditylensTransactorSession struct {
	Contract     *AbiliquiditylensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// AbiliquiditylensRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbiliquiditylensRaw struct {
	Contract *Abiliquiditylens // Generic contract binding to access the raw methods on
}

// AbiliquiditylensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbiliquiditylensCallerRaw struct {
	Contract *AbiliquiditylensCaller // Generic read-only contract binding to access the raw methods on
}

// AbiliquiditylensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbiliquiditylensTransactorRaw struct {
	Contract *AbiliquiditylensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbiliquiditylens creates a new instance of Abiliquiditylens, bound to a specific deployed contract.
func NewAbiliquiditylens(address common.Address, backend bind.ContractBackend) (*Abiliquiditylens, error) {
	contract, err := bindAbiliquiditylens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abiliquiditylens{AbiliquiditylensCaller: AbiliquiditylensCaller{contract: contract}, AbiliquiditylensTransactor: AbiliquiditylensTransactor{contract: contract}, AbiliquiditylensFilterer: AbiliquiditylensFilterer{contract: contract}}, nil
}

// NewAbiliquiditylensCaller creates a new read-only instance of Abiliquiditylens, bound to a specific deployed contract.
func NewAbiliquiditylensCaller(address common.Address, caller bind.ContractCaller) (*AbiliquiditylensCaller, error) {
	contract, err := bindAbiliquiditylens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbiliquiditylensCaller{contract: contract}, nil
}

// NewAbiliquiditylensTransactor creates a new write-only instance of Abiliquiditylens, bound to a specific deployed contract.
func NewAbiliquiditylensTransactor(address common.Address, transactor bind.ContractTransactor) (*AbiliquiditylensTransactor, error) {
	contract, err := bindAbiliquiditylens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbiliquiditylensTransactor{contract: contract}, nil
}

// NewAbiliquiditylensFilterer creates a new log filterer instance of Abiliquiditylens, bound to a specific deployed contract.
func NewAbiliquiditylensFilterer(address common.Address, filterer bind.ContractFilterer) (*AbiliquiditylensFilterer, error) {
	contract, err := bindAbiliquiditylens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbiliquiditylensFilterer{contract: contract}, nil
}

// bindAbiliquiditylens binds a generic wrapper to an already deployed contract.
func bindAbiliquiditylens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbiliquiditylensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abiliquiditylens *AbiliquiditylensRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abiliquiditylens.Contract.AbiliquiditylensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abiliquiditylens *AbiliquiditylensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abiliquiditylens.Contract.AbiliquiditylensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abiliquiditylens *AbiliquiditylensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abiliquiditylens.Contract.AbiliquiditylensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abiliquiditylens *AbiliquiditylensCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abiliquiditylens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abiliquiditylens *AbiliquiditylensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abiliquiditylens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abiliquiditylens *AbiliquiditylensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abiliquiditylens.Contract.contract.Transact(opts, method, params...)
}

// CTokenBalances is a paid mutator transaction binding the contract method 0x9cb37627.
//
// Solidity: function cTokenBalances(address comp, address account, uint8 featureflags) returns(((address,uint256,uint256,uint256,uint256,uint256,uint256)[],uint256,uint256))
func (_Abiliquiditylens *AbiliquiditylensTransactor) CTokenBalances(opts *bind.TransactOpts, comp common.Address, account common.Address, featureflags uint8) (*types.Transaction, error) {
	return _Abiliquiditylens.contract.Transact(opts, "cTokenBalances", comp, account, featureflags)
}

// CTokenBalances is a paid mutator transaction binding the contract method 0x9cb37627.
//
// Solidity: function cTokenBalances(address comp, address account, uint8 featureflags) returns(((address,uint256,uint256,uint256,uint256,uint256,uint256)[],uint256,uint256))
func (_Abiliquiditylens *AbiliquiditylensSession) CTokenBalances(comp common.Address, account common.Address, featureflags uint8) (*types.Transaction, error) {
	return _Abiliquiditylens.Contract.CTokenBalances(&_Abiliquiditylens.TransactOpts, comp, account, featureflags)
}

// CTokenBalances is a paid mutator transaction binding the contract method 0x9cb37627.
//
// Solidity: function cTokenBalances(address comp, address account, uint8 featureflags) returns(((address,uint256,uint256,uint256,uint256,uint256,uint256)[],uint256,uint256))
func (_Abiliquiditylens *AbiliquiditylensTransactorSession) CTokenBalances(comp common.Address, account common.Address, featureflags uint8) (*types.Transaction, error) {
	return _Abiliquiditylens.Contract.CTokenBalances(&_Abiliquiditylens.TransactOpts, comp, account, featureflags)
}

// GetBorrowBalance is a paid mutator transaction binding the contract method 0x118e31b7.
//
// Solidity: function getBorrowBalance(address asset, address account) returns(uint256)
func (_Abiliquiditylens *AbiliquiditylensTransactor) GetBorrowBalance(opts *bind.TransactOpts, asset common.Address, account common.Address) (*types.Transaction, error) {
	return _Abiliquiditylens.contract.Transact(opts, "getBorrowBalance", asset, account)
}

// GetBorrowBalance is a paid mutator transaction binding the contract method 0x118e31b7.
//
// Solidity: function getBorrowBalance(address asset, address account) returns(uint256)
func (_Abiliquiditylens *AbiliquiditylensSession) GetBorrowBalance(asset common.Address, account common.Address) (*types.Transaction, error) {
	return _Abiliquiditylens.Contract.GetBorrowBalance(&_Abiliquiditylens.TransactOpts, asset, account)
}

// GetBorrowBalance is a paid mutator transaction binding the contract method 0x118e31b7.
//
// Solidity: function getBorrowBalance(address asset, address account) returns(uint256)
func (_Abiliquiditylens *AbiliquiditylensTransactorSession) GetBorrowBalance(asset common.Address, account common.Address) (*types.Transaction, error) {
	return _Abiliquiditylens.Contract.GetBorrowBalance(&_Abiliquiditylens.TransactOpts, asset, account)
}

// GetPrice is a paid mutator transaction binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address oracle, address addr) returns(uint256)
func (_Abiliquiditylens *AbiliquiditylensTransactor) GetPrice(opts *bind.TransactOpts, oracle common.Address, addr common.Address) (*types.Transaction, error) {
	return _Abiliquiditylens.contract.Transact(opts, "getPrice", oracle, addr)
}

// GetPrice is a paid mutator transaction binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address oracle, address addr) returns(uint256)
func (_Abiliquiditylens *AbiliquiditylensSession) GetPrice(oracle common.Address, addr common.Address) (*types.Transaction, error) {
	return _Abiliquiditylens.Contract.GetPrice(&_Abiliquiditylens.TransactOpts, oracle, addr)
}

// GetPrice is a paid mutator transaction binding the contract method 0xac41865a.
//
// Solidity: function getPrice(address oracle, address addr) returns(uint256)
func (_Abiliquiditylens *AbiliquiditylensTransactorSession) GetPrice(oracle common.Address, addr common.Address) (*types.Transaction, error) {
	return _Abiliquiditylens.Contract.GetPrice(&_Abiliquiditylens.TransactOpts, oracle, addr)
}
