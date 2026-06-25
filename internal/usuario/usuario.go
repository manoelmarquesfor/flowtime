package usuario

import (
	"time"

	"github.com/google/uuid"
)

type UsuarioCreate struct {
	Nome     string `json:"nome"     validate:"required,min=1,max=100"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Perfil   string `json:"perfil"    validate:"required,oneof=ADMIN USER"`
}

type UsuarioCreateRepository struct {
	ID        uuid.UUID `db:"id"`
	Nome      string    `db:"nome"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Perfil    string    `db:"perfil"`
	Ativo     bool      `db:"ativo"`
	DtCreated time.Time `db:"dt_criacao"`
}

type UsuarioRepository struct {
	ID     uuid.UUID `db:"id"`
	Nome   string    `db:"nome"`
	Email  string    `db:"email"`
	Perfil string    `db:"perfil"`
	Ativo  bool      `db:"ativo"`
}

type Usuario struct {
	ID     uuid.UUID `json:"id"`
	Nome   string    `json:"nome"`
	Email  string    `json:"email"`
	Perfil string    `json:"perfil"`
	Ativo  bool      `json:"ativo"`
}

type DeleteUsuarioResponse struct {
	ID     string `json:"id"`
	Detail string `json:"detail"`
}
