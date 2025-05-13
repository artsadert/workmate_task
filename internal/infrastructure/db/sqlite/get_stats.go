package sqlite

import (
	"database/sql"
)

func GetStats(db *sql.DB, acc string) (trades int, profit float64, err error) {
	row := db.QueryRow(`SELECT trades, profit FROM account WHERE account = ?`, acc)
	err = row.Scan(&trades, &profit)
	if err == sql.ErrNoRows {
		return 0, 0, nil
	}
	return
}
