package usuario

import (
	"context"
	"errors"
	"strconv"

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
	usuarioAutenticado helpauth.UsuarioAutenticado,
) (Usuario, error) {
	var result Usuario

	if !helpauth.PerfilIsAdmin(usuarioAutenticado) {
		return result, errs.NewForbiddenError()
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
	usuarioAutenticado helpauth.UsuarioAutenticado,
	ativo string,
) ([]Usuario, error) {
	result := []Usuario{}

	if !helpauth.PerfilIsAdmin(usuarioAutenticado) {
		return result, errs.NewForbiddenError()
	}

	ativoBool := true

	if ativo != "" {
		var err error
		ativoBool, err = strconv.ParseBool(ativo)
		if err != nil {
			return result, errs.NewValidationError("Erro ao converter parâmetro 'ativo' para booleano")
		}
	}

	usuariosRepository, err := s.repository.UsuarioAll(ctx, &ativoBool)
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
	usuarioAutenticado helpauth.UsuarioAutenticado,
) (Usuario, error) {
	var result Usuario

	if !helpauth.PerfilIsAdmin(usuarioAutenticado) {
		return result, errs.NewForbiddenError()
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
	usuarioAutenticado helpauth.UsuarioAutenticado,
) error {
	if !helpauth.PerfilIsAdmin(usuarioAutenticado) {
		return errs.NewForbiddenError()
	}

	userID, err := uuid.Parse(idUsuario)
	if err != nil {
		return errs.NewValidationError("ID inválido")
	}

	usuarios, err := s.repository.UsuarioAll(ctx, nil)
	if err != nil {
		return err
	}

	var idUsuarioMaster *uuid.UUID

	for _, u := range usuarios {
		if u.Perfil == constantes.PerfilMaster {
			idUsuarioMaster = &u.ID
		}
	}

	if idUsuarioMaster != nil && *idUsuarioMaster == userID {
		return errs.NewValidationError("Não é permitido excluir usuário master, apenas inativar")
	}

	if !s.possuiUsuariosAdminAtivo(usuarios) {
		return errs.NewValidationError("Não é permitido excluir o último usuário admin ativo")
	}

	err = s.repository.UsuarioDelete(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) alterarSituacao(
	ctx context.Context,
	idUsuario string,
	usuarioAutenticado helpauth.UsuarioAutenticado,
) (Usuario, error) {
	var result Usuario

	if !helpauth.PerfilIsAdmin(usuarioAutenticado) {
		return result, errs.NewForbiddenError()
	}

	userID, err := uuid.Parse(idUsuario)
	if err != nil {
		return result, errs.NewValidationError("ID inválido")
	}

	usuarios, err := s.repository.UsuarioAll(ctx, nil)
	if err != nil {
		return result, err
	}

	if !s.possuiUsuariosAdminAtivo(usuarios) {
		return result, errs.NewValidationError("Não é permitido alterar a situação do último usuário admin ativo")
	}

	usuarioRepository, err := s.repository.UsuarioByID(ctx, userID)
	if err != nil {
		return result, err
	}

	usuarioRepository.Ativo = !usuarioRepository.Ativo

	err = s.repository.UsuarioUpdate(ctx, usuarioRepository)
	if err != nil {
		return result, err
	}

	return usuarioRepository.ToUsuario(), nil
}

func (s *Service) possuiUsuariosAdminAtivo(usuarios []UsuarioRepository) bool {
	qtdUsuariosAdminAtivos := 0

	for _, u := range usuarios {
		if u.Perfil == constantes.PerfilAdmin && u.Ativo {
			qtdUsuariosAdminAtivos++
		}
	}

	return qtdUsuariosAdminAtivos > 0
}
