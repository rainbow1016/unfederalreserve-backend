package main

import (
	"fmt"
	"math/big"
	"net/http"
	"os"

	"github.com/UnFederalReserve/compound-server-api/internal/api"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

type config struct {
	Host               string `default:""`
	Port               string `default:"8085"`
	ComptrollerAddress string `required:"true" split_words:"true"`
	EthereumNode       string `required:"true" split_words:"true"`
	PrivateKey         string `required:"true" split_words:"true"`
	AuthHost           string `required:"true" split_words:"true"`
	NetworkId          int64  `required:"true" split_words:"true"`
	ApiKey             string `split_words:"true"`
}

func main() {
	fmt.Println("starting...")
	log := zerolog.New(os.Stdout).With().Timestamp().Logger()
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
	signer, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(cfg.NetworkId))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create transaction signer")
	}
	server, err := api.New(log, cfg.AuthHost, cfg.ComptrollerAddress, cfg.ApiKey, eth, signer)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create api server")
	}
	router := mux.NewRouter()
	router.Use(allowCORS)
	server.AttachRoutes(router)
	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler: router,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to listen")
	}

}

func allowCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
		next.ServeHTTP(w, r)
	})
}
