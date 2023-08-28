package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"strings"
	"time"

	abicomptroller "github.com/UnFederalReserve/compound-server-api/abi/comptroller"
	abignosissafe "github.com/UnFederalReserve/compound-server-api/abi/gnosissafe"
	abioracle "github.com/UnFederalReserve/compound-server-api/abi/oracle"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

type config struct {
	ComptrollerAddress    string `required:"true" split_words:"true"`
	EthereumNode          string `required:"true" split_words:"true"`
	PrivateKey            string `required:"true" split_words:"true"`
	NetworkId             int64  `required:"true" split_words:"true"`
	GnosisSafeAddress     string `required:"true" split_words:"true"`
	GnosisRelayHost       string `required:"true" split_words:"true"`
	GnosisTransactionHost string `required:"true" split_words:"true"`
	TokenAddress          string `required:"true" split_words:"true"`
}

func main() {
	log := zerolog.New(os.Stdout).With().Timestamp().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout})
	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		_ = envconfig.Usage("", &cfg)
		log.Fatal().Err(err).Msg("failed to parse config")
	}
	eth, err := ethclient.Dial(cfg.EthereumNode)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to ethereum node")
	}
	key, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to decode private key")
	}

	oracleAddr, err := getOracleAddr(common.HexToAddress(cfg.ComptrollerAddress), eth)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get oracle address")
	}

	ethPrice, err := getEthPrice(oracleAddr, eth)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get eth price")
	}

	ersdlEthPrice, err := getErsdlEthPrice()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get ersdl-eth price")
	}

	price, _ := new(big.Float).Mul(ersdlEthPrice, ethPrice).Int(nil)
	log.Info().Str("price", price.String()).Msg("calculated ersdl/usd price")

	txData, err := generateTransactionData(common.HexToAddress(cfg.TokenAddress), price)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to generate transaction data")
	}

	gas, nonce, err := estimateGnosisTx(cfg.GnosisRelayHost, common.HexToAddress(cfg.GnosisSafeAddress), oracleAddr, txData)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to estimate gas limit")
	}

	sender := crypto.PubkeyToAddress(key.PublicKey)
	txHash, err := getContractTransactionHash(common.HexToAddress(cfg.GnosisSafeAddress), eth, oracleAddr, txData, gas, nonce, sender)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to generate contract transaction hash")
	}

	log.Info().Str("hash", txHash.String()).Msg("generated contract transaction hash")

	sig, err := crypto.Sign(txHash.Bytes(), key)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to sign contract transaction hash")
	}
	sig[64] = 0x1c
	log.Info().Str("sign", common.Bytes2Hex(sig)).Msg("got signature")
	err = sendGnosisTx(cfg.GnosisTransactionHost, common.HexToAddress(cfg.GnosisSafeAddress), oracleAddr, txData, gas, nonce, txHash, sender, sig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to send multisign tx")
	}
}

func getErsdlEthPrice() (*big.Float, error) {
	httpClient := http.Client{Timeout: time.Minute}
	res, err := httpClient.Post("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2", "application/json", strings.NewReader(`{"query":"{token(id: \"0x5218e472cfcfe0b64a064f055b43b4cdc9efd3a6\"){derivedETH}}","variables":null}`))
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer res.Body.Close()
	var uniswapData struct {
		Data struct {
			Token struct {
				DerivedETH string
			}
		}
	}
	err = json.NewDecoder(res.Body).Decode(&uniswapData)
	if err != nil {
		return nil, fmt.Errorf("decode response body: %w", err)
	}
	ersdlEthPrice, _, err := new(big.Float).Parse(uniswapData.Data.Token.DerivedETH, 10)
	if err != nil {
		return nil, fmt.Errorf("parse price: %w", err)
	}
	return ersdlEthPrice, nil
}

func getOracleAddr(comptrollerAddr common.Address, client *ethclient.Client) (common.Address, error) {
	return common.HexToAddress("0x8c703994C32163872221EAB065B00e0Fa7d6f44A"), nil
	comptroller, err := abicomptroller.NewAbicomptroller(comptrollerAddr, client)
	if err != nil {
		return common.Address{}, fmt.Errorf("get comptroller abi caller: %w", err)
	}
	oracleAddr, err := comptroller.Oracle(nil)
	if err != nil {
		return common.Address{}, fmt.Errorf("get oracle addr: %w", err)
	}
	return oracleAddr, nil
}

func getEthPrice(oracleAddr common.Address, client *ethclient.Client) (*big.Float, error) {
	oracle, err := abioracle.NewAbioracleCaller(oracleAddr, client)
	if err != nil {
		return nil, fmt.Errorf("get oracle abi caller: %w", err)
	}
	price, err := oracle.AssetPrices(nil, common.Address{})
	if err != nil {
		return nil, fmt.Errorf("get asset price: %w", err)
	}
	return new(big.Float).SetInt(price), nil
}

