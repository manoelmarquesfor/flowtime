package feriado

import (
	"context"
	"errors"
	"time"

	"github/manoelmarquesfor/flowtime/internal/errs"
	"github/manoelmarquesfor/flowtime/internal/helpauth"
	"github/manoelmarquesfor/flowtime/internal/webutil"
)

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

type Service struct {
	repository *Repository
}

func (s *Service) CreateFeriado(
	context context.Context,
	create CreateRequest,
	usuarioAutenticado helpauth.UsuarioAutenticado,
) error {
	if !helpauth.PerfilIsAdmin(usuarioAutenticado) {
		return errs.NewForbiddenError()
	}

	err := webutil.ValidateStruct(create)
	if err != nil {
		return errs.NewValidationError(err.Error())
	}

	data, err := time.Parse(time.DateOnly, create.Data)
	if err != nil {
		return errs.NewValidationError("Data inválida")
	}

	var notFoundErr *errs.NotFoundError

	feriado, err := s.repository.FeriadoByData(context, data)
	if err != nil && !errors.As(err, &notFoundErr) {
		return err
	}

	if feriado.Data.Equal(data) {
		return errs.NewValidationError("Feriado já cadastrado")
	}

	err = s.repository.FeriadoCreate(context, RepositoryModel{
		Data:      data,
		Descricao: create.Descricao,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetFeriados(context context.Context, ano string) ([]Response, error) {
	result := []Response{}

	inicial, err := time.Parse("2006", ano)
	if err != nil {
		return nil, errs.NewValidationError("Ano inválido")
	}

	final := inicial.AddDate(1, 0, 0)

	feriados, err := s.repository.FeriadoByPeriodo(context, inicial, final)
	if err != nil {
		return nil, err
	}

	for _, feriado := range feriados {
		result = append(result, feriado.ToResponse())
	}

	return result, nil
}

func (s *Service) DeleteFeriado(
	context context.Context,
	data string,
	usuarioAutenticado helpauth.UsuarioAutenticado,
) error {
	if !helpauth.PerfilIsAdmin(usuarioAutenticado) {
		return errs.NewForbiddenError()
	}

	parsedData, err := time.Parse(time.DateOnly, data)
	if err != nil {
		return errs.NewValidationError("Data inválida")
	}

	err = s.repository.FeriadoDelete(context, parsedData)
	if err != nil {
		return err
	}

	return nil
}
