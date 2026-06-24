package auth

import (
	"context"

	"github/manoelmarquesfor/flowtime/internal/errs"

	"github.com/google/uuid"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Login(
	ctx context.Context,
	email string,
	password string,
) (string, error) {
	if email == "" || password == "" {
		return "", errs.NewValidationError("Email e senha são obrigatórios")
	}

	return uuid.NewString(), nil
}

func (s *Service) Logout(
	ctx context.Context,
	sessionID string,
) error {
	// remover sessão

	return nil
}
