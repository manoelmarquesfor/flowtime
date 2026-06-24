package errs

type UnauthorizedError struct{}

func (e *UnauthorizedError) Error() string {
	return "Acesso não autorizado"
}

var ErrUnauthorized = &UnauthorizedError{}

func NewUnauthorizedError() *UnauthorizedError {
	return ErrUnauthorized
}

type InvalidTokenError struct {
	Message string
}

func (e *InvalidTokenError) Error() string {
	if e.Message != "" {
		return e.Message
	}

	return "Token inválido"
}

func NewInvalidTokenError(message string) *InvalidTokenError {
	return &InvalidTokenError{
		Message: message,
	}
}
