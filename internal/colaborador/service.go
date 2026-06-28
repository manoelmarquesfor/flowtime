package colaborador

import (
	"context"
	"errors"
	"strconv"

	"github/manoelmarquesfor/flowtime/internal/errs"
	"github/manoelmarquesfor/flowtime/internal/helpauth"
	"github/manoelmarquesfor/flowtime/internal/webutil"

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

func (s *Service) get(
	ctx context.Context,
	idColaborador string,
	usuarioAutenticado helpauth.UsuarioAutenticado,
) (Colaborador, error) {
	var result Colaborador

	if !helpauth.PerfilIsAdmin(usuarioAutenticado) {
		return result, errs.NewForbiddenError()
	}

	colaboradorID, err := uuid.Parse(idColaborador)
	if err != nil {
		return result, errs.NewValidationError("ID inválido")
	}

	colaboradorRepository, err := s.repository.ByID(ctx, colaboradorID)
	if err != nil {
		return result, err
	}

	return colaboradorRepository.toColaborador(), nil
}

func (s *Service) getAll(
	ctx context.Context,
	usuarioAutenticado helpauth.UsuarioAutenticado,
	ativo string,
) ([]Colaborador, error) {
	result := []Colaborador{}

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

	colaboradoresRepository, err := s.repository.All(ctx, &ativoBool)
	if err != nil {
		return result, err
	}

	for _, colaboradorRepository := range colaboradoresRepository {
		result = append(result, colaboradorRepository.toColaborador())
	}

	return result, nil
}

func (s *Service) create(
	ctx context.Context,
	colaborador CreateRequest,
	usuarioAutenticado helpauth.UsuarioAutenticado,
) (Colaborador, error) {
	var result Colaborador

	if !helpauth.PerfilIsAdmin(usuarioAutenticado) {
		return result, errs.NewForbiddenError()
	}

	err := webutil.ValidateStruct(colaborador)
	if err != nil {
		return result, errs.NewValidationError(err.Error())
	}

	var errNotFound *errs.NotFoundError

	colaboradoresRepository, err := s.repository.All(ctx, nil)
	if err != nil && !errors.As(err, &errNotFound) {
		return result, err
	}

	colaboradorCreate := colaborador.toRepositoryCreate()

	existTagCadastrada := false
	existMatriculaCadastrada := false

	for _, colab := range colaboradoresRepository {
		if colab.Tag == colaboradorCreate.Tag {
			existTagCadastrada = true
			break
		}

		if colab.Matricula == colaboradorCreate.Matricula {
			existMatriculaCadastrada = true
			break
		}
	}

	if existMatriculaCadastrada {
		return result, errs.NewValidationError("Matricula ja cadastrada.")
	}

	if existTagCadastrada {
		return result, errs.NewValidationError("Tag ja cadastrada")
	}

	err = s.repository.Create(ctx, colaboradorCreate)
	if err != nil {
		return result, err
	}
	return colaboradorCreate.toColaborador(), nil
}

func (s *Service) alterarSituacao(
	ctx context.Context,
	idColaborador string,
	usuarioAutenticado helpauth.UsuarioAutenticado,
) (Colaborador, error) {
	var result Colaborador

	if !helpauth.PerfilIsAdmin(usuarioAutenticado) {
		return result, errs.NewForbiddenError()
	}

	colaboradorID, err := uuid.Parse(idColaborador)
	if err != nil {
		return result, errs.NewValidationError("ID inválido")
	}

	colaboradorRepository, err := s.repository.ByID(ctx, colaboradorID)
	if err != nil {
		return result, err
	}

	colaboradorRepository.Ativo = !colaboradorRepository.Ativo

	err = s.repository.Update(ctx, colaboradorRepository)
	if err != nil {
		return result, err
	}

	return colaboradorRepository.toColaborador(), nil
}
