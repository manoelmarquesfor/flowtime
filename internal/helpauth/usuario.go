package helpauth

import "github.com/google/uuid"

type UsuarioAutencicado struct {
	ID    uuid.UUID
	Nome  string
	Email string
	Regra string
	Ativo bool
}

func (u *UsuarioRepository) ToUsuarioAutenticado() *UsuarioAutencicado {
	return &UsuarioAutencicado{
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
