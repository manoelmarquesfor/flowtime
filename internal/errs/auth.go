package errs

type UnauthorizedError struct{}

func (e *UnauthorizedError) Error() string {
	return "Acesso não autorizado"
}

func NewUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{}
}

type ForbiddenError struct{}

func (e *ForbiddenError) Error() string {
	return "Acesso proibido"
}

func NewForbiddenError() *ForbiddenError {
	return &ForbiddenError{}
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
