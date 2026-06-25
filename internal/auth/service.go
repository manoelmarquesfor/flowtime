package auth

import (
	"context"
	"errors"

	"github/manoelmarquesfor/flowtime/internal/errs"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
) (uuid.UUID, error) {
	if email == "" || password == "" {
		return uuid.Nil, errs.NewValidationError("Email e senha são obrigatórios")
	}

	usuario, err := s.repository.UsuarioByEmail(ctx, email)
	if err != nil {
		if errors.As(err, &errs.ErrNotFound) {
			return uuid.Nil, errs.NewValidationError("Email ou senha incorreto")
		}

		return uuid.Nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(password))
	if err != nil {
		return uuid.Nil, errs.NewValidationError("Email ou senha incorreto")
	}

	if !usuario.Ativo {
		return uuid.Nil, errs.NewValidationError("Usuário inativo")
	}

	sessaoID := uuid.New()

	err = s.repository.SaveSessao(ctx, &sessaoID, usuario.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return sessaoID, nil
}

func (s *Service) Logout(
	ctx context.Context,
	sessionID uuid.UUID,
) error {
	err := s.repository.SaveSessao(ctx, nil, sessionID)
	if err != nil {
		return err
	}

	return nil
}
