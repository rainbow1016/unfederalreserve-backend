package deployer

import (
	"github.com/ethereum/go-ethereum/common"
)

type State struct {
	Unitroller        Unitroller
	Comptroller       SettableContract
	Oracle            SettableContract
	Timelock          Timelock
	InterestRateModel InterestRateModel
	JumpRateModelV2   JumpRateModelV2
	Tokens            []Token
	EthToken          EthToken
	RewardToken       SettableContract
	Lens              Contract
	LiquidityLens     Contract
}

type Unitroller struct {
	Contract                     Contract
	LiquidationIncentiveMantissa SettableString
	CloseFactorMantissa          SettableString
}

type Contract struct {
	IsDeployed bool
	Address    common.Address
}

type Timelock struct {
	Contract Contract
	Delay    string
}

type SettableContract struct {
	Contract Contract
	IsSet    bool
}

type SettableString struct {
	Value string
	IsSet bool
}

type EthToken struct {
	Contract                    Contract
	IsSetSupportedMarket        bool
	CollateralFactor            CollateralFactor
	Price                       Price
	InterestRateModel           common.Address
	InterestRateModelType       string
	InitialExchangeRateMantissa string
	Name                        string
	Symbol                      string
	Decimals                    uint8
	Admin                       common.Address
	RewardSpeed                 string
	IsSetRewardSpeed            bool
}

type Token struct {
	Delegator                   Contract
	Delegate                    SettableContract
	IsSetSupportedMarket        bool
	CollateralFactor            CollateralFactor
	Price                       Price
	InterestRateModel           common.Address
	InterestRateModelType       string
	InitialExchangeRateMantissa string
	Name                        string
	Symbol                      string
	Decimals                    uint8
	Admin                       common.Address
	Underlying                  UnderlyingToken
	RewardSpeed                 string
	IsSetRewardSpeed            bool
}

type UnderlyingToken struct {
	Contract      Contract
	InitialAmount string
	Name          string
	Symbol        string
	Decimals      uint8
	Admin         string
}

type InterestRateModel struct {
	Contract          Contract
	BaseRatePerYear   string
	MultiplierPerYear string
}

type JumpRateModelV2 struct {
	Contract              Contract
	BaseRatePerYear       string
	MultiplierPerYear     string
	JumpMultiplierPerYear string
	Kink                  string
	Admin                 common.Address
}

type Price struct {
	Feed        common.Address
	StaticPrice string
	IsSet       bool
}

type CollateralFactor struct {
	Mantissa string
	IsSet    bool
}
