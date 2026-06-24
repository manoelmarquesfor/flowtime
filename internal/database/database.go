package database

import (
	"errors"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var ErrDatabaseNotConnected = errors.New("database not connected")

func Conect(databaseURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite", databaseURL)
	if err != nil {
		return nil, ErrDatabaseNotConnected
	}

	return db, nil
}
