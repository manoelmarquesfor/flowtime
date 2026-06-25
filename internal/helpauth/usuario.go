package helpauth

import "github.com/google/uuid"

type UsuarioAutenticado struct {
	ID     uuid.UUID
	Nome   string
	Email  string
	Perfil string
	Ativo  bool
}

func (u *UsuarioRepository) ToUsuarioAutenticado() *UsuarioAutenticado {
	return &UsuarioAutenticado{
		ID:     u.ID,
		Nome:   u.Nome,
		Email:  u.Email,
		Perfil: u.Perfil,
		Ativo:  u.Ativo,
	}
}

type UsuarioRepository struct {
	ID     uuid.UUID `db:"id"`
	Nome   string    `db:"nome"`
	Email  string    `db:"email"`
	Perfil string    `db:"perfil"`
	Ativo  bool      `db:"ativo"`
}
