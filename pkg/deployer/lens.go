package deployer

import (
	"fmt"

	abilens "github.com/UnFederalReserve/compound-server-api/abi/lens"
	abiliquiditylens "github.com/UnFederalReserve/compound-server-api/abi/liquiditylens"
)

func (d *Deployer) deployLens() error {
	d.log.Info().Msg("Deploying Lens")
	addr, tx, _, err := abilens.DeployAbilens(d.signer, d.eth)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	d.state.Lens.Address = addr
	d.state.Lens.IsDeployed = true

	return nil
}

func (d *Deployer) deployLiquidityLens() error {
	d.log.Info().Msg("Deploying LiquidityLens")
	addr, tx, _, err := abiliquiditylens.DeployAbiliquiditylens(d.signer, d.eth)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	d.state.LiquidityLens.Address = addr
	d.state.LiquidityLens.IsDeployed = true

	return nil
}
