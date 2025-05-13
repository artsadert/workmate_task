package main

import (
	"database/sql"
	"flag"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/digineat/go-broker-test/internal/infrastructure/db/sqlite"
	"gitlab.com/digineat/go-broker-test/internal/interface/worker"
)

func main() {
	// Command line flags
	dbPath := flag.String("db", "data.db", "path to SQLite database")
	pollInterval := flag.Duration("poll", 100*time.Millisecond, "polling interval")
	flag.Parse()

	// Initialize database connection
	db, err := sql.Open("sqlite3", *dbPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	if err := sqlite.Init_db(db); err != nil {
		log.Fatalf("Failed to create DB: %v", err)
	}
	// Test database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Printf("Worker started with polling interval: %v", *pollInterval)
	worker := worker.NewWorker(db, *pollInterval)

	// Main worker loop
	worker.Run()
}