func generateTransactionData(ctokenAddr common.Address, price *big.Int) ([]byte, error) {
	oracleABI, err := abi.JSON(strings.NewReader(abioracle.AbioracleABI))
	if err != nil {
		return nil, fmt.Errorf("read oracle ABI: %w", err)
	}
	input, err := oracleABI.Pack("setUnderlyingPrice", ctokenAddr, price)
	if err != nil {
		return nil, fmt.Errorf("generate transaction input: %w", err)
	}
	return input, nil
}

func estimateGnosisTx(apiHost string, safeAddr, oracleAddr common.Address, txInputData []byte) (nonce *big.Int, gasLimit *big.Int, err error) {
	httpClient := http.Client{Timeout: time.Minute}
	body := new(bytes.Buffer)

	txData := struct {
		To        string      `json:"to"`
		Value     int         `json:"value"`
		Data      string      `json:"data"`
		Operation int         `json:"operation"`
		GasToken  interface{} `json:"gasToken"`
	}{
		To:        oracleAddr.String(),
		Value:     0,
		Data:      hexutil.Bytes(txInputData).String(),
		Operation: 0,
		GasToken:  nil,
	}
	err = json.NewEncoder(body).Encode(txData)
	if err != nil {
		return nil, nil, fmt.Errorf("encode transaction data: %w", err)
	}
	res, err := httpClient.Post(apiHost+"/api/v2/safes/"+safeAddr.String()+"/transactions/estimate/", "application/json", body)
	if err != nil {
		return nil, nil, fmt.Errorf("send request: %w", err)
	}
	defer res.Body.Close()
	var response struct {
		SafeTxGas     string `json:"safeTxGas"`
		LastUsedNonce int64  `json:"lastUsedNonce"`
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, nil, fmt.Errorf("decode response body: %w", err)
	}
	gasLimit, ok := new(big.Int).SetString(response.SafeTxGas, 10)
	if !ok {
		return nil, nil, fmt.Errorf("parse gas limit: %s", response.SafeTxGas)
	}

	return gasLimit, big.NewInt(response.LastUsedNonce + 1), nil
}

func getContractTransactionHash(gnosisSafeAddr common.Address, client *ethclient.Client, oracleAddr common.Address, input []byte, gas *big.Int, nonce *big.Int, sender common.Address) (common.Hash, error) {
	gnosisSafe, err := abignosissafe.NewAbignosissafeCaller(gnosisSafeAddr, client)
	if err != nil {
		return common.Hash{}, fmt.Errorf("read gnosis safe abi: %w", err)
	}
	return gnosisSafe.GetTransactionHash(&bind.CallOpts{From: sender}, oracleAddr, big.NewInt(0), input, 0, gas, big.NewInt(0), big.NewInt(0), common.Address{}, common.Address{}, nonce)
}

func sendGnosisTx(apiHost string, safeAddr, oracleAddr common.Address, txInputData []byte, gas, nonce *big.Int, contractTransactionHash common.Hash, sender common.Address, sig []byte) error {
	httpClient := http.Client{Timeout: time.Minute}
	body := new(bytes.Buffer)

	txData := struct {
		To                      string      `json:"to"`
		Value                   string      `json:"value"`
		Data                    interface{} `json:"data"`
		Operation               int         `json:"operation"`
		GasToken                interface{} `json:"gasToken"`
		SafeTxGas               int64       `json:"safeTxGas"`
		BaseGas                 int         `json:"baseGas"`
		GasPrice                string      `json:"gasPrice"`
		RefundReceiver          interface{} `json:"refundReceiver"`
		Nonce                   int64       `json:"nonce"`
		ContractTransactionHash string      `json:"contractTransactionHash"`
		Sender                  string      `json:"sender"`
		Signature               string      `json:"signature"`
		Origin                  interface{} `json:"origin"`
		TransactionHash         interface{} `json:"transactionHash"`
	}{
		To:                      oracleAddr.String(),
		Value:                   "0",
		GasPrice:                "0",
		Data:                    hexutil.Bytes(txInputData).String(),
		Operation:               0,
		GasToken:                common.Address{},
		RefundReceiver:          common.Address{},
		SafeTxGas:               gas.Int64(),
		Nonce:                   nonce.Int64(),
		ContractTransactionHash: contractTransactionHash.String(),
		Sender:                  sender.String(),
		Signature:               common.Bytes2Hex(sig),
		Origin:                  nil,
		TransactionHash:         nil,
	}
	err := json.NewEncoder(body).Encode(txData)
	if err != nil {
		return fmt.Errorf("encode transaction data: %w", err)
	}
	res, err := httpClient.Post(apiHost+"/api/v1/safes/"+safeAddr.String()+"/multisig-transactions/", "application/json", body)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer res.Body.Close()
	var response struct {
		Errors []string `json:"nonFieldErrors"`
	}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}
	if len(response.Errors) > 0 {
		return fmt.Errorf("errors occured: %+v", response.Errors)
	}
	return nil
}
