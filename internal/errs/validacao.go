package errs

type ValidationError struct {
	Message string
}

var ErrValidation *ValidationError

func (e *ValidationError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return "Falha na validação dos dados"
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		Message: message,
	}
}
