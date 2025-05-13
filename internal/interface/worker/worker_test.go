package worker

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	sqlite "gitlab.com/digineat/go-broker-test/internal/infrastructure/db/sqlite"
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
func TestProcessOneTrade_NoNewTrades(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	w := NewWorker(db, time.Second)

	err := w.processOneTrade()
	if err != nil {
		t.Fatalf("expected no error when no new trades, got %v", err)
	}
}
func TestWorkerRunStop(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	w := NewWorker(db, 10*time.Millisecond)

	_, err := db.Exec(`INSERT INTO trades (account, symbol, volume, open, close, side, processed) VALUES (?, ?, ?, ?, ?, ?, 0)`,
		"acc1", "EURUSD", 1.0, 1.1, 1.2, "buy")
	if err != nil {
		t.Fatalf("failed to insert trade: %v", err)
	}

	go w.Run()

	time.Sleep(50 * time.Millisecond)

	w.Stop()

	time.Sleep(10 * time.Millisecond)
}
func TestWorkerProcessOneTrade(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	w := NewWorker(db, time.Second)

	_, err := db.Exec(`INSERT INTO trades (account, symbol, volume, open, close, side, processed) VALUES (?, ?, ?, ?, ?, ?, 0)`,
		"acc1", "EURUSD", 1.0, 1.1, 1.2, "buy")
	if err != nil {
		t.Fatalf("failed to insert trade: %v", err)
	}

	err = w.processOneTrade()
	if err != nil {
		t.Fatalf("processOneTrade failed: %v", err)
	}

	var processed int
	err = db.QueryRow(`SELECT processed FROM trades WHERE account = ?`, "acc1").Scan(&processed)
	if err != nil {
		t.Fatalf("failed to query trade: %v", err)
	}
	if processed != 1 {
		t.Errorf("expected processed=1, got %d", processed)
	}

	trades, profit, err := sqlite.GetStats(db, "acc1")
	if err != nil {
		t.Fatalf("GetAccountStats failed: %v", err)
	}
	if trades != 1 {
		t.Errorf("expected trades=1, got %d", trades)
	}
	expectedProfit := (1.2 - 1.1) * 1.0 * 100000.0
	const epsilon = 1e-9
	if diff := profit - expectedProfit; diff < -epsilon || diff > epsilon {
		t.Errorf("expected profit=%f, got %f", expectedProfit, profit)
	}

	err = w.processOneTrade()
	if err != nil {
		t.Fatalf("processOneTrade second call failed: %v", err)
	}
}
