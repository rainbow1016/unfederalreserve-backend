package deployer

import (
	"errors"
	"fmt"

	abicether "github.com/UnFederalReserve/compound-server-api/abi/cether"
	"github.com/ethereum/go-ethereum/common"
)

func (d *Deployer) interestRateModel(t string) common.Address {
	switch t {
	case "jumpV2":
		return d.state.JumpRateModelV2.Contract.Address
	default:
		return d.state.InterestRateModel.Contract.Address
	}
}

func (d *Deployer) deployEthToken() error {
	d.log.Info().Msg("Deploying ETH Token")

	initialExchangeRateMantissa, ok := parseBigInt(d.state.EthToken.InitialExchangeRateMantissa)
	if !ok {
		return errors.New("failed to parse InitialAmount number")
	}

	addr, tx, _, err := abicether.DeployAbicether(
		d.signer, d.eth,
		d.state.Unitroller.Contract.Address, d.interestRateModel(d.state.EthToken.InterestRateModelType), initialExchangeRateMantissa,
		d.state.EthToken.Name, d.state.EthToken.Symbol, d.state.EthToken.Decimals, d.publicKey,
	)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	d.state.EthToken.Contract.Address = addr
	d.state.EthToken.Contract.IsDeployed = true

	return nil
}
