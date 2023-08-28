package deployer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog"
)

type Deployer struct {
	log       zerolog.Logger
	eth       *ethclient.Client
	signer    *bind.TransactOpts
	state     State
	publicKey common.Address
}

func New(log zerolog.Logger, eth *ethclient.Client, signer *bind.TransactOpts, state State, publicKey common.Address) *Deployer {
	d := &Deployer{
		log:       log,
		eth:       eth,
		signer:    signer,
		state:     state,
		publicKey: publicKey,
	}
	return d
}

func (d *Deployer) Deploy() error {

	if !d.state.Lens.IsDeployed {
		if err := d.deployLens(); err != nil {
			return fmt.Errorf("deploy lens: %w", err)
		}
	}

	if !d.state.LiquidityLens.IsDeployed {
		if err := d.deployLiquidityLens(); err != nil {
			return fmt.Errorf("deploy liquidity lens: %w", err)
		}
	}

	if !d.state.Unitroller.Contract.IsDeployed {
		if err := d.deployUnitroller(); err != nil {
			return fmt.Errorf("deploy unitroller: %w", err)
		}
	}

	if !d.state.Comptroller.Contract.IsDeployed {
		if err := d.deployComptroller(); err != nil {
			return fmt.Errorf("deploy comptroller: %w", err)
		}
	}

	if !d.state.Comptroller.IsSet {
		if err := d.setUnitrollerImplementation(); err != nil {
			return fmt.Errorf("set unitroller implementation: %w", err)
		}
	}

	if !d.state.Oracle.Contract.IsDeployed {
		if err := d.deployPriceOracle(); err != nil {
			return fmt.Errorf("deploy price oracle: %w", err)
		}
	}

	if !d.state.Oracle.IsSet {
		if err := d.setPriceOracle(); err != nil {
			return fmt.Errorf("set price oracle: %w", err)
		}
	}

	if !d.state.Unitroller.CloseFactorMantissa.IsSet {
		if err := d.setCloseFactor(); err != nil {
			return fmt.Errorf("set CloseFactor: %w", err)
		}
	}

	if !d.state.Unitroller.LiquidationIncentiveMantissa.IsSet {
		if err := d.setLiquidationIncentive(); err != nil {
			return fmt.Errorf("set LiquidationIncentive: %w", err)
		}
	}

	if !d.state.InterestRateModel.Contract.IsDeployed {
		if err := d.deployInterestRateModel(); err != nil {
			return fmt.Errorf("deploy interest rate model: %w", err)
		}
	}

	if !d.state.JumpRateModelV2.Contract.IsDeployed {
		if err := d.deployJumpRateModelV2(); err != nil {
			return fmt.Errorf("deploy JumpRateModelV2: %w", err)
		}
	}

	{
		// ETH
		if !d.state.EthToken.Contract.IsDeployed {
			if err := d.deployEthToken(); err != nil {
				return fmt.Errorf("deploy ETH token: %w", err)
			}
		}

		if !d.state.EthToken.Price.IsSet {
			if err := d.setTokenPrice(d.state.EthToken.Symbol, d.state.EthToken.Contract.Address, d.state.EthToken.Price, 18); err != nil {
				return fmt.Errorf("set ETH token price: %w", err)
			}
			d.state.EthToken.Price.IsSet = true
		}

		if !d.state.EthToken.IsSetSupportedMarket {
			if err := d.setSupportMarket(d.state.EthToken.Symbol, d.state.EthToken.Contract.Address); err != nil {
				return fmt.Errorf("set ETH supported market: %w", err)
			}
			d.state.EthToken.IsSetSupportedMarket = true
		}

		if !d.state.EthToken.CollateralFactor.IsSet {
			if err := d.setCollateralFactor(d.state.EthToken.Symbol, d.state.EthToken.Contract.Address, d.state.EthToken.CollateralFactor.Mantissa); err != nil {
				return fmt.Errorf("set ETH collateral factor: %w", err)
			}
			d.state.EthToken.CollateralFactor.IsSet = true
		}

		if !d.state.EthToken.IsSetRewardSpeed {
			if err := d.setRewardSpeed(d.state.EthToken.Symbol, d.state.EthToken.Contract.Address, d.state.EthToken.RewardSpeed); err != nil {
				return fmt.Errorf("set ETH reward speed: %w", err)
			}
			d.state.EthToken.IsSetRewardSpeed = true
		}
	}

	for i, token := range d.state.Tokens {
		if !token.Underlying.Contract.IsDeployed {
			if err := d.deployUnderlyingToken(&token.Underlying); err != nil {
				d.state.Tokens[i] = token
				return fmt.Errorf("deploy underlying token %+v: %w", token.Underlying, err)
			}
		}

		if !token.Delegate.Contract.IsDeployed {
			if err := d.deployTokenDelegate(&token); err != nil {
				d.state.Tokens[i] = token
				return fmt.Errorf("deploy delegate token %+v: %w", token, err)
			}
		}
		if token.Delegator.IsDeployed && !token.Delegate.IsSet {
			if err := d.setTokenDelegate(&token); err != nil {
				d.state.Tokens[i] = token
				return fmt.Errorf("set delegate token %+v: %w", token, err)
			}
		} else if !token.Delegator.IsDeployed {
			if err := d.deployTokenDelegator(&token); err != nil {
				d.state.Tokens[i] = token
				return fmt.Errorf("deploy delegator token %+v: %w", token, err)
			}
		}

		if !token.Price.IsSet {
			if err := d.setTokenPrice(token.Symbol, token.Delegator.Address, token.Price, token.Underlying.Decimals); err != nil {
				d.state.Tokens[i] = token
				return fmt.Errorf("set token price for %+v: %w", token, err)
			}
			token.Price.IsSet = true
			d.state.Tokens[i] = token
		}

		if !token.IsSetSupportedMarket {
			if err := d.setSupportMarket(token.Symbol, token.Delegator.Address); err != nil {
				d.state.Tokens[i] = token
				return fmt.Errorf("set supported market for %+v: %w", token, err)
			}
			token.IsSetSupportedMarket = true
			d.state.Tokens[i] = token
		}

		if !token.CollateralFactor.IsSet {
			if err := d.setCollateralFactor(token.Symbol, token.Delegator.Address, token.CollateralFactor.Mantissa); err != nil {
				d.state.Tokens[i] = token
				return fmt.Errorf("set collateral factor for %+v: %w", token, err)
			}
			token.CollateralFactor.IsSet = true
			d.state.Tokens[i] = token
		}

		if !token.IsSetRewardSpeed {
			if err := d.setRewardSpeed(token.Symbol, token.Delegator.Address, token.RewardSpeed); err != nil {
				d.state.Tokens[i] = token
				return fmt.Errorf("set reward speed for %+v: %w", token, err)
			}
			token.IsSetRewardSpeed = true
			d.state.Tokens[i] = token
		}
		d.state.Tokens[i] = token
	}
	if !d.state.RewardToken.IsSet {
		if err := d.setRewardToken(); err != nil {
			return fmt.Errorf("set reward token: %w", err)
		}
	}
	if !d.state.Timelock.Contract.IsDeployed {
		if err := d.deployTimelock(); err != nil {
			return fmt.Errorf("deploy timelock: %w", err)
		}
	}
	return nil
}

func (d *Deployer) State() State {
	return d.state
}

func (d *Deployer) waitForTx(hash common.Hash) error {
	start := time.Now()
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	errs := 0
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()
	for {
		<-ticker.C
		if errs > 20 {
			return errors.New("retry limit exceeded")
		}
		tx, isPending, err := d.eth.TransactionByHash(ctx, hash)
		if !isPending && err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return fmt.Errorf("unrecoverable error: %w", err)
			}
			d.log.Warn().Err(err).Msg("failed to check tx")
			errs++
			continue
		}
		if isPending || tx == nil {
			continue
		}
		d.log.Info().Str("tx", hash.String()).Str("timer", time.Since(start).String()).Msg("transaction mined")
		break
	}
	return nil
}

func parseBigInt(s string) (*big.Int, bool) {
	if s == "" || s == "0" {
		return big.NewInt(0), true
	}
	return new(big.Int).SetString(s, 10)
}
