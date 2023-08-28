// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abicerc20delegator

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

// Abicerc20delegatorABI is the input ABI used to generate the binding from.
const Abicerc20delegatorABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying_\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowResign\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"name\":\"_setImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToViewImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isCToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractCTokenInterface\",\"name\":\"cTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Abicerc20delegatorBin is the compiled bytecode used for deploying new contracts.
var Abicerc20delegatorBin = "0x60806040523480156200001157600080fd5b506040516200251b3803806200251b83398181016040526101408110156200003857600080fd5b81516020830151604080850151606086015160808701805193519597949692959194919392820192846401000000008211156200007457600080fd5b9083019060208201858111156200008a57600080fd5b8251640100000000811182820188101715620000a557600080fd5b82525081516020918201929091019080838360005b83811015620000d4578181015183820152602001620000ba565b50505050905090810190601f168015620001025780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200012657600080fd5b9083019060208201858111156200013c57600080fd5b82516401000000008111828201881017156200015757600080fd5b82525081516020918201929091019080838360005b83811015620001865781810151838201526020016200016c565b50505050905090810190601f168015620001b45780820380516001836020036101000a031916815260200191505b50604081815260208301519083015160608401516080909401805192969195919284640100000000821115620001e957600080fd5b908301906020820185811115620001ff57600080fd5b82516401000000008111828201881017156200021a57600080fd5b82525081516020918201929091019080838360005b83811015620002495781810151838201526020016200022f565b50505050905090810190601f168015620002775780820380516001836020036101000a031916815260200191505b5060405250505033600360016101000a8154816001600160a01b0302191690836001600160a01b0316021790555062000424828b8b8b8b8b8b8b60405160240180886001600160a01b03166001600160a01b03168152602001876001600160a01b03166001600160a01b03168152602001866001600160a01b03166001600160a01b0316815260200185815260200180602001806020018460ff1660ff168152602001838103835286818151815260200191508051906020019080838360005b838110156200035157818101518382015260200162000337565b50505050905090810190601f1680156200037f5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015620003b45781810151838201526020016200039a565b50505050905090810190601f168015620003e25780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529190526020810180516001600160e01b03908116631a31d46560e01b17909152909a50620004721698505050505050505050565b506200043c826000836001600160e01b036200053916565b5050600380546001600160a01b0390921661010002610100600160a81b0319909216919091179055506200071a95505050505050565b606060006060846001600160a01b0316846040518082805190602001908083835b60208310620004b45780518252601f19909201916020918201910162000493565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d806000811462000516576040519150601f19603f3d011682016040523d82523d6000602084013e6200051b565b606091505b5091509150600082141562000531573d60208201fd5b949350505050565b60035461010090046001600160a01b03163314620005895760405162461bcd60e51b8152600401808060200182810382526039815260200180620024e26039913960400191505060405180910390fd5b8115620005cb576040805160048152602481019091526020810180516001600160e01b0390811663153ab50560e01b17909152620005c99190620006f016565b505b601280546001600160a01b038581166001600160a01b03198316179092556040516020602482018181528551604484015285519490931693620006a1938693909283926064909201919085019080838360005b83811015620006385781810151838201526020016200061e565b50505050905090810190601f168015620006665780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529190526020810180516001600160e01b03908116630adccee560e31b17909152909350620006f016915050565b50601254604080516001600160a01b038085168252909216602083015280517fd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a9281900390910190a150505050565b60125460609062000714906001600160a01b0316836001600160e01b036200047216565b92915050565b611db8806200072a6000396000f3fe6080604052600436106102c95760003560e01c806373acee9811610175578063bd6d894d116100dc578063f2b3abbd11610095578063f851a4401161006f578063f851a44014610cce578063f8f9da2814610ce3578063fca7820b14610cf8578063fe9c44ae14610d22576102c9565b8063f2b3abbd14610c43578063f3fdb15a14610c76578063f5e3c46214610c8b576102c9565b8063bd6d894d14610b31578063c37f68e214610b46578063c5ebeaec14610b9f578063db006a7514610bc9578063dd62ed3e14610bf3578063e9c714f214610c2e576102c9565b8063a6afed951161012e578063a6afed9514610a43578063a9059cbb14610a58578063aa5af0fd14610a91578063ae9d70b014610aa6578063b2a02ff114610abb578063b71d1a0c14610afe576102c9565b806373acee981461097d578063852a12e3146109925780638f840ddd146109bc57806395d89b41146109d157806395dd9193146109e6578063a0712d6814610a19576102c9565b80633af9e66911610234578063555bcc40116101ed578063601a0bf1116101c7578063601a0bf1146108f65780636c540baf146109205780636f307dc31461093557806370a082311461094a576102c9565b8063555bcc40146108025780635c60da1b146108cc5780635fe3b567146108e1576102c9565b80633af9e669146106975780633b1d21a2146106ca5780633e941010146106df5780634487152f146107095780634576b5db146107ba57806347bd3718146107ed576102c9565b806318160ddd1161028657806318160ddd14610595578063182df0f5146105aa57806323b872dd146105bf5780632608f81814610602578063267822471461063b578063313ce5671461066c576102c9565b806306fdde03146103895780630933c1ed14610413578063095ea7b3146104c45780630e75270214610511578063173b99041461054d57806317bfdfbc14610562575b34156103065760405162461bcd60e51b8152600401808060200182810382526037815260200180611d146037913960400191505060405180910390fd5b6012546040516000916001600160a01b031690829036908083838082843760405192019450600093509091505080830381855af49150503d8060008114610369576040519150601f19603f3d011682016040523d82523d6000602084013e61036e565b606091505b505090506040513d6000823e818015610385573d82f35b3d82fd5b34801561039557600080fd5b5061039e610d37565b6040805160208082528351818301528351919283929083019185019080838360005b838110156103d85781810151838201526020016103c0565b50505050905090810190601f1680156104055780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561041f57600080fd5b5061039e6004803603602081101561043657600080fd5b810190602081018135600160201b81111561045057600080fd5b82018360208201111561046257600080fd5b803590602001918460018302840111600160201b8311171561048357600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610dc4945050505050565b3480156104d057600080fd5b506104fd600480360360408110156104e757600080fd5b506001600160a01b038135169060200135610de3565b604080519115158252519081900360200190f35b34801561051d57600080fd5b5061053b6004803603602081101561053457600080fd5b5035610e5a565b60408051918252519081900360200190f35b34801561055957600080fd5b5061053b610ec1565b34801561056e57600080fd5b5061053b6004803603602081101561058557600080fd5b50356001600160a01b0316610ec7565b3480156105a157600080fd5b5061053b610f19565b3480156105b657600080fd5b5061053b610f1f565b3480156105cb57600080fd5b506104fd600480360360608110156105e257600080fd5b506001600160a01b03813581169160208101359091169060400135610f76565b34801561060e57600080fd5b5061053b6004803603604081101561062557600080fd5b506001600160a01b038135169060200135610ff6565b34801561064757600080fd5b5061065061104c565b604080516001600160a01b039092168252519081900360200190f35b34801561067857600080fd5b5061068161105b565b6040805160ff9092168252519081900360200190f35b3480156106a357600080fd5b5061053b600480360360208110156106ba57600080fd5b50356001600160a01b0316611064565b3480156106d657600080fd5b5061053b6110b6565b3480156106eb57600080fd5b5061053b6004803603602081101561070257600080fd5b50356110ee565b34801561071557600080fd5b5061039e6004803603602081101561072c57600080fd5b810190602081018135600160201b81111561074657600080fd5b82018360208201111561075857600080fd5b803590602001918460018302840111600160201b8311171561077957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611135945050505050565b3480156107c657600080fd5b5061053b600480360360208110156107dd57600080fd5b50356001600160a01b0316611354565b3480156107f957600080fd5b5061053b6113a6565b34801561080e57600080fd5b506108ca6004803603606081101561082557600080fd5b6001600160a01b03823516916020810135151591810190606081016040820135600160201b81111561085657600080fd5b82018360208201111561086857600080fd5b803590602001918460018302840111600160201b8311171561088957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506113ac945050505050565b005b3480156108d857600080fd5b5061065061154f565b3480156108ed57600080fd5b5061065061155e565b34801561090257600080fd5b5061053b6004803603602081101561091957600080fd5b503561156d565b34801561092c57600080fd5b5061053b6115b4565b34801561094157600080fd5b506106506115ba565b34801561095657600080fd5b5061053b6004803603602081101561096d57600080fd5b50356001600160a01b03166115c9565b34801561098957600080fd5b5061053b61161b565b34801561099e57600080fd5b5061053b600480360360208110156109b557600080fd5b5035611653565b3480156109c857600080fd5b5061053b61169a565b3480156109dd57600080fd5b5061039e6116a0565b3480156109f257600080fd5b5061053b60048036036020811015610a0957600080fd5b50356001600160a01b03166116f8565b348015610a2557600080fd5b5061053b60048036036020811015610a3c57600080fd5b503561174a565b348015610a4f57600080fd5b5061053b611791565b348015610a6457600080fd5b506104fd60048036036040811015610a7b57600080fd5b506001600160a01b0381351690602001356117c9565b348015610a9d57600080fd5b5061053b61181f565b348015610ab257600080fd5b5061053b611825565b348015610ac757600080fd5b5061053b60048036036060811015610ade57600080fd5b506001600160a01b0381358116916020810135909116906040013561185d565b348015610b0a57600080fd5b5061053b60048036036020811015610b2157600080fd5b50356001600160a01b03166118bb565b348015610b3d57600080fd5b5061053b61190d565b348015610b5257600080fd5b50610b7960048036036020811015610b6957600080fd5b50356001600160a01b0316611945565b604080519485526020850193909352838301919091526060830152519081900360800190f35b348015610bab57600080fd5b5061053b60048036036020811015610bc257600080fd5b50356119d7565b348015610bd557600080fd5b5061053b60048036036020811015610bec57600080fd5b5035611a1e565b348015610bff57600080fd5b5061053b60048036036040811015610c1657600080fd5b506001600160a01b0381358116916020013516611a65565b348015610c3a57600080fd5b5061053b611abf565b348015610c4f57600080fd5b5061053b60048036036020811015610c6657600080fd5b50356001600160a01b0316611af7565b348015610c8257600080fd5b50610650611b49565b348015610c9757600080fd5b5061053b60048036036060811015610cae57600080fd5b506001600160a01b03813581169160208101359160409091013516611b58565b348015610cda57600080fd5b50610650611bb9565b348015610cef57600080fd5b5061053b611bcd565b348015610d0457600080fd5b5061053b60048036036020811015610d1b57600080fd5b5035611c05565b348015610d2e57600080fd5b506104fd611c4c565b60018054604080516020600284861615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610dbc5780601f10610d9157610100808354040283529160200191610dbc565b820191906000526020600020905b815481529060010190602001808311610d9f57829003601f168201915b505050505081565b601254606090610ddd906001600160a01b031683611c51565b92915050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663095ea7b360e01b179052600090606090610e3990610dc4565b9050808060200190516020811015610e5057600080fd5b5051949350505050565b6040805160248082018490528251808303909101815260449091019091526020810180516001600160e01b031663073a938160e11b179052600090606090610ea190610dc4565b9050808060200190516020811015610eb857600080fd5b50519392505050565b60085481565b604080516001600160a01b0383166024808301919091528251808303909101815260449091019091526020810180516001600160e01b03166305eff7ef60e21b179052600090606090610ea190610dc4565b600d5481565b6040805160048152602481019091526020810180516001600160e01b031663182df0f560e01b179052600090606090610f5790611135565b9050808060200190516020811015610f6e57600080fd5b505191505090565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b03166323b872dd60e01b179052600090606090610fd490610dc4565b9050808060200190516020811015610feb57600080fd5b505195945050505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b03166304c11f0360e31b179052600090606090610e3990610dc4565b6004546001600160a01b031681565b60035460ff1681565b604080516001600160a01b0383166024808301919091528251808303909101815260449091019091526020810180516001600160e01b0316633af9e66960e01b179052600090606090610ea190610dc4565b6040805160048152602481019091526020810180516001600160e01b0316631d8e90d160e11b179052600090606090610f5790611135565b6040805160248082018490528251808303909101815260449091019091526020810180516001600160e01b03166303e9410160e41b179052600090606090610ea190610dc4565b606060006060306001600160a01b0316846040516024018080602001828103825283818151815260200191508051906020019080838360005b8381101561118657818101518382015260200161116e565b50505050905090810190601f1680156111b35780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529181526020820180516001600160e01b0316630933c1ed60e01b178152905182519295509350839250908083835b6020831061120e5780518252601f1990920191602091820191016111ef565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855afa9150503d806000811461126e576040519150601f19603f3d011682016040523d82523d6000602084013e611273565b606091505b50915091506000821415611288573d60208201fd5b80806020019051602081101561129d57600080fd5b8101908080516040519392919084600160201b8211156112bc57600080fd5b9083019060208201858111156112d157600080fd5b8251600160201b8111828201881017156112ea57600080fd5b82525081516020918201929091019080838360005b838110156113175781810151838201526020016112ff565b50505050905090810190601f1680156113445780820380516001836020036101000a031916815260200191505b5060405250505092505050919050565b604080516001600160a01b0383166024808301919091528251808303909101815260449091019091526020810180516001600160e01b0316634576b5db60e01b179052600090606090610ea190610dc4565b600b5481565b60035461010090046001600160a01b031633146113fa5760405162461bcd60e51b8152600401808060200182810382526039815260200180611d4b6039913960400191505060405180910390fd5b8115611434576040805160048152602481019091526020810180516001600160e01b031663153ab50560e01b17905261143290610dc4565b505b601280546001600160a01b038581166001600160a01b03198316179092556040516020602482018181528551604484015285519490931693611500938693909283926064909201919085019080838360005b8381101561149e578181015183820152602001611486565b50505050905090810190601f1680156114cb5780820380516001836020036101000a031916815260200191505b5060408051601f198184030181529190526020810180516001600160e01b0316630adccee560e31b1790529250610dc4915050565b50601254604080516001600160a01b038085168252909216602083015280517fd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a9281900390910190a150505050565b6012546001600160a01b031681565b6005546001600160a01b031681565b6040805160248082018490528251808303909101815260449091019091526020810180516001600160e01b031663601a0bf160e01b179052600090606090610ea190610dc4565b60095481565b6011546001600160a01b031681565b604080516001600160a01b0383166024808301919091528251808303909101815260449091019091526020810180516001600160e01b03166370a0823160e01b179052600090606090610ea190611135565b6040805160048152602481019091526020810180516001600160e01b0316630e759dd360e31b179052600090606090610f5790610dc4565b6040805160248082018490528251808303909101815260449091019091526020810180516001600160e01b031663852a12e360e01b179052600090606090610ea190610dc4565b600c5481565b6002805460408051602060018416156101000260001901909316849004601f81018490048402820184019092528181529291830182828015610dbc5780601f10610d9157610100808354040283529160200191610dbc565b604080516001600160a01b0383166024808301919091528251808303909101815260449091019091526020810180516001600160e01b03166395dd919360e01b179052600090606090610ea190611135565b6040805160248082018490528251808303909101815260449091019091526020810180516001600160e01b031663140e25ad60e31b179052600090606090610ea190610dc4565b6040805160048152602481019091526020810180516001600160e01b031663a6afed9560e01b179052600090606090610f5790610dc4565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052600090606090610e3990610dc4565b600a5481565b6040805160048152602481019091526020810180516001600160e01b0316630ae9d70b60e41b179052600090606090610f5790611135565b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180516001600160e01b031663b2a02ff160e01b179052600090606090610fd490610dc4565b604080516001600160a01b0383166024808301919091528251808303909101815260449091019091526020810180516001600160e01b0316632dc7468360e21b179052600090606090610ea190610dc4565b6040805160048152602481019091526020810180516001600160e01b031663bd6d894d60e01b179052600090606090610f5790610dc4565b604080516001600160a01b0383166024808301919091528251808303909101815260449091019091526020810180516001600160e01b03166361bfb47160e11b17905260009081908190819060609061199d90611135565b90508080602001905160808110156119b457600080fd5b508051602082015160408301516060909301519199909850919650945092505050565b6040805160248082018490528251808303909101815260449091019091526020810180516001600160e01b031663317afabb60e21b179052600090606090610ea190610dc4565b6040805160248082018490528251808303909101815260449091019091526020810180516001600160e01b031663db006a7560e01b179052600090606090610ea190610dc4565b604080516001600160a01b03808516602483015283166044808301919091528251808303909101815260649091019091526020810180516001600160e01b0316636eb1769f60e11b179052600090606090610e3990611135565b6040805160048152602481019091526020810180516001600160e01b03166374e38a7960e11b179052600090606090610f5790610dc4565b604080516001600160a01b0383166024808301919091528251808303909101815260449091019091526020810180516001600160e01b031663f2b3abbd60e01b179052600090606090610ea190610dc4565b6006546001600160a01b031681565b604080516001600160a01b0380861660248301526044820185905283166064808301919091528251808303909101815260849091019091526020810180516001600160e01b0316637af1e23160e11b179052600090606090610fd490610dc4565b60035461010090046001600160a01b031681565b6040805160048152602481019091526020810180516001600160e01b0316631f1f3b4560e31b179052600090606090610f5790611135565b6040805160248082018490528251808303909101815260449091019091526020810180516001600160e01b031663fca7820b60e01b179052600090606090610ea190610dc4565b600181565b606060006060846001600160a01b0316846040518082805190602001908083835b60208310611c915780518252601f199092019160209182019101611c72565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d8060008114611cf1576040519150601f19603f3d011682016040523d82523d6000602084013e611cf6565b606091505b50915091506000821415611d0b573d60208201fd5b94935050505056fe43457263323044656c656761746f723a66616c6c6261636b3a2063616e6e6f742073656e642076616c756520746f2066616c6c6261636b43457263323044656c656761746f723a3a5f736574496d706c656d656e746174696f6e3a2043616c6c6572206d7573742062652061646d696ea265627a7a723158207f4dbd8ddf9d69162bd76669835a3603119d33eb6dc600af90b93f3e022d234f64736f6c6343000510003243457263323044656c656761746f723a3a5f736574496d706c656d656e746174696f6e3a2043616c6c6572206d7573742062652061646d696e"

