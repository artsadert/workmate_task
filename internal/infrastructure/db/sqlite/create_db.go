package sqlite

import (
	"database/sql"
	"fmt"
)

func Init_db(db *sql.DB) error {
	trade := `CREATE TABLE IF NOT EXISTS trades (
						id INTEGER PRIMARY KEY AUTOINCREMENT,
						account TEXT NOT NULL CHECK (account != ''),
						symbol TEXT NOT NULL,
						volume REAL NOT NULL CHECK (volume > 0),
						open REAL NOT NULL CHECK (open > 0),
						close REAL NOT NULL CHECK (close > 0),
						side TEXT NOT NULL CHECK (side in ('buy', 'sell')),
						processed INTEGER NOT NULL DEFAULT 0
					);`
	account := `CREATE TABLE IF NOT EXISTS account (
							account TEXT PRIMARY KEY,
							trades INTEGER NOT NULL DEFAULT 0,
							profit REAL NOT NULL DEFAULT 0
						);`

	if _, err := db.Exec(trade); err != nil {
		return fmt.Errorf("Failed to create table trade: %v", err)
	}

	if _, err := db.Exec(account); err != nil {
		return fmt.Errorf("Failed to create table account: %v", err)
	}

	return nil

}
