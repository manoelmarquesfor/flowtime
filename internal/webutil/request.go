package webutil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
)

var (
	ErrInvalidJSON      = errors.New("json inválido")
	ErrEmptyBody        = errors.New("corpo da requisição vazio")
	ErrUnknownField     = errors.New("campo desconhecido no json")
	ErrInvalidFieldType = errors.New("tipo de campo inválido")
)

func DecodeJSON(body io.Reader, v any) error {
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(v); err != nil {
		return parseJSONError(v, err)
	}

	if decoder.Decode(&struct{}{}) != io.EOF {
		return ErrInvalidJSON
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
		return fmt.Errorf(
			"%w: campo '%s' tem tipo inválido, esperado '%s'",
			ErrInvalidFieldType,
			jsonFieldName(value, typeErr.Field),
			typeErr.Type.String(),
		)
	}

	var syntaxErr *json.SyntaxError
	if errors.As(err, &syntaxErr) {
		return ErrInvalidJSON
	}

	if errors.Is(err, io.EOF) {
		return ErrEmptyBody
	}

	return err
}
