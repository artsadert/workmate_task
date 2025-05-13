package commands

import (
	"database/sql"

	"gitlab.com/digineat/go-broker-test/internal/infrastructure/db/sqlite"
)

func GetStats(db *sql.DB, acc string) (int, float64, error) {
	return sqlite.GetStats(db, acc)
}
