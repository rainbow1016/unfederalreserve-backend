// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abicomptroller

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

// AbicomptrollerABI is the input ABI used to generate the binding from.
const AbicomptrollerABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CompGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"CompSpeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"ContributorCompSpeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedBorrowerComp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compSupplyIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedSupplierComp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"MarketListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowCap\",\"type\":\"uint256\"}],\"name\":\"NewBorrowCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBorrowCapGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"NewBorrowCapGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRewardTokenAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRewardTokenAddr\",\"type\":\"address\"}],\"name\":\"NewCompAddress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_grantComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"_setBorrowCapGuardian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setBorrowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRewardTokenAddr\",\"type\":\"address\"}],\"name\":\"_setCompAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"compSpeed\",\"type\":\"uint256\"}],\"name\":\"_setCompSpeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"compSpeed\",\"type\":\"uint256\"}],\"name\":\"_setContributorCompSpeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBorrowCaps\",\"type\":\"uint256[]\"}],\"name\":\"_setMarketBorrowCaps\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setMintPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"_setPauseGuardian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setSeizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setTransferPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"_supportMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"claimRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compContributorSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCompAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastContributorBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isComped\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rewardTokenAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"}],\"name\":\"updateContributorRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AbicomptrollerBin is the compiled bytecode used for deploying new contracts.
var AbicomptrollerBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b0319163317905561593680620000336000396000f3fe608060405234801561001057600080fd5b506004361061047f5760003560e01c80636d35bf9111610257578063bb82aa5e11610146578063dcfbc0c7116100c3578063e9af029211610087578063e9af029214611437578063eabe7d911461145d578063ede4edd014611493578063ef5cfb8c146114b9578063f851a440146114df5761047f565b8063dcfbc0c7146113eb578063e4028eee146113f3578063e6653f3d1461141f578063e875544614611427578063e89dec6d1461142f5761047f565b8063ca0af0431161010a578063ca0af043146112ef578063cc7ebdc41461131d578063d02f735114611343578063da3d454c14611389578063dce15449146113bf5761047f565b8063bb82aa5e14611195578063bdcdc2581461119d578063bea6b8b8146111d9578063c2998238146111ff578063c488847b146112a05761047f565b806394b2294b116101d4578063aa90075411610198578063aa900754146110d9578063abfceffc146110e1578063ac0b0bb714611157578063b0772d0b1461115f578063b21be7fd146111675761047f565b806394b2294b14611059578063986ab838146110615780639d1b5a0a14611087578063a76b3fda1461108f578063a7f0e231146110b55761047f565b806387f763031161021b57806387f7630314610f965780638c57804e14610f9e5780638e8f294b14610fc45780638ebf63641461100c578063929fe9a11461102b5761047f565b80636d35bf9114610ed6578063731f0c2b14610f1c578063741b252514610f425780637dc0d1d014610f685780638435be4614610f705761047f565b806347ef3b3b11610373578063598ee1cb116102f0578063607ef6c1116102b4578063607ef6c114610c3c5780636810dfa614610cfa5780636a56947e14610e265780636b79c38d14610e625780636d154ea514610eb05761047f565b8063598ee1cb14610b485780635c77860514610b745780635ec88c7914610baa5780635f5af1aa14610bd05780635fc7e71e14610bf65761047f565b80634fd42e17116103375780634fd42e171461098057806351dff9891461099d57806352d84d1e146109d957806354eb76fa146109f657806355ee1fe114610b225761047f565b806347ef3b3b146108765780634a584432146108c25780634ada90af146108e85780634e79238f146108f05780634ef4c3e11461094a5761047f565b806326782247116104015780633bcf7ec1116103c55780633bcf7ec1146107d05780633c94786f146107fe57806341c728b91461080657806342cbb15c14610842578063434caf251461084a5761047f565b8063267822471461073a57806327efe3cb146107425780632d70db781461076e578063317b0b771461078d578063391957d7146107aa5761047f565b80631ededc91116104485780631ededc91146105df5780632026ffa31461062157806321af4569146106d257806324008a62146106f657806324a3d622146107325761047f565b80627e3dd21461048457806318c882a5146104a05780631c3db2e0146104ce5780631d504dc6146105815780631d7b33d7146105a7575b600080fd5b61048c6114e7565b604080519115158252519081900360200190f35b61048c600480360360408110156104b657600080fd5b506001600160a01b03813516906020013515156114ec565b61057f600480360360408110156104e457600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561050e57600080fd5b82018360208201111561052057600080fd5b803590602001918460208302840111600160201b8311171561054157600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061168c945050505050565b005b61057f6004803603602081101561059757600080fd5b50356001600160a01b03166116ee565b6105cd600480360360208110156105bd57600080fd5b50356001600160a01b031661184d565b60408051918252519081900360200190f35b61057f600480360360a08110156105f557600080fd5b506001600160a01b0381358116916020810135821691604082013516906060810135906080013561185f565b61057f6004803603604081101561063757600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561066157600080fd5b82018360208201111561067357600080fd5b803590602001918460208302840111600160201b8311171561069457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611866945050505050565b6106da611874565b604080516001600160a01b039092168252519081900360200190f35b6105cd6004803603608081101561070c57600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135611883565b6106da61194a565b6106da611959565b61057f6004803603604081101561075857600080fd5b506001600160a01b038135169060200135611968565b61048c6004803603602081101561078457600080fd5b50351515611a6b565b6105cd600480360360208110156107a357600080fd5b5035611ba5565b61057f600480360360208110156107c057600080fd5b50356001600160a01b0316611c52565b61048c600480360360408110156107e657600080fd5b506001600160a01b0381351690602001351515611cfe565b61048c611e99565b61057f6004803603608081101561081c57600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135611ea9565b6105cd611eaf565b61057f6004803603604081101561086057600080fd5b506001600160a01b038135169060200135611eb4565b61057f600480360360c081101561088c57600080fd5b506001600160a01b0381358116916020810135821691604082013581169160608101359091169060808101359060a00135611f17565b6105cd600480360360208110156108d857600080fd5b50356001600160a01b0316611f1f565b6105cd611f31565b61092c6004803603608081101561090657600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135611f37565b60408051938452602084019290925282820152519081900360600190f35b6105cd6004803603606081101561096057600080fd5b506001600160a01b03813581169160208101359091169060400135611f71565b6105cd6004803603602081101561099657600080fd5b5035612017565b61057f600480360360808110156109b357600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135612087565b6106da600480360360208110156109ef57600080fd5b50356120db565b61057f60048036036080811015610a0c57600080fd5b810190602081018135600160201b811115610a2657600080fd5b820183602082011115610a3857600080fd5b803590602001918460208302840111600160201b83111715610a5957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610aa857600080fd5b820183602082011115610aba57600080fd5b803590602001918460208302840111600160201b83111715610adb57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550505050803515159150602001351515612102565b6105cd60048036036020811015610b3857600080fd5b50356001600160a01b031661210e565b61057f60048036036040811015610b5e57600080fd5b506001600160a01b038135169060200135612193565b61057f60048036036060811015610b8a57600080fd5b506001600160a01b0381358116916020810135909116906040013561228e565b61092c60048036036020811015610bc057600080fd5b50356001600160a01b0316612293565b6105cd60048036036020811015610be657600080fd5b50356001600160a01b03166122c8565b6105cd600480360360a0811015610c0c57600080fd5b506001600160a01b038135811691602081013582169160408201358116916060810135909116906080013561234c565b61057f60048036036040811015610c5257600080fd5b810190602081018135600160201b811115610c6c57600080fd5b820183602082011115610c7e57600080fd5b803590602001918460208302840111600160201b83111715610c9f57600080fd5b919390929091602081019035600160201b811115610cbc57600080fd5b820183602082011115610cce57600080fd5b803590602001918460208302840111600160201b83111715610cef57600080fd5b5090925090506124b1565b61057f60048036036080811015610d1057600080fd5b810190602081018135600160201b811115610d2a57600080fd5b820183602082011115610d3c57600080fd5b803590602001918460208302840111600160201b83111715610d5d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610dac57600080fd5b820183602082011115610dbe57600080fd5b803590602001918460208302840111600160201b83111715610ddf57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550505050803515159150602001351515612641565b61057f60048036036080811015610e3c57600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135611ea9565b610e8860048036036020811015610e7857600080fd5b50356001600160a01b03166128ca565b604080516001600160e01b03909316835263ffffffff90911660208301528051918290030190f35b61048c60048036036020811015610ec657600080fd5b50356001600160a01b03166128f4565b61057f600480360360a0811015610eec57600080fd5b506001600160a01b038135811691602081013582169160408201358116916060810135909116906080013561185f565b61048c60048036036020811015610f3257600080fd5b50356001600160a01b0316612909565b61057f60048036036020811015610f5857600080fd5b50356001600160a01b031661291e565b6106da6129e1565b6105cd60048036036020811015610f8657600080fd5b50356001600160a01b03166129f0565b61048c612ad0565b610e8860048036036020811015610fb457600080fd5b50356001600160a01b0316612ae0565b610fea60048036036020811015610fda57600080fd5b50356001600160a01b0316612b0a565b6040805193151584526020840192909252151582820152519081900360600190f35b61048c6004803603602081101561102257600080fd5b50351515612b30565b61048c6004803603604081101561104157600080fd5b506001600160a01b0381358116916020013516612c69565b6105cd612c9c565b6105cd6004803603602081101561107757600080fd5b50356001600160a01b0316612ca2565b6106da612cb4565b6105cd600480360360208110156110a557600080fd5b50356001600160a01b0316612cc3565b6110bd612e20565b604080516001600160e01b039092168252519081900360200190f35b6105cd612e33565b611107600480360360208110156110f757600080fd5b50356001600160a01b0316612e39565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561114357818101518382015260200161112b565b505050509050019250505060405180910390f35b61048c612ec2565b611107612ed2565b6105cd6004803603604081101561117d57600080fd5b506001600160a01b0381358116916020013516612f34565b6106da612f51565b6105cd600480360360808110156111b357600080fd5b506001600160a01b03813581169160208101358216916040820135169060600135612f60565b6105cd600480360360208110156111ef57600080fd5b50356001600160a01b0316612ff0565b6111076004803603602081101561121557600080fd5b810190602081018135600160201b81111561122f57600080fd5b82018360208201111561124157600080fd5b803590602001918460208302840111600160201b8311171561126257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550613002945050505050565b6112d6600480360360608110156112b657600080fd5b506001600160a01b03813581169160208101359091169060400135613099565b6040805192835260208301919091528051918290030190f35b6105cd6004803603604081101561130557600080fd5b506001600160a01b03813581169160200135166132c1565b6105cd6004803603602081101561133357600080fd5b50356001600160a01b03166132de565b6105cd600480360360a081101561135957600080fd5b506001600160a01b03813581169160208101358216916040820135811691606081013590911690608001356132f0565b6105cd6004803603606081101561139f57600080fd5b506001600160a01b038135811691602081013590911690604001356134a4565b6106da600480360360408110156113d557600080fd5b506001600160a01b03813516906020013561387e565b6106da6138b3565b6105cd6004803603604081101561140957600080fd5b506001600160a01b0381351690602001356138c2565b61048c613a72565b6105cd613a82565b6106da613a88565b61057f6004803603602081101561144d57600080fd5b50356001600160a01b0316613a97565b6105cd6004803603606081101561147357600080fd5b506001600160a01b03813581169160208101359091169060400135613afb565b6105cd600480360360208110156114a957600080fd5b50356001600160a01b0316613b36565b61057f600480360360208110156114cf57600080fd5b50356001600160a01b0316613e49565b6106da613e52565b600181565b6001600160a01b03821660009081526009602052604081205460ff166115435760405162461bcd60e51b81526004018080602001828103825260288152602001806157ec6028913960400191505060405180910390fd5b600a546001600160a01b031633148061156657506000546001600160a01b031633145b6115a15760405162461bcd60e51b81526004018080602001828103825260278152602001806158146027913960400191505060405180910390fd5b6000546001600160a01b03163314806115bc57506001821515145b611606576040805162461bcd60e51b81526020600482015260166024820152756f6e6c792061646d696e2063616e20756e706175736560501b604482015290519081900360640190fd5b6001600160a01b0383166000818152600c6020908152604091829020805486151560ff199091168117909155825193845283830152606090830181905260069083015265426f72726f7760d01b6080830152517f71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b09181900360a00190a150805b92915050565b6040805160018082528183019092526060916020808301908038833901905050905082816000815181106116bc57fe5b60200260200101906001600160a01b031690816001600160a01b0316815250506116e98183600180612641565b505050565b806001600160a01b031663f851a4406040518163ffffffff1660e01b815260040160206040518083038186803b15801561172757600080fd5b505afa15801561173b573d6000803e3d6000fd5b505050506040513d602081101561175157600080fd5b50516001600160a01b031633146117995760405162461bcd60e51b81526004018080602001828103825260278152602001806158db6027913960400191505060405180910390fd5b806001600160a01b031663c1e803346040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156117d457600080fd5b505af11580156117e8573d6000803e3d6000fd5b505050506040513d60208110156117fe57600080fd5b50511561184a576040805162461bcd60e51b815260206004820152601560248201527418da185b99d9481b9bdd08185d5d1a1bdc9a5e9959605a1b604482015290519081900360640190fd5b50565b600f6020526000908152604090205481565b5050505050565b611870828261168c565b5050565b6015546001600160a01b031681565b6001600160a01b03841660009081526009602052604081205460ff166118ab57506009611942565b6118b361572c565b6040518060200160405280876001600160a01b031663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156118f757600080fd5b505afa15801561190b573d6000803e3d6000fd5b505050506040513d602081101561192157600080fd5b5051905290506119318682613e61565b61193c8685836140e9565b60009150505b949350505050565b600a546001600160a01b031681565b6001546001600160a01b031681565b611970614280565b6119c1576040805162461bcd60e51b815260206004820152601960248201527f6f6e6c792061646d696e2063616e206772616e7420636f6d7000000000000000604482015290519081900360640190fd5b60006119cd83836142a9565b90508015611a22576040805162461bcd60e51b815260206004820152601b60248201527f696e73756666696369656e7420636f6d7020666f72206772616e740000000000604482015290519081900360640190fd5b604080516001600160a01b03851681526020810184905281517f98b2f82a3a07f223a0be64b3d0f47711c64dccd1feafb94aa28156b38cd9695c929181900390910190a1505050565b600a546000906001600160a01b0316331480611a9157506000546001600160a01b031633145b611acc5760405162461bcd60e51b81526004018080602001828103825260278152602001806158146027913960400191505060405180910390fd5b6000546001600160a01b0316331480611ae757506001821515145b611b31576040805162461bcd60e51b81526020600482015260166024820152756f6e6c792061646d696e2063616e20756e706175736560501b604482015290519081900360640190fd5b600a8054831515600160b81b810260ff60b81b1990921691909117909155604080516020810192909252808252600582820152645365697a6560d81b6060830152517fef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de09181900360800190a150805b919050565b600080546001600160a01b03163314611c05576040805162461bcd60e51b815260206004820152601f60248201527f6f6e6c792061646d696e2063616e2073657420636c6f736520666163746f7200604482015290519081900360640190fd5b6005805490839055604080518281526020810185905281517f3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9929181900390910190a160005b9392505050565b6000546001600160a01b03163314611c9b5760405162461bcd60e51b815260040180806020018281038252602681526020018061583b6026913960400191505060405180910390fd5b601580546001600160a01b038381166001600160a01b0319831681179093556040805191909216808252602082019390935281517feda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29929181900390910190a15050565b6001600160a01b03821660009081526009602052604081205460ff16611d555760405162461bcd60e51b81526004018080602001828103825260288152602001806157ec6028913960400191505060405180910390fd5b600a546001600160a01b0316331480611d7857506000546001600160a01b031633145b611db35760405162461bcd60e51b81526004018080602001828103825260278152602001806158146027913960400191505060405180910390fd5b6000546001600160a01b0316331480611dce57506001821515145b611e18576040805162461bcd60e51b81526020600482015260166024820152756f6e6c792061646d696e2063616e20756e706175736560501b604482015290519081900360640190fd5b6001600160a01b0383166000818152600b6020908152604091829020805486151560ff199091168117909155825193845283830152606090830181905260049083015263135a5b9d60e21b6080830152517f71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b09181900360a00190a150919050565b600a54600160a01b900460ff1681565b50505050565b435b90565b611ebc614280565b611f0d576040805162461bcd60e51b815260206004820152601d60248201527f6f6e6c792061646d696e2063616e2073657420636f6d70207370656564000000604482015290519081900360640190fd5b61187082826143e3565b505050505050565b60166020526000908152604090205481565b60065481565b600080600080600080611f4c8a8a8a8a614701565b925092509250826011811115611f5e57fe5b95509093509150505b9450945094915050565b6001600160a01b0383166000908152600b602052604081205460ff1615611fd0576040805162461bcd60e51b815260206004820152600e60248201526d1b5a5b9d081a5cc81c185d5cd95960921b604482015290519081900360640190fd5b6001600160a01b03841660009081526009602052604090205460ff16611ffa5760095b9050611c4b565b61200384614a39565b61200d8484614cb7565b6000949350505050565b600080546001600160a01b0316331461203d576120366001600b614e96565b9050611ba0565b6006805490839055604080518281526020810185905281517faeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316929181900390910190a16000611c4b565b801580156120955750600082115b15611ea9576040805162461bcd60e51b815260206004820152601160248201527072656465656d546f6b656e73207a65726f60781b604482015290519081900360640190fd5b600d81815481106120e857fe5b6000918252602090912001546001600160a01b0316905081565b611ea984848484612641565b600080546001600160a01b0316331461212d5761203660016010614e96565b600480546001600160a01b038481166001600160a01b0319831681179093556040805191909216808252602082019390935281517fd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22929181900390910190a16000611c4b565b61219b614280565b6121ec576040805162461bcd60e51b815260206004820152601d60248201527f6f6e6c792061646d696e2063616e2073657420636f6d70207370656564000000604482015290519081900360640190fd5b6121f58261291e565b80612218576001600160a01b03821660009081526018602052604081205561223a565b612220611eaf565b6001600160a01b0383166000908152601860205260409020555b6001600160a01b038216600081815260176020908152604091829020849055815184815291517f386537fa92edc3319af95f1f904dcf1900021e4f3f4e08169a577a09076e66b39281900390910190a25050565b6116e9565b6000806000806000806122aa876000806000614701565b9250925092508260118111156122bc57fe5b97919650945092505050565b600080546001600160a01b031633146122e75761203660016013614e96565b600a80546001600160a01b038481166001600160a01b0319831617928390556040805192821680845293909116602083015280517f0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e9281900390910190a16000611c4b565b6001600160a01b03851660009081526009602052604081205460ff16158061238d57506001600160a01b03851660009081526009602052604090205460ff16155b1561239c5760095b90506124a8565b6000806123a885614efc565b919350909150600090508260118111156123be57fe5b146123d8578160118111156123cf57fe5b925050506124a8565b806123e45760036123cf565b6000886001600160a01b03166395dd9193876040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b15801561243c57600080fd5b505afa158015612450573d6000803e3d6000fd5b505050506040513d602081101561246657600080fd5b5051604080516020810190915260055481529091506000906124889083614f1c565b90508086111561249f5760119450505050506124a8565b60009450505050505b95945050505050565b6000546001600160a01b03163314806124d457506015546001600160a01b031633145b61250f5760405162461bcd60e51b81526004018080602001828103825260358152602001806158616035913960400191505060405180910390fd5b8281811580159061251f57508082145b612560576040805162461bcd60e51b815260206004820152600d60248201526c1a5b9d985b1a59081a5b9c1d5d609a1b604482015290519081900360640190fd5b60005b828110156126385784848281811061257757fe5b905060200201356016600089898581811061258e57fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b03168152602001908152602001600020819055508686828181106125ce57fe5b905060200201356001600160a01b03166001600160a01b03167f6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f686868481811061261457fe5b905060200201356040518082815260200191505060405180910390a2600101612563565b50505050505050565b60005b835181101561185f57600084828151811061265b57fe5b6020908102919091018101516001600160a01b0381166000908152600990925260409091205490915060ff166126d0576040805162461bcd60e51b81526020600482015260156024820152741b585c9ad95d081b5d5cdd081899481b1a5cdd1959605a1b604482015290519081900360640190fd5b6001841515141561281f576126e361572c565b6040518060200160405280836001600160a01b031663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561272757600080fd5b505afa15801561273b573d6000803e3d6000fd5b505050506040513d602081101561275157600080fd5b5051905290506127618282613e61565b60005b875181101561281c5761278b8389838151811061277d57fe5b6020026020010151846140e9565b6127e088828151811061279a57fe5b6020026020010151601460008b85815181106127b257fe5b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020546142a9565b601460008a84815181106127f057fe5b6020908102919091018101516001600160a01b0316825281019190915260400160002055600101612764565b50505b600183151514156128c15761283381614a39565b60005b86518110156128bf5761285c8288838151811061284f57fe5b6020026020010151614cb7565b61288387828151811061286b57fe5b6020026020010151601460008a85815181106127b257fe5b6014600089848151811061289357fe5b6020908102919091018101516001600160a01b0316825281019190915260400160002055600101612836565b505b50600101612644565b6010602052600090815260409020546001600160e01b03811690600160e01b900463ffffffff1682565b600c6020526000908152604090205460ff1681565b600b6020526000908152604090205460ff1681565b6001600160a01b03811660009081526017602052604081205490612940611eaf565b6001600160a01b03841660009081526018602052604081205491925090612968908390614f3b565b905060008111801561297a5750600083115b15611ea957600061298b8285614f75565b6001600160a01b038616600090815260146020526040812054919250906129b29083614fb7565b6001600160a01b0387166000908152601460209081526040808320939093556018905220849055505050505050565b6004546001600160a01b031681565b60006129fb82614fed565b612a4c576040805162461bcd60e51b815260206004820181905260248201527f6f6e6c7920636f6e7472616374206164647265737320697320616c6c6f776564604482015290519081900360640190fd5b6000546001600160a01b03163314612a6a5761203660016014614e96565b601980546001600160a01b038481166001600160a01b0319831681179093556040805191909216808252602082019390935281517fdde58acf741e6f5e57788ee0f64a68576e5eadf169e78251f32824653dd308bf929181900390910190a16000611c4b565b600a54600160b01b900460ff1681565b6011602052600090815260409020546001600160e01b03811690600160e01b900463ffffffff1682565b60096020526000908152604090208054600182015460039092015460ff91821692911683565b600a546000906001600160a01b0316331480612b5657506000546001600160a01b031633145b612b915760405162461bcd60e51b81526004018080602001828103825260278152602001806158146027913960400191505060405180910390fd5b6000546001600160a01b0316331480612bac57506001821515145b612bf6576040805162461bcd60e51b81526020600482015260166024820152756f6e6c792061646d696e2063616e20756e706175736560501b604482015290519081900360640190fd5b600a8054831515600160b01b810260ff60b01b1990921691909117909155604080516020810192909252808252600882820152672a3930b739b332b960c11b6060830152517fef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de09181900360800190a15090565b6001600160a01b038082166000908152600960209081526040808320938616835260029093019052205460ff1692915050565b60075481565b60176020526000908152604090205481565b6019546001600160a01b031690565b600080546001600160a01b03163314612ce25761203660016012614e96565b6001600160a01b03821660009081526009602052604090205460ff1615612d0f57612036600a6011614e96565b816001600160a01b031663fe9c44ae6040518163ffffffff1660e01b815260040160206040518083038186803b158015612d4857600080fd5b505afa158015612d5c573d6000803e3d6000fd5b505050506040513d6020811015612d7257600080fd5b5050604080516060810182526001808252600060208381018281528486018381526001600160a01b03891684526009909252949091209251835490151560ff19918216178455935191830191909155516003909101805491151591909216179055612ddc82614ff3565b604080516001600160a01b038416815290517fcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f9181900360200190a1600092915050565b6ec097ce7bc90715b34b9f100000000081565b600e5481565b60608060086000846001600160a01b03166001600160a01b03168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015612eb557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612e97575b5093979650505050505050565b600a54600160b81b900460ff1681565b6060600d805480602002602001604051908101604052809291908181526020018280548015612f2a57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612f0c575b5050505050905090565b601260209081526000928352604080842090915290825290205481565b6002546001600160a01b031681565b600a54600090600160b01b900460ff1615612fb7576040805162461bcd60e51b81526020600482015260126024820152711d1c985b9cd9995c881a5cc81c185d5cd95960721b604482015290519081900360640190fd5b6000612fc48686856150d1565b90508015612fd3579050611942565b612fdc86614a39565b612fe68686614cb7565b61193c8685614cb7565b60186020526000908152604090205481565b6060600082519050606081604051908082528060200260200182016040528015613036578160200160208202803883390190505b50905060005b8281101561309157600085828151811061305257fe5b60200260200101519050613066813361517d565b601181111561307157fe5b83838151811061307d57fe5b60209081029190910101525060010161303c565b509392505050565b600480546040805163fc57d4df60e01b81526001600160a01b038781169482019490945290516000938493849391169163fc57d4df91602480820192602092909190829003018186803b1580156130ef57600080fd5b505afa158015613103573d6000803e3d6000fd5b505050506040513d602081101561311957600080fd5b5051600480546040805163fc57d4df60e01b81526001600160a01b038a8116948201949094529051939450600093929091169163fc57d4df91602480820192602092909190829003018186803b15801561317257600080fd5b505afa158015613186573d6000803e3d6000fd5b505050506040513d602081101561319c57600080fd5b505190508115806131ab575080155b156131c057600d9350600092506132b9915050565b6000866001600160a01b031663182df0f56040518163ffffffff1660e01b815260040160206040518083038186803b1580156131fb57600080fd5b505afa15801561320f573d6000803e3d6000fd5b505050506040513d602081101561322557600080fd5b50519050600061323361572c565b61323b61572c565b61324361572c565b61326b604051806020016040528060065481525060405180602001604052808a815250615273565b9250613293604051806020016040528088815250604051806020016040528088815250615273565b915061329f83836152b2565b90506132ab818b614f1c565b600099509750505050505050505b935093915050565b601360209081526000928352604080842090915290825290205481565b60146020526000908152604090205481565b600a54600090600160b81b900460ff1615613344576040805162461bcd60e51b815260206004820152600f60248201526e1cd95a5e99481a5cc81c185d5cd959608a1b604482015290519081900360640190fd5b6001600160a01b03861660009081526009602052604090205460ff16158061338557506001600160a01b03851660009081526009602052604090205460ff16155b15613391576009612395565b846001600160a01b0316635fe3b5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156133ca57600080fd5b505afa1580156133de573d6000803e3d6000fd5b505050506040513d60208110156133f457600080fd5b505160408051635fe3b56760e01b815290516001600160a01b0392831692891691635fe3b567916004808301926020929190829003018186803b15801561343a57600080fd5b505afa15801561344e573d6000803e3d6000fd5b505050506040513d602081101561346457600080fd5b50516001600160a01b03161461347b576002612395565b61348486614a39565b61348e8684614cb7565b6134988685614cb7565b60009695505050505050565b6001600160a01b0383166000908152600c602052604081205460ff1615613505576040805162461bcd60e51b815260206004820152601060248201526f189bdc9c9bddc81a5cc81c185d5cd95960821b604482015290519081900360640190fd5b6001600160a01b03841660009081526009602052604090205460ff1661352c576009611ff3565b6001600160a01b038085166000908152600960209081526040808320938716835260029093019052205460ff1661361c57336001600160a01b038516146135b2576040805162461bcd60e51b815260206004820152601560248201527439b2b73232b91036bab9ba1031329031aa37b5b2b760591b604482015290519081900360640190fd5b60006135be338561517d565b905060008160118111156135ce57fe5b146135e7578060118111156135df57fe5b915050611c4b565b6001600160a01b038086166000908152600960209081526040808320938816835260029093019052205460ff1661361a57fe5b505b600480546040805163fc57d4df60e01b81526001600160a01b03888116948201949094529051929091169163fc57d4df91602480820192602092909190829003018186803b15801561366d57600080fd5b505afa158015613681573d6000803e3d6000fd5b505050506040513d602081101561369757600080fd5b50516136a457600d611ff3565b6001600160a01b0384166000908152601660205260409020548015613791576000856001600160a01b03166347bd37186040518163ffffffff1660e01b815260040160206040518083038186803b1580156136fe57600080fd5b505afa158015613712573d6000803e3d6000fd5b505050506040513d602081101561372857600080fd5b5051905060006137388286614fb7565b905082811061378e576040805162461bcd60e51b815260206004820152601960248201527f6d61726b657420626f72726f7720636170207265616368656400000000000000604482015290519081900360640190fd5b50505b6000806137a18688600088614701565b919350909150600090508260118111156137b757fe5b146137d2578160118111156137c857fe5b9350505050611c4b565b80156137df5760046137c8565b6137e761572c565b6040518060200160405280896001600160a01b031663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561382b57600080fd5b505afa15801561383f573d6000803e3d6000fd5b505050506040513d602081101561385557600080fd5b5051905290506138658882613e61565b6138708888836140e9565b600098975050505050505050565b6008602052816000526040600020818154811061389757fe5b6000918252602090912001546001600160a01b03169150829050565b6003546001600160a01b031681565b600080546001600160a01b031633146138e8576138e160016006614e96565b9050611686565b6001600160a01b0383166000908152600960205260409020805460ff1661391d5761391560096007614e96565b915050611686565b61392561572c565b50604080516020810190915283815261393c61572c565b506040805160208101909152670c7d713b49da0000815261395d81836152ee565b156139785761396e60066008614e96565b9350505050611686565b8415801590613a015750600480546040805163fc57d4df60e01b81526001600160a01b038a8116948201949094529051929091169163fc57d4df91602480820192602092909190829003018186803b1580156139d357600080fd5b505afa1580156139e7573d6000803e3d6000fd5b505050506040513d60208110156139fd57600080fd5b5051155b15613a125761396e600d6009614e96565b60018301805490869055604080516001600160a01b03891681526020810183905280820188905290517f70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc59181900360600190a16000979650505050505050565b600a54600160a81b900460ff1681565b60055481565b6019546001600160a01b031681565b61184a81600d805480602002602001604051908101604052809291908181526020018280548015613af157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613ad3575b505050505061168c565b600080613b098585856150d1565b90508015613b18579050611c4b565b613b2185614a39565b613b2b8585614cb7565b600095945050505050565b6000808290506000806000836001600160a01b031663c37f68e2336040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060806040518083038186803b158015613b9757600080fd5b505afa158015613bab573d6000803e3d6000fd5b505050506040513d6080811015613bc157600080fd5b508051602082015160409092015190945090925090508215613c145760405162461bcd60e51b81526004018080602001828103825260258152602001806158966025913960400191505060405180910390fd5b8015613c3157613c26600c6002614e96565b945050505050611ba0565b6000613c3e8733856150d1565b90508015613c5f57613c53600e6003836152f5565b95505050505050611ba0565b6001600160a01b0385166000908152600960209081526040808320338452600281019092529091205460ff16613c9e5760009650505050505050611ba0565b3360009081526002820160209081526040808320805460ff191690556008825291829020805483518184028101840190945280845260609392830182828015613d1057602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311613cf2575b5050835193945083925060009150505b82811015613d6557896001600160a01b0316848281518110613d3e57fe5b60200260200101516001600160a01b03161415613d5d57809150613d65565b600101613d20565b50818110613d6f57fe5b336000908152600860205260409020805481906000198101908110613d9057fe5b9060005260206000200160009054906101000a90046001600160a01b0316818381548110613dba57fe5b600091825260209091200180546001600160a01b0319166001600160a01b03929092169190911790558054613df382600019830161573f565b50604080516001600160a01b038c16815233602082015281517fe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d929181900390910190a160009c9b505050505050505050505050565b61184a81613a97565b6000546001600160a01b031681565b6001600160a01b0382166000908152601160209081526040808320600f9092528220549091613e8e611eaf565b8354909150600090613eae908390600160e01b900463ffffffff16614f3b565b9050600081118015613ec05750600083115b1561408f576000613f35876001600160a01b03166347bd37186040518163ffffffff1660e01b815260040160206040518083038186803b158015613f0357600080fd5b505afa158015613f17573d6000803e3d6000fd5b505050506040513d6020811015613f2d57600080fd5b50518761535b565b90506000613f438386614f75565b9050613f4d61572c565b60008311613f6a5760405180602001604052806000815250613f74565b613f748284615379565b9050613f7e61572c565b604080516020810190915288546001600160e01b03168152613fa090836153ae565b90506040518060400160405280613ff083600001516040518060400160405280601a81526020017f6e657720696e64657820657863656564732032323420626974730000000000008152506153d3565b6001600160e01b0316815260200161402b886040518060400160405280601c81526020016000805160206158bb83398151915281525061546d565b63ffffffff9081169091526001600160a01b038c166000908152601160209081526040909120835181549490920151909216600160e01b026001600160e01b039182166001600160e01b0319909416939093171691909117905550611f1792505050565b8015611f17576140c2826040518060400160405280601c81526020016000805160206158bb83398151915281525061546d565b845463ffffffff91909116600160e01b026001600160e01b03909116178455505050505050565b6001600160a01b038316600090815260116020526040902061410961572c565b50604080516020810190915281546001600160e01b0316815261412a61572c565b5060408051602080820183526001600160a01b0380891660009081526013835284812091891680825282845294812080548552865195909152915291909155805115611f175761417861572c565b61418283836154c2565b905060006141df886001600160a01b03166395dd9193896040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015613f0357600080fd5b905060006141ed82846154e7565b6001600160a01b038916600090815260146020526040812054919250906142149083614fb7565b6001600160a01b03808b166000818152601460209081526040918290208590558a5182518881529182015281519495509193928e16927f1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a69281900390910190a350505050505050505050565b600080546001600160a01b03163314806142a457506002546001600160a01b031633145b905090565b6000806142b4612cb4565b604080516370a0823160e01b815230600482015290519192506000916001600160a01b038416916370a08231916024808301926020929190829003018186803b15801561430057600080fd5b505afa158015614314573d6000803e3d6000fd5b505050506040513d602081101561432a57600080fd5b50519050831580159061433d5750808411155b156143da57816001600160a01b031663a9059cbb86866040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b1580156143a257600080fd5b505af11580156143b6573d6000803e3d6000fd5b505050506040513d60208110156143cc57600080fd5b506000935061168692505050565b50919392505050565b6001600160a01b0382166000908152600f602052604090205480156144975761440a61572c565b6040518060200160405280856001600160a01b031663aa5af0fd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561444e57600080fd5b505afa158015614462573d6000803e3d6000fd5b505050506040513d602081101561447857600080fd5b50519052905061448784614a39565b6144918482613e61565b506146a5565b81156146a5576001600160a01b0383166000908152600960205260409020805460ff161515600114614510576040805162461bcd60e51b815260206004820152601960248201527f636f6d70206d61726b6574206973206e6f74206c697374656400000000000000604482015290519081900360640190fd5b6001600160a01b0384166000908152601060205260409020546001600160e01b03166145ee5760405180604001604052806ec097ce7bc90715b34b9f10000000006001600160e01b0316815260200161459361456a611eaf565b6040518060400160405280601c81526020016000805160206158bb83398151915281525061546d565b63ffffffff9081169091526001600160a01b0386166000908152601060209081526040909120835181549490920151909216600160e01b026001600160e01b039182166001600160e01b031990941693909317169190911790555b6001600160a01b0384166000908152601160205260409020546001600160e01b03166146a35760405180604001604052806ec097ce7bc90715b34b9f10000000006001600160e01b0316815260200161464861456a611eaf565b63ffffffff9081169091526001600160a01b0386166000908152601160209081526040909120835181549490920151909216600160e01b026001600160e01b039182166001600160e01b031990941693909317169190911790555b505b8181146116e9576001600160a01b0383166000818152600f6020908152604091829020859055815185815291517f2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b8079281900390910190a2505050565b600080600061470e615763565b6001600160a01b0388166000908152600860209081526040808320805482518185028101850190935280835260609383018282801561477657602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311614758575b50939450600093505050505b81518110156149fa57600082828151811061479957fe5b60200260200101519050806001600160a01b031663c37f68e28d6040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060806040518083038186803b1580156147f957600080fd5b505afa15801561480d573d6000803e3d6000fd5b505050506040513d608081101561482357600080fd5b508051602082015160408084015160609485015160808b01529389019390935291870191909152935083156148685750600f965060009550859450611f679350505050565b60408051602080820183526001600160a01b0380851660008181526009845285902060010154845260c08a01939093528351808301855260808a0151815260e08a015260048054855163fc57d4df60e01b815291820194909452935192169263fc57d4df9260248083019392829003018186803b1580156148e857600080fd5b505afa1580156148fc573d6000803e3d6000fd5b505050506040513d602081101561491257600080fd5b505160a086018190526149355750600d965060009550859450611f679350505050565b604080516020810190915260a0860151815261010086015260c085015160e086015161496f9161496491615273565b866101000151615273565b610120860181905260408601518651614989929190615516565b8552610100850151606086015160208701516149a6929190615516565b60208601526001600160a01b03818116908c1614156149f1576149d38561012001518b8760200151615516565b602086018190526101008601516149eb918b90615516565b60208601525b50600101614782565b50602083015183511115614a205750506020810151905160009450039150829050611f67565b5050805160209091015160009450849350039050611f67565b6001600160a01b0381166000908152601060209081526040808320600f9092528220549091614a66611eaf565b8354909150600090614a86908390600160e01b900463ffffffff16614f3b565b9050600081118015614a985750600083115b15614c5e576000856001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015614ad857600080fd5b505afa158015614aec573d6000803e3d6000fd5b505050506040513d6020811015614b0257600080fd5b505190506000614b128386614f75565b9050614b1c61572c565b60008311614b395760405180602001604052806000815250614b43565b614b438284615379565b9050614b4d61572c565b604080516020810190915288546001600160e01b03168152614b6f90836153ae565b90506040518060400160405280614bbf83600001516040518060400160405280601a81526020017f6e657720696e64657820657863656564732032323420626974730000000000008152506153d3565b6001600160e01b03168152602001614bfa886040518060400160405280601c81526020016000805160206158bb83398151915281525061546d565b63ffffffff9081169091526001600160a01b038b166000908152601060209081526040909120835181549490920151909216600160e01b026001600160e01b039182166001600160e01b031990941693909317169190911790555061185f92505050565b801561185f57614c91826040518060400160405280601c81526020016000805160206158bb83398151915281525061546d565b845463ffffffff91909116600160e01b026001600160e01b039091161784555050505050565b6001600160a01b0382166000908152601060205260409020614cd761572c565b50604080516020810190915281546001600160e01b03168152614cf861572c565b5060408051602080820183526001600160a01b03808816600090815260128352848120918816808252828452948120805485528651959091529152919091558051158015614d465750815115155b15614d5e576ec097ce7bc90715b34b9f100000000081525b614d6661572c565b614d7083836154c2565b90506000866001600160a01b03166370a08231876040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b158015614dca57600080fd5b505afa158015614dde573d6000803e3d6000fd5b505050506040513d6020811015614df457600080fd5b505190506000614e0482846154e7565b6001600160a01b03881660009081526014602052604081205491925090614e2b9083614fb7565b6001600160a01b03808a166000818152601460209081526040918290208590558a5182518881529182015281519495509193928d16927f2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a9281900390910190a3505050505050505050565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0836011811115614ec557fe5b836014811115614ed157fe5b604080519283526020830191909152600082820152519081900360600190a1826011811115611c4b57fe5b6000806000614f0f846000806000614701565b9250925092509193909250565b6000614f2661572c565b614f30848461553e565b90506119428161555f565b6000611c4b8383604051806040016040528060158152602001747375627472616374696f6e20756e646572666c6f7760581b81525061556e565b6000611c4b83836040518060400160405280601781526020017f6d756c7469706c69636174696f6e206f766572666c6f770000000000000000008152506155c8565b6000611c4b8383604051806040016040528060118152602001706164646974696f6e206f766572666c6f7760781b815250615647565b3b151590565b60005b600d5481101561507e57816001600160a01b0316600d828154811061501757fe5b6000918252602090912001546001600160a01b03161415615076576040805162461bcd60e51b81526020600482015260146024820152731b585c9ad95d08185b1c9958591e48185919195960621b604482015290519081900360640190fd5b600101614ff6565b50600d80546001810182556000919091527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb50180546001600160a01b0319166001600160a01b0392909216919091179055565b6001600160a01b03831660009081526009602052604081205460ff166150f8576009611ff3565b6001600160a01b038085166000908152600960209081526040808320938716835260029093019052205460ff16615130576000611ff3565b6000806151408587866000614701565b9193509091506000905082601181111561515657fe5b146151705781601181111561516757fe5b92505050611c4b565b8015613498576004615167565b6001600160a01b0382166000908152600960205260408120805460ff166151a8576009915050611686565b6001600160a01b038316600090815260028201602052604090205460ff161515600114156151da576000915050611686565b6001600160a01b0380841660008181526002840160209081526040808320805460ff19166001908117909155600883528184208054918201815584529282902090920180549489166001600160a01b031990951685179055815193845283019190915280517f3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a59281900390910190a15060009392505050565b61527b61572c565b6040518060200160405280670de0b6b3a76400006152a186600001518660000151614f75565b816152a857fe5b0490529392505050565b6152ba61572c565b60405180602001604052806152e56152de8660000151670de0b6b3a7640000614f75565b855161569c565b90529392505050565b5190511090565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa084601181111561532457fe5b84601481111561533057fe5b604080519283526020830191909152818101859052519081900360600190a183601181111561194257fe5b6000611c4b61537284670de0b6b3a7640000614f75565b835161569c565b61538161572c565b60405180602001604052806152e56153a8866ec097ce7bc90715b34b9f1000000000614f75565b8561569c565b6153b661572c565b60405180602001604052806152e585600001518560000151614fb7565b600081600160e01b84106154655760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561542a578181015183820152602001615412565b50505050905090810190601f1680156154575780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b509192915050565b600081600160201b84106154655760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561542a578181015183820152602001615412565b6154ca61572c565b60405180602001604052806152e585600001518560000151614f3b565b60006ec097ce7bc90715b34b9f1000000000615507848460000151614f75565b8161550e57fe5b049392505050565b600061552061572c565b61552a858561553e565b90506124a86155388261555f565b84614fb7565b61554661572c565b60405180602001604052806152e5856000015185614f75565b51670de0b6b3a7640000900490565b600081848411156155c05760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561542a578181015183820152602001615412565b505050900390565b60008315806155d5575082155b156155e257506000611c4b565b838302838582816155ef57fe5b0414839061563e5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561542a578181015183820152602001615412565b50949350505050565b6000838301828582101561563e5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561542a578181015183820152602001615412565b6000611c4b83836040518060400160405280600e81526020016d646976696465206279207a65726f60901b815250600081836157195760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561542a578181015183820152602001615412565b5082848161572357fe5b04949350505050565b6040518060200160405280600081525090565b8154818355818111156116e9576000838152602090206116e99181019083016157cd565b6040518061014001604052806000815260200160008152602001600081526020016000815260200160008152602001600081526020016157a161572c565b81526020016157ae61572c565b81526020016157bb61572c565b81526020016157c861572c565b905290565b611eb191905b808211156157e757600081556001016157d3565b509056fe63616e6e6f742070617573652061206d61726b65742074686174206973206e6f74206c69737465646f6e6c7920706175736520677561726469616e20616e642061646d696e2063616e2070617573656f6e6c792061646d696e2063616e2073657420626f72726f772063617020677561726469616e6f6e6c792061646d696e206f7220626f72726f772063617020677561726469616e2063616e2073657420626f72726f772063617073657869744d61726b65743a206765744163636f756e74536e617073686f74206661696c6564626c6f636b206e756d62657220657863656564732033322062697473000000006f6e6c7920756e6974726f6c6c65722061646d696e2063616e206368616e676520627261696e73a265627a7a723158203ea9f270a9c8d433fe9684989d148ea95f7b1054b23dec53e4d017a91bf4ab8064736f6c63430005100032"

