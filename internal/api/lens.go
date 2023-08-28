package api

import (
	"encoding/json"
	"net/http"

	abilens "github.com/UnFederalReserve/compound-server-api/abi/lens"
)

func (s *Server) handleLensDeploy(w http.ResponseWriter, _ *http.Request) {
	log := s.log.With().Str("method", "LensDeploy").Logger()
	addr, tx, _, err := abilens.DeployAbilens(
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
