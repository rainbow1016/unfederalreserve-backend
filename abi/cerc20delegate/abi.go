// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abicerc20delegate

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

// Abicerc20delegateABI is the input ABI used to generate the binding from.
const Abicerc20delegateABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"_becomeImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"_resignImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying_\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isCToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractCTokenInterface\",\"name\":\"cTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Abicerc20delegateBin is the compiled bytecode used for deploying new contracts.
var Abicerc20delegateBin = "0x608060405234801561001057600080fd5b50615144806100206000396000f3fe608060405234801561001057600080fd5b50600436106102f15760003560e01c806373acee981161019d578063bd6d894d116100e9578063f2b3abbd116100a2578063f851a4401161007c578063f851a44014610b32578063f8f9da2814610b3a578063fca7820b14610b42578063fe9c44ae14610b5f576102f1565b8063f2b3abbd14610ace578063f3fdb15a14610af4578063f5e3c46214610afc576102f1565b8063bd6d894d14610a0a578063c37f68e214610a12578063c5ebeaec14610a5e578063db006a7514610a7b578063dd62ed3e14610a98578063e9c714f214610ac6576102f1565b8063a0712d6811610156578063aa5af0fd11610130578063aa5af0fd1461099e578063ae9d70b0146109a6578063b2a02ff1146109ae578063b71d1a0c146109e4576102f1565b8063a0712d681461094d578063a6afed951461096a578063a9059cbb14610972576102f1565b806373acee98146107a4578063852a12e3146107ac5780638f840ddd146107c957806395d89b41146107d157806395dd9193146107d957806399d8c1b4146107ff576102f1565b8063313ce5671161025c57806356e6772811610215578063601a0bf1116101ef578063601a0bf1146107515780636c540baf1461076e5780636f307dc31461077657806370a082311461077e576102f1565b806356e677281461069d5780635c60da1b146107415780635fe3b56714610749576102f1565b8063313ce567146106065780633af9e669146106245780633b1d21a21461064a5780633e941010146106525780634576b5db1461066f57806347bd371814610695576102f1565b806318160ddd116102ae57806318160ddd1461041a578063182df0f5146104225780631a31d4651461042a57806323b872dd146105805780632608f818146105b657806326782247146105e2576102f1565b806306fdde03146102f6578063095ea7b3146103735780630e752702146103b3578063153ab505146103e2578063173b9904146103ec57806317bfdfbc146103f4575b600080fd5b6102fe610b67565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610338578181015183820152602001610320565b50505050905090810190601f1680156103655780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61039f6004803603604081101561038957600080fd5b506001600160a01b038135169060200135610bf4565b604080519115158252519081900360200190f35b6103d0600480360360208110156103c957600080fd5b5035610c61565b60408051918252519081900360200190f35b6103ea610c77565b005b6103d0610cc7565b6103d06004803603602081101561040a57600080fd5b50356001600160a01b0316610ccd565b6103d0610d8d565b6103d0610d93565b6103ea600480360360e081101561044057600080fd5b6001600160a01b03823581169260208101358216926040820135909216916060820135919081019060a081016080820135600160201b81111561048257600080fd5b82018360208201111561049457600080fd5b803590602001918460018302840111600160201b831117156104b557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561050757600080fd5b82018360208201111561051957600080fd5b803590602001918460018302840111600160201b8311171561053a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff169150610df69050565b61039f6004803603606081101561059657600080fd5b506001600160a01b03813581169160208101359091169060400135610e95565b6103d0600480360360408110156105cc57600080fd5b506001600160a01b038135169060200135610f07565b6105ea610f1d565b604080516001600160a01b039092168252519081900360200190f35b61060e610f2c565b6040805160ff9092168252519081900360200190f35b6103d06004803603602081101561063a57600080fd5b50356001600160a01b0316610f35565b6103d0610feb565b6103d06004803603602081101561066857600080fd5b5035610ffa565b6103d06004803603602081101561068557600080fd5b50356001600160a01b0316611005565b6103d061115a565b6103ea600480360360208110156106b357600080fd5b810190602081018135600160201b8111156106cd57600080fd5b8201836020820111156106df57600080fd5b803590602001918460018302840111600160201b8311171561070057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611160945050505050565b6105ea6111b1565b6105ea6111c0565b6103d06004803603602081101561076757600080fd5b50356111cf565b6103d061126a565b6105ea611270565b6103d06004803603602081101561079457600080fd5b50356001600160a01b031661127f565b6103d061129a565b6103d0600480360360208110156107c257600080fd5b5035611350565b6103d061135b565b6102fe611361565b6103d0600480360360208110156107ef57600080fd5b50356001600160a01b03166113b9565b6103ea600480360360c081101561081557600080fd5b6001600160a01b03823581169260208101359091169160408201359190810190608081016060820135600160201b81111561084f57600080fd5b82018360208201111561086157600080fd5b803590602001918460018302840111600160201b8311171561088257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156108d457600080fd5b8201836020820111156108e657600080fd5b803590602001918460018302840111600160201b8311171561090757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff1691506114169050565b6103d06004803603602081101561096357600080fd5b50356115fd565b6103d0611609565b61039f6004803603604081101561098857600080fd5b506001600160a01b038135169060200135611961565b6103d06119d2565b6103d06119d8565b6103d0600480360360608110156109c457600080fd5b506001600160a01b03813581169160208101359091169060400135611a77565b6103d0600480360360208110156109fa57600080fd5b50356001600160a01b0316611ae8565b6103d0611b74565b610a3860048036036020811015610a2857600080fd5b50356001600160a01b0316611c30565b604080519485526020850193909352838301919091526060830152519081900360800190f35b6103d060048036036020811015610a7457600080fd5b5035611cc5565b6103d060048036036020811015610a9157600080fd5b5035611cd0565b6103d060048036036040811015610aae57600080fd5b506001600160a01b0381358116916020013516611cdb565b6103d0611d06565b6103d060048036036020811015610ae457600080fd5b50356001600160a01b0316611e09565b6105ea611e43565b6103d060048036036060811015610b1257600080fd5b506001600160a01b03813581169160208101359160409091013516611e52565b6105ea611e6a565b6103d0611e7e565b6103d060048036036020811015610b5857600080fd5b5035611ee2565b61039f611f60565b60018054604080516020600284861615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610bec5780601f10610bc157610100808354040283529160200191610bec565b820191906000526020600020905b815481529060010190602001808311610bcf57829003601f168201915b505050505081565b336000818152600f602090815260408083206001600160a01b03871680855290835281842086905581518681529151939493909284927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a360019150505b92915050565b600080610c6d83611f65565b509150505b919050565b60035461010090046001600160a01b03163314610cc55760405162461bcd60e51b815260040180806020018281038252602d815260200180614ee1602d913960400191505060405180910390fd5b565b60085481565b6000805460ff16610d12576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155610d24611609565b14610d6f576040805162461bcd60e51b81526020600482015260166024820152751858d8dc9d59481a5b9d195c995cdd0819985a5b195960521b604482015290519081900360640190fd5b610d78826113b9565b90505b6000805460ff19166001179055919050565b600d5481565b6000806000610da061200e565b90925090506000826003811115610db357fe5b14610def5760405162461bcd60e51b815260040180806020018281038252603581526020018061502e6035913960400191505060405180910390fd5b9150505b90565b610e04868686868686611416565b601180546001600160a01b0319166001600160a01b038981169190911791829055604080516318160ddd60e01b8152905192909116916318160ddd91600480820192602092909190829003018186803b158015610e6057600080fd5b505afa158015610e74573d6000803e3d6000fd5b505050506040513d6020811015610e8a57600080fd5b505050505050505050565b6000805460ff16610eda576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155610ef0338686866120bd565b1490506000805460ff191660011790559392505050565b600080610f1484846123cb565b50949350505050565b6004546001600160a01b031681565b60035460ff1681565b6000610f3f614cef565b6040518060200160405280610f52611b74565b90526001600160a01b0384166000908152600e6020526040812054919250908190610f7e908490612476565b90925090506000826003811115610f9157fe5b14610fe3576040805162461bcd60e51b815260206004820152601f60248201527f62616c616e636520636f756c64206e6f742062652063616c63756c6174656400604482015290519081900360640190fd5b949350505050565b6000610ff56124ca565b905090565b6000610c5b8261254a565b60035460009061010090046001600160a01b031633146110325761102b6001603f6125de565b9050610c72565b60055460408051623f1ee960e11b815290516001600160a01b0392831692851691627e3dd2916004808301926020929190829003018186803b15801561107757600080fd5b505afa15801561108b573d6000803e3d6000fd5b505050506040513d60208110156110a157600080fd5b50516110f4576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600580546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517f7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d9281900390910190a160005b9392505050565b600b5481565b60035461010090046001600160a01b031633146111ae5760405162461bcd60e51b815260040180806020018281038252602d8152602001806150e3602d913960400191505060405180910390fd5b50565b6012546001600160a01b031681565b6005546001600160a01b031681565b6000805460ff16611214576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155611226611609565b9050801561124c5761124481601081111561123d57fe5b60306125de565b915050610d7b565b61125583612644565b9150506000805460ff19166001179055919050565b60095481565b6011546001600160a01b031681565b6001600160a01b03166000908152600e602052604090205490565b6000805460ff166112df576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff191681556112f1611609565b1461133c576040805162461bcd60e51b81526020600482015260166024820152751858d8dc9d59481a5b9d195c995cdd0819985a5b195960521b604482015290519081900360640190fd5b50600b546000805460ff1916600117905590565b6000610c5b82612777565b600c5481565b6002805460408051602060018416156101000260001901909316849004601f81018490048402820184019092528181529291830182828015610bec5780601f10610bc157610100808354040283529160200191610bec565b60008060006113c7846127f8565b909250905060008260038111156113da57fe5b146111535760405162461bcd60e51b8152600401808060200182810382526037815260200180614f396037913960400191505060405180910390fd5b60035461010090046001600160a01b031633146114645760405162461bcd60e51b8152600401808060200182810382526024815260200180614e486024913960400191505060405180910390fd5b6009541580156114745750600a54155b6114af5760405162461bcd60e51b8152600401808060200182810382526023815260200180614e6c6023913960400191505060405180910390fd5b6007849055836114f05760405162461bcd60e51b8152600401808060200182810382526030815260200180614e8f6030913960400191505060405180910390fd5b60006114fb87611005565b90508015611550576040805162461bcd60e51b815260206004820152601a60248201527f73657474696e6720636f6d7074726f6c6c6572206661696c6564000000000000604482015290519081900360640190fd5b6115586128ac565b600955670de0b6b3a7640000600a55611570866128b0565b905080156115af5760405162461bcd60e51b8152600401808060200182810382526022815260200180614ebf6022913960400191505060405180910390fd5b83516115c2906001906020870190614d02565b5082516115d6906002906020860190614d02565b50506003805460ff90921660ff199283161790556000805490911660011790555050505050565b600080610c6d83612a25565b6000806116146128ac565b6009549091508082141561162d57600092505050610df3565b60006116376124ca565b600b54600c54600a54600654604080516315f2405360e01b815260048101879052602481018690526044810185905290519596509394929391926000926001600160a01b03909216916315f24053916064808301926020929190829003018186803b1580156116a557600080fd5b505afa1580156116b9573d6000803e3d6000fd5b505050506040513d60208110156116cf57600080fd5b5051905065048c2739500081111561172e576040805162461bcd60e51b815260206004820152601c60248201527f626f72726f772072617465206973206162737572646c79206869676800000000604482015290519081900360640190fd5b60008061173b8989612aa6565b9092509050600082600381111561174e57fe5b146117a0576040805162461bcd60e51b815260206004820152601f60248201527f636f756c64206e6f742063616c63756c61746520626c6f636b2064656c746100604482015290519081900360640190fd5b6117a8614cef565b6000806000806117c660405180602001604052808a81525087612ac9565b909750945060008760038111156117d957fe5b1461180b576117f6600960068960038111156117f157fe5b612b31565b9e505050505050505050505050505050610df3565b611815858c612476565b9097509350600087600381111561182857fe5b14611840576117f6600960018960038111156117f157fe5b61184a848c612b97565b9097509250600087600381111561185d57fe5b14611875576117f6600960048960038111156117f157fe5b6118906040518060200160405280600854815250858c612bbd565b909750915060008760038111156118a357fe5b146118bb576117f6600960058960038111156117f157fe5b6118c6858a8b612bbd565b909750905060008760038111156118d957fe5b146118f1576117f6600960038960038111156117f157fe5b60098e9055600a819055600b839055600c829055604080518d8152602081018690528082018390526060810185905290517f4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc049181900360800190a160009e50505050505050505050505050505090565b6000805460ff166119a6576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff191681556119bc333386866120bd565b1490506000805460ff1916600117905592915050565b600a5481565b6006546000906001600160a01b031663b81688166119f46124ca565b600b54600c546008546040518563ffffffff1660e01b81526004018085815260200184815260200183815260200182815260200194505050505060206040518083038186803b158015611a4657600080fd5b505afa158015611a5a573d6000803e3d6000fd5b505050506040513d6020811015611a7057600080fd5b5051905090565b6000805460ff16611abc576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19169055611ad233858585612c19565b90506000805460ff191660011790559392505050565b60035460009061010090046001600160a01b03163314611b0e5761102b600160456125de565b600480546001600160a01b038481166001600160a01b0319831681179093556040805191909216808252602082019390935281517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9929181900390910190a16000611153565b6000805460ff16611bb9576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155611bcb611609565b14611c16576040805162461bcd60e51b81526020600482015260166024820152751858d8dc9d59481a5b9d195c995cdd0819985a5b195960521b604482015290519081900360640190fd5b611c1e610d93565b90506000805460ff1916600117905590565b6001600160a01b0381166000908152600e6020526040812054819081908190818080611c5b896127f8565b935090506000816003811115611c6d57fe5b14611c8b5760095b975060009650869550859450611cbe9350505050565b611c9361200e565b925090506000816003811115611ca557fe5b14611cb1576009611c75565b5060009650919450925090505b9193509193565b6000610c5b82612e7f565b6000610c5b82612efe565b6001600160a01b039182166000908152600f6020908152604080832093909416825291909152205490565b6004546000906001600160a01b031633141580611d21575033155b15611d3957611d32600160006125de565b9050610df3565b60038054600480546001600160a01b03818116610100818102610100600160a81b0319871617968790556001600160a01b031990931690935560408051948390048216808652929095041660208401528351909391927ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc92908290030190a1600454604080516001600160a01b038085168252909216602083015280517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a99281900390910190a160009250505090565b600080611e14611609565b90508015611e3a57611e32816010811115611e2b57fe5b60406125de565b915050610c72565b611153836128b0565b6006546001600160a01b031681565b600080611e60858585612f78565b5095945050505050565b60035461010090046001600160a01b031681565b6006546000906001600160a01b03166315f24053611e9a6124ca565b600b54600c546040518463ffffffff1660e01b815260040180848152602001838152602001828152602001935050505060206040518083038186803b158015611a4657600080fd5b6000805460ff16611f27576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155611f39611609565b90508015611f5757611244816010811115611f5057fe5b60466125de565b611255836130aa565b600181565b60008054819060ff16611fac576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155611fbe611609565b90508015611fe957611fdc816010811115611fd557fe5b60366125de565b925060009150611ffa9050565b611ff4333386613152565b92509250505b6000805460ff191660011790559092909150565b600d54600090819080612029575050600754600091506120b9565b60006120336124ca565b9050600061203f614cef565b600061205084600b54600c54613537565b93509050600081600381111561206257fe5b14612077579550600094506120b99350505050565b6120818386613575565b92509050600081600381111561209357fe5b146120a8579550600094506120b99350505050565b50516000955093506120b992505050565b9091565b600554604080516317b9b84b60e31b81523060048201526001600160a01b03868116602483015285811660448301526064820185905291516000938493169163bdcdc25891608480830192602092919082900301818787803b15801561212257600080fd5b505af1158015612136573d6000803e3d6000fd5b505050506040513d602081101561214c57600080fd5b50519050801561216b576121636003604a83612b31565b915050610fe3565b836001600160a01b0316856001600160a01b03161415612191576121636002604b6125de565b60006001600160a01b0387811690871614156121b057506000196121d8565b506001600160a01b038086166000908152600f60209081526040808320938a16835292905220545b6000806000806121e88589612aa6565b909450925060008460038111156121fb57fe5b146122195761220c6009604b6125de565b9650505050505050610fe3565b6001600160a01b038a166000908152600e602052604090205461223c9089612aa6565b9094509150600084600381111561224f57fe5b146122605761220c6009604c6125de565b6001600160a01b0389166000908152600e60205260409020546122839089612b97565b9094509050600084600381111561229657fe5b146122a75761220c6009604d6125de565b6001600160a01b03808b166000908152600e6020526040808220859055918b1681522081905560001985146122ff576001600160a01b03808b166000908152600f60209081526040808320938f168352929052208390555b886001600160a01b03168a6001600160a01b0316600080516020614faa8339815191528a6040518082815260200191505060405180910390a36005546040805163352b4a3f60e11b81523060048201526001600160a01b038d811660248301528c81166044830152606482018c905291519190921691636a56947e91608480830192600092919082900301818387803b15801561239b57600080fd5b505af11580156123af573d6000803e3d6000fd5b50600092506123bc915050565b9b9a5050505050505050505050565b60008054819060ff16612412576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155612424611609565b9050801561244f5761244281601081111561243b57fe5b60356125de565b9250600091506124609050565b61245a338686613152565b92509250505b6000805460ff1916600117905590939092509050565b6000806000612483614cef565b61248d8686612ac9565b909250905060008260038111156124a057fe5b146124b157509150600090506124c3565b60006124bc82613625565b9350935050505b9250929050565b601154604080516370a0823160e01b815230600482015290516000926001600160a01b03169182916370a0823191602480820192602092909190829003018186803b15801561251857600080fd5b505afa15801561252c573d6000803e3d6000fd5b505050506040513d602081101561254257600080fd5b505191505090565b6000805460ff1661258f576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff191681556125a1611609565b905080156125bf576112448160108111156125b857fe5b604e6125de565b6125c883613634565b509150506000805460ff19166001179055919050565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa083601081111561260d57fe5b83605081111561261957fe5b604080519283526020830191909152600082820152519081900360600190a182601081111561115357fe5b600354600090819061010090046001600160a01b0316331461266c57611e32600160316125de565b6126746128ac565b6009541461268857611e32600a60336125de565b826126916124ca565b10156126a357611e32600e60326125de565b600c548311156126b957611e32600260346125de565b50600c54828103908111156126ff5760405162461bcd60e51b81526004018080602001828103825260248152602001806150bf6024913960400191505060405180910390fd5b600c81905560035461271f9061010090046001600160a01b03168461371c565b600354604080516101009092046001600160a01b0316825260208201859052818101839052517f3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e916060908290030190a16000611153565b6000805460ff166127bc576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff191681556127ce611609565b905080156127ec576112448160108111156127e557fe5b60276125de565b61125533600085613813565b6001600160a01b03811660009081526010602052604081208054829182918291829161282f5750600094508493506128a792505050565b61283f8160000154600a54613cda565b9094509250600084600381111561285257fe5b146128675750919350600092506128a7915050565b612875838260010154613d19565b9094509150600084600381111561288857fe5b1461289d5750919350600092506128a7915050565b5060009450925050505b915091565b4390565b600354600090819061010090046001600160a01b031633146128d857611e32600160426125de565b6128e06128ac565b600954146128f457611e32600a60416125de565b600660009054906101000a90046001600160a01b03169050826001600160a01b0316632191f92a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561294557600080fd5b505afa158015612959573d6000803e3d6000fd5b505050506040513d602081101561296f57600080fd5b50516129c2576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517fedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f9269281900390910190a16000611153565b60008054819060ff16612a6c576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155612a7e611609565b90508015612a9c57611fdc816010811115612a9557fe5b601e6125de565b611ff43385613d44565b600080838311612abd5750600090508183036124c3565b506003905060006124c3565b6000612ad3614cef565b600080612ae4866000015186613cda565b90925090506000826003811115612af757fe5b14612b16575060408051602081019091526000815290925090506124c3565b60408051602081019091529081526000969095509350505050565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0846010811115612b6057fe5b846050811115612b6c57fe5b604080519283526020830191909152818101859052519081900360600190a1836010811115610fe357fe5b600080838301848110612baf576000925090506124c3565b5060029150600090506124c3565b6000806000612bca614cef565b612bd48787612ac9565b90925090506000826003811115612be757fe5b14612bf85750915060009050612c11565b612c0a612c0482613625565b86612b97565b9350935050505b935093915050565b6005546040805163d02f735160e01b81523060048201526001600160a01b038781166024830152868116604483015285811660648301526084820185905291516000938493169163d02f73519160a480830192602092919082900301818787803b158015612c8657600080fd5b505af1158015612c9a573d6000803e3d6000fd5b505050506040513d6020811015612cb057600080fd5b505190508015612cc7576121636003601b83612b31565b846001600160a01b0316846001600160a01b03161415612ced576121636006601c6125de565b6001600160a01b0384166000908152600e602052604081205481908190612d149087612aa6565b90935091506000836003811115612d2757fe5b14612d4a57612d3f6009601a8560038111156117f157fe5b945050505050610fe3565b6001600160a01b0388166000908152600e6020526040902054612d6d9087612b97565b90935090506000836003811115612d8057fe5b14612d9857612d3f600960198560038111156117f157fe5b6001600160a01b038088166000818152600e60209081526040808320879055938c168083529184902085905583518a815293519193600080516020614faa833981519152929081900390910190a360055460408051636d35bf9160e01b81523060048201526001600160a01b038c811660248301528b811660448301528a81166064830152608482018a905291519190921691636d35bf919160a480830192600092919082900301818387803b158015612e5157600080fd5b505af1158015612e65573d6000803e3d6000fd5b5060009250612e72915050565b9998505050505050505050565b6000805460ff16612ec4576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155612ed6611609565b90508015612ef457611244816010811115612eed57fe5b60086125de565b61125533846141a3565b6000805460ff16612f43576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155612f55611609565b90508015612f6c576112448160108111156127e557fe5b61125533846000613813565b60008054819060ff16612fbf576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6000805460ff19168155612fd1611609565b90508015612ffc57612fef816010811115612fe857fe5b600f6125de565b9250600091506130939050565b836001600160a01b031663a6afed956040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561303757600080fd5b505af115801561304b573d6000803e3d6000fd5b505050506040513d602081101561306157600080fd5b50519050801561308157612fef81601081111561307a57fe5b60106125de565b61308d338787876144b1565b92509250505b6000805460ff191660011790559094909350915050565b60035460009061010090046001600160a01b031633146130d05761102b600160476125de565b6130d86128ac565b600954146130ec5761102b600a60486125de565b670de0b6b3a76400008211156131085761102b600260496125de565b6008805490839055604080518281526020810185905281517faaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460929181900390910190a16000611153565b60055460408051631200453160e11b81523060048201526001600160a01b0386811660248301528581166044830152606482018590529151600093849384939116916324008a629160848082019260209290919082900301818787803b1580156131bb57600080fd5b505af11580156131cf573d6000803e3d6000fd5b505050506040513d60208110156131e557600080fd5b505190508015613209576131fc6003603883612b31565b925060009150612c119050565b6132116128ac565b60095414613225576131fc600a60396125de565b61322d614d80565b6001600160a01b0386166000908152601060205260409020600101546060820152613257866127f8565b608083018190526020830182600381111561326e57fe5b600381111561327957fe5b905250600090508160200151600381111561329057fe5b146132ba576132ac60096037836020015160038111156117f157fe5b935060009250612c11915050565b6000198514156132d357608081015160408201526132db565b604081018590525b6132e9878260400151614a34565b60e0820181905260808201516132fe91612aa6565b60a083018190526020830182600381111561331557fe5b600381111561332057fe5b905250600090508160200151600381111561333757fe5b146133735760405162461bcd60e51b815260040180806020018281038252603a815260200180614f70603a913960400191505060405180910390fd5b613383600b548260e00151612aa6565b60c083018190526020830182600381111561339a57fe5b60038111156133a557fe5b90525060009050816020015160038111156133bc57fe5b146133f85760405162461bcd60e51b8152600401808060200182810382526031815260200180614fca6031913960400191505060405180910390fd5b60a080820180516001600160a01b03808a16600081815260106020908152604091829020948555600a5460019095019490945560c0870151600b81905560e088015195518251948f16855294840192909252828101949094526060820192909252608081019190915290517f1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1929181900390910190a160055460e0820151606083015160408051631ededc9160e01b81523060048201526001600160a01b038c811660248301528b8116604483015260648201949094526084810192909252519190921691631ededc919160a480830192600092919082900301818387803b15801561350357600080fd5b505af1158015613517573d6000803e3d6000fd5b5060009250613524915050565b8160e00151935093505050935093915050565b6000806000806135478787612b97565b9092509050600082600381111561355a57fe5b1461356b5750915060009050612c11565b612c0a8186612aa6565b600061357f614cef565b60008061359486670de0b6b3a7640000613cda565b909250905060008260038111156135a757fe5b146135c6575060408051602081019091526000815290925090506124c3565b6000806135d38388613d19565b909250905060008260038111156135e657fe5b14613608575060408051602081019091526000815290945092506124c3915050565b604080516020810190915290815260009890975095505050505050565b51670de0b6b3a7640000900490565b6000806000806136426128ac565b6009541461366157613656600a604f6125de565b935091506128a79050565b61366b3386614a34565b905080600c54019150600c548210156136cb576040805162461bcd60e51b815260206004820181905260248201527f61646420726573657276657320756e6578706563746564206f766572666c6f77604482015290519081900360640190fd5b600c829055604080513381526020810183905280820184905290517fa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc59181900360600190a160009350915050915091565b6011546040805163a9059cbb60e01b81526001600160a01b0385811660048301526024820185905291519190921691829163a9059cbb9160448082019260009290919082900301818387803b15801561377457600080fd5b505af1158015613788573d6000803e3d6000fd5b5050505060003d600081146137a457602081146137ae57600080fd5b60001991506137ba565b60206000803e60005191505b508061380d576040805162461bcd60e51b815260206004820152601960248201527f544f4b454e5f5452414e534645525f4f55545f4641494c454400000000000000604482015290519081900360640190fd5b50505050565b6000821580613820575081155b61385b5760405162461bcd60e51b815260040180806020018281038252603481526020018061508b6034913960400191505060405180910390fd5b613863614dc6565b61386b61200e565b604083018190526020830182600381111561388257fe5b600381111561388d57fe5b90525060009050816020015160038111156138a457fe5b146138c8576138c06009602b836020015160038111156117f157fe5b915050611153565b83156139495760608101849052604080516020810182529082015181526138ef9085612476565b608083018190526020830182600381111561390657fe5b600381111561391157fe5b905250600090508160200151600381111561392857fe5b14613944576138c060096029836020015160038111156117f157fe5b6139c2565b6139658360405180602001604052808460400151815250614c7e565b606083018190526020830182600381111561397c57fe5b600381111561398757fe5b905250600090508160200151600381111561399e57fe5b146139ba576138c06009602a836020015160038111156117f157fe5b608081018390525b60055460608201516040805163eabe7d9160e01b81523060048201526001600160a01b03898116602483015260448201939093529051600093929092169163eabe7d919160648082019260209290919082900301818787803b158015613a2757600080fd5b505af1158015613a3b573d6000803e3d6000fd5b505050506040513d6020811015613a5157600080fd5b505190508015613a7157613a686003602883612b31565b92505050611153565b613a796128ac565b60095414613a8d57613a68600a602c6125de565b613a9d600d548360600151612aa6565b60a0840181905260208401826003811115613ab457fe5b6003811115613abf57fe5b9052506000905082602001516003811115613ad657fe5b14613af257613a686009602e846020015160038111156117f157fe5b6001600160a01b0386166000908152600e60205260409020546060830151613b1a9190612aa6565b60c0840181905260208401826003811115613b3157fe5b6003811115613b3c57fe5b9052506000905082602001516003811115613b5357fe5b14613b6f57613a686009602d846020015160038111156117f157fe5b8160800151613b7c6124ca565b1015613b8e57613a68600e602f6125de565b613b9c86836080015161371c565b60a0820151600d5560c08201516001600160a01b0387166000818152600e6020908152604091829020939093556060850151815190815290513093600080516020614faa833981519152928290030190a36080820151606080840151604080516001600160a01b038b168152602081019490945283810191909152517fe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a9299281900390910190a160055460808301516060840151604080516351dff98960e01b81523060048201526001600160a01b038b81166024830152604482019490945260648101929092525191909216916351dff98991608480830192600092919082900301818387803b158015613caf57600080fd5b505af1158015613cc3573d6000803e3d6000fd5b5060009250613cd0915050565b9695505050505050565b60008083613ced575060009050806124c3565b83830283858281613cfa57fe5b0414613d0e575060029150600090506124c3565b6000925090506124c3565b60008082613d2d57506001905060006124c3565b6000838581613d3857fe5b04915091509250929050565b60055460408051634ef4c3e160e01b81523060048201526001600160a01b03858116602483015260448201859052915160009384938493911691634ef4c3e19160648082019260209290919082900301818787803b158015613da557600080fd5b505af1158015613db9573d6000803e3d6000fd5b505050506040513d6020811015613dcf57600080fd5b505190508015613df357613de66003601f83612b31565b9250600091506124c39050565b613dfb6128ac565b60095414613e0f57613de6600a60226125de565b613e17614dc6565b613e1f61200e565b6040830181905260208301826003811115613e3657fe5b6003811115613e4157fe5b9052506000905081602001516003811115613e5857fe5b14613e8257613e7460096021836020015160038111156117f157fe5b9350600092506124c3915050565b613e8c8686614a34565b60c0820181905260408051602081018252908301518152613ead9190614c7e565b6060830181905260208301826003811115613ec457fe5b6003811115613ecf57fe5b9052506000905081602001516003811115613ee657fe5b14613f38576040805162461bcd60e51b815260206004820181905260248201527f4d494e545f45584348414e47455f43414c43554c4154494f4e5f4641494c4544604482015290519081900360640190fd5b613f48600d548260600151612b97565b6080830181905260208301826003811115613f5f57fe5b6003811115613f6a57fe5b9052506000905081602001516003811115613f8157fe5b14613fbd5760405162461bcd60e51b81526004018080602001828103825260288152602001806150636028913960400191505060405180910390fd5b6001600160a01b0386166000908152600e60205260409020546060820151613fe59190612b97565b60a0830181905260208301826003811115613ffc57fe5b600381111561400757fe5b905250600090508160200151600381111561401e57fe5b1461405a5760405162461bcd60e51b815260040180806020018281038252602b815260200180614f0e602b913960400191505060405180910390fd5b6080810151600d5560a08101516001600160a01b0387166000818152600e60209081526040918290209390935560c084015160608086015183519485529484019190915282820193909352517f4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f929181900390910190a1606081015160408051918252516001600160a01b038816913091600080516020614faa8339815191529181900360200190a360055460c08201516060830151604080516341c728b960e01b81523060048201526001600160a01b038b81166024830152604482019490945260648101929092525191909216916341c728b991608480830192600092919082900301818387803b15801561417057600080fd5b505af1158015614184573d6000803e3d6000fd5b5060009250614191915050565b8160c001519350935050509250929050565b6005546040805163368f515360e21b81523060048201526001600160a01b0385811660248301526044820185905291516000938493169163da3d454c91606480830192602092919082900301818787803b15801561420057600080fd5b505af1158015614214573d6000803e3d6000fd5b505050506040513d602081101561422a57600080fd5b505190508015614249576142416003600e83612b31565b915050610c5b565b6142516128ac565b6009541461426457614241600a806125de565b8261426d6124ca565b101561427f57614241600e60096125de565b614287614e04565b614290856127f8565b60208301819052828260038111156142a457fe5b60038111156142af57fe5b90525060009050815160038111156142c357fe5b146142e8576142df60096007836000015160038111156117f157fe5b92505050610c5b565b6142f6816020015185612b97565b604083018190528282600381111561430a57fe5b600381111561431557fe5b905250600090508151600381111561432957fe5b14614345576142df6009600c836000015160038111156117f157fe5b614351600b5485612b97565b606083018190528282600381111561436557fe5b600381111561437057fe5b905250600090508151600381111561438457fe5b146143a0576142df6009600b836000015160038111156117f157fe5b6143aa858561371c565b604080820180516001600160a01b03881660008181526010602090815290859020928355600a54600190930192909255606080860151600b81905593518551928352928201899052818501929092529081019190915290517f13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab809181900360800190a160055460408051635c77860560e01b81523060048201526001600160a01b0388811660248301526044820188905291519190921691635c77860591606480830192600092919082900301818387803b15801561448757600080fd5b505af115801561449b573d6000803e3d6000fd5b50600092506144a8915050565b95945050505050565b60055460408051632fe3f38f60e11b81523060048201526001600160a01b0384811660248301528781166044830152868116606483015260848201869052915160009384938493911691635fc7e71e9160a48082019260209290919082900301818787803b15801561452257600080fd5b505af1158015614536573d6000803e3d6000fd5b505050506040513d602081101561454c57600080fd5b505190508015614570576145636003601283612b31565b925060009150614a2b9050565b6145786128ac565b6009541461458c57614563600a60166125de565b6145946128ac565b846001600160a01b0316636c540baf6040518163ffffffff1660e01b815260040160206040518083038186803b1580156145cd57600080fd5b505afa1580156145e1573d6000803e3d6000fd5b505050506040513d60208110156145f757600080fd5b50511461460a57614563600a60116125de565b866001600160a01b0316866001600160a01b0316141561463057614563600660176125de565b8461464157614563600760156125de565b60001985141561465757614563600760146125de565b600080614665898989613152565b909250905081156146955761468682601081111561467f57fe5b60186125de565b945060009350614a2b92505050565b6005546040805163c488847b60e01b81523060048201526001600160a01b038981166024830152604482018590528251600094859492169263c488847b926064808301939192829003018186803b1580156146ef57600080fd5b505afa158015614703573d6000803e3d6000fd5b505050506040513d604081101561471957600080fd5b508051602090910151909250905081156147645760405162461bcd60e51b8152600401808060200182810382526033815260200180614ffb6033913960400191505060405180910390fd5b80886001600160a01b03166370a082318c6040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156147bb57600080fd5b505afa1580156147cf573d6000803e3d6000fd5b505050506040513d60208110156147e557600080fd5b5051101561483a576040805162461bcd60e51b815260206004820152601860248201527f4c49515549444154455f5345495a455f544f4f5f4d5543480000000000000000604482015290519081900360640190fd5b60006001600160a01b03891630141561486057614859308d8d85612c19565b90506148ea565b6040805163b2a02ff160e01b81526001600160a01b038e811660048301528d81166024830152604482018590529151918b169163b2a02ff1916064808201926020929091908290030181600087803b1580156148bb57600080fd5b505af11580156148cf573d6000803e3d6000fd5b505050506040513d60208110156148e557600080fd5b505190505b8015614934576040805162461bcd60e51b81526020600482015260146024820152731d1bdad95b881cd95a5e9d5c994819985a5b195960621b604482015290519081900360640190fd5b604080516001600160a01b03808f168252808e1660208301528183018790528b1660608201526080810184905290517f298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb529181900360a00190a1600554604080516347ef3b3b60e01b81523060048201526001600160a01b038c811660248301528f811660448301528e811660648301526084820188905260a48201869052915191909216916347ef3b3b9160c480830192600092919082900301818387803b1580156149ff57600080fd5b505af1158015614a13573d6000803e3d6000fd5b5060009250614a20915050565b975092955050505050505b94509492505050565b601154604080516370a0823160e01b815230600482015290516000926001600160a01b031691839183916370a08231916024808301926020929190829003018186803b158015614a8357600080fd5b505afa158015614a97573d6000803e3d6000fd5b505050506040513d6020811015614aad57600080fd5b5051604080516323b872dd60e01b81526001600160a01b038881166004830152306024830152604482018890529151929350908416916323b872dd9160648082019260009290919082900301818387803b158015614b0a57600080fd5b505af1158015614b1e573d6000803e3d6000fd5b5050505060003d60008114614b3a5760208114614b4457600080fd5b6000199150614b50565b60206000803e60005191505b5080614ba3576040805162461bcd60e51b815260206004820152601860248201527f544f4b454e5f5452414e534645525f494e5f4641494c45440000000000000000604482015290519081900360640190fd5b601154604080516370a0823160e01b815230600482015290516000926001600160a01b0316916370a08231916024808301926020929190829003018186803b158015614bee57600080fd5b505afa158015614c02573d6000803e3d6000fd5b505050506040513d6020811015614c1857600080fd5b5051905082811015614c71576040805162461bcd60e51b815260206004820152601a60248201527f544f4b454e5f5452414e534645525f494e5f4f564552464c4f57000000000000604482015290519081900360640190fd5b9190910395945050505050565b6000806000614c8b614cef565b61248d86866000614c9a614cef565b600080614caf670de0b6b3a764000087613cda565b90925090506000826003811115614cc257fe5b14614ce1575060408051602081019091526000815290925090506124c3565b6124bc818660000151613575565b6040518060200160405280600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614d4357805160ff1916838001178555614d70565b82800160010185558215614d70579182015b82811115614d70578251825591602001919060010190614d55565b50614d7c929150614e2d565b5090565b6040805161010081019091528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6040805160e0810190915280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b604080516080810190915280600081526020016000815260200160008152602001600081525090565b610df391905b80821115614d7c5760008155600101614e3356fe6f6e6c792061646d696e206d617920696e697469616c697a6520746865206d61726b65746d61726b6574206d6179206f6e6c7920626520696e697469616c697a6564206f6e6365696e697469616c2065786368616e67652072617465206d7573742062652067726561746572207468616e207a65726f2e73657474696e6720696e7465726573742072617465206d6f64656c206661696c65646f6e6c79207468652061646d696e206d61792063616c6c205f72657369676e496d706c656d656e746174696f6e4d494e545f4e45575f4143434f554e545f42414c414e43455f43414c43554c4154494f4e5f4641494c4544626f72726f7742616c616e636553746f7265643a20626f72726f7742616c616e636553746f726564496e7465726e616c206661696c656452455041595f424f52524f575f4e45575f4143434f554e545f424f52524f575f42414c414e43455f43414c43554c4154494f4e5f4641494c4544ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef52455041595f424f52524f575f4e45575f544f54414c5f42414c414e43455f43414c43554c4154494f4e5f4641494c45444c49515549444154455f434f4d5054524f4c4c45525f43414c43554c4154455f414d4f554e545f5345495a455f4641494c454465786368616e67655261746553746f7265643a2065786368616e67655261746553746f726564496e7465726e616c206661696c65644d494e545f4e45575f544f54414c5f535550504c595f43414c43554c4154494f4e5f4641494c45446f6e65206f662072656465656d546f6b656e73496e206f722072656465656d416d6f756e74496e206d757374206265207a65726f72656475636520726573657276657320756e657870656374656420756e646572666c6f776f6e6c79207468652061646d696e206d61792063616c6c205f6265636f6d65496d706c656d656e746174696f6ea265627a7a723158205d5f96c30495cbc2e281abf3f8c2e33ab9c7597f596024aa8c8713bfe64ee3c364736f6c63430005100032"

