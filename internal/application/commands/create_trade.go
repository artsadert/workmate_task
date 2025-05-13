package commands

import (
	"database/sql"

	"gitlab.com/digineat/go-broker-test/internal/infrastructure/db/sqlite"
)

func CreateTrade(db *sql.DB, account string, symbol string, volume float64, open float64, close_ float64, side string) error {
	err := sqlite.CreateTrade(db, account, symbol, volume, open, close_, side)

	if err != nil {
		return err
	}
	return nil
}
