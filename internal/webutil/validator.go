package webutil

import (
	"errors"
	"reflect"

	"github/manoelmarquesfor/flowtime/internal/errs"

	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_br_translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	custonValidator = newValidator()
	uni             *ut.UniversalTranslator
	translator      ut.Translator
)

func newValidator() *validator.Validate {
	val := validator.New()

	ptBr := pt_BR.New()
	uni = ut.New(ptBr, ptBr)

	translator, _ = uni.GetTranslator(ptBr.Locale())

	pt_br_translations.RegisterDefaultTranslations(val, translator)

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
		return errs.NewValidationError(vErr.Translate(translator))
	}

	return errs.NewValidationError("Erro ao validar os dados: " + err.Error())
}
