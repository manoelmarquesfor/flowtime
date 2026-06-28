package colaborador

import "github.com/google/uuid"

type CreateRequest struct {
	Nome      string `json:"nome"      validate:"required,min=1,max=255"`
	Matricula string `json:"matricula" validate:"required,min=1,max=15"`
	Tag       string `json:"tag"       validate:"required,min=1,max=15"`
	Setor     string `json:"setor"     validate:"required,min=1,max=30"`
}

type Colaborador struct {
	ID        uuid.UUID `json:"id"`
	Nome      string    `json:"nome"`
	Matricula string    `json:"matricula"`
	Tag       string    `json:"tag"`
	Setor     string    `json:"setor"`
	Ativo     bool      `json:"ativo"`
}

type ColaboradorRepository struct {
	ID        uuid.UUID `db:"id"`
	Nome      string    `db:"nome"`
	Matricula string    `db:"matricula"`
	Tag       string    `db:"tag_id"`
	Setor     string    `db:"setor"`
	Ativo     bool      `db:"ativo"`
}
