package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	abicether "github.com/UnFederalReserve/compound-server-api/abi/cether"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Server) handleCetherDeploy(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "CetherDeploy").Logger()
	var request struct {
		InterestRateModel           string
		InitialExchangeRateMantissa string
		Name                        string
		Symbol                      string
		Admin                       string
		Decimals                    uint8
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if request.InterestRateModel == "" || request.InitialExchangeRateMantissa == "" || request.Name == "" || request.Symbol == "" || request.Admin == "" || request.Decimals == 0 {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}
	initialExchangeRateMantissa, ok := new(big.Int).SetString(request.InitialExchangeRateMantissa, 10)
	if !ok {
		log.Warn().Err(err).Str("input", request.InitialExchangeRateMantissa).Msg("Failed to decode InitialExchangeRateMantissa")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	addr, tx, _, err := abicether.DeployAbicether(
		s.signer, s.eth,
		common.HexToAddress(s.comptrollerAddress), common.HexToAddress(request.InterestRateModel),
		initialExchangeRateMantissa, request.Name, request.Symbol, request.Decimals, common.HexToAddress(request.Admin),
	)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to send tx")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(struct {
		TxHash string
		Addr   string
	}{TxHash: tx.Hash().String(), Addr: addr.String()}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonError(w, err)
	}
}

func (s *Server) handleCetherSetComptroller(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "CetherSetComptroller").Logger()
	var request struct {
		Ctoken string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	ctoken, err := abicether.NewAbicetherTransactor(common.HexToAddress(request.Ctoken), s.eth)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to ctoken contract")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	tx, err := ctoken.SetComptroller(s.signer, common.HexToAddress(s.comptrollerAddress))
	if err != nil {
		log.Warn().Err(err).Msg("Failed to send tx")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(struct{ TxHash string }{TxHash: tx.Hash().String()}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		jsonError(w, err)
	}
}
