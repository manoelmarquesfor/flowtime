package feriado

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github/manoelmarquesfor/flowtime/internal/errs"

	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) FeriadoByData(ctx context.Context, data time.Time) (RepositoryModel, error) {
	query := `
		SELECT
			data,
			descricao
		FROM feriados WHERE data = ?`

	var feriado RepositoryModel

	err := r.db.GetContext(ctx, &feriado, query, data.Format(time.DateOnly))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return feriado, errs.NewNotFoundError("Feriado não encontrado")
		}

		return feriado, errs.NewRepositoryError("Erro ao buscar feriado: " + err.Error())
	}

	return feriado, nil
}

func (r *Repository) FeriadoByPeriodo(ctx context.Context, inicial, final time.Time) ([]RepositoryModel, error) {
	query := `
		SELECT
			data,
			descricao
		FROM feriados
		WHERE data >= ? AND data < ?`

	var feriados []RepositoryModel

	err := r.db.SelectContext(ctx, &feriados, query, inicial.Format(time.DateOnly), final.Format(time.DateOnly))
	if err != nil {
		return nil, errs.NewRepositoryError("Erro ao buscar feriados: " + err.Error())
	}

	return feriados, nil
}

func (r *Repository) FeriadoDelete(ctx context.Context, data time.Time) error {
	query := `
		DELETE FROM feriados WHERE data = ?;`

	result, err := r.db.ExecContext(ctx, query, data.Format(time.DateOnly))
	if err != nil {
		return errs.NewRepositoryError("Erro ao deletar feriado: " + err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errs.NewRepositoryError("Erro ao verificar deleção de feriado: " + err.Error())
	}

	if rowsAffected == 0 {
		return errs.NewNotFoundError("Feriado não encontrado")
	}

	return nil
}

func (r *Repository) FeriadoCreate(ctx context.Context, feriado RepositoryModel) error {
	query := `
		INSERT INTO feriados (data, descricao) VALUES (?, ?);`

	_, err := r.db.ExecContext(ctx, query, feriado.Data.Format(time.DateOnly), feriado.Descricao)
	if err != nil {
		return errs.NewRepositoryError("Erro ao criar feriado: " + err.Error())
	}

	return nil
}
