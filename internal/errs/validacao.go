package errs

type ValidationError struct {
	Message string
}

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

type EntidadeError struct {
	Message string
}

func (e *EntidadeError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return "Falha no decode dos dados"
}

func NewEntidadeError(message string) *EntidadeError {
	return &EntidadeError{
		Message: message,
	}
}
