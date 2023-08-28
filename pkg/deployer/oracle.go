package deployer

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	abichainlink "github.com/UnFederalReserve/compound-server-api/abi/chainlink"
	abicomptroller "github.com/UnFederalReserve/compound-server-api/abi/comptroller"
	abioracle "github.com/UnFederalReserve/compound-server-api/abi/oracle"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

func (d *Deployer) deployPriceOracle() error {
	d.log.Info().Msg("Deploying Oracle")
	addr, tx, _, err := abioracle.DeployAbioracle(d.signer, d.eth)
	if err != nil {
		return fmt.Errorf("send tx: %w", err)
	}

	if err := d.waitForTx(tx.Hash()); err != nil {
		return fmt.Errorf("wait for tx: %w", err)
	}

	d.state.Oracle.Contract.Address = addr
	d.state.Oracle.Contract.IsDeployed = true

	return nil
}

func (d *Deployer) setPriceOracle() error {
	d.log.Info().Msg("Setting price oracle")

	if !d.state.Comptroller.IsSet {
		return errors.New("comptroller implementation is not set")
	}

	comptroller, err := abicomptroller.NewAbicomptrollerTransactor(d.state.Unitroller.Contract.Address, d.eth)
	if err != nil {
		return fmt.Errorf("get comptroller abi: %w", err)
	}

	{
		tx, err := comptroller.SetPriceOracle(d.signer, d.state.Oracle.Contract.Address)
		if err != nil {
			return fmt.Errorf("send tx: %w", err)
		}
		if err := d.waitForTx(tx.Hash()); err != nil {
			return fmt.Errorf("wait for tx: %w", err)
		}
	}

	d.state.Oracle.IsSet = true

	return nil
}

func (d *Deployer) setTokenPrice(symbol string, addr common.Address, price Price, underlyingDecimals uint8) error {
	d.log.Info().Str("token", symbol).Msg("Setting token price")

	if !d.state.Oracle.IsSet {
		return errors.New("oracle is not set")
	}
	oracle, err := abioracle.NewAbioracleTransactor(d.state.Oracle.Contract.Address, d.eth)
	if err != nil {
		return fmt.Errorf("get oracle abi: %w", err)
	}

	if price.StaticPrice != "" {
		price, ok := parseBigInt(price.StaticPrice)
		if !ok {
			return fmt.Errorf("parse StaticPrice number: %w", err)
		}
		var tx *types.Transaction
		if symbol == "unETH" {
			tx, err = oracle.SetDirectPrice(d.signer, common.Address{}, price)
		} else {
			tx, err = oracle.SetUnderlyingPrice(d.signer, addr, price)
		}
		if err := d.waitForTx(tx.Hash()); err != nil {
			return fmt.Errorf("wait for tx: %w", err)
		}
	} else {
		feed, err := abichainlink.NewAbichainlinkCaller(price.Feed, d.eth)
		if err != nil {
			return fmt.Errorf("get chainlink feed abi: %w", err)
		}
		feedDecimals, err := feed.Decimals(nil)
		if err != nil {
			return fmt.Errorf("get feed decimals: %w", err)
		}

		multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18+18-underlyingDecimals-feedDecimals)), nil)
		var tx *types.Transaction
		if symbol == "unETH" {
			tx, err = oracle.SetDirectChainlinkFeed(d.signer, common.Address{}, price.Feed, multiplier)
			if err != nil {
				return fmt.Errorf("send tx: %w", err)
			}
		} else {
			tx, err = oracle.SetUnderlyingChainlinkFeed(d.signer, addr, price.Feed, multiplier)
			if err != nil {

				oracleABI, e := abi.JSON(strings.NewReader(abioracle.AbioracleABI))
				if e != nil {
					return fmt.Errorf("read oracle abi while sending tx: %w (%v)", e, err)
				}
				input, e := oracleABI.Pack("setUnderlyingChainlinkFeed", addr, price.Feed, multiplier)
				if e != nil {
					return fmt.Errorf("generate tx input while sending tx: %w (%v)", e, err)
				}
				return fmt.Errorf("send tx: %w (%s)", err, hexutil.Encode(input))
			}
		}
		if err := d.waitForTx(tx.Hash()); err != nil {
			return fmt.Errorf("wait for tx: %w", err)
		}
	}

	return nil
}