// DeployAbicomptroller deploys a new Ethereum contract, binding an instance of Abicomptroller to it.
func DeployAbicomptroller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Abicomptroller, error) {
	parsed, err := abi.JSON(strings.NewReader(AbicomptrollerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbicomptrollerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Abicomptroller{AbicomptrollerCaller: AbicomptrollerCaller{contract: contract}, AbicomptrollerTransactor: AbicomptrollerTransactor{contract: contract}, AbicomptrollerFilterer: AbicomptrollerFilterer{contract: contract}}, nil
}

// Abicomptroller is an auto generated Go binding around an Ethereum contract.
type Abicomptroller struct {
	AbicomptrollerCaller     // Read-only binding to the contract
	AbicomptrollerTransactor // Write-only binding to the contract
	AbicomptrollerFilterer   // Log filterer for contract events
}

// AbicomptrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbicomptrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbicomptrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbicomptrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbicomptrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbicomptrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbicomptrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbicomptrollerSession struct {
	Contract     *Abicomptroller   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbicomptrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbicomptrollerCallerSession struct {
	Contract *AbicomptrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AbicomptrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbicomptrollerTransactorSession struct {
	Contract     *AbicomptrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AbicomptrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbicomptrollerRaw struct {
	Contract *Abicomptroller // Generic contract binding to access the raw methods on
}

// AbicomptrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbicomptrollerCallerRaw struct {
	Contract *AbicomptrollerCaller // Generic read-only contract binding to access the raw methods on
}

// AbicomptrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbicomptrollerTransactorRaw struct {
	Contract *AbicomptrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbicomptroller creates a new instance of Abicomptroller, bound to a specific deployed contract.
func NewAbicomptroller(address common.Address, backend bind.ContractBackend) (*Abicomptroller, error) {
	contract, err := bindAbicomptroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abicomptroller{AbicomptrollerCaller: AbicomptrollerCaller{contract: contract}, AbicomptrollerTransactor: AbicomptrollerTransactor{contract: contract}, AbicomptrollerFilterer: AbicomptrollerFilterer{contract: contract}}, nil
}

// NewAbicomptrollerCaller creates a new read-only instance of Abicomptroller, bound to a specific deployed contract.
func NewAbicomptrollerCaller(address common.Address, caller bind.ContractCaller) (*AbicomptrollerCaller, error) {
	contract, err := bindAbicomptroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerCaller{contract: contract}, nil
}

// NewAbicomptrollerTransactor creates a new write-only instance of Abicomptroller, bound to a specific deployed contract.
func NewAbicomptrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*AbicomptrollerTransactor, error) {
	contract, err := bindAbicomptroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerTransactor{contract: contract}, nil
}

// NewAbicomptrollerFilterer creates a new log filterer instance of Abicomptroller, bound to a specific deployed contract.
func NewAbicomptrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*AbicomptrollerFilterer, error) {
	contract, err := bindAbicomptroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerFilterer{contract: contract}, nil
}

// bindAbicomptroller binds a generic wrapper to an already deployed contract.
func bindAbicomptroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbicomptrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abicomptroller *AbicomptrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abicomptroller.Contract.AbicomptrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abicomptroller *AbicomptrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicomptroller.Contract.AbicomptrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abicomptroller *AbicomptrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abicomptroller.Contract.AbicomptrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abicomptroller *AbicomptrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abicomptroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abicomptroller *AbicomptrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abicomptroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abicomptroller *AbicomptrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abicomptroller.Contract.contract.Transact(opts, method, params...)
}

// VarBorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerCaller) VarBorrowGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "_borrowGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VarBorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerSession) VarBorrowGuardianPaused() (bool, error) {
	return _Abicomptroller.Contract.VarBorrowGuardianPaused(&_Abicomptroller.CallOpts)
}

// VarBorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerCallerSession) VarBorrowGuardianPaused() (bool, error) {
	return _Abicomptroller.Contract.VarBorrowGuardianPaused(&_Abicomptroller.CallOpts)
}

// VarMintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerCaller) VarMintGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "_mintGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VarMintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerSession) VarMintGuardianPaused() (bool, error) {
	return _Abicomptroller.Contract.VarMintGuardianPaused(&_Abicomptroller.CallOpts)
}

// VarMintGuardianPaused is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerCallerSession) VarMintGuardianPaused() (bool, error) {
	return _Abicomptroller.Contract.VarMintGuardianPaused(&_Abicomptroller.CallOpts)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Abicomptroller *AbicomptrollerCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "accountAssets", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Abicomptroller *AbicomptrollerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Abicomptroller.Contract.AccountAssets(&_Abicomptroller.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Abicomptroller *AbicomptrollerCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Abicomptroller.Contract.AccountAssets(&_Abicomptroller.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abicomptroller *AbicomptrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abicomptroller *AbicomptrollerSession) Admin() (common.Address, error) {
	return _Abicomptroller.Contract.Admin(&_Abicomptroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Abicomptroller *AbicomptrollerCallerSession) Admin() (common.Address, error) {
	return _Abicomptroller.Contract.Admin(&_Abicomptroller.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Abicomptroller *AbicomptrollerCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Abicomptroller *AbicomptrollerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Abicomptroller.Contract.AllMarkets(&_Abicomptroller.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Abicomptroller *AbicomptrollerCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Abicomptroller.Contract.AllMarkets(&_Abicomptroller.CallOpts, arg0)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Abicomptroller *AbicomptrollerCaller) BorrowCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "borrowCapGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Abicomptroller *AbicomptrollerSession) BorrowCapGuardian() (common.Address, error) {
	return _Abicomptroller.Contract.BorrowCapGuardian(&_Abicomptroller.CallOpts)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Abicomptroller *AbicomptrollerCallerSession) BorrowCapGuardian() (common.Address, error) {
	return _Abicomptroller.Contract.BorrowCapGuardian(&_Abicomptroller.CallOpts)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) BorrowCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "borrowCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.BorrowCaps(&_Abicomptroller.CallOpts, arg0)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.BorrowCaps(&_Abicomptroller.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Abicomptroller *AbicomptrollerCaller) BorrowGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "borrowGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Abicomptroller *AbicomptrollerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _Abicomptroller.Contract.BorrowGuardianPaused(&_Abicomptroller.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Abicomptroller *AbicomptrollerCallerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _Abicomptroller.Contract.BorrowGuardianPaused(&_Abicomptroller.CallOpts, arg0)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Abicomptroller *AbicomptrollerCaller) CheckMembership(opts *bind.CallOpts, account common.Address, cToken common.Address) (bool, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "checkMembership", account, cToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Abicomptroller *AbicomptrollerSession) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _Abicomptroller.Contract.CheckMembership(&_Abicomptroller.CallOpts, account, cToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Abicomptroller *AbicomptrollerCallerSession) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _Abicomptroller.Contract.CheckMembership(&_Abicomptroller.CallOpts, account, cToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "closeFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Abicomptroller.Contract.CloseFactorMantissa(&_Abicomptroller.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Abicomptroller.Contract.CloseFactorMantissa(&_Abicomptroller.CallOpts)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) CompAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "compAccrued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.CompAccrued(&_Abicomptroller.CallOpts, arg0)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.CompAccrued(&_Abicomptroller.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Abicomptroller *AbicomptrollerCaller) CompBorrowState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "compBorrowState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Abicomptroller *AbicomptrollerSession) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Abicomptroller.Contract.CompBorrowState(&_Abicomptroller.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Abicomptroller *AbicomptrollerCallerSession) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Abicomptroller.Contract.CompBorrowState(&_Abicomptroller.CallOpts, arg0)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) CompBorrowerIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "compBorrowerIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.CompBorrowerIndex(&_Abicomptroller.CallOpts, arg0, arg1)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.CompBorrowerIndex(&_Abicomptroller.CallOpts, arg0, arg1)
}

// CompContributorSpeeds is a free data retrieval call binding the contract method 0x986ab838.
//
// Solidity: function compContributorSpeeds(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) CompContributorSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "compContributorSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompContributorSpeeds is a free data retrieval call binding the contract method 0x986ab838.
//
// Solidity: function compContributorSpeeds(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) CompContributorSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.CompContributorSpeeds(&_Abicomptroller.CallOpts, arg0)
}

// CompContributorSpeeds is a free data retrieval call binding the contract method 0x986ab838.
//
// Solidity: function compContributorSpeeds(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) CompContributorSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.CompContributorSpeeds(&_Abicomptroller.CallOpts, arg0)
}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Abicomptroller *AbicomptrollerCaller) CompInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "compInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Abicomptroller *AbicomptrollerSession) CompInitialIndex() (*big.Int, error) {
	return _Abicomptroller.Contract.CompInitialIndex(&_Abicomptroller.CallOpts)
}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Abicomptroller *AbicomptrollerCallerSession) CompInitialIndex() (*big.Int, error) {
	return _Abicomptroller.Contract.CompInitialIndex(&_Abicomptroller.CallOpts)
}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) CompRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "compRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) CompRate() (*big.Int, error) {
	return _Abicomptroller.Contract.CompRate(&_Abicomptroller.CallOpts)
}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) CompRate() (*big.Int, error) {
	return _Abicomptroller.Contract.CompRate(&_Abicomptroller.CallOpts)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) CompSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "compSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.CompSpeeds(&_Abicomptroller.CallOpts, arg0)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.CompSpeeds(&_Abicomptroller.CallOpts, arg0)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) CompSupplierIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "compSupplierIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.CompSupplierIndex(&_Abicomptroller.CallOpts, arg0, arg1)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.CompSupplierIndex(&_Abicomptroller.CallOpts, arg0, arg1)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Abicomptroller *AbicomptrollerCaller) CompSupplyState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "compSupplyState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Abicomptroller *AbicomptrollerSession) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Abicomptroller.Contract.CompSupplyState(&_Abicomptroller.CallOpts, arg0)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Abicomptroller *AbicomptrollerCallerSession) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Abicomptroller.Contract.CompSupplyState(&_Abicomptroller.CallOpts, arg0)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Abicomptroller *AbicomptrollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Abicomptroller *AbicomptrollerSession) ComptrollerImplementation() (common.Address, error) {
	return _Abicomptroller.Contract.ComptrollerImplementation(&_Abicomptroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Abicomptroller *AbicomptrollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _Abicomptroller.Contract.ComptrollerImplementation(&_Abicomptroller.CallOpts)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Abicomptroller *AbicomptrollerCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "getAccountLiquidity", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Abicomptroller *AbicomptrollerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Abicomptroller.Contract.GetAccountLiquidity(&_Abicomptroller.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Abicomptroller.Contract.GetAccountLiquidity(&_Abicomptroller.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Abicomptroller *AbicomptrollerCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Abicomptroller *AbicomptrollerSession) GetAllMarkets() ([]common.Address, error) {
	return _Abicomptroller.Contract.GetAllMarkets(&_Abicomptroller.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Abicomptroller *AbicomptrollerCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _Abicomptroller.Contract.GetAllMarkets(&_Abicomptroller.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Abicomptroller *AbicomptrollerCaller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "getAssetsIn", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Abicomptroller *AbicomptrollerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Abicomptroller.Contract.GetAssetsIn(&_Abicomptroller.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Abicomptroller *AbicomptrollerCallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Abicomptroller.Contract.GetAssetsIn(&_Abicomptroller.CallOpts, account)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) GetBlockNumber() (*big.Int, error) {
	return _Abicomptroller.Contract.GetBlockNumber(&_Abicomptroller.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Abicomptroller.Contract.GetBlockNumber(&_Abicomptroller.CallOpts)
}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Abicomptroller *AbicomptrollerCaller) GetCompAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "getCompAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Abicomptroller *AbicomptrollerSession) GetCompAddress() (common.Address, error) {
	return _Abicomptroller.Contract.GetCompAddress(&_Abicomptroller.CallOpts)
}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Abicomptroller *AbicomptrollerCallerSession) GetCompAddress() (common.Address, error) {
	return _Abicomptroller.Contract.GetCompAddress(&_Abicomptroller.CallOpts)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Abicomptroller *AbicomptrollerCaller) GetHypotheticalAccountLiquidity(opts *bind.CallOpts, account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "getHypotheticalAccountLiquidity", account, cTokenModify, redeemTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Abicomptroller *AbicomptrollerSession) GetHypotheticalAccountLiquidity(account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Abicomptroller.Contract.GetHypotheticalAccountLiquidity(&_Abicomptroller.CallOpts, account, cTokenModify, redeemTokens, borrowAmount)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) GetHypotheticalAccountLiquidity(account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Abicomptroller.Contract.GetHypotheticalAccountLiquidity(&_Abicomptroller.CallOpts, account, cTokenModify, redeemTokens, borrowAmount)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Abicomptroller *AbicomptrollerCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "isComptroller")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Abicomptroller *AbicomptrollerSession) IsComptroller() (bool, error) {
	return _Abicomptroller.Contract.IsComptroller(&_Abicomptroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Abicomptroller *AbicomptrollerCallerSession) IsComptroller() (bool, error) {
	return _Abicomptroller.Contract.IsComptroller(&_Abicomptroller.CallOpts)
}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) LastContributorBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "lastContributorBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) LastContributorBlock(arg0 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.LastContributorBlock(&_Abicomptroller.CallOpts, arg0)
}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) LastContributorBlock(arg0 common.Address) (*big.Int, error) {
	return _Abicomptroller.Contract.LastContributorBlock(&_Abicomptroller.CallOpts, arg0)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Abicomptroller *AbicomptrollerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", cTokenBorrowed, cTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Abicomptroller *AbicomptrollerSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Abicomptroller.Contract.LiquidateCalculateSeizeTokens(&_Abicomptroller.CallOpts, cTokenBorrowed, cTokenCollateral, actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Abicomptroller.Contract.LiquidateCalculateSeizeTokens(&_Abicomptroller.CallOpts, cTokenBorrowed, cTokenCollateral, actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "liquidationIncentiveMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Abicomptroller.Contract.LiquidationIncentiveMantissa(&_Abicomptroller.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Abicomptroller.Contract.LiquidationIncentiveMantissa(&_Abicomptroller.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_Abicomptroller *AbicomptrollerCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
		IsComped                 bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsListed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CollateralFactorMantissa = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsComped = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_Abicomptroller *AbicomptrollerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	return _Abicomptroller.Contract.Markets(&_Abicomptroller.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_Abicomptroller *AbicomptrollerCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	return _Abicomptroller.Contract.Markets(&_Abicomptroller.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Abicomptroller *AbicomptrollerCaller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "maxAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) MaxAssets() (*big.Int, error) {
	return _Abicomptroller.Contract.MaxAssets(&_Abicomptroller.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Abicomptroller *AbicomptrollerCallerSession) MaxAssets() (*big.Int, error) {
	return _Abicomptroller.Contract.MaxAssets(&_Abicomptroller.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Abicomptroller *AbicomptrollerCaller) MintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "mintGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Abicomptroller *AbicomptrollerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Abicomptroller.Contract.MintGuardianPaused(&_Abicomptroller.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Abicomptroller *AbicomptrollerCallerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Abicomptroller.Contract.MintGuardianPaused(&_Abicomptroller.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Abicomptroller *AbicomptrollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Abicomptroller *AbicomptrollerSession) Oracle() (common.Address, error) {
	return _Abicomptroller.Contract.Oracle(&_Abicomptroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Abicomptroller *AbicomptrollerCallerSession) Oracle() (common.Address, error) {
	return _Abicomptroller.Contract.Oracle(&_Abicomptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Abicomptroller *AbicomptrollerCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Abicomptroller *AbicomptrollerSession) PauseGuardian() (common.Address, error) {
	return _Abicomptroller.Contract.PauseGuardian(&_Abicomptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Abicomptroller *AbicomptrollerCallerSession) PauseGuardian() (common.Address, error) {
	return _Abicomptroller.Contract.PauseGuardian(&_Abicomptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abicomptroller *AbicomptrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abicomptroller *AbicomptrollerSession) PendingAdmin() (common.Address, error) {
	return _Abicomptroller.Contract.PendingAdmin(&_Abicomptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Abicomptroller *AbicomptrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _Abicomptroller.Contract.PendingAdmin(&_Abicomptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Abicomptroller *AbicomptrollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Abicomptroller *AbicomptrollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Abicomptroller.Contract.PendingComptrollerImplementation(&_Abicomptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Abicomptroller *AbicomptrollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Abicomptroller.Contract.PendingComptrollerImplementation(&_Abicomptroller.CallOpts)
}

// RewardTokenAddr is a free data retrieval call binding the contract method 0xe89dec6d.
//
// Solidity: function rewardTokenAddr() view returns(address)
func (_Abicomptroller *AbicomptrollerCaller) RewardTokenAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "rewardTokenAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardTokenAddr is a free data retrieval call binding the contract method 0xe89dec6d.
//
// Solidity: function rewardTokenAddr() view returns(address)
func (_Abicomptroller *AbicomptrollerSession) RewardTokenAddr() (common.Address, error) {
	return _Abicomptroller.Contract.RewardTokenAddr(&_Abicomptroller.CallOpts)
}

// RewardTokenAddr is a free data retrieval call binding the contract method 0xe89dec6d.
//
// Solidity: function rewardTokenAddr() view returns(address)
func (_Abicomptroller *AbicomptrollerCallerSession) RewardTokenAddr() (common.Address, error) {
	return _Abicomptroller.Contract.RewardTokenAddr(&_Abicomptroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "seizeGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerSession) SeizeGuardianPaused() (bool, error) {
	return _Abicomptroller.Contract.SeizeGuardianPaused(&_Abicomptroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerCallerSession) SeizeGuardianPaused() (bool, error) {
	return _Abicomptroller.Contract.SeizeGuardianPaused(&_Abicomptroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Abicomptroller.contract.Call(opts, &out, "transferGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerSession) TransferGuardianPaused() (bool, error) {
	return _Abicomptroller.Contract.TransferGuardianPaused(&_Abicomptroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Abicomptroller *AbicomptrollerCallerSession) TransferGuardianPaused() (bool, error) {
	return _Abicomptroller.Contract.TransferGuardianPaused(&_Abicomptroller.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Abicomptroller *AbicomptrollerTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Abicomptroller *AbicomptrollerSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.Become(&_Abicomptroller.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.Become(&_Abicomptroller.TransactOpts, unitroller)
}

// GrantComp is a paid mutator transaction binding the contract method 0x27efe3cb.
//
// Solidity: function _grantComp(address recipient, uint256 amount) returns()
func (_Abicomptroller *AbicomptrollerTransactor) GrantComp(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_grantComp", recipient, amount)
}

// GrantComp is a paid mutator transaction binding the contract method 0x27efe3cb.
//
// Solidity: function _grantComp(address recipient, uint256 amount) returns()
func (_Abicomptroller *AbicomptrollerSession) GrantComp(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.GrantComp(&_Abicomptroller.TransactOpts, recipient, amount)
}

// GrantComp is a paid mutator transaction binding the contract method 0x27efe3cb.
//
// Solidity: function _grantComp(address recipient, uint256 amount) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) GrantComp(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.GrantComp(&_Abicomptroller.TransactOpts, recipient, amount)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Abicomptroller *AbicomptrollerTransactor) SetBorrowCapGuardian(opts *bind.TransactOpts, newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setBorrowCapGuardian", newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Abicomptroller *AbicomptrollerSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetBorrowCapGuardian(&_Abicomptroller.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetBorrowCapGuardian(&_Abicomptroller.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerTransactor) SetBorrowPaused(opts *bind.TransactOpts, cToken common.Address, state bool) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setBorrowPaused", cToken, state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerSession) SetBorrowPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetBorrowPaused(&_Abicomptroller.TransactOpts, cToken, state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerTransactorSession) SetBorrowPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetBorrowPaused(&_Abicomptroller.TransactOpts, cToken, state)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) SetCloseFactor(opts *bind.TransactOpts, newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setCloseFactor", newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetCloseFactor(&_Abicomptroller.TransactOpts, newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetCloseFactor(&_Abicomptroller.TransactOpts, newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) SetCollateralFactor(opts *bind.TransactOpts, cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setCollateralFactor", cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetCollateralFactor(&_Abicomptroller.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetCollateralFactor(&_Abicomptroller.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetCompAddress is a paid mutator transaction binding the contract method 0x8435be46.
//
// Solidity: function _setCompAddress(address newRewardTokenAddr) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) SetCompAddress(opts *bind.TransactOpts, newRewardTokenAddr common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setCompAddress", newRewardTokenAddr)
}

// SetCompAddress is a paid mutator transaction binding the contract method 0x8435be46.
//
// Solidity: function _setCompAddress(address newRewardTokenAddr) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) SetCompAddress(newRewardTokenAddr common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetCompAddress(&_Abicomptroller.TransactOpts, newRewardTokenAddr)
}

// SetCompAddress is a paid mutator transaction binding the contract method 0x8435be46.
//
// Solidity: function _setCompAddress(address newRewardTokenAddr) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) SetCompAddress(newRewardTokenAddr common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetCompAddress(&_Abicomptroller.TransactOpts, newRewardTokenAddr)
}

// SetCompSpeed is a paid mutator transaction binding the contract method 0x434caf25.
//
// Solidity: function _setCompSpeed(address cToken, uint256 compSpeed) returns()
func (_Abicomptroller *AbicomptrollerTransactor) SetCompSpeed(opts *bind.TransactOpts, cToken common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setCompSpeed", cToken, compSpeed)
}

// SetCompSpeed is a paid mutator transaction binding the contract method 0x434caf25.
//
// Solidity: function _setCompSpeed(address cToken, uint256 compSpeed) returns()
func (_Abicomptroller *AbicomptrollerSession) SetCompSpeed(cToken common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetCompSpeed(&_Abicomptroller.TransactOpts, cToken, compSpeed)
}

// SetCompSpeed is a paid mutator transaction binding the contract method 0x434caf25.
//
// Solidity: function _setCompSpeed(address cToken, uint256 compSpeed) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) SetCompSpeed(cToken common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetCompSpeed(&_Abicomptroller.TransactOpts, cToken, compSpeed)
}

// SetContributorCompSpeed is a paid mutator transaction binding the contract method 0x598ee1cb.
//
// Solidity: function _setContributorCompSpeed(address contributor, uint256 compSpeed) returns()
func (_Abicomptroller *AbicomptrollerTransactor) SetContributorCompSpeed(opts *bind.TransactOpts, contributor common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setContributorCompSpeed", contributor, compSpeed)
}

// SetContributorCompSpeed is a paid mutator transaction binding the contract method 0x598ee1cb.
//
// Solidity: function _setContributorCompSpeed(address contributor, uint256 compSpeed) returns()
func (_Abicomptroller *AbicomptrollerSession) SetContributorCompSpeed(contributor common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetContributorCompSpeed(&_Abicomptroller.TransactOpts, contributor, compSpeed)
}

// SetContributorCompSpeed is a paid mutator transaction binding the contract method 0x598ee1cb.
//
// Solidity: function _setContributorCompSpeed(address contributor, uint256 compSpeed) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) SetContributorCompSpeed(contributor common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetContributorCompSpeed(&_Abicomptroller.TransactOpts, contributor, compSpeed)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) SetLiquidationIncentive(opts *bind.TransactOpts, newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setLiquidationIncentive", newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetLiquidationIncentive(&_Abicomptroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetLiquidationIncentive(&_Abicomptroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Abicomptroller *AbicomptrollerTransactor) SetMarketBorrowCaps(opts *bind.TransactOpts, cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setMarketBorrowCaps", cTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Abicomptroller *AbicomptrollerSession) SetMarketBorrowCaps(cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetMarketBorrowCaps(&_Abicomptroller.TransactOpts, cTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) SetMarketBorrowCaps(cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetMarketBorrowCaps(&_Abicomptroller.TransactOpts, cTokens, newBorrowCaps)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerTransactor) SetMintPaused(opts *bind.TransactOpts, cToken common.Address, state bool) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setMintPaused", cToken, state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerSession) SetMintPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetMintPaused(&_Abicomptroller.TransactOpts, cToken, state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerTransactorSession) SetMintPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetMintPaused(&_Abicomptroller.TransactOpts, cToken, state)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) SetPauseGuardian(opts *bind.TransactOpts, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setPauseGuardian", newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetPauseGuardian(&_Abicomptroller.TransactOpts, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetPauseGuardian(&_Abicomptroller.TransactOpts, newPauseGuardian)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetPriceOracle(&_Abicomptroller.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetPriceOracle(&_Abicomptroller.TransactOpts, newOracle)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerTransactor) SetSeizePaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setSeizePaused", state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerSession) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetSeizePaused(&_Abicomptroller.TransactOpts, state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerTransactorSession) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetSeizePaused(&_Abicomptroller.TransactOpts, state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerTransactor) SetTransferPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_setTransferPaused", state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerSession) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetTransferPaused(&_Abicomptroller.TransactOpts, state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Abicomptroller *AbicomptrollerTransactorSession) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SetTransferPaused(&_Abicomptroller.TransactOpts, state)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) SupportMarket(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "_supportMarket", cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SupportMarket(&_Abicomptroller.TransactOpts, cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SupportMarket(&_Abicomptroller.TransactOpts, cToken)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) BorrowAllowed(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "borrowAllowed", cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.BorrowAllowed(&_Abicomptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.BorrowAllowed(&_Abicomptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Abicomptroller *AbicomptrollerTransactor) BorrowVerify(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "borrowVerify", cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Abicomptroller *AbicomptrollerSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.BorrowVerify(&_Abicomptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.BorrowVerify(&_Abicomptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactor) ClaimComp(opts *bind.TransactOpts, holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "claimComp", holder, cTokens)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_Abicomptroller *AbicomptrollerSession) ClaimComp(holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimComp(&_Abicomptroller.TransactOpts, holder, cTokens)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) ClaimComp(holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimComp(&_Abicomptroller.TransactOpts, holder, cTokens)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Abicomptroller *AbicomptrollerTransactor) ClaimComp0(opts *bind.TransactOpts, holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "claimComp0", holders, cTokens, borrowers, suppliers)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Abicomptroller *AbicomptrollerSession) ClaimComp0(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimComp0(&_Abicomptroller.TransactOpts, holders, cTokens, borrowers, suppliers)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) ClaimComp0(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimComp0(&_Abicomptroller.TransactOpts, holders, cTokens, borrowers, suppliers)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Abicomptroller *AbicomptrollerTransactor) ClaimComp1(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "claimComp1", holder)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Abicomptroller *AbicomptrollerSession) ClaimComp1(holder common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimComp1(&_Abicomptroller.TransactOpts, holder)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) ClaimComp1(holder common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimComp1(&_Abicomptroller.TransactOpts, holder)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x2026ffa3.
//
// Solidity: function claimRewards(address holder, address[] cTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactor) ClaimRewards(opts *bind.TransactOpts, holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "claimRewards", holder, cTokens)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x2026ffa3.
//
// Solidity: function claimRewards(address holder, address[] cTokens) returns()
func (_Abicomptroller *AbicomptrollerSession) ClaimRewards(holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimRewards(&_Abicomptroller.TransactOpts, holder, cTokens)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x2026ffa3.
//
// Solidity: function claimRewards(address holder, address[] cTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) ClaimRewards(holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimRewards(&_Abicomptroller.TransactOpts, holder, cTokens)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x54eb76fa.
//
// Solidity: function claimRewards(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Abicomptroller *AbicomptrollerTransactor) ClaimRewards0(opts *bind.TransactOpts, holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "claimRewards0", holders, cTokens, borrowers, suppliers)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x54eb76fa.
//
// Solidity: function claimRewards(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Abicomptroller *AbicomptrollerSession) ClaimRewards0(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimRewards0(&_Abicomptroller.TransactOpts, holders, cTokens, borrowers, suppliers)
}

// ClaimRewards0 is a paid mutator transaction binding the contract method 0x54eb76fa.
//
// Solidity: function claimRewards(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) ClaimRewards0(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimRewards0(&_Abicomptroller.TransactOpts, holders, cTokens, borrowers, suppliers)
}

// ClaimRewards1 is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address holder) returns()
func (_Abicomptroller *AbicomptrollerTransactor) ClaimRewards1(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "claimRewards1", holder)
}

// ClaimRewards1 is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address holder) returns()
func (_Abicomptroller *AbicomptrollerSession) ClaimRewards1(holder common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimRewards1(&_Abicomptroller.TransactOpts, holder)
}

// ClaimRewards1 is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address holder) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) ClaimRewards1(holder common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ClaimRewards1(&_Abicomptroller.TransactOpts, holder)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Abicomptroller *AbicomptrollerTransactor) EnterMarkets(opts *bind.TransactOpts, cTokens []common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "enterMarkets", cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Abicomptroller *AbicomptrollerSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.EnterMarkets(&_Abicomptroller.TransactOpts, cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Abicomptroller *AbicomptrollerTransactorSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.EnterMarkets(&_Abicomptroller.TransactOpts, cTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) ExitMarket(opts *bind.TransactOpts, cTokenAddress common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "exitMarket", cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ExitMarket(&_Abicomptroller.TransactOpts, cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.ExitMarket(&_Abicomptroller.TransactOpts, cTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "liquidateBorrowAllowed", cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.LiquidateBorrowAllowed(&_Abicomptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.LiquidateBorrowAllowed(&_Abicomptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "liquidateBorrowVerify", cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Abicomptroller *AbicomptrollerSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.LiquidateBorrowVerify(&_Abicomptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.LiquidateBorrowVerify(&_Abicomptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) MintAllowed(opts *bind.TransactOpts, cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "mintAllowed", cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.MintAllowed(&_Abicomptroller.TransactOpts, cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.MintAllowed(&_Abicomptroller.TransactOpts, cToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactor) MintVerify(opts *bind.TransactOpts, cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "mintVerify", cToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Abicomptroller *AbicomptrollerSession) MintVerify(cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.MintVerify(&_Abicomptroller.TransactOpts, cToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) MintVerify(cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.MintVerify(&_Abicomptroller.TransactOpts, cToken, minter, actualMintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) RedeemAllowed(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "redeemAllowed", cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.RedeemAllowed(&_Abicomptroller.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.RedeemAllowed(&_Abicomptroller.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactor) RedeemVerify(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "redeemVerify", cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Abicomptroller *AbicomptrollerSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.RedeemVerify(&_Abicomptroller.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.RedeemVerify(&_Abicomptroller.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "repayBorrowAllowed", cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.RepayBorrowAllowed(&_Abicomptroller.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.RepayBorrowAllowed(&_Abicomptroller.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Abicomptroller *AbicomptrollerTransactor) RepayBorrowVerify(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "repayBorrowVerify", cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Abicomptroller *AbicomptrollerSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.RepayBorrowVerify(&_Abicomptroller.TransactOpts, cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.RepayBorrowVerify(&_Abicomptroller.TransactOpts, cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) SeizeAllowed(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "seizeAllowed", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SeizeAllowed(&_Abicomptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SeizeAllowed(&_Abicomptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactor) SeizeVerify(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "seizeVerify", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Abicomptroller *AbicomptrollerSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SeizeVerify(&_Abicomptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.SeizeVerify(&_Abicomptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactor) TransferAllowed(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "transferAllowed", cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Abicomptroller *AbicomptrollerSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.TransferAllowed(&_Abicomptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Abicomptroller *AbicomptrollerTransactorSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.TransferAllowed(&_Abicomptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactor) TransferVerify(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "transferVerify", cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Abicomptroller *AbicomptrollerSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.TransferVerify(&_Abicomptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Abicomptroller.Contract.TransferVerify(&_Abicomptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// UpdateContributorRewards is a paid mutator transaction binding the contract method 0x741b2525.
//
// Solidity: function updateContributorRewards(address contributor) returns()
func (_Abicomptroller *AbicomptrollerTransactor) UpdateContributorRewards(opts *bind.TransactOpts, contributor common.Address) (*types.Transaction, error) {
	return _Abicomptroller.contract.Transact(opts, "updateContributorRewards", contributor)
}

// UpdateContributorRewards is a paid mutator transaction binding the contract method 0x741b2525.
//
// Solidity: function updateContributorRewards(address contributor) returns()
func (_Abicomptroller *AbicomptrollerSession) UpdateContributorRewards(contributor common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.UpdateContributorRewards(&_Abicomptroller.TransactOpts, contributor)
}

// UpdateContributorRewards is a paid mutator transaction binding the contract method 0x741b2525.
//
// Solidity: function updateContributorRewards(address contributor) returns()
func (_Abicomptroller *AbicomptrollerTransactorSession) UpdateContributorRewards(contributor common.Address) (*types.Transaction, error) {
	return _Abicomptroller.Contract.UpdateContributorRewards(&_Abicomptroller.TransactOpts, contributor)
}

// AbicomptrollerActionPausedIterator is returned from FilterActionPaused and is used to iterate over the raw logs and unpacked data for ActionPaused events raised by the Abicomptroller contract.
type AbicomptrollerActionPausedIterator struct {
	Event *AbicomptrollerActionPaused // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerActionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerActionPaused)
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
		it.Event = new(AbicomptrollerActionPaused)
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
func (it *AbicomptrollerActionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerActionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerActionPaused represents a ActionPaused event raised by the Abicomptroller contract.
type AbicomptrollerActionPaused struct {
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused is a free log retrieval operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Abicomptroller *AbicomptrollerFilterer) FilterActionPaused(opts *bind.FilterOpts) (*AbicomptrollerActionPausedIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerActionPausedIterator{contract: _Abicomptroller.contract, event: "ActionPaused", logs: logs, sub: sub}, nil
}

// WatchActionPaused is a free log subscription operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Abicomptroller *AbicomptrollerFilterer) WatchActionPaused(opts *bind.WatchOpts, sink chan<- *AbicomptrollerActionPaused) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerActionPaused)
				if err := _Abicomptroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
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

// ParseActionPaused is a log parse operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Abicomptroller *AbicomptrollerFilterer) ParseActionPaused(log types.Log) (*AbicomptrollerActionPaused, error) {
	event := new(AbicomptrollerActionPaused)
	if err := _Abicomptroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerActionPaused0Iterator is returned from FilterActionPaused0 and is used to iterate over the raw logs and unpacked data for ActionPaused0 events raised by the Abicomptroller contract.
type AbicomptrollerActionPaused0Iterator struct {
	Event *AbicomptrollerActionPaused0 // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerActionPaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerActionPaused0)
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
		it.Event = new(AbicomptrollerActionPaused0)
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
func (it *AbicomptrollerActionPaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerActionPaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerActionPaused0 represents a ActionPaused0 event raised by the Abicomptroller contract.
type AbicomptrollerActionPaused0 struct {
	CToken     common.Address
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused0 is a free log retrieval operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_Abicomptroller *AbicomptrollerFilterer) FilterActionPaused0(opts *bind.FilterOpts) (*AbicomptrollerActionPaused0Iterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerActionPaused0Iterator{contract: _Abicomptroller.contract, event: "ActionPaused0", logs: logs, sub: sub}, nil
}

// WatchActionPaused0 is a free log subscription operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_Abicomptroller *AbicomptrollerFilterer) WatchActionPaused0(opts *bind.WatchOpts, sink chan<- *AbicomptrollerActionPaused0) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerActionPaused0)
				if err := _Abicomptroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
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

// ParseActionPaused0 is a log parse operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_Abicomptroller *AbicomptrollerFilterer) ParseActionPaused0(log types.Log) (*AbicomptrollerActionPaused0, error) {
	event := new(AbicomptrollerActionPaused0)
	if err := _Abicomptroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerCompGrantedIterator is returned from FilterCompGranted and is used to iterate over the raw logs and unpacked data for CompGranted events raised by the Abicomptroller contract.
type AbicomptrollerCompGrantedIterator struct {
	Event *AbicomptrollerCompGranted // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerCompGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerCompGranted)
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
		it.Event = new(AbicomptrollerCompGranted)
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
func (it *AbicomptrollerCompGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerCompGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerCompGranted represents a CompGranted event raised by the Abicomptroller contract.
type AbicomptrollerCompGranted struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCompGranted is a free log retrieval operation binding the contract event 0x98b2f82a3a07f223a0be64b3d0f47711c64dccd1feafb94aa28156b38cd9695c.
//
// Solidity: event CompGranted(address recipient, uint256 amount)
func (_Abicomptroller *AbicomptrollerFilterer) FilterCompGranted(opts *bind.FilterOpts) (*AbicomptrollerCompGrantedIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "CompGranted")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerCompGrantedIterator{contract: _Abicomptroller.contract, event: "CompGranted", logs: logs, sub: sub}, nil
}

// WatchCompGranted is a free log subscription operation binding the contract event 0x98b2f82a3a07f223a0be64b3d0f47711c64dccd1feafb94aa28156b38cd9695c.
//
// Solidity: event CompGranted(address recipient, uint256 amount)
func (_Abicomptroller *AbicomptrollerFilterer) WatchCompGranted(opts *bind.WatchOpts, sink chan<- *AbicomptrollerCompGranted) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "CompGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerCompGranted)
				if err := _Abicomptroller.contract.UnpackLog(event, "CompGranted", log); err != nil {
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

// ParseCompGranted is a log parse operation binding the contract event 0x98b2f82a3a07f223a0be64b3d0f47711c64dccd1feafb94aa28156b38cd9695c.
//
// Solidity: event CompGranted(address recipient, uint256 amount)
func (_Abicomptroller *AbicomptrollerFilterer) ParseCompGranted(log types.Log) (*AbicomptrollerCompGranted, error) {
	event := new(AbicomptrollerCompGranted)
	if err := _Abicomptroller.contract.UnpackLog(event, "CompGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerCompSpeedUpdatedIterator is returned from FilterCompSpeedUpdated and is used to iterate over the raw logs and unpacked data for CompSpeedUpdated events raised by the Abicomptroller contract.
type AbicomptrollerCompSpeedUpdatedIterator struct {
	Event *AbicomptrollerCompSpeedUpdated // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerCompSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerCompSpeedUpdated)
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
		it.Event = new(AbicomptrollerCompSpeedUpdated)
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
func (it *AbicomptrollerCompSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerCompSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerCompSpeedUpdated represents a CompSpeedUpdated event raised by the Abicomptroller contract.
type AbicomptrollerCompSpeedUpdated struct {
	CToken   common.Address
	NewSpeed *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCompSpeedUpdated is a free log retrieval operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Abicomptroller *AbicomptrollerFilterer) FilterCompSpeedUpdated(opts *bind.FilterOpts, cToken []common.Address) (*AbicomptrollerCompSpeedUpdatedIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "CompSpeedUpdated", cTokenRule)
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerCompSpeedUpdatedIterator{contract: _Abicomptroller.contract, event: "CompSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchCompSpeedUpdated is a free log subscription operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Abicomptroller *AbicomptrollerFilterer) WatchCompSpeedUpdated(opts *bind.WatchOpts, sink chan<- *AbicomptrollerCompSpeedUpdated, cToken []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "CompSpeedUpdated", cTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerCompSpeedUpdated)
				if err := _Abicomptroller.contract.UnpackLog(event, "CompSpeedUpdated", log); err != nil {
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

// ParseCompSpeedUpdated is a log parse operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Abicomptroller *AbicomptrollerFilterer) ParseCompSpeedUpdated(log types.Log) (*AbicomptrollerCompSpeedUpdated, error) {
	event := new(AbicomptrollerCompSpeedUpdated)
	if err := _Abicomptroller.contract.UnpackLog(event, "CompSpeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerContributorCompSpeedUpdatedIterator is returned from FilterContributorCompSpeedUpdated and is used to iterate over the raw logs and unpacked data for ContributorCompSpeedUpdated events raised by the Abicomptroller contract.
type AbicomptrollerContributorCompSpeedUpdatedIterator struct {
	Event *AbicomptrollerContributorCompSpeedUpdated // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerContributorCompSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerContributorCompSpeedUpdated)
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
		it.Event = new(AbicomptrollerContributorCompSpeedUpdated)
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
func (it *AbicomptrollerContributorCompSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerContributorCompSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerContributorCompSpeedUpdated represents a ContributorCompSpeedUpdated event raised by the Abicomptroller contract.
type AbicomptrollerContributorCompSpeedUpdated struct {
	Contributor common.Address
	NewSpeed    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContributorCompSpeedUpdated is a free log retrieval operation binding the contract event 0x386537fa92edc3319af95f1f904dcf1900021e4f3f4e08169a577a09076e66b3.
//
// Solidity: event ContributorCompSpeedUpdated(address indexed contributor, uint256 newSpeed)
func (_Abicomptroller *AbicomptrollerFilterer) FilterContributorCompSpeedUpdated(opts *bind.FilterOpts, contributor []common.Address) (*AbicomptrollerContributorCompSpeedUpdatedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "ContributorCompSpeedUpdated", contributorRule)
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerContributorCompSpeedUpdatedIterator{contract: _Abicomptroller.contract, event: "ContributorCompSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchContributorCompSpeedUpdated is a free log subscription operation binding the contract event 0x386537fa92edc3319af95f1f904dcf1900021e4f3f4e08169a577a09076e66b3.
//
// Solidity: event ContributorCompSpeedUpdated(address indexed contributor, uint256 newSpeed)
func (_Abicomptroller *AbicomptrollerFilterer) WatchContributorCompSpeedUpdated(opts *bind.WatchOpts, sink chan<- *AbicomptrollerContributorCompSpeedUpdated, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "ContributorCompSpeedUpdated", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerContributorCompSpeedUpdated)
				if err := _Abicomptroller.contract.UnpackLog(event, "ContributorCompSpeedUpdated", log); err != nil {
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

// ParseContributorCompSpeedUpdated is a log parse operation binding the contract event 0x386537fa92edc3319af95f1f904dcf1900021e4f3f4e08169a577a09076e66b3.
//
// Solidity: event ContributorCompSpeedUpdated(address indexed contributor, uint256 newSpeed)
func (_Abicomptroller *AbicomptrollerFilterer) ParseContributorCompSpeedUpdated(log types.Log) (*AbicomptrollerContributorCompSpeedUpdated, error) {
	event := new(AbicomptrollerContributorCompSpeedUpdated)
	if err := _Abicomptroller.contract.UnpackLog(event, "ContributorCompSpeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerDistributedBorrowerCompIterator is returned from FilterDistributedBorrowerComp and is used to iterate over the raw logs and unpacked data for DistributedBorrowerComp events raised by the Abicomptroller contract.
type AbicomptrollerDistributedBorrowerCompIterator struct {
	Event *AbicomptrollerDistributedBorrowerComp // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerDistributedBorrowerCompIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerDistributedBorrowerComp)
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
		it.Event = new(AbicomptrollerDistributedBorrowerComp)
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
func (it *AbicomptrollerDistributedBorrowerCompIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerDistributedBorrowerCompIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerDistributedBorrowerComp represents a DistributedBorrowerComp event raised by the Abicomptroller contract.
type AbicomptrollerDistributedBorrowerComp struct {
	CToken          common.Address
	Borrower        common.Address
	CompDelta       *big.Int
	CompBorrowIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributedBorrowerComp is a free log retrieval operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Abicomptroller *AbicomptrollerFilterer) FilterDistributedBorrowerComp(opts *bind.FilterOpts, cToken []common.Address, borrower []common.Address) (*AbicomptrollerDistributedBorrowerCompIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "DistributedBorrowerComp", cTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerDistributedBorrowerCompIterator{contract: _Abicomptroller.contract, event: "DistributedBorrowerComp", logs: logs, sub: sub}, nil
}

// WatchDistributedBorrowerComp is a free log subscription operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Abicomptroller *AbicomptrollerFilterer) WatchDistributedBorrowerComp(opts *bind.WatchOpts, sink chan<- *AbicomptrollerDistributedBorrowerComp, cToken []common.Address, borrower []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "DistributedBorrowerComp", cTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerDistributedBorrowerComp)
				if err := _Abicomptroller.contract.UnpackLog(event, "DistributedBorrowerComp", log); err != nil {
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

// ParseDistributedBorrowerComp is a log parse operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Abicomptroller *AbicomptrollerFilterer) ParseDistributedBorrowerComp(log types.Log) (*AbicomptrollerDistributedBorrowerComp, error) {
	event := new(AbicomptrollerDistributedBorrowerComp)
	if err := _Abicomptroller.contract.UnpackLog(event, "DistributedBorrowerComp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerDistributedSupplierCompIterator is returned from FilterDistributedSupplierComp and is used to iterate over the raw logs and unpacked data for DistributedSupplierComp events raised by the Abicomptroller contract.
type AbicomptrollerDistributedSupplierCompIterator struct {
	Event *AbicomptrollerDistributedSupplierComp // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerDistributedSupplierCompIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerDistributedSupplierComp)
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
		it.Event = new(AbicomptrollerDistributedSupplierComp)
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
func (it *AbicomptrollerDistributedSupplierCompIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerDistributedSupplierCompIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerDistributedSupplierComp represents a DistributedSupplierComp event raised by the Abicomptroller contract.
type AbicomptrollerDistributedSupplierComp struct {
	CToken          common.Address
	Supplier        common.Address
	CompDelta       *big.Int
	CompSupplyIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributedSupplierComp is a free log retrieval operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Abicomptroller *AbicomptrollerFilterer) FilterDistributedSupplierComp(opts *bind.FilterOpts, cToken []common.Address, supplier []common.Address) (*AbicomptrollerDistributedSupplierCompIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "DistributedSupplierComp", cTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerDistributedSupplierCompIterator{contract: _Abicomptroller.contract, event: "DistributedSupplierComp", logs: logs, sub: sub}, nil
}

// WatchDistributedSupplierComp is a free log subscription operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Abicomptroller *AbicomptrollerFilterer) WatchDistributedSupplierComp(opts *bind.WatchOpts, sink chan<- *AbicomptrollerDistributedSupplierComp, cToken []common.Address, supplier []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "DistributedSupplierComp", cTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerDistributedSupplierComp)
				if err := _Abicomptroller.contract.UnpackLog(event, "DistributedSupplierComp", log); err != nil {
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

// ParseDistributedSupplierComp is a log parse operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Abicomptroller *AbicomptrollerFilterer) ParseDistributedSupplierComp(log types.Log) (*AbicomptrollerDistributedSupplierComp, error) {
	event := new(AbicomptrollerDistributedSupplierComp)
	if err := _Abicomptroller.contract.UnpackLog(event, "DistributedSupplierComp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Abicomptroller contract.
type AbicomptrollerFailureIterator struct {
	Event *AbicomptrollerFailure // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerFailure)
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
		it.Event = new(AbicomptrollerFailure)
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
func (it *AbicomptrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerFailure represents a Failure event raised by the Abicomptroller contract.
type AbicomptrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Abicomptroller *AbicomptrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*AbicomptrollerFailureIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerFailureIterator{contract: _Abicomptroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Abicomptroller *AbicomptrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *AbicomptrollerFailure) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerFailure)
				if err := _Abicomptroller.contract.UnpackLog(event, "Failure", log); err != nil {
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
func (_Abicomptroller *AbicomptrollerFilterer) ParseFailure(log types.Log) (*AbicomptrollerFailure, error) {
	event := new(AbicomptrollerFailure)
	if err := _Abicomptroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerMarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the Abicomptroller contract.
type AbicomptrollerMarketEnteredIterator struct {
	Event *AbicomptrollerMarketEntered // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerMarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerMarketEntered)
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
		it.Event = new(AbicomptrollerMarketEntered)
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
func (it *AbicomptrollerMarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerMarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerMarketEntered represents a MarketEntered event raised by the Abicomptroller contract.
type AbicomptrollerMarketEntered struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Abicomptroller *AbicomptrollerFilterer) FilterMarketEntered(opts *bind.FilterOpts) (*AbicomptrollerMarketEnteredIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerMarketEnteredIterator{contract: _Abicomptroller.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Abicomptroller *AbicomptrollerFilterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *AbicomptrollerMarketEntered) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerMarketEntered)
				if err := _Abicomptroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
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

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Abicomptroller *AbicomptrollerFilterer) ParseMarketEntered(log types.Log) (*AbicomptrollerMarketEntered, error) {
	event := new(AbicomptrollerMarketEntered)
	if err := _Abicomptroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerMarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the Abicomptroller contract.
type AbicomptrollerMarketExitedIterator struct {
	Event *AbicomptrollerMarketExited // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerMarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerMarketExited)
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
		it.Event = new(AbicomptrollerMarketExited)
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
func (it *AbicomptrollerMarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerMarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerMarketExited represents a MarketExited event raised by the Abicomptroller contract.
type AbicomptrollerMarketExited struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Abicomptroller *AbicomptrollerFilterer) FilterMarketExited(opts *bind.FilterOpts) (*AbicomptrollerMarketExitedIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerMarketExitedIterator{contract: _Abicomptroller.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Abicomptroller *AbicomptrollerFilterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *AbicomptrollerMarketExited) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerMarketExited)
				if err := _Abicomptroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
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

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Abicomptroller *AbicomptrollerFilterer) ParseMarketExited(log types.Log) (*AbicomptrollerMarketExited, error) {
	event := new(AbicomptrollerMarketExited)
	if err := _Abicomptroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerMarketListedIterator is returned from FilterMarketListed and is used to iterate over the raw logs and unpacked data for MarketListed events raised by the Abicomptroller contract.
type AbicomptrollerMarketListedIterator struct {
	Event *AbicomptrollerMarketListed // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerMarketListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerMarketListed)
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
		it.Event = new(AbicomptrollerMarketListed)
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
func (it *AbicomptrollerMarketListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerMarketListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerMarketListed represents a MarketListed event raised by the Abicomptroller contract.
type AbicomptrollerMarketListed struct {
	CToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketListed is a free log retrieval operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Abicomptroller *AbicomptrollerFilterer) FilterMarketListed(opts *bind.FilterOpts) (*AbicomptrollerMarketListedIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerMarketListedIterator{contract: _Abicomptroller.contract, event: "MarketListed", logs: logs, sub: sub}, nil
}

// WatchMarketListed is a free log subscription operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Abicomptroller *AbicomptrollerFilterer) WatchMarketListed(opts *bind.WatchOpts, sink chan<- *AbicomptrollerMarketListed) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerMarketListed)
				if err := _Abicomptroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
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

// ParseMarketListed is a log parse operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Abicomptroller *AbicomptrollerFilterer) ParseMarketListed(log types.Log) (*AbicomptrollerMarketListed, error) {
	event := new(AbicomptrollerMarketListed)
	if err := _Abicomptroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerNewBorrowCapIterator is returned from FilterNewBorrowCap and is used to iterate over the raw logs and unpacked data for NewBorrowCap events raised by the Abicomptroller contract.
type AbicomptrollerNewBorrowCapIterator struct {
	Event *AbicomptrollerNewBorrowCap // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerNewBorrowCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerNewBorrowCap)
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
		it.Event = new(AbicomptrollerNewBorrowCap)
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
func (it *AbicomptrollerNewBorrowCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerNewBorrowCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerNewBorrowCap represents a NewBorrowCap event raised by the Abicomptroller contract.
type AbicomptrollerNewBorrowCap struct {
	CToken       common.Address
	NewBorrowCap *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCap is a free log retrieval operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Abicomptroller *AbicomptrollerFilterer) FilterNewBorrowCap(opts *bind.FilterOpts, cToken []common.Address) (*AbicomptrollerNewBorrowCapIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "NewBorrowCap", cTokenRule)
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerNewBorrowCapIterator{contract: _Abicomptroller.contract, event: "NewBorrowCap", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCap is a free log subscription operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Abicomptroller *AbicomptrollerFilterer) WatchNewBorrowCap(opts *bind.WatchOpts, sink chan<- *AbicomptrollerNewBorrowCap, cToken []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "NewBorrowCap", cTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerNewBorrowCap)
				if err := _Abicomptroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
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

// ParseNewBorrowCap is a log parse operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Abicomptroller *AbicomptrollerFilterer) ParseNewBorrowCap(log types.Log) (*AbicomptrollerNewBorrowCap, error) {
	event := new(AbicomptrollerNewBorrowCap)
	if err := _Abicomptroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerNewBorrowCapGuardianIterator is returned from FilterNewBorrowCapGuardian and is used to iterate over the raw logs and unpacked data for NewBorrowCapGuardian events raised by the Abicomptroller contract.
type AbicomptrollerNewBorrowCapGuardianIterator struct {
	Event *AbicomptrollerNewBorrowCapGuardian // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerNewBorrowCapGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerNewBorrowCapGuardian)
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
		it.Event = new(AbicomptrollerNewBorrowCapGuardian)
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
func (it *AbicomptrollerNewBorrowCapGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerNewBorrowCapGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerNewBorrowCapGuardian represents a NewBorrowCapGuardian event raised by the Abicomptroller contract.
type AbicomptrollerNewBorrowCapGuardian struct {
	OldBorrowCapGuardian common.Address
	NewBorrowCapGuardian common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCapGuardian is a free log retrieval operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Abicomptroller *AbicomptrollerFilterer) FilterNewBorrowCapGuardian(opts *bind.FilterOpts) (*AbicomptrollerNewBorrowCapGuardianIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerNewBorrowCapGuardianIterator{contract: _Abicomptroller.contract, event: "NewBorrowCapGuardian", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCapGuardian is a free log subscription operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Abicomptroller *AbicomptrollerFilterer) WatchNewBorrowCapGuardian(opts *bind.WatchOpts, sink chan<- *AbicomptrollerNewBorrowCapGuardian) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerNewBorrowCapGuardian)
				if err := _Abicomptroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
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

// ParseNewBorrowCapGuardian is a log parse operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Abicomptroller *AbicomptrollerFilterer) ParseNewBorrowCapGuardian(log types.Log) (*AbicomptrollerNewBorrowCapGuardian, error) {
	event := new(AbicomptrollerNewBorrowCapGuardian)
	if err := _Abicomptroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerNewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the Abicomptroller contract.
type AbicomptrollerNewCloseFactorIterator struct {
	Event *AbicomptrollerNewCloseFactor // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerNewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerNewCloseFactor)
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
		it.Event = new(AbicomptrollerNewCloseFactor)
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
func (it *AbicomptrollerNewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerNewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerNewCloseFactor represents a NewCloseFactor event raised by the Abicomptroller contract.
type AbicomptrollerNewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Abicomptroller *AbicomptrollerFilterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*AbicomptrollerNewCloseFactorIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerNewCloseFactorIterator{contract: _Abicomptroller.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Abicomptroller *AbicomptrollerFilterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *AbicomptrollerNewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerNewCloseFactor)
				if err := _Abicomptroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
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

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Abicomptroller *AbicomptrollerFilterer) ParseNewCloseFactor(log types.Log) (*AbicomptrollerNewCloseFactor, error) {
	event := new(AbicomptrollerNewCloseFactor)
	if err := _Abicomptroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerNewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the Abicomptroller contract.
type AbicomptrollerNewCollateralFactorIterator struct {
	Event *AbicomptrollerNewCollateralFactor // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerNewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerNewCollateralFactor)
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
		it.Event = new(AbicomptrollerNewCollateralFactor)
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
func (it *AbicomptrollerNewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerNewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerNewCollateralFactor represents a NewCollateralFactor event raised by the Abicomptroller contract.
type AbicomptrollerNewCollateralFactor struct {
	CToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Abicomptroller *AbicomptrollerFilterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*AbicomptrollerNewCollateralFactorIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerNewCollateralFactorIterator{contract: _Abicomptroller.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Abicomptroller *AbicomptrollerFilterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *AbicomptrollerNewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerNewCollateralFactor)
				if err := _Abicomptroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
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

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Abicomptroller *AbicomptrollerFilterer) ParseNewCollateralFactor(log types.Log) (*AbicomptrollerNewCollateralFactor, error) {
	event := new(AbicomptrollerNewCollateralFactor)
	if err := _Abicomptroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerNewCompAddressIterator is returned from FilterNewCompAddress and is used to iterate over the raw logs and unpacked data for NewCompAddress events raised by the Abicomptroller contract.
type AbicomptrollerNewCompAddressIterator struct {
	Event *AbicomptrollerNewCompAddress // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerNewCompAddressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerNewCompAddress)
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
		it.Event = new(AbicomptrollerNewCompAddress)
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
func (it *AbicomptrollerNewCompAddressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerNewCompAddressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerNewCompAddress represents a NewCompAddress event raised by the Abicomptroller contract.
type AbicomptrollerNewCompAddress struct {
	OldRewardTokenAddr common.Address
	NewRewardTokenAddr common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNewCompAddress is a free log retrieval operation binding the contract event 0xdde58acf741e6f5e57788ee0f64a68576e5eadf169e78251f32824653dd308bf.
//
// Solidity: event NewCompAddress(address oldRewardTokenAddr, address newRewardTokenAddr)
func (_Abicomptroller *AbicomptrollerFilterer) FilterNewCompAddress(opts *bind.FilterOpts) (*AbicomptrollerNewCompAddressIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "NewCompAddress")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerNewCompAddressIterator{contract: _Abicomptroller.contract, event: "NewCompAddress", logs: logs, sub: sub}, nil
}

// WatchNewCompAddress is a free log subscription operation binding the contract event 0xdde58acf741e6f5e57788ee0f64a68576e5eadf169e78251f32824653dd308bf.
//
// Solidity: event NewCompAddress(address oldRewardTokenAddr, address newRewardTokenAddr)
func (_Abicomptroller *AbicomptrollerFilterer) WatchNewCompAddress(opts *bind.WatchOpts, sink chan<- *AbicomptrollerNewCompAddress) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "NewCompAddress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerNewCompAddress)
				if err := _Abicomptroller.contract.UnpackLog(event, "NewCompAddress", log); err != nil {
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

// ParseNewCompAddress is a log parse operation binding the contract event 0xdde58acf741e6f5e57788ee0f64a68576e5eadf169e78251f32824653dd308bf.
//
// Solidity: event NewCompAddress(address oldRewardTokenAddr, address newRewardTokenAddr)
func (_Abicomptroller *AbicomptrollerFilterer) ParseNewCompAddress(log types.Log) (*AbicomptrollerNewCompAddress, error) {
	event := new(AbicomptrollerNewCompAddress)
	if err := _Abicomptroller.contract.UnpackLog(event, "NewCompAddress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerNewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the Abicomptroller contract.
type AbicomptrollerNewLiquidationIncentiveIterator struct {
	Event *AbicomptrollerNewLiquidationIncentive // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerNewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerNewLiquidationIncentive)
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
		it.Event = new(AbicomptrollerNewLiquidationIncentive)
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
func (it *AbicomptrollerNewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerNewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerNewLiquidationIncentive represents a NewLiquidationIncentive event raised by the Abicomptroller contract.
type AbicomptrollerNewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Abicomptroller *AbicomptrollerFilterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*AbicomptrollerNewLiquidationIncentiveIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerNewLiquidationIncentiveIterator{contract: _Abicomptroller.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Abicomptroller *AbicomptrollerFilterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *AbicomptrollerNewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerNewLiquidationIncentive)
				if err := _Abicomptroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
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

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Abicomptroller *AbicomptrollerFilterer) ParseNewLiquidationIncentive(log types.Log) (*AbicomptrollerNewLiquidationIncentive, error) {
	event := new(AbicomptrollerNewLiquidationIncentive)
	if err := _Abicomptroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerNewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the Abicomptroller contract.
type AbicomptrollerNewPauseGuardianIterator struct {
	Event *AbicomptrollerNewPauseGuardian // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerNewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerNewPauseGuardian)
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
		it.Event = new(AbicomptrollerNewPauseGuardian)
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
func (it *AbicomptrollerNewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerNewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerNewPauseGuardian represents a NewPauseGuardian event raised by the Abicomptroller contract.
type AbicomptrollerNewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Abicomptroller *AbicomptrollerFilterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*AbicomptrollerNewPauseGuardianIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerNewPauseGuardianIterator{contract: _Abicomptroller.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Abicomptroller *AbicomptrollerFilterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *AbicomptrollerNewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerNewPauseGuardian)
				if err := _Abicomptroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
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

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Abicomptroller *AbicomptrollerFilterer) ParseNewPauseGuardian(log types.Log) (*AbicomptrollerNewPauseGuardian, error) {
	event := new(AbicomptrollerNewPauseGuardian)
	if err := _Abicomptroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbicomptrollerNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the Abicomptroller contract.
type AbicomptrollerNewPriceOracleIterator struct {
	Event *AbicomptrollerNewPriceOracle // Event containing the contract specifics and raw log

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
func (it *AbicomptrollerNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbicomptrollerNewPriceOracle)
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
		it.Event = new(AbicomptrollerNewPriceOracle)
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
func (it *AbicomptrollerNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbicomptrollerNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbicomptrollerNewPriceOracle represents a NewPriceOracle event raised by the Abicomptroller contract.
type AbicomptrollerNewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Abicomptroller *AbicomptrollerFilterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*AbicomptrollerNewPriceOracleIterator, error) {

	logs, sub, err := _Abicomptroller.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &AbicomptrollerNewPriceOracleIterator{contract: _Abicomptroller.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Abicomptroller *AbicomptrollerFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *AbicomptrollerNewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _Abicomptroller.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbicomptrollerNewPriceOracle)
				if err := _Abicomptroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
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

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Abicomptroller *AbicomptrollerFilterer) ParseNewPriceOracle(log types.Log) (*AbicomptrollerNewPriceOracle, error) {
	event := new(AbicomptrollerNewPriceOracle)
	if err := _Abicomptroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
