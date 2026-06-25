package webutil

import (
	"errors"
	"log"
	"net/http"

	"github/manoelmarquesfor/flowtime/internal/errs"
)

func ResponseError(writer http.ResponseWriter, erro error) {
	var statusCode int

	msg := "Erro interno do servidor"

	var errValidation *errs.ValidationError

	var errUnauthorized *errs.UnauthorizedError

	var errNotFound *errs.NotFoundError

	var errInternal *errs.InternalError

	var errRepository *errs.RepositoryError

	switch {
	case errors.As(erro, &errValidation):
		statusCode = http.StatusBadRequest
		msg = erro.Error()

	case errors.As(erro, &errUnauthorized):
		statusCode = http.StatusUnauthorized
		msg = erro.Error()

	case errors.As(erro, &errNotFound):
		statusCode = http.StatusNotFound
		msg = erro.Error()

	case errors.As(erro, &errInternal) || errors.As(erro, &errRepository):
		statusCode = http.StatusInternalServerError

		log.Println(erro)

	default:
		statusCode = http.StatusInternalServerError
	}

	detailResponse := struct {
		Detail string `json:"detail"`
	}{
		Detail: msg,
	}

	Response(writer, statusCode, detailResponse)
}
