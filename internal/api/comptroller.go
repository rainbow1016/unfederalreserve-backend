package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	abicomptroller "github.com/UnFederalReserve/compound-server-api/abi/comptroller"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Server) handleComptrollerDeploy(w http.ResponseWriter, _ *http.Request) {
	log := s.log.With().Str("method", "ComptrollerDeploy").Logger()
	addr, tx, _, err := abicomptroller.DeployAbicomptroller(
		s.signer, s.eth,
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

func (s *Server) handleComptrollerSetBorrowPaused(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetBorrowPaused").Logger()
	var request struct {
		Ctoken string
		State  bool
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Ctoken == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}

	tx, err := s.comptroller.SetBorrowPaused(s.signer, common.HexToAddress(request.Ctoken), request.State)
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

func (s *Server) handleComptrollerSetMintPaused(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetMintPaused").Logger()
	var request struct {
		Ctoken string
		State  bool
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Ctoken == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}

	tx, err := s.comptroller.SetMintPaused(s.signer, common.HexToAddress(request.Ctoken), request.State)
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

func (s *Server) handleComptrollerSetSeizePaused(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetSeizePaused").Logger()
	var request struct {
		State bool
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	tx, err := s.comptroller.SetSeizePaused(s.signer, request.State)
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

func (s *Server) handleComptrollerSetTransferPaused(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetTransferPaused").Logger()
	var request struct {
		State bool
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	tx, err := s.comptroller.SetTransferPaused(s.signer, request.State)
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

func (s *Server) handleComptrollerSupportMarket(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSupportMarket").Logger()
	var request struct {
		Ctoken string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Ctoken == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}
	signer := &bind.TransactOpts{
		From:     s.signer.From,
		Signer:   s.signer.Signer,
		GasLimit: 150000,
	}
	tx, err := s.comptroller.SupportMarket(signer, common.HexToAddress(request.Ctoken))
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

func (s *Server) handleComptrollerSetPriceOracle(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetPriceOracle").Logger()
	var request struct {
		Address string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Address == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}

	tx, err := s.comptroller.SetPriceOracle(s.signer, common.HexToAddress(request.Address))
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

func (s *Server) handleComptrollerSetCollateralFactor(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetCollateralFactor").Logger()
	var request struct {
		Ctoken                      string
		NewCollateralFactorMantissa string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Ctoken == "" || request.NewCollateralFactorMantissa == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}

	newCollateralFactorMantissa, ok := new(big.Int).SetString(request.NewCollateralFactorMantissa, 10)
	if !ok {
		log.Warn().Err(err).Str("input", request.NewCollateralFactorMantissa).Msg("Failed to decode newCollateralFactorMantissa")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	signer := &bind.TransactOpts{
		From:     s.signer.From,
		Signer:   s.signer.Signer,
		GasLimit: 150000,
	}
	tx, err := s.comptroller.SetCollateralFactor(signer, common.HexToAddress(request.Ctoken), newCollateralFactorMantissa)
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

func (s *Server) handleComptrollerGrantComp(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerGrantComp").Logger()
	var request struct {
		Recipient string
		Amount    string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Recipient == "" || request.Amount == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}

	amount, ok := new(big.Int).SetString(request.Amount, 10)
	if !ok {
		log.Warn().Err(err).Str("input", request.Amount).Msg("Failed to decode amount")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	tx, err := s.comptroller.GrantComp(s.signer, common.HexToAddress(request.Recipient), amount)
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

func (s *Server) handleComptrollerSetCompSpeed(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetCompSpeed").Logger()
	var request struct {
		Ctoken string
		Speed  string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Ctoken == "" || request.Speed == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}

	speed, ok := new(big.Int).SetString(request.Speed, 10)
	if !ok {
		log.Warn().Err(err).Str("input", request.Speed).Msg("Failed to decode speed")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	tx, err := s.comptroller.SetCompSpeed(s.signer, common.HexToAddress(request.Ctoken), speed)
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

func (s *Server) handleComptrollerSetContributorCompSpeed(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetContributorCompSpeed").Logger()
	var request struct {
		Contributor string
		Speed       string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Contributor == "" || request.Speed == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}

	speed, ok := new(big.Int).SetString(request.Speed, 10)
	if !ok {
		log.Warn().Err(err).Str("input", request.Speed).Msg("Failed to decode speed")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	tx, err := s.comptroller.SetContributorCompSpeed(s.signer, common.HexToAddress(request.Contributor), speed)
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

func (s *Server) handleComptrollerSetLiquidationIncentive(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetLiquidationIncentive").Logger()
	var request struct {
		Incentive string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Incentive == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}

	incentive, ok := new(big.Int).SetString(request.Incentive, 10)
	if !ok {
		log.Warn().Err(err).Str("input", request.Incentive).Msg("Failed to decode incentive")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	tx, err := s.comptroller.SetLiquidationIncentive(s.signer, incentive)
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

func (s *Server) handleComptrollerSetMarketBorrowCaps(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetMarketBorrowCaps").Logger()
	var request struct {
		Ctokens    []string
		BorrowCaps []string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if len(request.Ctokens) == 0 || len(request.BorrowCaps) == 0 {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}
	var ctokens []common.Address
	var caps []*big.Int
	for _, ctoken := range request.Ctokens {
		ctokens = append(ctokens, common.HexToAddress(ctoken))
	}
	for _, capString := range request.BorrowCaps {

		borrowCap, ok := new(big.Int).SetString(capString, 10)
		if !ok {
			log.Warn().Err(err).Str("input", capString).Msg("Failed to decode borrowCap")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		caps = append(caps, borrowCap)
	}

	tx, err := s.comptroller.SetMarketBorrowCaps(s.signer, ctokens, caps)
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

func (s *Server) handleComptrollerSetPauseGuardian(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetPauseGuardian").Logger()
	var request struct {
		Address string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Address == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}

	tx, err := s.comptroller.SetPauseGuardian(s.signer, common.HexToAddress(request.Address))
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

func (s *Server) handleComptrollerSetCompAddress(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "ComptrollerSetCompAddress").Logger()
	var request struct {
		Address string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if request.Address == "" {
		http.Error(w, fmt.Sprintf("all fields are required: %+v", request), http.StatusBadRequest)
		return
	}

	tx, err := s.comptroller.SetCompAddress(s.signer, common.HexToAddress(request.Address))
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
