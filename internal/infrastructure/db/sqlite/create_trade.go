package sqlite

import "database/sql"

func CreateTrade(db *sql.DB, account string, symbol string, volume float64, open float64, close_ float64, side string) error {
	_, err := db.Exec(`INSERT INTO trades (account, symbol, volume, open, close, side, processed) VALUES (?, ?, ?, ?, ?, ?, 0)`,
		account, symbol, volume, open, close_, side)
	if err != nil {
		return err
	}
	return nil
}
