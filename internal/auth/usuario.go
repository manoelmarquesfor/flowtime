package auth

import (
	"time"

	"github.com/google/uuid"
)

type Usuario struct {
	ID        uuid.UUID  `sqlx:"id"`
	Nome      string     `sqlx:"nome"`
	Email     string     `sqlx:"email"`
	Password  string     `sqlx:"password"`
	SessaoID  *uuid.UUID `sqlx:"sessao_id"`
	Regra     string     `sqlx:"regra"`
	Ativo     bool       `sqlx:"ativo"`
	DtCriacao time.Time  `sqlx:"dt_criacao"`
}
