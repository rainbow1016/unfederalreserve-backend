package deployer

import (
	"errors"
	"fmt"

	abitimelock "github.com/UnFederalReserve/compound-server-api/abi/timelock"
)

func (d *Deployer) deployTimelock() error {
	d.log.Info().Msg("Deploying Timelock")

	delay, ok := parseBigInt(d.state.Timelock.Delay)
	if !ok {
		return errors.New("failed to parse Delay number")
	}

	addr, tx, _, err := abitimelock.DeployAbitimelock(d.signer, d.eth, d.publicKey, delay)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	d.state.Timelock.Contract.Address = addr
	d.state.Timelock.Contract.IsDeployed = true

	return nil
}
