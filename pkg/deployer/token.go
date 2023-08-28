package deployer

import (
	"errors"
	"fmt"

	abicerc20delegate "github.com/UnFederalReserve/compound-server-api/abi/cerc20delegate"
	abicerc20delegator "github.com/UnFederalReserve/compound-server-api/abi/cerc20delegator"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (d *Deployer) deployTokenDelegate(token *Token) error {
	d.log.Info().Str("token", token.Symbol).Msg("Deploying Token Delegate")

	addr, tx, _, err := abicerc20delegate.DeployAbicerc20delegate(
		d.signer, d.eth,
	)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	token.Delegate.Contract.Address = addr
	token.Delegate.Contract.IsDeployed = true

	return nil
}

func (d *Deployer) deployTokenDelegator(token *Token) error {
	d.log.Info().Str("token", token.Symbol).Msg("Deploying Token Delegator")

	initialExchangeRateMantissa, ok := parseBigInt(token.InitialExchangeRateMantissa)
	if !ok {
		return errors.New("failed to parse InitialAmount number")
	}
	signer := &bind.TransactOpts{
		From:   d.signer.From,
		Signer: d.signer.Signer,
		// override gas limit, because the estimator is failing quite often
		GasLimit: 2200000,
		//GasPrice: big.NewInt(10 * 1e9),
	}
	addr, tx, _, err := abicerc20delegator.DeployAbicerc20delegator(
		signer, d.eth,
		token.Underlying.Contract.Address, d.state.Unitroller.Contract.Address, d.interestRateModel(token.InterestRateModelType), initialExchangeRateMantissa,
		token.Name, token.Symbol, token.Decimals, d.publicKey, token.Delegate.Contract.Address, []byte{},
	)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	token.Delegator.Address = addr
	token.Delegator.IsDeployed = true
	token.Delegate.IsSet = true

	return nil
}

func (d *Deployer) setTokenDelegate(token *Token) error {
	d.log.Info().Str("token", token.Symbol).Msg("Setting token delegate")

	delegator, err := abicerc20delegator.NewAbicerc20delegatorTransactor(token.Delegator.Address, d.eth)
	if err != nil {
		return fmt.Errorf("get delegator abi: %w", err)
	}

	{
		tx, err := delegator.SetImplementation(d.signer, token.Delegate.Contract.Address, true, []byte{})
		if err != nil {
			return fmt.Errorf("send SetImplementation tx: %w", err)
		}
		if err := d.waitForTx(tx.Hash()); err != nil {
			return fmt.Errorf("wait for SetImplementation tx: %w", err)
		}
	}

	token.Delegate.IsSet = true

	return nil
}
