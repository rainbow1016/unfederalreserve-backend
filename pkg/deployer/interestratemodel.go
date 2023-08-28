package deployer

import (
	"errors"
	"fmt"
	"math/big"

	abiinterestratemodel "github.com/UnFederalReserve/compound-server-api/abi/interestratemodel"
	abijumpratemodelv2 "github.com/UnFederalReserve/compound-server-api/abi/jumpratemodelv2"
	"github.com/ethereum/go-ethereum/common"
)

var zeroAddr = common.Address{}

func (d *Deployer) deployInterestRateModel() error {
	d.log.Info().Msg("Deploying InterestRateModel")

	baseRatePerYear, ok := new(big.Int).SetString(d.state.InterestRateModel.BaseRatePerYear, 10)
	if !ok {
		return errors.New("failed to parse BaseRatePerYear number")
	}

	multiplierPerYear, ok := new(big.Int).SetString(d.state.InterestRateModel.MultiplierPerYear, 10)
	if !ok {
		return errors.New("failed to parse MultiplierPerYear number")
	}

	addr, tx, _, err := abiinterestratemodel.DeployAbiinterestratemodel(d.signer, d.eth, baseRatePerYear, multiplierPerYear)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	d.state.InterestRateModel.Contract.Address = addr
	d.state.InterestRateModel.Contract.IsDeployed = true

	return nil
}

func (d *Deployer) deployJumpRateModelV2() error {
	d.log.Info().Msg("Deploying JumpRateModelV2")

	baseRatePerYear, ok := new(big.Int).SetString(d.state.JumpRateModelV2.BaseRatePerYear, 10)
	if !ok {
		return errors.New("failed to parse BaseRatePerYear number")
	}

	multiplierPerYear, ok := new(big.Int).SetString(d.state.JumpRateModelV2.MultiplierPerYear, 10)
	if !ok {
		return errors.New("failed to parse MultiplierPerYear number")
	}

	jumpMultiplierPerYear, ok := new(big.Int).SetString(d.state.JumpRateModelV2.JumpMultiplierPerYear, 10)
	if !ok {
		return errors.New("failed to parse JumpMultiplierPerYear number")
	}

	kink, ok := new(big.Int).SetString(d.state.JumpRateModelV2.Kink, 10)
	if !ok {
		return errors.New("failed to parse Kink number")
	}
	admin := d.publicKey
	if d.state.JumpRateModelV2.Admin != zeroAddr {
		admin = d.state.JumpRateModelV2.Admin
	}
	addr, tx, _, err := abijumpratemodelv2.DeployAbijumpratemodelv2(d.signer, d.eth, baseRatePerYear, multiplierPerYear, jumpMultiplierPerYear, kink, admin)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	d.state.JumpRateModelV2.Contract.Address = addr
	d.state.JumpRateModelV2.Contract.IsDeployed = true

	return nil
}
