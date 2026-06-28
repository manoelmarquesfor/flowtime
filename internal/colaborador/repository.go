package colaborador

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github/manoelmarquesfor/flowtime/internal/errs"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) ByID(ctx context.Context, colaboradorID uuid.UUID) (ColaboradorRepository, error) {
	var colaborador ColaboradorRepository

	query := `
		SELECT
			id, 
			matricula, 
			tag_id, 
			nome, 
			setor, 
			ativo					
		FROM colaboradores WHERE id = ?`

	err := r.db.GetContext(ctx, &colaborador, query, colaboradorID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return colaborador, errs.NewNotFoundError("Colaborador não encontrado")
		}

		return colaborador, errs.NewRepositoryError("Erro ao buscar colaborador por id: " + err.Error())
	}

	return colaborador, nil
}

func (r *Repository) All(ctx context.Context, ativo *bool) ([]ColaboradorRepository, error) {
	var colaboradores []ColaboradorRepository

	where := ""
	args := []interface{}{}
	if ativo != nil {
		where = "WHERE ativo = ?"
		args = append(args, *ativo)
	}

	query := `
		SELECT
			id, 
			matricula, 
			tag_id, 
			nome, 
			setor, 
			ativo					
		FROM colaboradores
		` + where + `order by nome;`

	err := r.db.SelectContext(ctx, &colaboradores, query, args...)
	if err != nil {
		return colaboradores, errs.NewRepositoryError("Erro ao buscar colaboradores: " + err.Error())
	}

	return colaboradores, nil
}

func (r *Repository) Create(ctx context.Context, colaborador ColaboradorRepository) error {
	query := `
		INSERT INTO colaboradores 
		(id, matricula, tag_id, nome, setor, ativo,dt_criacao)
		VALUES (?, ?, ?, ?, ?, ?, ?);`

	_, err := r.db.ExecContext(ctx, query,
		colaborador.ID, colaborador.Matricula,
		colaborador.Tag, colaborador.Nome,
		colaborador.Setor, colaborador.Ativo, time.Now())
	if err != nil {
		return errs.NewRepositoryError("Erro ao criar colaborador: " + err.Error())
	}

	return nil
}

func (r *Repository) Update(ctx context.Context, colaborador ColaboradorRepository) error {
	query := `
		UPDATE colaboradores
			SET nome = ?,
				tag_id = ?,
			    matricula = ?,
				setor = ?,
			    ativo = ?
		WHERE id = ?;`

	result, err := r.db.ExecContext(ctx, query,
		colaborador.Nome, colaborador.Tag, colaborador.Matricula,
		colaborador.Setor, colaborador.Ativo, colaborador.ID)
	if err != nil {
		return errs.NewRepositoryError("Erro ao atualizar colaborador: " + err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errs.NewRepositoryError("Erro ao verificar atualização do colaborador: " + err.Error())
	}

	if rowsAffected == 0 {
		return errs.NewNotFoundError("Colaborador não encontrado")
	}

	return nil
}
