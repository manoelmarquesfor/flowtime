package errs

func NewInternalError(message string) *InternalError {
	return &InternalError{
		Message: message,
	}
}

type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return "Erro interno"
}
