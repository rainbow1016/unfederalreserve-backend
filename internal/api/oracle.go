package api

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	abichainlink "github.com/UnFederalReserve/compound-server-api/abi/chainlink"
	abictoken "github.com/UnFederalReserve/compound-server-api/abi/ctoken"
	abioracle "github.com/UnFederalReserve/compound-server-api/abi/oracle"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (s *Server) handleOracleDeploy(w http.ResponseWriter, _ *http.Request) {
	log := s.log.With().Str("method", "OracleDeploy").Logger()
	addr, tx, _, err := abioracle.DeployAbioracle(
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

func (s *Server) handleOracleSetPrice(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "OracleSetPrice").Logger()
	var request struct {
		Ctoken   string
		Asset    string
		Price    string
		PriceUSD float64
		Oracle   string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if (request.Ctoken == "" && request.Asset == "") || (request.Price == "" && request.PriceUSD == 0) {
		http.Error(w, fmt.Sprintf("Price and (Ctoken|Asset) required: %+v", request), http.StatusBadRequest)
		return
	}

	var oracleAddr common.Address
	if request.Oracle != "" {
		oracleAddr = common.HexToAddress(request.Oracle)
	} else {
		oracleAddr, err = s.comptroller.Oracle(nil)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get price oracle address")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	oracle, err := abioracle.NewAbioracleTransactor(oracleAddr, s.eth)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to price oracle")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var price *big.Int

	if request.Price != "" {
		var ok bool
		price, ok = new(big.Int).SetString(request.Price, 10)
		if !ok {
			log.Warn().Err(err).Str("input", request.Price).Msg("Failed to decode price")
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	} else {
		var decimals uint8
		switch {
		case request.Asset != "":
			token, err := abictoken.NewAbictokenCaller(common.HexToAddress(request.Asset), s.eth)
			if err != nil {
				log.Warn().Err(err).Str("input", request.Asset).Msg("Failed to get token abi")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			decimals, err = token.Decimals(nil)
			if err != nil {
				log.Warn().Err(err).Str("input", request.Asset).Msg("Failed to get token decimals")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		case request.Ctoken != "":
			ctoken, err := abictoken.NewAbictokenCaller(common.HexToAddress(request.Ctoken), s.eth)
			if err != nil {
				log.Warn().Err(err).Str("input", request.Ctoken).Msg("Failed to get ctoken abi")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			symbol, err := ctoken.Symbol(nil)
			if err != nil {
				log.Warn().Err(err).Str("input", request.Ctoken).Msg("Failed to get ctoken symbol")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			if symbol == "cETH" || symbol == "unETH" {
				decimals = 18
				request.Asset = "0x0000000000000000000000000000000000000000"
			} else {
				underlyingAddr, err := ctoken.Underlying(nil)
				if err != nil {
					log.Warn().Err(err).Str("input", request.Ctoken).Msg("Failed to get underlying asset address")
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}

				token, err := abictoken.NewAbictokenCaller(underlyingAddr, s.eth)
				if err != nil {
					log.Warn().Err(err).Str("input", request.Asset).Msg("Failed to get token abi")
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
				decimals, err = token.Decimals(nil)
				if err != nil {
					log.Warn().Err(err).Str("input", request.Asset).Msg("Failed to get token decimals")
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					return
				}
			}
		}

		p := new(big.Float).SetFloat64(request.PriceUSD)
		p = p.Mul(p, new(big.Float).SetInt(big.NewInt(1e18)))
		p = p.Mul(p, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18-decimals)), nil)))
		price, _ = p.Int(nil)
	}
	var tx *types.Transaction
	switch {
	case request.Asset != "":
		tx, err = oracle.SetDirectPrice(s.signer, common.HexToAddress(request.Asset), price)
	case request.Ctoken != "":
		tx, err = oracle.SetUnderlyingPrice(s.signer, common.HexToAddress(request.Ctoken), price)
	default:
		http.Error(w, "either asset or ctoken must not be empty", http.StatusBadRequest)
		return
	}
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

func (s *Server) handleOracleSetPriceFeed(w http.ResponseWriter, r *http.Request) {
	log := s.log.With().Str("method", "OracleSetPriceFeed").Logger()
	var request struct {
		Ctoken string
		Asset  string
		Feed   string
		Oracle string
	}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to decode request body")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	if (request.Ctoken == "" && request.Asset == "") || request.Feed == "" {
		http.Error(w, fmt.Sprintf("Feed and (Ctoken|Asset) required: %+v", request), http.StatusBadRequest)
		return
	}
	var oracleAddr common.Address
	if request.Oracle != "" {
		oracleAddr = common.HexToAddress(request.Oracle)
	} else {
		oracleAddr, err = s.comptroller.Oracle(nil)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get price oracle address")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	oracle, err := abioracle.NewAbioracleTransactor(oracleAddr, s.eth)
	if err != nil {
		log.Error().Err(err).Msg("Failed to connect to price oracle")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var decimals uint8
	switch {
	case request.Asset != "":
		token, err := abictoken.NewAbictokenCaller(common.HexToAddress(request.Asset), s.eth)
		if err != nil {
			log.Warn().Err(err).Str("input", request.Asset).Msg("Failed to get token abi")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		decimals, err = token.Decimals(nil)
		if err != nil {
			log.Warn().Err(err).Str("input", request.Asset).Msg("Failed to get token decimals")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	case request.Ctoken != "":
		ctoken, err := abictoken.NewAbictokenCaller(common.HexToAddress(request.Ctoken), s.eth)
		if err != nil {
			log.Warn().Err(err).Str("input", request.Ctoken).Msg("Failed to get ctoken abi")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		symbol, err := ctoken.Symbol(nil)
		if err != nil {
			log.Warn().Err(err).Str("input", request.Ctoken).Msg("Failed to get ctoken symbol")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if symbol == "cETH" || symbol == "unETH" {
			decimals = 18
			request.Asset = "0x0000000000000000000000000000000000000000"
		} else {
			underlyingAddr, err := ctoken.Underlying(nil)
			if err != nil {
				log.Warn().Err(err).Str("input", request.Ctoken).Msg("Failed to get underlying asset address")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			token, err := abictoken.NewAbictokenCaller(underlyingAddr, s.eth)
			if err != nil {
				log.Warn().Err(err).Str("input", request.Asset).Msg("Failed to get token abi")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			decimals, err = token.Decimals(nil)
			if err != nil {
				log.Warn().Err(err).Str("input", request.Asset).Msg("Failed to get token decimals")
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}
	}
	feed, err := abichainlink.NewAbichainlinkCaller(common.HexToAddress(request.Feed), s.eth)
	if err != nil {
		log.Warn().Err(err).Str("input", request.Feed).Msg("Failed to get feed abi")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	feedDecimals, err := feed.Decimals(nil)
	if err != nil {
		log.Warn().Err(err).Str("input", request.Feed).Msg("Failed to get feed decimals")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18+18-decimals-feedDecimals)), nil)
	var tx *types.Transaction
	switch {
	case request.Asset != "":
		log.Info().Str("asset", request.Asset).Str("feed", request.Feed).Str("mult", multiplier.String()).Msg("Setting direct chainlink feed")
		tx, err = oracle.SetDirectChainlinkFeed(s.signer, common.HexToAddress(request.Asset), common.HexToAddress(request.Feed), multiplier)
	case request.Ctoken != "":
		log.Info().Str("ctoken", request.Ctoken).Str("feed", request.Feed).Str("mult", multiplier.String()).Msg("Setting ctoken chainlink feed")
		tx, err = oracle.SetUnderlyingChainlinkFeed(s.signer, common.HexToAddress(request.Ctoken), common.HexToAddress(request.Feed), multiplier)
	default:
		http.Error(w, "either asset or ctoken must not be empty", http.StatusBadRequest)
		return
	}
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
