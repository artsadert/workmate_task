package worker

import (
	"database/sql"
	"log"
	"time"

	"gitlab.com/digineat/go-broker-test/internal/infrastructure/db/sqlite"
)

type Worker struct {
	DB           *sql.DB
	PollInterval time.Duration
	quit         chan struct{}
}

func NewWorker(db *sql.DB, pollInterval time.Duration) *Worker {
	return &Worker{
		DB:           db,
		PollInterval: pollInterval,
		quit:         make(chan struct{}),
	}
}

func (w *Worker) Run() {
	ticker := time.NewTicker(w.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := w.processOneTrade(); err != nil {
				log.Printf("Worker error: %v", err)
			}
		case <-w.quit:
			log.Println("Worker stopping")
			return
		}
	}
}

func (w *Worker) Stop() {
	close(w.quit)
}

func (w *Worker) processOneTrade() error {
	return sqlite.ProcessOneTradeTx(w.DB)
}
