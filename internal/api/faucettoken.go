package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	abifaucettoken "github.com/UnFederalReserve/compound-server-api/abi/faucettoken"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Server) handleFaucetTokenDeploy(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "FaucetTokenDeploy").Logger()
	var request struct {
		InitialAmount string
		Name          string
		Decimals      uint8
		Symbol        string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if request.InitialAmount == "" || request.Name == "" || request.Decimals == 0 || request.Symbol == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}
	initialAmount, ok := new(big.Int).SetString(request.InitialAmount, 10)
	if !ok {
		log.Warn().Err(err).Str("input", request.InitialAmount).Msg("Failed to decode InitialAmount")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	addr, tx, _, err := abifaucettoken.DeployAbictoken(
		s.signer, s.eth,
		initialAmount, request.Name, request.Decimals, request.Symbol,
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

func (s *Server) handleFaucetTokenAllocate(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "FaucetTokenAllocate").Logger()
	var request struct {
		Asset  string
		To     string
		Amount string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.To == "" || request.Asset == "" || request.Amount == "" {
		http.Error(w, fmt.Sprintf("All fields are required: %+v", request), http.StatusBadRequest)
		return
	}

	token, err := abifaucettoken.NewAbictokenTransactor(common.HexToAddress(request.Asset), s.eth)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to token contract")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	amount, ok := new(big.Int).SetString(request.Amount, 10)
	if !ok {
		log.Warn().Err(err).Str("input", request.Amount).Msg("Failed to decode amount")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	tx, err := token.AllocateTo(s.signer, common.HexToAddress(request.To), amount)
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