// DeployAbicerc20delegator deploys a new Ethereum contract, binding an instance of Abicerc20delegator to it.
func DeployAbicerc20delegator(auth *bind.TransactOpts, backend bind.ContractBackend, underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8, admin_ common.Address, implementation_ common.Address, becomeImplementationData []byte) (common.Address, *types.Transaction, *Abicerc20delegator, error) {
	parsed, err := abi.JSON(strings.NewReader(Abicerc20delegatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Abicerc20delegatorBin), backend, underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_, admin_, implementation_, becomeImplementationData)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Abicerc20delegator{Abicerc20delegatorCaller: Abicerc20delegatorCaller{contract: contract}, Abicerc20delegatorTransactor: Abicerc20delegatorTransactor{contract: contract}, Abicerc20delegatorFilterer: Abicerc20delegatorFilterer{contract: contract}}, nil
}

// Abicerc20delegator is an auto generated Go binding around an Ethereum contract.
type Abicerc20delegator struct {
	Abicerc20delegatorCaller     // Read-only binding to the contract
	Abicerc20delegatorTransactor // Write-only binding to the contract
	Abicerc20delegatorFilterer   // Log filterer for contract events
}

// Abicerc20delegatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type Abicerc20delegatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abicerc20delegatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Abicerc20delegatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abicerc20delegatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Abicerc20delegatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abicerc20delegatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Abicerc20delegatorSession struct {
	Contract     *Abicerc20delegator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// Abicerc20delegatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Abicerc20delegatorCallerSession struct {
	Contract *Abicerc20delegatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// Abicerc20delegatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Abicerc20delegatorTransactorSession struct {
	Contract     *Abicerc20delegatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// Abicerc20delegatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type Abicerc20delegatorRaw struct {
	Contract *Abicerc20delegator // Generic contract binding to access the raw methods on
}

