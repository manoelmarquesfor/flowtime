package webutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github/manoelmarquesfor/flowtime/internal/errs"
)

const (
	jsonInvalido          string = "json inválido"
	corpoRequisicaoVazio  string = "corpo da requisição vazio"
	campoDesconhecidoJson string = "campo desconhecido no json"
	tipoCampoInvalido     string = "tipo de campo inválido"
)

func DecodeJSON(body io.Reader, v any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(v); err != nil {
		return parseJSONError(v, err)
	}

	if decoder.Decode(&struct{}{}) != io.EOF {
		return errs.NewEntidadeError(jsonInvalido)
	}

	return nil
}

func jsonFieldName(value any, fieldName string) string {
	typeOf := reflect.TypeOf(value)

	if typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
	}

	if typeOf.Kind() != reflect.Struct {
		return fieldName
	}

	field, ok := typeOf.FieldByName(fieldName)
	if !ok {
		return fieldName
	}

	tag := field.Tag.Get("json")
	if tag == "" {
		return fieldName
	}

	name := strings.Split(tag, ",")[0]
	if name == "" {
		return fieldName
	}

	return name
}

func parseJSONError(value any, err error) error {
	var typeErr *json.UnmarshalTypeError
	if errors.As(err, &typeErr) {
		msg := fmt.Sprintf(
			"%s: campo '%s' tem tipo inválido, esperado '%s'",
			tipoCampoInvalido,
			jsonFieldName(value, typeErr.Field),
			typeErr.Type.String(),
		)
		return errs.NewEntidadeError(msg)
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		return errs.NewEntidadeError(jsonInvalido)
	}

	if errors.Is(err, io.EOF) {
		return errs.NewEntidadeError(corpoRequisicaoVazio)
	}

	return err
}
