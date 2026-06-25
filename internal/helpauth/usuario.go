package helpauth

import "github.com/google/uuid"

type UsuarioAutenticado struct {
	ID    uuid.UUID
	Nome  string
	Email string
	Regra string
	Ativo bool
}

func (u *UsuarioRepository) ToUsuarioAutenticado() *UsuarioAutenticado {
	return &UsuarioAutenticado{
		ID:    u.ID,
		Nome:  u.Nome,
		Email: u.Email,
		Regra: u.Regra,
		Ativo: u.Ativo,
	}
}

type UsuarioRepository struct {
	ID    uuid.UUID `db:"id"`
	Nome  string    `db:"nome"`
	Email string    `db:"email"`
	Regra string    `db:"regra"`
	Ativo bool      `db:"ativo"`
}
