package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/big"
	"os"

	"github.com/UnFederalReserve/compound-server-api/pkg/deployer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

var (
	stateFile = flag.String("state-file", "", "JSNON file with state")
)

type config struct {
	EthereumNode string `required:"true" split_words:"true"`
	PrivateKey   string `required:"true" split_words:"true"`
	NetworkId    int64  `required:"true" split_words:"true"`
}

func main() {
	flag.Parse()
	if *stateFile == "" {
		flag.Usage()
		return
	}
	log := zerolog.New(os.Stdout).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout})
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		_ = envconfig.Usage("", &cfg)
		log.Fatal().Err(err).Msg("Failed to parse config")
	}
	eth, err := ethclient.Dial(cfg.EthereumNode)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to ethereum node")
	}
	key, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to decode private key")
	}

	signer, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(cfg.NetworkId))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create transaction signer")
	}
	//signer.GasPrice = big.NewInt(108 * 1e9)
	state, err := readState(*stateFile)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to read state")
	}
	d := deployer.New(log, eth, signer, state, crypto.PubkeyToAddress(key.PublicKey))
	defer func() {
		if err := saveState(*stateFile, d.State()); err != nil {
			log.Error().Err(err).Msg("Failed to save state")
		}
	}()
	if err := d.Deploy(); err != nil {
		log.Error().Err(err).Msg("Failed to deploy")
		return
	}
	if err := saveState(*stateFile, d.State()); err != nil {
		log.Error().Err(err).Msg("Failed to save state")
	}
	log.Info().Msg("Done")
}

func readState(filename string) (deployer.State, error) {
	var state deployer.State
	f, err := os.Open(filename)
	if err != nil {
		return state, fmt.Errorf("open file: %w", err)
	}
	defer f.Close()
	err = json.NewDecoder(f).Decode(&state)
	if err != nil {
		return state, fmt.Errorf("decode: %w", err)
	}
	return state, nil
}

func saveState(filename string, state deployer.State) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC|os.O_SYNC, 0644)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer f.Close()
	err = json.NewEncoder(f).Encode(state)
	return err
}
