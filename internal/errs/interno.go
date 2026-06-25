package errs

func NewInternalError(message string) *InternalError {
	return &InternalError{
		Message: message,
	}
}

var ErrInternal *InternalError

type InternalError struct {
	Message string
}

func (e *InternalError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return "Erro interno"
}