// Abicerc20delegatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Abicerc20delegatorCallerRaw struct {
	Contract *Abicerc20delegatorCaller // Generic read-only contract binding to access the raw methods on
}

// Abicerc20delegatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Abicerc20delegatorTransactorRaw struct {
	Contract *Abicerc20delegatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbicerc20delegator creates a new instance of Abicerc20delegator, bound to a specific deployed contract.
func NewAbicerc20delegator(address common.Address, backend bind.ContractBackend) (*Abicerc20delegator, error) {
	contract, err := bindAbicerc20delegator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegator{Abicerc20delegatorCaller: Abicerc20delegatorCaller{contract: contract}, Abicerc20delegatorTransactor: Abicerc20delegatorTransactor{contract: contract}, Abicerc20delegatorFilterer: Abicerc20delegatorFilterer{contract: contract}}, nil
}

// NewAbicerc20delegatorCaller creates a new read-only instance of Abicerc20delegator, bound to a specific deployed contract.
func NewAbicerc20delegatorCaller(address common.Address, caller bind.ContractCaller) (*Abicerc20delegatorCaller, error) {
	contract, err := bindAbicerc20delegator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorCaller{contract: contract}, nil
}

// NewAbicerc20delegatorTransactor creates a new write-only instance of Abicerc20delegator, bound to a specific deployed contract.
func NewAbicerc20delegatorTransactor(address common.Address, transactor bind.ContractTransactor) (*Abicerc20delegatorTransactor, error) {
	contract, err := bindAbicerc20delegator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorTransactor{contract: contract}, nil
}

