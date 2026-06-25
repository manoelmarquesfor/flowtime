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

	switch {
	case errors.As(erro, &errs.ErrValidation):
		statusCode = http.StatusBadRequest
		msg = erro.Error()

	case errors.As(erro, &errs.ErrUnauthorized):
		statusCode = http.StatusUnauthorized
		msg = erro.Error()

	case errors.As(erro, &errs.ErrNotFound):
		statusCode = http.StatusNotFound
		msg = erro.Error()

	case errors.As(erro, &errs.ErrInternal) || errors.As(erro, &errs.ErrRepository):
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
