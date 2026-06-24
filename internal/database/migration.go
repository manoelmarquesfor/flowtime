package database

import (
	"database/sql"
	"embed"
	"errors"

	"github.com/pressly/goose/v3"
)

const (
	migrationsDir = "migrations"
	dialect       = "sqlite"
)

var (
	ErrFalhaSetDialect         = errors.New("falha ao setar o dialect")
	ErrFalhaExecutarMigrations = errors.New("falha ao executar as migrations")
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func Migrate(database *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect(dialect); err != nil {
		return ErrFalhaSetDialect
	}

	if err := goose.Up(database, migrationsDir); err != nil {
		return ErrFalhaExecutarMigrations
	}

	return nil
}
