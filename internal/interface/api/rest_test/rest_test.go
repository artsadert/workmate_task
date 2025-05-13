package rest_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/digineat/go-broker-test/internal/infrastructure/db/sqlite"
	"gitlab.com/digineat/go-broker-test/internal/interface/api/rest"
	trade_request "gitlab.com/digineat/go-broker-test/internal/interface/api/rest/dto/request"
)

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	if err := sqlite.Init_db(db); err != nil {
		t.Fatalf("failed to init db: %v", err)
	}

	return db
}

func TestTradesHandler(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	server := rest.NewServer(db)

	// Prepare a valid trade request JSON
	tradeReq := trade_request.TradeRequest{
		Account: "acc123",
		Symbol:  "AAPLDS",
		Volume:  10,
		Open:    150.0,
		Close:   155.0,
		Side:    "buy",
	}

	body, err := json.Marshal(tradeReq)
	if err != nil {
		t.Fatalf("failed to marshal trade request: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/trades", bytes.NewReader(body))
	w := httptest.NewRecorder()

	server.Trades(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", resp.Status)
	}
}

func TestHealthzHandler(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	server := rest.NewServer(db)

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	w := httptest.NewRecorder()

	server.Healthz(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", resp.Status)
	}

	body := w.Body.String()
	if strings.TrimSpace(body) != "OK" {
		t.Errorf("expected body OK, got %s", body)
	}
}
