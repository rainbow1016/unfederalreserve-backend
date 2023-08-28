package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	abicomptroller "github.com/UnFederalReserve/compound-server-api/abi/comptroller"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

type Server struct {
	log                zerolog.Logger
	authHost           string
	apiKey             string
	eth                *ethclient.Client
	comptroller        *abicomptroller.Abicomptroller
	comptrollerAddress string
	signer             *bind.TransactOpts
}

func New(log zerolog.Logger, authHost, comptrollerAddr, apiKey string, eth *ethclient.Client, signer *bind.TransactOpts) (*Server, error) {
	s := &Server{
		log:                log,
		authHost:           authHost,
		apiKey:             apiKey,
		eth:                eth,
		signer:             signer,
		comptrollerAddress: comptrollerAddr,
	}
	var err error
	s.comptroller, err = abicomptroller.NewAbicomptroller(common.HexToAddress(comptrollerAddr), eth)
	if err != nil {
		return nil, fmt.Errorf("connect to comptroller abi: %w", err)
	}
	return s, nil
}
func (s *Server) AttachRoutes(r *mux.Router) {
	r.Path("/healthcheck").Methods(http.MethodGet).HandlerFunc(s.handleCORS(http.MethodGet))

	r.Path("/oracle/deploy").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPost))
	r.Path("/oracle/deploy").Methods(http.MethodPost).Handler(s.authMiddleware(s.handleOracleDeploy))

	r.Path("/oracle/set-price").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/oracle/set-price").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleOracleSetPrice))

	r.Path("/oracle/set-price-feed").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/oracle/set-price-feed").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleOracleSetPriceFeed))

	r.Path("/unitroller/deploy").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPost))
	r.Path("/unitroller/deploy").Methods(http.MethodPost).Handler(s.authMiddleware(s.handleUnitrollerDeploy))

	r.Path("/comptroller/deploy").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPost))
	r.Path("/comptroller/deploy").Methods(http.MethodPost).Handler(s.authMiddleware(s.handleComptrollerDeploy))

	r.Path("/comptroller/set-borrow-paused").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-borrow-paused").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetBorrowPaused))

	r.Path("/comptroller/set-mint-paused").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-mint-paused").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetMintPaused))

	r.Path("/comptroller/set-seize-paused").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-seize-paused").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetSeizePaused))

	r.Path("/comptroller/set-transfer-paused").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-transfer-paused").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetTransferPaused))

	r.Path("/comptroller/support-market").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/support-market").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSupportMarket))

	r.Path("/comptroller/set-price-oracle").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-price-oracle").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetPriceOracle))

	r.Path("/comptroller/set-collateral-factor").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-collateral-factor").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetCollateralFactor))

	r.Path("/comptroller/grant-comp").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/grant-comp").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerGrantComp))

	r.Path("/comptroller/set-comp-speed").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-comp-speed").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetCompSpeed))

	r.Path("/comptroller/set-contributor-comp-speed").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-contributor-comp-speed").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetContributorCompSpeed))

	r.Path("/comptroller/set-liquidation-incentive").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-liquidation-incentive").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetLiquidationIncentive))

	r.Path("/comptroller/set-market-borrow-caps").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-market-borrow-caps").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetMarketBorrowCaps))

	r.Path("/comptroller/set-pause-guardian").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-pause-guardian").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetPauseGuardian))

	r.Path("/comptroller/set-comp-address").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/comptroller/set-comp-address").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleComptrollerSetCompAddress))

	r.Path("/ctoken/deploy").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPost))
	r.Path("/ctoken/deploy").Methods(http.MethodPost).Handler(s.authMiddleware(s.handleCtokenDeploy))

	r.Path("/ctoken/set-comptroller").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/ctoken/set-comptroller").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleCtokenSetComptroller))

	r.Path("/cether/deploy").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPost))
	r.Path("/cether/deploy").Methods(http.MethodPost).Handler(s.authMiddleware(s.handleCetherDeploy))

	r.Path("/cether/set-comptroller").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPut))
	r.Path("/cether/set-comptroller").Methods(http.MethodPut).Handler(s.authMiddleware(s.handleCetherSetComptroller))

	r.Path("/faucet-token/deploy").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPost))
	r.Path("/faucet-token/deploy").Methods(http.MethodPost).Handler(s.authMiddleware(s.handleFaucetTokenDeploy))

	r.Path("/faucet-token/allocate-to").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPost))
	r.Path("/faucet-token/allocate-to").Methods(http.MethodPost).Handler(s.authMiddleware(s.handleFaucetTokenAllocate))

	r.Path("/lens/deploy").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPost))
	r.Path("/lens/deploy").Methods(http.MethodPost).Handler(s.authMiddleware(s.handleLensDeploy))

	r.Path("/interestratemodel/deploy").Methods(http.MethodOptions).HandlerFunc(s.handleCORS(http.MethodPost))
	r.Path("/interestratemodel/deploy").Methods(http.MethodPost).Handler(s.authMiddleware(s.handleInterestRateModelDeploy))

}

func (s *Server) authMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		if s.apiKey != "" && r.Header.Get("Authorization") == s.apiKey {
			next(w, r)
			return
		}
		client := http.Client{
			Timeout: time.Minute,
		}
		req, err := http.NewRequest(http.MethodGet, s.authHost+"/validate", nil)
		if err != nil {
			s.log.Error().Err(err).Msg("Failed to create auth request")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		req.Header.Set("Authorization", r.Header.Get("Authorization"))
		res, err := client.Do(req)
		if err != nil {
			s.log.Error().Err(err).Msg("Failed to send auth request")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		_ = res.Body.Close()
		if res.StatusCode != http.StatusNoContent {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		next(w, r)
	})
}

func (s *Server) handleCORS(methods ...string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Accept-Encoding, Authorization")
	}
}

func jsonError(w http.ResponseWriter, err error) {
	_ = json.NewEncoder(w).Encode(struct{ Error string }{Error: err.Error()})
}