// DeployAbicerc20delegate deploys a new Ethereum contract, binding an instance of Abicerc20delegate to it.
func DeployAbicerc20delegate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Abicerc20delegate, error) {
	parsed, err := abi.JSON(strings.NewReader(Abicerc20delegateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(Abicerc20delegateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Abicerc20delegate{Abicerc20delegateCaller: Abicerc20delegateCaller{contract: contract}, Abicerc20delegateTransactor: Abicerc20delegateTransactor{contract: contract}, Abicerc20delegateFilterer: Abicerc20delegateFilterer{contract: contract}}, nil
}

// Abicerc20delegate is an auto generated Go binding around an Ethereum contract.
type Abicerc20delegate struct {
	Abicerc20delegateCaller     // Read-only binding to the contract
	Abicerc20delegateTransactor // Write-only binding to the contract
	Abicerc20delegateFilterer   // Log filterer for contract events
}

// Abicerc20delegateCaller is an auto generated read-only Go binding around an Ethereum contract.
type Abicerc20delegateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abicerc20delegateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Abicerc20delegateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abicerc20delegateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Abicerc20delegateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Abicerc20delegateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Abicerc20delegateSession struct {
	Contract     *Abicerc20delegate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// Abicerc20delegateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Abicerc20delegateCallerSession struct {
	Contract *Abicerc20delegateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// Abicerc20delegateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Abicerc20delegateTransactorSession struct {
	Contract     *Abicerc20delegateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// Abicerc20delegateRaw is an auto generated low-level Go binding around an Ethereum contract.
type Abicerc20delegateRaw struct {
	Contract *Abicerc20delegate // Generic contract binding to access the raw methods on
}

// Abicerc20delegateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Abicerc20delegateCallerRaw struct {
	Contract *Abicerc20delegateCaller // Generic read-only contract binding to access the raw methods on
}

// Abicerc20delegateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Abicerc20delegateTransactorRaw struct {
	Contract *Abicerc20delegateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbicerc20delegate creates a new instance of Abicerc20delegate, bound to a specific deployed contract.
func NewAbicerc20delegate(address common.Address, backend bind.ContractBackend) (*Abicerc20delegate, error) {
	contract, err := bindAbicerc20delegate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegate{Abicerc20delegateCaller: Abicerc20delegateCaller{contract: contract}, Abicerc20delegateTransactor: Abicerc20delegateTransactor{contract: contract}, Abicerc20delegateFilterer: Abicerc20delegateFilterer{contract: contract}}, nil
}

// NewAbicerc20delegateCaller creates a new read-only instance of Abicerc20delegate, bound to a specific deployed contract.
func NewAbicerc20delegateCaller(address common.Address, caller bind.ContractCaller) (*Abicerc20delegateCaller, error) {
	contract, err := bindAbicerc20delegate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateCaller{contract: contract}, nil
}

// NewAbicerc20delegateTransactor creates a new write-only instance of Abicerc20delegate, bound to a specific deployed contract.
func NewAbicerc20delegateTransactor(address common.Address, transactor bind.ContractTransactor) (*Abicerc20delegateTransactor, error) {
	contract, err := bindAbicerc20delegate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateTransactor{contract: contract}, nil
}

// NewAbicerc20delegateFilterer creates a new log filterer instance of Abicerc20delegate, bound to a specific deployed contract.
func NewAbicerc20delegateFilterer(address common.Address, filterer bind.ContractFilterer) (*Abicerc20delegateFilterer, error) {
	contract, err := bindAbicerc20delegate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateFilterer{contract: contract}, nil
}

// bindAbicerc20delegate binds a generic wrapper to an already deployed contract.
func bindAbicerc20delegate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Abicerc20delegateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abicerc20delegate *Abicerc20delegateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abicerc20delegate.Contract.Abicerc20delegateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abicerc20delegate *Abicerc20delegateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Abicerc20delegateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abicerc20delegate *Abicerc20delegateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Abicerc20delegateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abicerc20delegate *Abicerc20delegateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abicerc20delegate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abicerc20delegate *Abicerc20delegateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abicerc20delegate *Abicerc20delegateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) AccrualBlockNumber() (*big.Int, error) {
	return _Abicerc20delegate.Contract.AccrualBlockNumber(&_Abicerc20delegate.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _Abicerc20delegate.Contract.AccrualBlockNumber(&_Abicerc20delegate.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateSession) Admin() (common.Address, error) {
	return _Abicerc20delegate.Contract.Admin(&_Abicerc20delegate.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) Admin() (common.Address, error) {
	return _Abicerc20delegate.Contract.Admin(&_Abicerc20delegate.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Abicerc20delegate.Contract.Allowance(&_Abicerc20delegate.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Abicerc20delegate.Contract.Allowance(&_Abicerc20delegate.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Abicerc20delegate.Contract.BalanceOf(&_Abicerc20delegate.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Abicerc20delegate.Contract.BalanceOf(&_Abicerc20delegate.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Abicerc20delegate.Contract.BorrowBalanceStored(&_Abicerc20delegate.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _Abicerc20delegate.Contract.BorrowBalanceStored(&_Abicerc20delegate.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) BorrowIndex() (*big.Int, error) {
	return _Abicerc20delegate.Contract.BorrowIndex(&_Abicerc20delegate.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) BorrowIndex() (*big.Int, error) {
	return _Abicerc20delegate.Contract.BorrowIndex(&_Abicerc20delegate.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Abicerc20delegate.Contract.BorrowRatePerBlock(&_Abicerc20delegate.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _Abicerc20delegate.Contract.BorrowRatePerBlock(&_Abicerc20delegate.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateSession) Comptroller() (common.Address, error) {
	return _Abicerc20delegate.Contract.Comptroller(&_Abicerc20delegate.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) Comptroller() (common.Address, error) {
	return _Abicerc20delegate.Contract.Comptroller(&_Abicerc20delegate.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abicerc20delegate *Abicerc20delegateCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abicerc20delegate *Abicerc20delegateSession) Decimals() (uint8, error) {
	return _Abicerc20delegate.Contract.Decimals(&_Abicerc20delegate.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) Decimals() (uint8, error) {
	return _Abicerc20delegate.Contract.Decimals(&_Abicerc20delegate.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) ExchangeRateStored() (*big.Int, error) {
	return _Abicerc20delegate.Contract.ExchangeRateStored(&_Abicerc20delegate.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _Abicerc20delegate.Contract.ExchangeRateStored(&_Abicerc20delegate.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "getAccountSnapshot", account)

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
func (_Abicerc20delegate *Abicerc20delegateSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Abicerc20delegate.Contract.GetAccountSnapshot(&_Abicerc20delegate.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Abicerc20delegate.Contract.GetAccountSnapshot(&_Abicerc20delegate.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) GetCash() (*big.Int, error) {
	return _Abicerc20delegate.Contract.GetCash(&_Abicerc20delegate.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) GetCash() (*big.Int, error) {
	return _Abicerc20delegate.Contract.GetCash(&_Abicerc20delegate.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateSession) Implementation() (common.Address, error) {
	return _Abicerc20delegate.Contract.Implementation(&_Abicerc20delegate.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) Implementation() (common.Address, error) {
	return _Abicerc20delegate.Contract.Implementation(&_Abicerc20delegate.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateSession) InterestRateModel() (common.Address, error) {
	return _Abicerc20delegate.Contract.InterestRateModel(&_Abicerc20delegate.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) InterestRateModel() (common.Address, error) {
	return _Abicerc20delegate.Contract.InterestRateModel(&_Abicerc20delegate.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Abicerc20delegate *Abicerc20delegateCaller) IsCToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "isCToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Abicerc20delegate *Abicerc20delegateSession) IsCToken() (bool, error) {
	return _Abicerc20delegate.Contract.IsCToken(&_Abicerc20delegate.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) IsCToken() (bool, error) {
	return _Abicerc20delegate.Contract.IsCToken(&_Abicerc20delegate.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abicerc20delegate *Abicerc20delegateCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abicerc20delegate *Abicerc20delegateSession) Name() (string, error) {
	return _Abicerc20delegate.Contract.Name(&_Abicerc20delegate.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) Name() (string, error) {
	return _Abicerc20delegate.Contract.Name(&_Abicerc20delegate.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateSession) PendingAdmin() (common.Address, error) {
	return _Abicerc20delegate.Contract.PendingAdmin(&_Abicerc20delegate.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) PendingAdmin() (common.Address, error) {
	return _Abicerc20delegate.Contract.PendingAdmin(&_Abicerc20delegate.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Abicerc20delegate.Contract.ReserveFactorMantissa(&_Abicerc20delegate.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _Abicerc20delegate.Contract.ReserveFactorMantissa(&_Abicerc20delegate.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Abicerc20delegate.Contract.SupplyRatePerBlock(&_Abicerc20delegate.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _Abicerc20delegate.Contract.SupplyRatePerBlock(&_Abicerc20delegate.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abicerc20delegate *Abicerc20delegateCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abicerc20delegate *Abicerc20delegateSession) Symbol() (string, error) {
	return _Abicerc20delegate.Contract.Symbol(&_Abicerc20delegate.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) Symbol() (string, error) {
	return _Abicerc20delegate.Contract.Symbol(&_Abicerc20delegate.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) TotalBorrows() (*big.Int, error) {
	return _Abicerc20delegate.Contract.TotalBorrows(&_Abicerc20delegate.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) TotalBorrows() (*big.Int, error) {
	return _Abicerc20delegate.Contract.TotalBorrows(&_Abicerc20delegate.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) TotalReserves() (*big.Int, error) {
	return _Abicerc20delegate.Contract.TotalReserves(&_Abicerc20delegate.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) TotalReserves() (*big.Int, error) {
	return _Abicerc20delegate.Contract.TotalReserves(&_Abicerc20delegate.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) TotalSupply() (*big.Int, error) {
	return _Abicerc20delegate.Contract.TotalSupply(&_Abicerc20delegate.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) TotalSupply() (*big.Int, error) {
	return _Abicerc20delegate.Contract.TotalSupply(&_Abicerc20delegate.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicerc20delegate.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateSession) Underlying() (common.Address, error) {
	return _Abicerc20delegate.Contract.Underlying(&_Abicerc20delegate.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_Abicerc20delegate *Abicerc20delegateCallerSession) Underlying() (common.Address, error) {
	return _Abicerc20delegate.Contract.Underlying(&_Abicerc20delegate.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) AcceptAdmin() (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.AcceptAdmin(&_Abicerc20delegate.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.AcceptAdmin(&_Abicerc20delegate.TransactOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) AddReserves(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "_addReserves", addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.AddReserves(&_Abicerc20delegate.TransactOpts, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.AddReserves(&_Abicerc20delegate.TransactOpts, addAmount)
}

// BecomeImplementation is a paid mutator transaction binding the contract method 0x56e67728.
//
// Solidity: function _becomeImplementation(bytes data) returns()
func (_Abicerc20delegate *Abicerc20delegateTransactor) BecomeImplementation(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "_becomeImplementation", data)
}

// BecomeImplementation is a paid mutator transaction binding the contract method 0x56e67728.
//
// Solidity: function _becomeImplementation(bytes data) returns()
func (_Abicerc20delegate *Abicerc20delegateSession) BecomeImplementation(data []byte) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.BecomeImplementation(&_Abicerc20delegate.TransactOpts, data)
}

// BecomeImplementation is a paid mutator transaction binding the contract method 0x56e67728.
//
// Solidity: function _becomeImplementation(bytes data) returns()
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) BecomeImplementation(data []byte) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.BecomeImplementation(&_Abicerc20delegate.TransactOpts, data)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.ReduceReserves(&_Abicerc20delegate.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.ReduceReserves(&_Abicerc20delegate.TransactOpts, reduceAmount)
}

// ResignImplementation is a paid mutator transaction binding the contract method 0x153ab505.
//
// Solidity: function _resignImplementation() returns()
func (_Abicerc20delegate *Abicerc20delegateTransactor) ResignImplementation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "_resignImplementation")
}

// ResignImplementation is a paid mutator transaction binding the contract method 0x153ab505.
//
// Solidity: function _resignImplementation() returns()
func (_Abicerc20delegate *Abicerc20delegateSession) ResignImplementation() (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.ResignImplementation(&_Abicerc20delegate.TransactOpts)
}

// ResignImplementation is a paid mutator transaction binding the contract method 0x153ab505.
//
// Solidity: function _resignImplementation() returns()
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) ResignImplementation() (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.ResignImplementation(&_Abicerc20delegate.TransactOpts)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.SetComptroller(&_Abicerc20delegate.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.SetComptroller(&_Abicerc20delegate.TransactOpts, newComptroller)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.SetInterestRateModel(&_Abicerc20delegate.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.SetInterestRateModel(&_Abicerc20delegate.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.SetPendingAdmin(&_Abicerc20delegate.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.SetPendingAdmin(&_Abicerc20delegate.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.SetReserveFactor(&_Abicerc20delegate.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.SetReserveFactor(&_Abicerc20delegate.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) AccrueInterest() (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.AccrueInterest(&_Abicerc20delegate.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.AccrueInterest(&_Abicerc20delegate.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abicerc20delegate *Abicerc20delegateTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abicerc20delegate *Abicerc20delegateSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Approve(&_Abicerc20delegate.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Approve(&_Abicerc20delegate.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.BalanceOfUnderlying(&_Abicerc20delegate.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.BalanceOfUnderlying(&_Abicerc20delegate.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Borrow(&_Abicerc20delegate.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Borrow(&_Abicerc20delegate.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.BorrowBalanceCurrent(&_Abicerc20delegate.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.BorrowBalanceCurrent(&_Abicerc20delegate.TransactOpts, account)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.ExchangeRateCurrent(&_Abicerc20delegate.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.ExchangeRateCurrent(&_Abicerc20delegate.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x1a31d465.
//
// Solidity: function initialize(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Abicerc20delegate *Abicerc20delegateTransactor) Initialize(opts *bind.TransactOpts, underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "initialize", underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1a31d465.
//
// Solidity: function initialize(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Abicerc20delegate *Abicerc20delegateSession) Initialize(underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Initialize(&_Abicerc20delegate.TransactOpts, underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1a31d465.
//
// Solidity: function initialize(address underlying_, address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) Initialize(underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Initialize(&_Abicerc20delegate.TransactOpts, underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Abicerc20delegate *Abicerc20delegateTransactor) Initialize0(opts *bind.TransactOpts, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "initialize0", comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Abicerc20delegate *Abicerc20delegateSession) Initialize0(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Initialize0(&_Abicerc20delegate.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x99d8c1b4.
//
// Solidity: function initialize(address comptroller_, address interestRateModel_, uint256 initialExchangeRateMantissa_, string name_, string symbol_, uint8 decimals_) returns()
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) Initialize0(comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ uint8) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Initialize0(&_Abicerc20delegate.TransactOpts, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.LiquidateBorrow(&_Abicerc20delegate.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.LiquidateBorrow(&_Abicerc20delegate.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Mint(&_Abicerc20delegate.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Mint(&_Abicerc20delegate.TransactOpts, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Redeem(&_Abicerc20delegate.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Redeem(&_Abicerc20delegate.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.RedeemUnderlying(&_Abicerc20delegate.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.RedeemUnderlying(&_Abicerc20delegate.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.RepayBorrow(&_Abicerc20delegate.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.RepayBorrow(&_Abicerc20delegate.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.RepayBorrowBehalf(&_Abicerc20delegate.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.RepayBorrowBehalf(&_Abicerc20delegate.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Seize(&_Abicerc20delegate.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Seize(&_Abicerc20delegate.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.TotalBorrowsCurrent(&_Abicerc20delegate.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.TotalBorrowsCurrent(&_Abicerc20delegate.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Abicerc20delegate *Abicerc20delegateTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Abicerc20delegate *Abicerc20delegateSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Transfer(&_Abicerc20delegate.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.Transfer(&_Abicerc20delegate.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Abicerc20delegate *Abicerc20delegateTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Abicerc20delegate *Abicerc20delegateSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.TransferFrom(&_Abicerc20delegate.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_Abicerc20delegate *Abicerc20delegateTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicerc20delegate.Contract.TransferFrom(&_Abicerc20delegate.TransactOpts, src, dst, amount)
}

// Abicerc20delegateAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the Abicerc20delegate contract.
type Abicerc20delegateAccrueInterestIterator struct {
	Event *Abicerc20delegateAccrueInterest // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateAccrueInterest)
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
		it.Event = new(Abicerc20delegateAccrueInterest)
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
func (it *Abicerc20delegateAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateAccrueInterest represents a AccrueInterest event raised by the Abicerc20delegate contract.
type Abicerc20delegateAccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*Abicerc20delegateAccrueInterestIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateAccrueInterestIterator{contract: _Abicerc20delegate.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateAccrueInterest)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseAccrueInterest(log types.Log) (*Abicerc20delegateAccrueInterest, error) {
	event := new(Abicerc20delegateAccrueInterest)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Abicerc20delegate contract.
type Abicerc20delegateApprovalIterator struct {
	Event *Abicerc20delegateApproval // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateApproval)
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
		it.Event = new(Abicerc20delegateApproval)
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
func (it *Abicerc20delegateApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateApproval represents a Approval event raised by the Abicerc20delegate contract.
type Abicerc20delegateApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Abicerc20delegateApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateApprovalIterator{contract: _Abicerc20delegate.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateApproval)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseApproval(log types.Log) (*Abicerc20delegateApproval, error) {
	event := new(Abicerc20delegateApproval)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the Abicerc20delegate contract.
type Abicerc20delegateBorrowIterator struct {
	Event *Abicerc20delegateBorrow // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateBorrow)
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
		it.Event = new(Abicerc20delegateBorrow)
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
func (it *Abicerc20delegateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateBorrow represents a Borrow event raised by the Abicerc20delegate contract.
type Abicerc20delegateBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterBorrow(opts *bind.FilterOpts) (*Abicerc20delegateBorrowIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateBorrowIterator{contract: _Abicerc20delegate.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateBorrow) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateBorrow)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "Borrow", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseBorrow(log types.Log) (*Abicerc20delegateBorrow, error) {
	event := new(Abicerc20delegateBorrow)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Abicerc20delegate contract.
type Abicerc20delegateFailureIterator struct {
	Event *Abicerc20delegateFailure // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateFailure)
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
		it.Event = new(Abicerc20delegateFailure)
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
func (it *Abicerc20delegateFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateFailure represents a Failure event raised by the Abicerc20delegate contract.
type Abicerc20delegateFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterFailure(opts *bind.FilterOpts) (*Abicerc20delegateFailureIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateFailureIterator{contract: _Abicerc20delegate.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateFailure) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateFailure)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "Failure", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseFailure(log types.Log) (*Abicerc20delegateFailure, error) {
	event := new(Abicerc20delegateFailure)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the Abicerc20delegate contract.
type Abicerc20delegateLiquidateBorrowIterator struct {
	Event *Abicerc20delegateLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateLiquidateBorrow)
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
		it.Event = new(Abicerc20delegateLiquidateBorrow)
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
func (it *Abicerc20delegateLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateLiquidateBorrow represents a LiquidateBorrow event raised by the Abicerc20delegate contract.
type Abicerc20delegateLiquidateBorrow struct {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*Abicerc20delegateLiquidateBorrowIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateLiquidateBorrowIterator{contract: _Abicerc20delegate.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateLiquidateBorrow)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseLiquidateBorrow(log types.Log) (*Abicerc20delegateLiquidateBorrow, error) {
	event := new(Abicerc20delegateLiquidateBorrow)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Abicerc20delegate contract.
type Abicerc20delegateMintIterator struct {
	Event *Abicerc20delegateMint // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateMint)
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
		it.Event = new(Abicerc20delegateMint)
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
func (it *Abicerc20delegateMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateMint represents a Mint event raised by the Abicerc20delegate contract.
type Abicerc20delegateMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterMint(opts *bind.FilterOpts) (*Abicerc20delegateMintIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateMintIterator{contract: _Abicerc20delegate.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateMint) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateMint)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "Mint", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseMint(log types.Log) (*Abicerc20delegateMint, error) {
	event := new(Abicerc20delegateMint)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the Abicerc20delegate contract.
type Abicerc20delegateNewAdminIterator struct {
	Event *Abicerc20delegateNewAdmin // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateNewAdmin)
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
		it.Event = new(Abicerc20delegateNewAdmin)
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
func (it *Abicerc20delegateNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateNewAdmin represents a NewAdmin event raised by the Abicerc20delegate contract.
type Abicerc20delegateNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*Abicerc20delegateNewAdminIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateNewAdminIterator{contract: _Abicerc20delegate.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateNewAdmin) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateNewAdmin)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseNewAdmin(log types.Log) (*Abicerc20delegateNewAdmin, error) {
	event := new(Abicerc20delegateNewAdmin)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the Abicerc20delegate contract.
type Abicerc20delegateNewComptrollerIterator struct {
	Event *Abicerc20delegateNewComptroller // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateNewComptroller)
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
		it.Event = new(Abicerc20delegateNewComptroller)
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
func (it *Abicerc20delegateNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateNewComptroller represents a NewComptroller event raised by the Abicerc20delegate contract.
type Abicerc20delegateNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*Abicerc20delegateNewComptrollerIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateNewComptrollerIterator{contract: _Abicerc20delegate.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateNewComptroller) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateNewComptroller)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseNewComptroller(log types.Log) (*Abicerc20delegateNewComptroller, error) {
	event := new(Abicerc20delegateNewComptroller)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateNewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the Abicerc20delegate contract.
type Abicerc20delegateNewMarketInterestRateModelIterator struct {
	Event *Abicerc20delegateNewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateNewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateNewMarketInterestRateModel)
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
		it.Event = new(Abicerc20delegateNewMarketInterestRateModel)
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
func (it *Abicerc20delegateNewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateNewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateNewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the Abicerc20delegate contract.
type Abicerc20delegateNewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*Abicerc20delegateNewMarketInterestRateModelIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateNewMarketInterestRateModelIterator{contract: _Abicerc20delegate.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateNewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateNewMarketInterestRateModel)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseNewMarketInterestRateModel(log types.Log) (*Abicerc20delegateNewMarketInterestRateModel, error) {
	event := new(Abicerc20delegateNewMarketInterestRateModel)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the Abicerc20delegate contract.
type Abicerc20delegateNewPendingAdminIterator struct {
	Event *Abicerc20delegateNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateNewPendingAdmin)
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
		it.Event = new(Abicerc20delegateNewPendingAdmin)
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
func (it *Abicerc20delegateNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateNewPendingAdmin represents a NewPendingAdmin event raised by the Abicerc20delegate contract.
type Abicerc20delegateNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*Abicerc20delegateNewPendingAdminIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateNewPendingAdminIterator{contract: _Abicerc20delegate.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateNewPendingAdmin)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseNewPendingAdmin(log types.Log) (*Abicerc20delegateNewPendingAdmin, error) {
	event := new(Abicerc20delegateNewPendingAdmin)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateNewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the Abicerc20delegate contract.
type Abicerc20delegateNewReserveFactorIterator struct {
	Event *Abicerc20delegateNewReserveFactor // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateNewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateNewReserveFactor)
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
		it.Event = new(Abicerc20delegateNewReserveFactor)
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
func (it *Abicerc20delegateNewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateNewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateNewReserveFactor represents a NewReserveFactor event raised by the Abicerc20delegate contract.
type Abicerc20delegateNewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*Abicerc20delegateNewReserveFactorIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateNewReserveFactorIterator{contract: _Abicerc20delegate.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateNewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateNewReserveFactor)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseNewReserveFactor(log types.Log) (*Abicerc20delegateNewReserveFactor, error) {
	event := new(Abicerc20delegateNewReserveFactor)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the Abicerc20delegate contract.
type Abicerc20delegateRedeemIterator struct {
	Event *Abicerc20delegateRedeem // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateRedeem)
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
		it.Event = new(Abicerc20delegateRedeem)
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
func (it *Abicerc20delegateRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateRedeem represents a Redeem event raised by the Abicerc20delegate contract.
type Abicerc20delegateRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterRedeem(opts *bind.FilterOpts) (*Abicerc20delegateRedeemIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateRedeemIterator{contract: _Abicerc20delegate.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateRedeem) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateRedeem)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "Redeem", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseRedeem(log types.Log) (*Abicerc20delegateRedeem, error) {
	event := new(Abicerc20delegateRedeem)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the Abicerc20delegate contract.
type Abicerc20delegateRepayBorrowIterator struct {
	Event *Abicerc20delegateRepayBorrow // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateRepayBorrow)
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
		it.Event = new(Abicerc20delegateRepayBorrow)
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
func (it *Abicerc20delegateRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateRepayBorrow represents a RepayBorrow event raised by the Abicerc20delegate contract.
type Abicerc20delegateRepayBorrow struct {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*Abicerc20delegateRepayBorrowIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateRepayBorrowIterator{contract: _Abicerc20delegate.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateRepayBorrow)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseRepayBorrow(log types.Log) (*Abicerc20delegateRepayBorrow, error) {
	event := new(Abicerc20delegateRepayBorrow)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the Abicerc20delegate contract.
type Abicerc20delegateReservesAddedIterator struct {
	Event *Abicerc20delegateReservesAdded // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateReservesAdded)
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
		it.Event = new(Abicerc20delegateReservesAdded)
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
func (it *Abicerc20delegateReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateReservesAdded represents a ReservesAdded event raised by the Abicerc20delegate contract.
type Abicerc20delegateReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterReservesAdded(opts *bind.FilterOpts) (*Abicerc20delegateReservesAddedIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateReservesAddedIterator{contract: _Abicerc20delegate.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateReservesAdded) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateReservesAdded)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseReservesAdded(log types.Log) (*Abicerc20delegateReservesAdded, error) {
	event := new(Abicerc20delegateReservesAdded)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the Abicerc20delegate contract.
type Abicerc20delegateReservesReducedIterator struct {
	Event *Abicerc20delegateReservesReduced // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateReservesReduced)
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
		it.Event = new(Abicerc20delegateReservesReduced)
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
func (it *Abicerc20delegateReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateReservesReduced represents a ReservesReduced event raised by the Abicerc20delegate contract.
type Abicerc20delegateReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*Abicerc20delegateReservesReducedIterator, error) {

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateReservesReducedIterator{contract: _Abicerc20delegate.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateReservesReduced) (event.Subscription, error) {

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateReservesReduced)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseReservesReduced(log types.Log) (*Abicerc20delegateReservesReduced, error) {
	event := new(Abicerc20delegateReservesReduced)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Abicerc20delegateTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Abicerc20delegate contract.
type Abicerc20delegateTransferIterator struct {
	Event *Abicerc20delegateTransfer // Event containing the contract specifics and raw log

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
func (it *Abicerc20delegateTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Abicerc20delegateTransfer)
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
		it.Event = new(Abicerc20delegateTransfer)
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
func (it *Abicerc20delegateTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Abicerc20delegateTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Abicerc20delegateTransfer represents a Transfer event raised by the Abicerc20delegate contract.
type Abicerc20delegateTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Abicerc20delegate *Abicerc20delegateFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Abicerc20delegateTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Abicerc20delegate.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Abicerc20delegateTransferIterator{contract: _Abicerc20delegate.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Abicerc20delegate *Abicerc20delegateFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Abicerc20delegateTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Abicerc20delegate.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Abicerc20delegateTransfer)
				if err := _Abicerc20delegate.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Abicerc20delegate *Abicerc20delegateFilterer) ParseTransfer(log types.Log) (*Abicerc20delegateTransfer, error) {
	event := new(Abicerc20delegateTransfer)
	if err := _Abicerc20delegate.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
