package auth

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

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) UsuarioByEmail(ctx context.Context, email string) (Usuario, error) {
	var usuario Usuario

	query := `
		SELECT
			id,
			nome,
			email,
			password,
			perfil,
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

func (r *Repository) SaveSessao(ctx context.Context, sessaoID *uuid.UUID, usuarioID uuid.UUID) error {
	query := `
		UPDATE usuario 
			SET sessao_id = ? 
		WHERE id = ?`

	_, err := r.db.ExecContext(ctx, query, sessaoID, usuarioID)
	if err != nil {
		return errs.NewRepositoryError("Erro ao salvar sessão do usuário: " + err.Error())
	}

	return nil
}
