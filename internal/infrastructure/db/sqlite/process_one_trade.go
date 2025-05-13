package sqlite

import (
	"database/sql"
)

func ProcessOneTradeTx(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	var (
		id      int64
		account string
		symbol  string
		volume  float64
		open    float64
		closeP  float64
		side    string
	)

	err = tx.QueryRow(`
		UPDATE trades
		SET processed = 2
		WHERE id = (
			SELECT id FROM trades WHERE processed = 0 LIMIT 1
		)
		RETURNING id, account, symbol, volume, open, close, side
	`).Scan(&id, &account, &symbol, &volume, &open, &closeP, &side)

	if err == sql.ErrNoRows {
		tx.Rollback()
		return nil
	}
	if err != nil {
		tx.Rollback()
		return err
	}

	const lot = 100000.0
	profit := (closeP - open) * volume * lot
	if side == "sell" {
		profit = -profit
	}

	_, err = tx.Exec(`
		INSERT INTO account (account, trades, profit)
		VALUES (?, 1, ?)
		ON CONFLICT(account) DO UPDATE SET
			trades = trades + 1,
			profit = profit + excluded.profit
	`, account, profit)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`UPDATE trades SET processed = 1 WHERE id = ?`, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
