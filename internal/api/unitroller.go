package api

import (
	"encoding/json"
	"math/big"
	"net/http"

	abiunitroller "github.com/UnFederalReserve/compound-server-api/abi/unitroller"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func (s *Server) handleUnitrollerDeploy(w http.ResponseWriter, _ *http.Request) {
	log := s.log.With().Str("method", "UnitrollerDeploy").Logger()
	signer := &bind.TransactOpts{
		From:     s.signer.From,
		Signer:   s.signer.Signer,
		GasPrice: big.NewInt(11 * 1e9),
	}
	addr, tx, _, err := abiunitroller.DeployAbiunitroller(
		signer, s.eth,
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
