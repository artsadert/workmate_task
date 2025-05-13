package rest

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"gitlab.com/digineat/go-broker-test/internal/application/commands"
	"gitlab.com/digineat/go-broker-test/internal/infrastructure/db/sqlite"
	"gitlab.com/digineat/go-broker-test/internal/interface/api/rest/dto/mapper"
	trade_request "gitlab.com/digineat/go-broker-test/internal/interface/api/rest/dto/request"
	"gitlab.com/digineat/go-broker-test/internal/interface/api/rest/dto/response"
)

type Server struct {
	database *sql.DB
}

func NewServer(database *sql.DB) *Server {
	return &Server{database: database}
}

func (s Server) Trades(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Falied to read request", http.StatusBadRequest)
		return
	}

	var request_trade trade_request.TradeRequest
	if err := json.Unmarshal(body, &request_trade); err != nil {
		http.Error(w, "Failed to unparse request body", http.StatusBadRequest)
		return
	}
	err = mapper.ValidateTrade(request_trade)
	if err != nil {
		http.Error(w, "Wrong fields in request", http.StatusBadRequest)
		return
	}

	err = commands.CreateTrade(s.database, request_trade.Account, request_trade.Symbol, request_trade.Volume, request_trade.Open, request_trade.Close, request_trade.Side)
	if err != nil {
		http.Error(w, "Failed to create trade", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s Server) Stats(w http.ResponseWriter, r *http.Request) {
	acc := r.PathValue("acc")
	if acc == "" {
		http.Error(w, "Failed to read accound field in query param", http.StatusBadRequest)
		return
	}

	trades_count, profit, err := commands.GetStats(s.database, acc)
	if err != nil {
		http.Error(w, "Failed to read stats", http.StatusInternalServerError)
	}

	stats := response.StatsResponse{
		Account: acc,
		Trades:  trades_count,
		Profit:  profit,
	}

	json.NewEncoder(w).Encode(stats)
	w.WriteHeader(http.StatusOK)
}

func (s Server) Healthz(w http.ResponseWriter, r *http.Request) {
	if err := sqlite.PingDB(s.database); err != nil {
		http.Error(w, "DB connection error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func NewHandler(database *sql.DB) http.Handler {
	mux := http.NewServeMux()
	server := NewServer(database)

	mux.HandleFunc("POST /trades", server.Trades)
	mux.HandleFunc("GET /stats/{acc}", server.Stats)
	mux.HandleFunc("GET /healthz", server.Healthz)

	return mux
}
