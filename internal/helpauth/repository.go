package helpauth

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

func (r *Repository) UsuarioBySessao(ctx context.Context, sessaoID uuid.UUID) (UsuarioRepository, error) {
	var usuario UsuarioRepository

	query := `
		SELECT
			id,
			nome,
			email,
			regra,
			ativo					
		FROM usuario WHERE sessao_id = ?`

	err := r.db.GetContext(ctx, &usuario, query, sessaoID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return usuario, errs.NewNotFoundError("Usuário não encontrado")
		}

		return usuario, errs.NewRepositoryError("Erro ao buscar usuário por sessão: " + err.Error())
	}

	return usuario, nil
}
