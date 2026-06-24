package database

import (
	"github/manoelmarquesfor/flowtime/internal/config"

	"github.com/jmoiron/sqlx"
)

func Setup(config *config.Config) (*sqlx.DB, error) {
	database, err := Conect(config.Database.Name)
	if err != nil {
		return nil, err
	}

	if err := Migrate(database.DB); err != nil {
		return nil, err
	}

	return database, nil
}