// NewAbicerc20delegatorFilterer creates a new log filterer instance of Abicerc20delegator, bound to a specific deployed contract.
func NewAbicerc20delegatorFilterer(address common.Address, filterer bind.ContractFilterer) (*Abicerc20delegatorFilterer, error) {
	contract, err := bindAbicerc20delegator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorFilterer{contract: contract}, nil
}

// bindAbicerc20delegator binds a generic wrapper to an already deployed contract.
func bindAbicerc20delegator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Abicerc20delegatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abicerc20delegator *Abicerc20delegatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abicerc20delegator.Contract.Abicerc20delegatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abicerc20delegator *Abicerc20delegatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Abicerc20delegatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abicerc20delegator *Abicerc20delegatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Abicerc20delegatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abicerc20delegator *Abicerc20delegatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abicerc20delegator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abicerc20delegator *Abicerc20delegatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abicerc20delegator *Abicerc20delegatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) AccrualBlockNumber() (*big.Int, error) {
	return _Abicerc20delegator.Contract.AccrualBlockNumber(&_Abicerc20delegator.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _Abicerc20delegator.Contract.AccrualBlockNumber(&_Abicerc20delegator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorSession) Admin() (common.Address, error) {
	return _Abicerc20delegator.Contract.Admin(&_Abicerc20delegator.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) Admin() (common.Address, error) {
	return _Abicerc20delegator.Contract.Admin(&_Abicerc20delegator.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Abicerc20delegator.Contract.Allowance(&_Abicerc20delegator.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Abicerc20delegator.Contract.Allowance(&_Abicerc20delegator.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Abicerc20delegator.Contract.BalanceOf(&_Abicerc20delegator.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Abicerc20delegator.Contract.BalanceOf(&_Abicerc20delegator.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Abicerc20delegator.Contract.BorrowBalanceStored(&_Abicerc20delegator.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Abicerc20delegator.Contract.BorrowBalanceStored(&_Abicerc20delegator.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) BorrowIndex() (*big.Int, error) {
	return _Abicerc20delegator.Contract.BorrowIndex(&_Abicerc20delegator.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) BorrowIndex() (*big.Int, error) {
	return _Abicerc20delegator.Contract.BorrowIndex(&_Abicerc20delegator.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Abicerc20delegator.Contract.BorrowRatePerBlock(&_Abicerc20delegator.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Abicerc20delegator.Contract.BorrowRatePerBlock(&_Abicerc20delegator.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorSession) Comptroller() (common.Address, error) {
	return _Abicerc20delegator.Contract.Comptroller(&_Abicerc20delegator.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) Comptroller() (common.Address, error) {
	return _Abicerc20delegator.Contract.Comptroller(&_Abicerc20delegator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abicerc20delegator *Abicerc20delegatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abicerc20delegator *Abicerc20delegatorSession) Decimals() (uint8, error) {
	return _Abicerc20delegator.Contract.Decimals(&_Abicerc20delegator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) Decimals() (uint8, error) {
	return _Abicerc20delegator.Contract.Decimals(&_Abicerc20delegator.CallOpts)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Abicerc20delegator *Abicerc20delegatorCaller) DelegateToViewImplementation(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "delegateToViewImplementation", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Abicerc20delegator *Abicerc20delegatorSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _Abicerc20delegator.Contract.DelegateToViewImplementation(&_Abicerc20delegator.CallOpts, data)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _Abicerc20delegator.Contract.DelegateToViewImplementation(&_Abicerc20delegator.CallOpts, data)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) ExchangeRateStored() (*big.Int, error) {
	return _Abicerc20delegator.Contract.ExchangeRateStored(&_Abicerc20delegator.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _Abicerc20delegator.Contract.ExchangeRateStored(&_Abicerc20delegator.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Abicerc20delegator.Contract.GetAccountSnapshot(&_Abicerc20delegator.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Abicerc20delegator.Contract.GetAccountSnapshot(&_Abicerc20delegator.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) GetCash() (*big.Int, error) {
	return _Abicerc20delegator.Contract.GetCash(&_Abicerc20delegator.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) GetCash() (*big.Int, error) {
	return _Abicerc20delegator.Contract.GetCash(&_Abicerc20delegator.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorSession) Implementation() (common.Address, error) {
	return _Abicerc20delegator.Contract.Implementation(&_Abicerc20delegator.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) Implementation() (common.Address, error) {
	return _Abicerc20delegator.Contract.Implementation(&_Abicerc20delegator.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorSession) InterestRateModel() (common.Address, error) {
	return _Abicerc20delegator.Contract.InterestRateModel(&_Abicerc20delegator.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) InterestRateModel() (common.Address, error) {
	return _Abicerc20delegator.Contract.InterestRateModel(&_Abicerc20delegator.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorCaller) IsCToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "isCToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorSession) IsCToken() (bool, error) {
	return _Abicerc20delegator.Contract.IsCToken(&_Abicerc20delegator.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) IsCToken() (bool, error) {
	return _Abicerc20delegator.Contract.IsCToken(&_Abicerc20delegator.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abicerc20delegator *Abicerc20delegatorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abicerc20delegator *Abicerc20delegatorSession) Name() (string, error) {
	return _Abicerc20delegator.Contract.Name(&_Abicerc20delegator.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) Name() (string, error) {
	return _Abicerc20delegator.Contract.Name(&_Abicerc20delegator.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorSession) PendingAdmin() (common.Address, error) {
	return _Abicerc20delegator.Contract.PendingAdmin(&_Abicerc20delegator.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) PendingAdmin() (common.Address, error) {
	return _Abicerc20delegator.Contract.PendingAdmin(&_Abicerc20delegator.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Abicerc20delegator.Contract.ReserveFactorMantissa(&_Abicerc20delegator.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Abicerc20delegator.Contract.ReserveFactorMantissa(&_Abicerc20delegator.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Abicerc20delegator.Contract.SupplyRatePerBlock(&_Abicerc20delegator.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Abicerc20delegator.Contract.SupplyRatePerBlock(&_Abicerc20delegator.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abicerc20delegator *Abicerc20delegatorCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abicerc20delegator *Abicerc20delegatorSession) Symbol() (string, error) {
	return _Abicerc20delegator.Contract.Symbol(&_Abicerc20delegator.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) Symbol() (string, error) {
	return _Abicerc20delegator.Contract.Symbol(&_Abicerc20delegator.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) TotalBorrows() (*big.Int, error) {
	return _Abicerc20delegator.Contract.TotalBorrows(&_Abicerc20delegator.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) TotalBorrows() (*big.Int, error) {
	return _Abicerc20delegator.Contract.TotalBorrows(&_Abicerc20delegator.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) TotalReserves() (*big.Int, error) {
	return _Abicerc20delegator.Contract.TotalReserves(&_Abicerc20delegator.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) TotalReserves() (*big.Int, error) {
	return _Abicerc20delegator.Contract.TotalReserves(&_Abicerc20delegator.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) TotalSupply() (*big.Int, error) {
	return _Abicerc20delegator.Contract.TotalSupply(&_Abicerc20delegator.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) TotalSupply() (*big.Int, error) {
	return _Abicerc20delegator.Contract.TotalSupply(&_Abicerc20delegator.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegator.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorSession) Underlying() (common.Address, error) {
	return _Abicerc20delegator.Contract.Underlying(&_Abicerc20delegator.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Abicerc20delegator *Abicerc20delegatorCallerSession) Underlying() (common.Address, error) {
	return _Abicerc20delegator.Contract.Underlying(&_Abicerc20delegator.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.AcceptAdmin(&_Abicerc20delegator.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.AcceptAdmin(&_Abicerc20delegator.TransactOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) AddReserves(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "_addReserves", addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.AddReserves(&_Abicerc20delegator.TransactOpts, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.AddReserves(&_Abicerc20delegator.TransactOpts, addAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.ReduceReserves(&_Abicerc20delegator.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.ReduceReserves(&_Abicerc20delegator.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.SetComptroller(&_Abicerc20delegator.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.SetComptroller(&_Abicerc20delegator.TransactOpts, newComptroller)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Abicerc20delegator *Abicerc20delegatorTransactor) SetImplementation(opts *bind.TransactOpts, implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "_setImplementation", implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Abicerc20delegator *Abicerc20delegatorSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.SetImplementation(&_Abicerc20delegator.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.SetImplementation(&_Abicerc20delegator.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.SetInterestRateModel(&_Abicerc20delegator.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.SetInterestRateModel(&_Abicerc20delegator.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.SetPendingAdmin(&_Abicerc20delegator.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.SetPendingAdmin(&_Abicerc20delegator.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.SetReserveFactor(&_Abicerc20delegator.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.SetReserveFactor(&_Abicerc20delegator.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) AccrueInterest() (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.AccrueInterest(&_Abicerc20delegator.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.AccrueInterest(&_Abicerc20delegator.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Approve(&_Abicerc20delegator.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Approve(&_Abicerc20delegator.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.BalanceOfUnderlying(&_Abicerc20delegator.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.BalanceOfUnderlying(&_Abicerc20delegator.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Borrow(&_Abicerc20delegator.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Borrow(&_Abicerc20delegator.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.BorrowBalanceCurrent(&_Abicerc20delegator.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.BorrowBalanceCurrent(&_Abicerc20delegator.TransactOpts, account)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) DelegateToImplementation(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "delegateToImplementation", data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Abicerc20delegator *Abicerc20delegatorSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.DelegateToImplementation(&_Abicerc20delegator.TransactOpts, data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.DelegateToImplementation(&_Abicerc20delegator.TransactOpts, data)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.ExchangeRateCurrent(&_Abicerc20delegator.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.ExchangeRateCurrent(&_Abicerc20delegator.TransactOpts)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.LiquidateBorrow(&_Abicerc20delegator.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.LiquidateBorrow(&_Abicerc20delegator.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Mint(&_Abicerc20delegator.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Mint(&_Abicerc20delegator.TransactOpts, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Redeem(&_Abicerc20delegator.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Redeem(&_Abicerc20delegator.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.RedeemUnderlying(&_Abicerc20delegator.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.RedeemUnderlying(&_Abicerc20delegator.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.RepayBorrow(&_Abicerc20delegator.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.RepayBorrow(&_Abicerc20delegator.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.RepayBorrowBehalf(&_Abicerc20delegator.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.RepayBorrowBehalf(&_Abicerc20delegator.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Seize(&_Abicerc20delegator.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Seize(&_Abicerc20delegator.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.TotalBorrowsCurrent(&_Abicerc20delegator.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.TotalBorrowsCurrent(&_Abicerc20delegator.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Transfer(&_Abicerc20delegator.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Transfer(&_Abicerc20delegator.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.TransferFrom(&_Abicerc20delegator.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.TransferFrom(&_Abicerc20delegator.TransactOpts, src, dst, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abicerc20delegator *Abicerc20delegatorTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Abicerc20delegator.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abicerc20delegator *Abicerc20delegatorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Fallback(&_Abicerc20delegator.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Abicerc20delegator *Abicerc20delegatorTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Abicerc20delegator.Contract.Fallback(&_Abicerc20delegator.TransactOpts, calldata)
}

// Abicerc20delegatorAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the Abicerc20delegator contract.
type Abicerc20delegatorAccrueInterestIterator struct {
	Event *Abicerc20delegatorAccrueInterest // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorAccrueInterest)
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
		it.Event = new(Abicerc20delegatorAccrueInterest)
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
func (it *Abicerc20delegatorAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorAccrueInterest represents a AccrueInterest event raised by the Abicerc20delegator contract.
type Abicerc20delegatorAccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*Abicerc20delegatorAccrueInterestIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorAccrueInterestIterator{contract: _Abicerc20delegator.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorAccrueInterest)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseAccrueInterest(log types.Log) (*Abicerc20delegatorAccrueInterest, error) {
	event := new(Abicerc20delegatorAccrueInterest)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Abicerc20delegator contract.
type Abicerc20delegatorApprovalIterator struct {
	Event *Abicerc20delegatorApproval // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorApproval)
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
		it.Event = new(Abicerc20delegatorApproval)
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
func (it *Abicerc20delegatorApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorApproval represents a Approval event raised by the Abicerc20delegator contract.
type Abicerc20delegatorApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Abicerc20delegatorApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorApprovalIterator{contract: _Abicerc20delegator.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorApproval)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseApproval(log types.Log) (*Abicerc20delegatorApproval, error) {
	event := new(Abicerc20delegatorApproval)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Abicerc20delegator contract.
type Abicerc20delegatorBorrowIterator struct {
	Event *Abicerc20delegatorBorrow // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorBorrow)
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
		it.Event = new(Abicerc20delegatorBorrow)
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
func (it *Abicerc20delegatorBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorBorrow represents a Borrow event raised by the Abicerc20delegator contract.
type Abicerc20delegatorBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterBorrow(opts *bind.FilterOpts) (*Abicerc20delegatorBorrowIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorBorrowIterator{contract: _Abicerc20delegator.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorBorrow) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorBorrow)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseBorrow(log types.Log) (*Abicerc20delegatorBorrow, error) {
	event := new(Abicerc20delegatorBorrow)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Abicerc20delegator contract.
type Abicerc20delegatorFailureIterator struct {
	Event *Abicerc20delegatorFailure // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorFailure)
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
		it.Event = new(Abicerc20delegatorFailure)
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
func (it *Abicerc20delegatorFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorFailure represents a Failure event raised by the Abicerc20delegator contract.
type Abicerc20delegatorFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterFailure(opts *bind.FilterOpts) (*Abicerc20delegatorFailureIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorFailureIterator{contract: _Abicerc20delegator.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorFailure) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorFailure)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "Failure", log); err != nil {
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
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseFailure(log types.Log) (*Abicerc20delegatorFailure, error) {
	event := new(Abicerc20delegatorFailure)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the Abicerc20delegator contract.
type Abicerc20delegatorLiquidateBorrowIterator struct {
	Event *Abicerc20delegatorLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorLiquidateBorrow)
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
		it.Event = new(Abicerc20delegatorLiquidateBorrow)
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
func (it *Abicerc20delegatorLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorLiquidateBorrow represents a LiquidateBorrow event raised by the Abicerc20delegator contract.
type Abicerc20delegatorLiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	CTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*Abicerc20delegatorLiquidateBorrowIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorLiquidateBorrowIterator{contract: _Abicerc20delegator.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorLiquidateBorrow)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseLiquidateBorrow(log types.Log) (*Abicerc20delegatorLiquidateBorrow, error) {
	event := new(Abicerc20delegatorLiquidateBorrow)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Abicerc20delegator contract.
type Abicerc20delegatorMintIterator struct {
	Event *Abicerc20delegatorMint // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorMint)
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
		it.Event = new(Abicerc20delegatorMint)
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
func (it *Abicerc20delegatorMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorMint represents a Mint event raised by the Abicerc20delegator contract.
type Abicerc20delegatorMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterMint(opts *bind.FilterOpts) (*Abicerc20delegatorMintIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorMintIterator{contract: _Abicerc20delegator.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorMint) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorMint)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseMint(log types.Log) (*Abicerc20delegatorMint, error) {
	event := new(Abicerc20delegatorMint)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewAdminIterator struct {
	Event *Abicerc20delegatorNewAdmin // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorNewAdmin)
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
		it.Event = new(Abicerc20delegatorNewAdmin)
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
func (it *Abicerc20delegatorNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorNewAdmin represents a NewAdmin event raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*Abicerc20delegatorNewAdminIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorNewAdminIterator{contract: _Abicerc20delegator.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorNewAdmin) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorNewAdmin)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseNewAdmin(log types.Log) (*Abicerc20delegatorNewAdmin, error) {
	event := new(Abicerc20delegatorNewAdmin)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewComptrollerIterator struct {
	Event *Abicerc20delegatorNewComptroller // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorNewComptroller)
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
		it.Event = new(Abicerc20delegatorNewComptroller)
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
func (it *Abicerc20delegatorNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorNewComptroller represents a NewComptroller event raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*Abicerc20delegatorNewComptrollerIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorNewComptrollerIterator{contract: _Abicerc20delegator.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorNewComptroller) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorNewComptroller)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseNewComptroller(log types.Log) (*Abicerc20delegatorNewComptroller, error) {
	event := new(Abicerc20delegatorNewComptroller)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorNewImplementationIterator is returned from FilterNewImplementation and is used to iterate over the raw logs and unpacked data for NewImplementation events raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewImplementationIterator struct {
	Event *Abicerc20delegatorNewImplementation // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorNewImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorNewImplementation)
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
		it.Event = new(Abicerc20delegatorNewImplementation)
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
func (it *Abicerc20delegatorNewImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorNewImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorNewImplementation represents a NewImplementation event raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewImplementation struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewImplementation is a free log retrieval operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterNewImplementation(opts *bind.FilterOpts) (*Abicerc20delegatorNewImplementationIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorNewImplementationIterator{contract: _Abicerc20delegator.contract, event: "NewImplementation", logs: logs, sub: sub}, nil
}

// WatchNewImplementation is a free log subscription operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchNewImplementation(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorNewImplementation) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorNewImplementation)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "NewImplementation", log); err != nil {
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
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseNewImplementation(log types.Log) (*Abicerc20delegatorNewImplementation, error) {
	event := new(Abicerc20delegatorNewImplementation)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "NewImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorNewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewMarketInterestRateModelIterator struct {
	Event *Abicerc20delegatorNewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorNewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorNewMarketInterestRateModel)
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
		it.Event = new(Abicerc20delegatorNewMarketInterestRateModel)
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
func (it *Abicerc20delegatorNewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorNewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorNewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*Abicerc20delegatorNewMarketInterestRateModelIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorNewMarketInterestRateModelIterator{contract: _Abicerc20delegator.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorNewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorNewMarketInterestRateModel)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseNewMarketInterestRateModel(log types.Log) (*Abicerc20delegatorNewMarketInterestRateModel, error) {
	event := new(Abicerc20delegatorNewMarketInterestRateModel)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewPendingAdminIterator struct {
	Event *Abicerc20delegatorNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorNewPendingAdmin)
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
		it.Event = new(Abicerc20delegatorNewPendingAdmin)
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
func (it *Abicerc20delegatorNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorNewPendingAdmin represents a NewPendingAdmin event raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*Abicerc20delegatorNewPendingAdminIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorNewPendingAdminIterator{contract: _Abicerc20delegator.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorNewPendingAdmin)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseNewPendingAdmin(log types.Log) (*Abicerc20delegatorNewPendingAdmin, error) {
	event := new(Abicerc20delegatorNewPendingAdmin)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorNewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewReserveFactorIterator struct {
	Event *Abicerc20delegatorNewReserveFactor // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorNewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorNewReserveFactor)
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
		it.Event = new(Abicerc20delegatorNewReserveFactor)
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
func (it *Abicerc20delegatorNewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorNewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorNewReserveFactor represents a NewReserveFactor event raised by the Abicerc20delegator contract.
type Abicerc20delegatorNewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*Abicerc20delegatorNewReserveFactorIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorNewReserveFactorIterator{contract: _Abicerc20delegator.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorNewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorNewReserveFactor)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseNewReserveFactor(log types.Log) (*Abicerc20delegatorNewReserveFactor, error) {
	event := new(Abicerc20delegatorNewReserveFactor)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Abicerc20delegator contract.
type Abicerc20delegatorRedeemIterator struct {
	Event *Abicerc20delegatorRedeem // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorRedeem)
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
		it.Event = new(Abicerc20delegatorRedeem)
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
func (it *Abicerc20delegatorRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorRedeem represents a Redeem event raised by the Abicerc20delegator contract.
type Abicerc20delegatorRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterRedeem(opts *bind.FilterOpts) (*Abicerc20delegatorRedeemIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorRedeemIterator{contract: _Abicerc20delegator.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorRedeem) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorRedeem)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseRedeem(log types.Log) (*Abicerc20delegatorRedeem, error) {
	event := new(Abicerc20delegatorRedeem)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the Abicerc20delegator contract.
type Abicerc20delegatorRepayBorrowIterator struct {
	Event *Abicerc20delegatorRepayBorrow // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorRepayBorrow)
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
		it.Event = new(Abicerc20delegatorRepayBorrow)
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
func (it *Abicerc20delegatorRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorRepayBorrow represents a RepayBorrow event raised by the Abicerc20delegator contract.
type Abicerc20delegatorRepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*Abicerc20delegatorRepayBorrowIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorRepayBorrowIterator{contract: _Abicerc20delegator.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorRepayBorrow)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseRepayBorrow(log types.Log) (*Abicerc20delegatorRepayBorrow, error) {
	event := new(Abicerc20delegatorRepayBorrow)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the Abicerc20delegator contract.
type Abicerc20delegatorReservesAddedIterator struct {
	Event *Abicerc20delegatorReservesAdded // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorReservesAdded)
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
		it.Event = new(Abicerc20delegatorReservesAdded)
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
func (it *Abicerc20delegatorReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorReservesAdded represents a ReservesAdded event raised by the Abicerc20delegator contract.
type Abicerc20delegatorReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterReservesAdded(opts *bind.FilterOpts) (*Abicerc20delegatorReservesAddedIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorReservesAddedIterator{contract: _Abicerc20delegator.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorReservesAdded) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorReservesAdded)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
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

// ParseReservesAdded is a log parse operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseReservesAdded(log types.Log) (*Abicerc20delegatorReservesAdded, error) {
	event := new(Abicerc20delegatorReservesAdded)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the Abicerc20delegator contract.
type Abicerc20delegatorReservesReducedIterator struct {
	Event *Abicerc20delegatorReservesReduced // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorReservesReduced)
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
		it.Event = new(Abicerc20delegatorReservesReduced)
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
func (it *Abicerc20delegatorReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorReservesReduced represents a ReservesReduced event raised by the Abicerc20delegator contract.
type Abicerc20delegatorReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*Abicerc20delegatorReservesReducedIterator, error) {

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorReservesReducedIterator{contract: _Abicerc20delegator.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorReservesReduced) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorReservesReduced)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseReservesReduced(log types.Log) (*Abicerc20delegatorReservesReduced, error) {
	event := new(Abicerc20delegatorReservesReduced)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegatorTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Abicerc20delegator contract.
type Abicerc20delegatorTransferIterator struct {
	Event *Abicerc20delegatorTransfer // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegatorTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegatorTransfer)
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
		it.Event = new(Abicerc20delegatorTransfer)
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
func (it *Abicerc20delegatorTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegatorTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegatorTransfer represents a Transfer event raised by the Abicerc20delegator contract.
type Abicerc20delegatorTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Abicerc20delegatorTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Abicerc20delegator.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegatorTransferIterator{contract: _Abicerc20delegator.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Abicerc20delegatorTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Abicerc20delegator.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegatorTransfer)
				if err := _Abicerc20delegator.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Abicerc20delegator *Abicerc20delegatorFilterer) ParseTransfer(log types.Log) (*Abicerc20delegatorTransfer, error) {
	event := new(Abicerc20delegatorTransfer)
	if err := _Abicerc20delegator.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
