package deployer

import (
	"errors"
	"fmt"

	abifaucettoken "github.com/UnFederalReserve/compound-server-api/abi/faucettoken"
)

func (d *Deployer) deployUnderlyingToken(token *UnderlyingToken) error {
	d.log.Info().Msg("Deploying Underlying token")

	initialAmount, ok := parseBigInt(token.InitialAmount)
	if !ok {
		return errors.New("failed to parse InitialAmount number")
	}

	addr, tx, _, err := abifaucettoken.DeployAbictoken(d.signer, d.eth, initialAmount, token.Name, token.Decimals, token.Symbol)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	token.Contract.Address = addr
	token.Contract.IsDeployed = true

	return nil
}
