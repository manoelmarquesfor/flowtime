package usuario

import (
	"context"
	"database/sql"
	"errors"

	"github/manoelmarquesfor/flowtime/internal/errs"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	db *sqlx.DB
}

func (r *Repository) UsuarioDelete(ctx context.Context, userID uuid.UUID) error {
	query := `
		DELETE FROM usuario WHERE id = ?;`

	result, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return errs.NewRepositoryError("Erro ao deletar usuário: " + err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errs.NewRepositoryError("Erro ao verificar deleção de usuário: " + err.Error())
	}

	if rowsAffected == 0 {
		return errs.NewNotFoundError("Usuário não encontrado")
	}

	return nil
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) UsuarioByEmail(ctx context.Context, email string) (UsuarioRepository, error) {
	var usuario UsuarioRepository

	query := `
		SELECT
			id,
			nome,
			email,
			regra,
			ativo					
		FROM usuario WHERE email = ?`

	err := r.db.GetContext(ctx, &usuario, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return usuario, errs.NewNotFoundError("Usuário não encontrado")
		}

		return usuario, errs.NewRepositoryError("Erro ao buscar usuário por email: " + err.Error())
	}

	return usuario, nil
}

func (r *Repository) UsuarioByID(ctx context.Context, userID uuid.UUID) (UsuarioRepository, error) {
	var usuario UsuarioRepository

	query := `
		SELECT
			id,
			nome,
			email,
			regra,
			ativo					
		FROM usuario WHERE id = ?`

	err := r.db.GetContext(ctx, &usuario, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return usuario, errs.NewNotFoundError("Usuário não encontrado")
		}

		return usuario, errs.NewRepositoryError("Erro ao buscar usuário por id: " + err.Error())
	}

	return usuario, nil
}

func (r *Repository) UsuarioAll(ctx context.Context) ([]UsuarioRepository, error) {
	var usuarios []UsuarioRepository

	query := `
		SELECT
			id,
			nome,
			email,
			regra,
			ativo					
		FROM usuario`

	err := r.db.SelectContext(ctx, &usuarios, query)
	if err != nil {
		return usuarios, errs.NewRepositoryError("Erro ao buscar usuários: " + err.Error())
	}

	return usuarios, nil
}

func (r *Repository) UsuarioCreate(ctx context.Context, user UsuarioCreateRepository) error {
	query := `
		INSERT INTO usuario 
		(id, nome, email, password, regra, ativo,dt_criacao)
		VALUES (?, ?, ?, ?, ?, ?, ?);`

	_, err := r.db.ExecContext(ctx, query,
		user.ID, user.Nome, user.Email, user.Password, user.Regra, user.Ativo, user.DtCreated)
	if err != nil {
		return errs.NewRepositoryError("Erro ao criar usuário: " + err.Error())
	}

	return nil
}

func (r *Repository) UsuarioUpdate(ctx context.Context, user UsuarioRepository) error {
	query := `
		UPDATE usuario
			SET nome = ?,
				email = ?,
			    regra = ?,
			    ativo = ?
		WHERE id = ?;`

	result, err := r.db.ExecContext(ctx, query, user.Nome, user.Email, user.Regra, user.Ativo, user.ID)
	if err != nil {
		return errs.NewRepositoryError("Erro ao atualizar usuário: " + err.Error())
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errs.NewRepositoryError("Erro ao verificar atualização de usuário: " + err.Error())
	}

	if rowsAffected == 0 {
		return errs.NewNotFoundError("Usuário não encontrado")
	}

	return nil
}
