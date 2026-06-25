package usuario

import (
	"time"

	"github.com/google/uuid"
)

type UsuarioCreate struct {
	Nome     string `json:"nome"     validate:"required,min=1,max=100"`
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Regra    string `json:"regra"    validate:"required,oneof=ADMIN USER"`
}

type UsuarioCreateRepository struct {
	ID        uuid.UUID `db:"id"`
	Nome      string    `db:"nome"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Regra     string    `db:"regra"`
	Ativo     bool      `db:"ativo"`
	DtCreated time.Time `db:"dt_criacao"`
}

type UsuarioRepository struct {
	ID    uuid.UUID `db:"id"`
	Nome  string    `db:"nome"`
	Email string    `db:"email"`
	Regra string    `db:"regra"`
	Ativo bool      `db:"ativo"`
}

type Usuario struct {
	ID    uuid.UUID `json:"id"`
	Nome  string    `json:"nome"`
	Email string    `json:"email"`
	Regra string    `json:"regra"`
	Ativo bool      `json:"ativo"`
}

type DeleteUsuarioResponse struct {
	ID     string `json:"id"`
	Detail string `json:"detail"`
}
