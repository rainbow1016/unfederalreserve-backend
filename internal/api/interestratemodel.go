package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	abiinterestratemodel "github.com/UnFederalReserve/compound-server-api/abi/interestratemodel"
)

func (s *Server) handleInterestRateModelDeploy(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "InterestRateModel").Logger()
	var request struct {
		BaseRatePerYear   string
		MultiplierPerYear string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if request.BaseRatePerYear == "" || request.MultiplierPerYear == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}
	baseRatePerYear, ok := new(big.Int).SetString(request.BaseRatePerYear, 10)
	if !ok {
		log.Warn().Err(err).Str("input", request.BaseRatePerYear).Msg("Failed to decode BaseRatePerYear")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	multiplierPerYear, ok := new(big.Int).SetString(request.MultiplierPerYear, 10)
	if !ok {
		log.Warn().Err(err).Str("input", request.MultiplierPerYear).Msg("Failed to decode MultiplierPerYear")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	addr, tx, _, err := abiinterestratemodel.DeployAbiinterestratemodel(
		s.signer, s.eth, baseRatePerYear, multiplierPerYear,
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
