package auth

import (
	"github.com/google/uuid"
)

type Usuario struct {
	ID       uuid.UUID `db:"id"`
	Nome     string    `db:"nome"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	Regra    string    `db:"regra"`
	Ativo    bool      `db:"ativo"`
}
