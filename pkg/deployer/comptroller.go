package deployer

import (
	"errors"
	"fmt"

	abicomptroller "github.com/UnFederalReserve/compound-server-api/abi/comptroller"
	abiunitroller "github.com/UnFederalReserve/compound-server-api/abi/unitroller"
	"github.com/ethereum/go-ethereum/common"
)

func (d *Deployer) deployUnitroller() error {
	d.log.Info().Msg("Deploying Unitroller")
	addr, tx, _, err := abiunitroller.DeployAbiunitroller(d.signer, d.eth)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	d.state.Unitroller.Contract.Address = addr
	d.state.Unitroller.Contract.IsDeployed = true

	return nil
}

func (d *Deployer) deployComptroller() error {
	d.log.Info().Msg("Deploying Comptroller")
	addr, tx, _, err := abicomptroller.DeployAbicomptroller(d.signer, d.eth)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	d.state.Comptroller.Contract.Address = addr
	d.state.Comptroller.Contract.IsDeployed = true

	return nil
}

func (d *Deployer) setUnitrollerImplementation() error {
	d.log.Info().Msg("Setting comptroller implementation")

	if !d.state.Unitroller.Contract.IsDeployed {
		return errors.New("unitroller not deployed")
	}
	if !d.state.Comptroller.Contract.IsDeployed {
		return errors.New("comptroller not deployed")
	}

	unitroller, err := abiunitroller.NewAbiunitrollerTransactor(d.state.Unitroller.Contract.Address, d.eth)
	if err != nil {
		return fmt.Errorf("get unitroller abi: %w", err)
	}

	comptroller, err := abicomptroller.NewAbicomptrollerTransactor(d.state.Comptroller.Contract.Address, d.eth)
	if err != nil {
		return fmt.Errorf("get comptroller abi: %w", err)
	}

	{
		tx, err := unitroller.SetPendingImplementation(d.signer, d.state.Comptroller.Contract.Address)
		if err != nil {
			return fmt.Errorf("send SetPendingImplementation tx: %w", err)
		}
		if err := d.waitForTx(tx.Hash()); err != nil {
			return fmt.Errorf("wait for SetPendingImplementation tx: %w", err)
		}
	}

	{
		tx, err := comptroller.Become(d.signer, d.state.Unitroller.Contract.Address)
		if err != nil {
			return fmt.Errorf("send Become tx: %w", err)
		}
		if err := d.waitForTx(tx.Hash()); err != nil {
			return fmt.Errorf("wait for Become tx: %w", err)
		}
	}

	d.state.Comptroller.IsSet = true

	return nil
}

func (d *Deployer) setCollateralFactor(symbol string, addr common.Address, mantissa string) error {
	d.log.Info().Str("token", symbol).Str("addr", addr.String()).Str("mantissa", mantissa).Msg("Setting collateral factor")

	comptroller, err := abicomptroller.NewAbicomptrollerTransactor(d.state.Unitroller.Contract.Address, d.eth)
	if err != nil {
		return fmt.Errorf("get comptroller abi: %w", err)
	}

	collateralFactor, ok := parseBigInt(mantissa)
	if !ok {
		return fmt.Errorf("parse CollateralFactor number: %w", err)
	}

	tx, err := comptroller.SetCollateralFactor(d.signer, addr, collateralFactor)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	return nil
}

func (d *Deployer) setSupportMarket(symbol string, addr common.Address) error {
	d.log.Info().Str("token", symbol).Msg("Setting supported market")

	comptroller, err := abicomptroller.NewAbicomptrollerTransactor(d.state.Unitroller.Contract.Address, d.eth)
	if err != nil {
		return fmt.Errorf("get comptroller abi: %w", err)
	}

	tx, err := comptroller.SupportMarket(d.signer, addr)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	return nil
}

func (d *Deployer) setRewardSpeed(symbol string, addr common.Address, rewardSpeed string) error {
	d.log.Info().Str("token", symbol).Msg("Setting reward speed")

	comptroller, err := abicomptroller.NewAbicomptrollerTransactor(d.state.Unitroller.Contract.Address, d.eth)
	if err != nil {
		return fmt.Errorf("get comptroller abi: %w", err)
	}

	speed, ok := parseBigInt(rewardSpeed)
	if !ok {
		return nil
	}

	tx, err := comptroller.SetCompSpeed(d.signer, addr, speed)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	return nil
}

func (d *Deployer) setRewardToken() error {
	d.log.Info().Msg("Setting reward token")

	comptroller, err := abicomptroller.NewAbicomptrollerTransactor(d.state.Unitroller.Contract.Address, d.eth)
	if err != nil {
		return fmt.Errorf("get comptroller abi: %w", err)
	}

	tx, err := comptroller.SetCompAddress(d.signer, d.state.RewardToken.Contract.Address)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}
	d.state.RewardToken.IsSet = true
	return nil
}

func (d *Deployer) setCloseFactor() error {
	d.log.Info().Msg("Setting CloseFactor token")

	comptroller, err := abicomptroller.NewAbicomptrollerTransactor(d.state.Unitroller.Contract.Address, d.eth)
	if err != nil {
		return fmt.Errorf("get comptroller abi: %w", err)
	}
	closeFactor, ok := parseBigInt(d.state.Unitroller.CloseFactorMantissa.Value)
	if !ok {
		return fmt.Errorf("parse CloseFactor number: %w", err)
	}
	tx, err := comptroller.SetCloseFactor(d.signer, closeFactor)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}
	d.state.Unitroller.CloseFactorMantissa.IsSet = true
	return nil

}

func (d *Deployer) setLiquidationIncentive() error {
	d.log.Info().Msg("Setting LiquidationIncentive token")

	comptroller, err := abicomptroller.NewAbicomptrollerTransactor(d.state.Unitroller.Contract.Address, d.eth)
	if err != nil {
		return fmt.Errorf("get comptroller abi: %w", err)
	}
	liquidationIncentive, ok := parseBigInt(d.state.Unitroller.LiquidationIncentiveMantissa.Value)
	if !ok {
		return fmt.Errorf("parse LiquidationIncentive number: %w", err)
	}
	tx, err := comptroller.SetLiquidationIncentive(d.signer, liquidationIncentive)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}
	d.state.Unitroller.LiquidationIncentiveMantissa.IsSet = true
	return nil
}
