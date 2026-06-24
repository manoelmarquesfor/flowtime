package webutil

import (
	"errors"
	"net/http"

	"github/manoelmarquesfor/flowtime/internal/errs"
)

var (
	errValidation *errs.ValidationError
	errInternal   *errs.InternalError
)

func ResponseError(writer http.ResponseWriter, erro error) {
	var statusCode int

	switch {
	case errors.As(erro, &errValidation):
		statusCode = http.StatusBadRequest
	case errors.As(erro, &errs.ErrUnauthorized):
		statusCode = http.StatusUnauthorized

	case errors.As(erro, &errInternal):
		statusCode = http.StatusInternalServerError
	default:
		statusCode = http.StatusInternalServerError
	}

	detailResponse := struct {
		Detail string `json:"detail"`
	}{
		Detail: erro.Error(),
	}

	Response(writer, statusCode, detailResponse)
}
