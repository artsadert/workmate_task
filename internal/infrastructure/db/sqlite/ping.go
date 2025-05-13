package sqlite

import (
	"database/sql"
	"errors"
)

func PingDB(database *sql.DB) error {
	if err := database.Ping(); err != nil {
		return errors.New("DB connection error")
	}

	return nil
}
