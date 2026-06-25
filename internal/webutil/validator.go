package webutil

import (
	"errors"
	"reflect"

	"github/manoelmarquesfor/flowtime/internal/errs"

	"github.com/go-playground/validator/v10"
)

var custonValidator = newValidator()

func newValidator() *validator.Validate {
	val := validator.New()

	val.RegisterTagNameFunc(func(field reflect.StructField) string {
		// mudando o nome do campo na mensagem de erro para o nome do campo json
		return field.Tag.Get("json")
	})

	return val
}

func ValidateStruct(model any) error {
	err := custonValidator.Struct(model)
	if err == nil {
		return nil
	}

	return tratarErroValidacao(err)
}

func tratarErroValidacao(err error) error {
	errValidator := &validator.ValidationErrors{}
	if errors.As(err, errValidator) {
		vErr := (*errValidator)[0]
		switch vErr.Tag() {
		case "required":
			return errs.NewValidationError("O campo " + vErr.Field() + " é obrigatório")
		case "max":
			return errs.NewValidationError("O campo " + vErr.Field() + " deve ter no máximo " + vErr.Param() + " caracteres")
		case "min":
			return errs.NewValidationError("O campo " + vErr.Field() + " deve ter no mínimo " + vErr.Param() + " caracteres")
		case "email":
			return errs.NewValidationError("O campo " + vErr.Field() + " não é um email válido")
		case "gte":
			return errs.NewValidationError("O campo " + vErr.Field() + " deve ser maior ou igual a " + vErr.Param())
		case "gt":
			return errs.NewValidationError("O campo " + vErr.Field() + " deve ser maior que " + vErr.Param())
		case "lte":
			return errs.NewValidationError("O campo " + vErr.Field() + " deve ser menor ou igual a " + vErr.Param())
		case "lt":
			return errs.NewValidationError("O campo " + vErr.Field() + " deve ser menor que " + vErr.Param())
		case "oneof":
			return errs.NewValidationError("O campo " + vErr.Field() + " deve ser: " + vErr.Param())
		default:
			return errs.NewValidationError("O campo " + vErr.Field() + " é inválido ")
		}
	}

	return errs.NewValidationError("Erro ao validar os dados: " + err.Error())
}
