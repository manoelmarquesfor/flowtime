package usuario

import (
	"context"
	"errors"

	"github/manoelmarquesfor/flowtime/internal/constantes"
	"github/manoelmarquesfor/flowtime/internal/errs"
	"github/manoelmarquesfor/flowtime/internal/helpauth"
	"github/manoelmarquesfor/flowtime/internal/webutil"

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

func (s *Service) get(
	ctx context.Context,
	idUsuario string,
	usuarioAutenticado helpauth.UsuarioAutencicado,
) (Usuario, error) {
	var result Usuario

	if err := s.validarPermissao(usuarioAutenticado); err != nil {
		return result, err
	}

	userID, err := uuid.Parse(idUsuario)
	if err != nil {
		return result, errs.NewValidationError("ID inválido")
	}

	usuarioRepository, err := s.repository.UsuarioByID(ctx, userID)
	if err != nil {
		return result, err
	}

	return usuarioRepository.ToUsuario(), nil
}

func (s *Service) getAll(
	ctx context.Context,
	usuarioAutenticado helpauth.UsuarioAutencicado,
) ([]Usuario, error) {
	result := []Usuario{}

	if err := s.validarPermissao(usuarioAutenticado); err != nil {
		return result, err
	}

	usuariosRepository, err := s.repository.UsuarioAll(ctx)
	if err != nil {
		return result, err
	}

	for _, usuarioRepository := range usuariosRepository {
		result = append(result, usuarioRepository.ToUsuario())
	}

	return result, nil
}

func (s *Service) create(
	ctx context.Context,
	usuario UsuarioCreate,
	usuarioAutenticado helpauth.UsuarioAutencicado,
) (Usuario, error) {
	var result Usuario

	if err := s.validarPermissao(usuarioAutenticado); err != nil {
		return result, err
	}

	err := webutil.ValidateStruct(usuario)
	if err != nil {
		return result, errs.NewValidationError(err.Error())
	}

	newSenha, err := bcrypt.GenerateFromPassword([]byte(usuario.Password), bcrypt.DefaultCost)
	if err != nil {
		return result, errs.NewValidationError("Erro ao gerar senha")
	}

	var errNotFound *errs.NotFoundError

	usuarioRepository, err := s.repository.UsuarioByEmail(ctx, usuario.Email)
	if err != nil && !errors.As(err, &errNotFound) {
		return result, err
	}

	if usuarioRepository.ID != uuid.Nil {
		return result, errs.NewValidationError("Email já cadastrado")
	}

	usuarioCreate := usuario.UsuarioCreateRepository(string(newSenha))

	err = s.repository.UsuarioCreate(ctx, usuarioCreate)
	if err != nil {
		return result, err
	}

	return usuarioCreate.ToUsuario(), nil
}

func (s *Service) delete(
	ctx context.Context,
	idUsuario string,
	usuarioAutenticado helpauth.UsuarioAutencicado,
) error {
	if err := s.validarPermissao(usuarioAutenticado); err != nil {
		return err
	}

	userID, err := uuid.Parse(idUsuario)
	if err != nil {
		return errs.NewValidationError("ID inválido")
	}

	err = s.repository.UsuarioDelete(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) validarPermissao(usuarioAutenticado helpauth.UsuarioAutencicado) error {
	if usuarioAutenticado.Regra != constantes.RegraAdmin {
		return errs.NewForbiddenError()
	}

	return nil
}
